// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ModelQualityJobDefinitionSpec defines the desired state of ModelQualityJobDefinition.
type ModelQualityJobDefinitionSpec struct {

	// The name of the monitoring job definition.
	// +kubebuilder:validation:Required
	JobDefinitionName *string `json:"jobDefinitionName"`
	// +kubebuilder:validation:Required
	JobResources *MonitoringResources `json:"jobResources"`
	// The container that runs the monitoring job.
	// +kubebuilder:validation:Required
	ModelQualityAppSpecification *ModelQualityAppSpecification `json:"modelQualityAppSpecification"`
	// Specifies the constraints and baselines for the monitoring job.
	ModelQualityBaselineConfig *ModelQualityBaselineConfig `json:"modelQualityBaselineConfig,omitempty"`
	// A list of the inputs that are monitored. Currently endpoints are supported.
	// +kubebuilder:validation:Required
	ModelQualityJobInput *ModelQualityJobInput `json:"modelQualityJobInput"`
	// +kubebuilder:validation:Required
	ModelQualityJobOutputConfig *MonitoringOutputConfig `json:"modelQualityJobOutputConfig"`
	// Specifies the network configuration for the monitoring job.
	NetworkConfig *MonitoringNetworkConfig `json:"networkConfig,omitempty"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume
	// to perform tasks on your behalf.
	// +kubebuilder:validation:Required
	RoleARN           *string                      `json:"roleARN"`
	StoppingCondition *MonitoringStoppingCondition `json:"stoppingCondition,omitempty"`
	// (Optional) An array of key-value pairs. For more information, see Using Cost
	// Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL)
	// in the Amazon Web Services Billing and Cost Management User Guide.
	Tags []*Tag `json:"tags,omitempty"`
}

// ModelQualityJobDefinitionStatus defines the observed state of ModelQualityJobDefinition
type ModelQualityJobDefinitionStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
}

// ModelQualityJobDefinition is the Schema for the ModelQualityJobDefinitions API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type ModelQualityJobDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ModelQualityJobDefinitionSpec   `json:"spec,omitempty"`
	Status            ModelQualityJobDefinitionStatus `json:"status,omitempty"`
}

// ModelQualityJobDefinitionList contains a list of ModelQualityJobDefinition
// +kubebuilder:object:root=true
type ModelQualityJobDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ModelQualityJobDefinition `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ModelQualityJobDefinition{}, &ModelQualityJobDefinitionList{})
}
