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

// PipelineSpec defines the desired state of Pipeline.
//
// A SageMaker Model Building Pipeline instance.
type PipelineSpec struct {
	// This is the configuration that controls the parallelism of the pipeline.
	// If specified, it applies to all runs of this pipeline by default.
	ParallelismConfiguration *ParallelismConfiguration `json:"parallelismConfiguration,omitempty"`
	// The JSON pipeline definition of the pipeline.
	PipelineDefinition *string `json:"pipelineDefinition,omitempty"`
	// A description of the pipeline.
	PipelineDescription *string `json:"pipelineDescription,omitempty"`
	// The display name of the pipeline.
	PipelineDisplayName *string `json:"pipelineDisplayName,omitempty"`
	// The name of the pipeline.
	// +kubebuilder:validation:Required
	PipelineName *string `json:"pipelineName"`
	// The Amazon Resource Name (ARN) of the role used by the pipeline to access
	// and create resources.
	// +kubebuilder:validation:Required
	RoleARN *string `json:"roleARN"`
	// A list of tags to apply to the created pipeline.
	Tags []*Tag `json:"tags,omitempty"`
}

// PipelineStatus defines the observed state of Pipeline
type PipelineStatus struct {
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
	// The time when the pipeline was created.
	// +kubebuilder:validation:Optional
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	// The time when the pipeline was last modified.
	// +kubebuilder:validation:Optional
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
	// The status of the pipeline execution.
	// +kubebuilder:validation:Optional
	PipelineStatus *string `json:"pipelineStatus,omitempty"`
}

// Pipeline is the Schema for the Pipelines API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type=string,priority=0,JSONPath=`.status.pipelineStatus`
type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PipelineSpec   `json:"spec,omitempty"`
	Status            PipelineStatus `json:"status,omitempty"`
}

// PipelineList contains a list of Pipeline
// +kubebuilder:object:root=true
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pipeline `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Pipeline{}, &PipelineList{})
}