# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.
"""Integration tests for the SageMaker pipelineExecution API.
"""

from http import client
import botocore
import pytest
import logging
from typing import Dict

from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from e2e import (
    service_marker,
    wait_for_status,
    create_sagemaker_resource,
    get_sagemaker_pipeline_execution,
    sagemaker_client,
    assert_tags_in_sync,
)
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e.common import config as cfg

RESOURCE_PLURAL = "pipelineexecutions"

DELETE_WAIT_PERIOD = 20
DELETE_WAIT_LENGTH = 30


@pytest.fixture(scope="function")
def pipeline_execution():
    resource_name = random_suffix_name("pipeline-execution", 28)
    client_request_token = random_suffix_name("client-request-token", 38)
    replacements = REPLACEMENT_VALUES.copy()
    replacements["PIPELINE_EXECUTION_RESOURCE_NAME"] = resource_name
    replacements["PIPELINE_EXECUTION_CLIENT_REQUEST_TOKEN"] = client_request_token
    (
        pipeline_execution_reference,
        pipeline_execution_spec,
        pipeline_execution_resource,
    ) = create_sagemaker_resource(
        resource_plural=cfg.pipeline_execution_resource_PLURAL,
        resource_name=resource_name,
        spec_file="pipeline_execution",
        replacements=replacements,
    )
    assert pipeline_execution_resource is not None
    if k8s.get_resource_arn(pipeline_execution_resource) is None:
        logging.error(
            f"ARN for this resource is None, resource status is: {pipeline_execution_resource['status']}"
        )
    assert k8s.get_resource_arn(pipeline_execution_resource) is not None

    yield (pipeline_execution_reference, pipeline_execution_resource)

    # Delete the k8s resource if not already deleted by tests
    if k8s.get_resource_exists(pipeline_execution_reference):
        _, deleted = k8s.delete_custom_resource(
            pipeline_execution_reference, DELETE_WAIT_PERIOD, DELETE_WAIT_LENGTH
        )
        assert deleted


def get_sagemaker_pipeline_execution_status(pipeline_execution_arn: str):
    sm_pipeline_execution_desc = get_sagemaker_pipeline_execution(
        pipeline_execution_arn
    )
    return sm_pipeline_execution_desc["PipelineExecutionStatus"]


def get_pipeline_execution_resource_status(reference: k8s.CustomResourceReference):
    resource = k8s.get_resource(reference)
    assert "pipelineExecutionStatus" in resource["status"]
    return resource["status"]["pipelineExecutionStatus"]


@pytest.mark.canary
@service_marker
class TestpipelineExecution:
    def _wait_resource_pipeline_execution_status(
        self,
        reference: k8s.CustomResourceReference,
        expected_status: str,
        wait_periods: int = 30,
        period_length: int = 30,
    ):
        return wait_for_status(
            expected_status,
            wait_periods,
            period_length,
            get_pipeline_execution_resource_status,
            reference,
        )

    def _wait_sagemaker_pipeline_execution_status(
        self,
        pipeline_execution_arn,
        expected_status: str,
        wait_periods: int = 30,
        period_length: int = 30,
    ):
        return wait_for_status(
            expected_status,
            wait_periods,
            period_length,
            get_sagemaker_pipeline_execution_status,
            pipeline_execution_arn,
        )

    def _assert_pipeline_execution_status_in_sync(
        self, pipeline_execution_arn, reference, expected_status
    ):
        assert (
            self._wait_sagemaker_pipeline_execution_status(
                pipeline_execution_arn, expected_status
            )
            == self._wait_resource_pipeline_execution_status(reference, expected_status)
            == expected_status
        )

    def test_pipeline_execution_succeeded(self, pipeline_execution):
        (reference, spec, resource) = pipeline_execution
        assert k8s.get_resource_exists(reference)

        pipeline_name = resource["spec"].get("pipelineName")
        # Need PipelineExecutionArn to reference the resource
        pipeline_execution_arn = sagemaker_client().list_pipeline_executions(
            PipelineName=pipeline_name
        )["PipelineExecutionSummaries"][0]["PipelineExecutionArn"]

        pipeline_execution_desc = get_sagemaker_pipeline_execution(
            pipeline_execution_arn
        )
        if k8s.get_resource_arn(resource) is None:
            logging.error(
                f"ARN for this resource is None, resource status is: {resource['status']}"
            )

        assert k8s.get_resource_arn(resource) == pipeline_execution_arn

        self._assert_pipeline_execution_status_in_sync(
            pipeline_execution_arn, reference, cfg.JOB_STATUS_INPROGRESS
        )
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "False")

        self._assert_pipeline_execution_status_in_sync(
            pipeline_execution_arn, reference, cfg.JOB_STATUS_COMPLETED
        )
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True")

        # Update the resource
        new_pipeline_execution_display_name = random_suffix_name(
            "updated-display-name", 38
        )
        spec["spec"][
            "pipelineExecutionDisplayName"
        ] = new_pipeline_execution_display_name
        resource = k8s.patch_custom_resource(reference, spec)
        resource = k8s.wait_resource_consumed_by_controller(reference)
        assert resource is not None

        self._assert_pipeline_excution_status_in_sync(
            pipeline_execution_arn, reference, cfg.JOB_STATUS_COMPLETED
        )
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True")

        pipeline_execution_desc = get_sagemaker_pipeline_execution(
            pipeline_execution_arn
        )

        assert (
            pipeline_execution_desc["PipelineExecutionDisplayName"]
            == new_pipeline_execution_display_name
        )
        assert (
            resource["spec"].get("pipelineExecutionDisplayName", None)
            == new_pipeline_execution_display_name
        )

        # Check that you can delete a completed resource from k8s
        _, deleted = k8s.delete_custom_resource(
            reference, DELETE_WAIT_PERIOD, DELETE_WAIT_LENGTH
        )
        assert deleted is True
