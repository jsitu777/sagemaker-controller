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

// TransformJobSpec defines the desired state of TransformJob.
//
// A batch transform job. For information about SageMaker batch transform, see
// Use Batch Transform (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform.html).
type TransformJobSpec struct {
	// Specifies the number of records to include in a mini-batch for an HTTP inference
	// request. A record is a single unit of input data that inference can be made
	// on. For example, a single line in a CSV file is a record.
	//
	// To enable the batch strategy, you must set the SplitType property to Line,
	// RecordIO, or TFRecord.
	//
	// To use only one record when making an HTTP invocation request to a container,
	// set BatchStrategy to SingleRecord and SplitType to Line.
	//
	// To fit as many records in a mini-batch as can fit within the MaxPayloadInMB
	// limit, set BatchStrategy to MultiRecord and SplitType to Line.
	BatchStrategy *string `json:"batchStrategy,omitempty"`
	// The data structure used to specify the data to be used for inference in a
	// batch transform job and to associate the data that is relevant to the prediction
	// results in the output. The input filter provided allows you to exclude input
	// data that is not needed for inference in a batch transform job. The output
	// filter provided allows you to include input data relevant to interpreting
	// the predictions in the output from the job. For more information, see Associate
	// Prediction Results with their Corresponding Input Records (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform-data-processing.html).
	DataProcessing *DataProcessing `json:"dataProcessing,omitempty"`
	// The environment variables to set in the Docker container. We support up to
	// 16 key and values entries in the map.
	Environment map[string]*string `json:"environment,omitempty"`

	ExperimentConfig *ExperimentConfig `json:"experimentConfig,omitempty"`
	// The maximum number of parallel requests that can be sent to each instance
	// in a transform job. If MaxConcurrentTransforms is set to 0 or left unset,
	// Amazon SageMaker checks the optional execution-parameters to determine the
	// settings for your chosen algorithm. If the execution-parameters endpoint
	// is not enabled, the default value is 1. For more information on execution-parameters,
	// see How Containers Serve Requests (https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-batch-code.html#your-algorithms-batch-code-how-containe-serves-requests).
	// For built-in algorithms, you don't need to set a value for MaxConcurrentTransforms.
	MaxConcurrentTransforms *int64 `json:"maxConcurrentTransforms,omitempty"`
	// The maximum allowed size of the payload, in MB. A payload is the data portion
	// of a record (without metadata). The value in MaxPayloadInMB must be greater
	// than, or equal to, the size of a single record. To estimate the size of a
	// record in MB, divide the size of your dataset by the number of records. To
	// ensure that the records fit within the maximum payload size, we recommend
	// using a slightly larger value. The default value is 6 MB.
	//
	// The value of MaxPayloadInMB cannot be greater than 100 MB. If you specify
	// the MaxConcurrentTransforms parameter, the value of (MaxConcurrentTransforms
	// * MaxPayloadInMB) also cannot exceed 100 MB.
	//
	// For cases where the payload might be arbitrarily large and is transmitted
	// using HTTP chunked encoding, set the value to 0. This feature works only
	// in supported algorithms. Currently, Amazon SageMaker built-in algorithms
	// do not support HTTP chunked encoding.
	MaxPayloadInMB *int64 `json:"maxPayloadInMB,omitempty"`
	// Configures the timeout and maximum number of retries for processing a transform
	// job invocation.
	ModelClientConfig *ModelClientConfig `json:"modelClientConfig,omitempty"`
	// The name of the model that you want to use for the transform job. ModelName
	// must be the name of an existing Amazon SageMaker model within an Amazon Web
	// Services Region in an Amazon Web Services account.
	// +kubebuilder:validation:Required
	ModelName *string `json:"modelName"`
	// (Optional) An array of key-value pairs. For more information, see Using Cost
	// Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what)
	// in the Amazon Web Services Billing and Cost Management User Guide.
	Tags []*Tag `json:"tags,omitempty"`
	// Describes the input source and the way the transform job consumes it.
	// +kubebuilder:validation:Required
	TransformInput *TransformInput `json:"transformInput"`
	// The name of the transform job. The name must be unique within an Amazon Web
	// Services Region in an Amazon Web Services account.
	// +kubebuilder:validation:Required
	TransformJobName *string `json:"transformJobName"`
	// Describes the results of the transform job.
	// +kubebuilder:validation:Required
	TransformOutput *TransformOutput `json:"transformOutput"`
	// Describes the resources, including ML instance types and ML instance count,
	// to use for the transform job.
	// +kubebuilder:validation:Required
	TransformResources *TransformResources `json:"transformResources"`
}

// TransformJobStatus defines the observed state of TransformJob
type TransformJobStatus struct {
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
	// If the transform job failed, FailureReason describes why it failed. A transform
	// job creates a log file, which includes error messages, and stores it as an
	// Amazon S3 object. For more information, see Log Amazon SageMaker Events with
	// Amazon CloudWatch (https://docs.aws.amazon.com/sagemaker/latest/dg/logging-cloudwatch.html).
	// +kubebuilder:validation:Optional
	FailureReason *string `json:"failureReason,omitempty"`
	// The status of the transform job. If the transform job failed, the reason
	// is returned in the FailureReason field.
	// +kubebuilder:validation:Optional
	TransformJobStatus *string `json:"transformJobStatus,omitempty"`
}

// TransformJob is the Schema for the TransformJobs API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="FAILURE-REASON",type=string,priority=1,JSONPath=`.status.failureReason`
// +kubebuilder:printcolumn:name="STATUS",type=string,priority=0,JSONPath=`.status.transformJobStatus`
type TransformJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransformJobSpec   `json:"spec,omitempty"`
	Status            TransformJobStatus `json:"status,omitempty"`
}

// TransformJobList contains a list of TransformJob
// +kubebuilder:object:root=true
type TransformJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransformJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TransformJob{}, &TransformJobList{})
}
