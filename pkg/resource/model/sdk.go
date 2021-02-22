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

package model

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.SageMaker{}
	_ = &svcapitypes.Model{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.DescribeModelWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeModel", respErr)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "ValidationException" && strings.HasPrefix(awsErr.Message(), "Could not find model") {
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.Containers != nil {
		f0 := []*svcapitypes.ContainerDefinition{}
		for _, f0iter := range resp.Containers {
			f0elem := &svcapitypes.ContainerDefinition{}
			if f0iter.ContainerHostname != nil {
				f0elem.ContainerHostname = f0iter.ContainerHostname
			}
			if f0iter.Environment != nil {
				f0elemf1 := map[string]*string{}
				for f0elemf1key, f0elemf1valiter := range f0iter.Environment {
					var f0elemf1val string
					f0elemf1val = *f0elemf1valiter
					f0elemf1[f0elemf1key] = &f0elemf1val
				}
				f0elem.Environment = f0elemf1
			}
			if f0iter.Image != nil {
				f0elem.Image = f0iter.Image
			}
			if f0iter.ImageConfig != nil {
				f0elemf3 := &svcapitypes.ImageConfig{}
				if f0iter.ImageConfig.RepositoryAccessMode != nil {
					f0elemf3.RepositoryAccessMode = f0iter.ImageConfig.RepositoryAccessMode
				}
				f0elem.ImageConfig = f0elemf3
			}
			if f0iter.Mode != nil {
				f0elem.Mode = f0iter.Mode
			}
			if f0iter.ModelDataUrl != nil {
				f0elem.ModelDataURL = f0iter.ModelDataUrl
			}
			if f0iter.ModelPackageName != nil {
				f0elem.ModelPackageName = f0iter.ModelPackageName
			}
			f0 = append(f0, f0elem)
		}
		ko.Spec.Containers = f0
	}
	if resp.EnableNetworkIsolation != nil {
		ko.Spec.EnableNetworkIsolation = resp.EnableNetworkIsolation
	}
	if resp.ExecutionRoleArn != nil {
		ko.Spec.ExecutionRoleARN = resp.ExecutionRoleArn
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.ModelArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.ModelArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.ModelName != nil {
		ko.Spec.ModelName = resp.ModelName
	}
	if resp.PrimaryContainer != nil {
		f6 := &svcapitypes.ContainerDefinition{}
		if resp.PrimaryContainer.ContainerHostname != nil {
			f6.ContainerHostname = resp.PrimaryContainer.ContainerHostname
		}
		if resp.PrimaryContainer.Environment != nil {
			f6f1 := map[string]*string{}
			for f6f1key, f6f1valiter := range resp.PrimaryContainer.Environment {
				var f6f1val string
				f6f1val = *f6f1valiter
				f6f1[f6f1key] = &f6f1val
			}
			f6.Environment = f6f1
		}
		if resp.PrimaryContainer.Image != nil {
			f6.Image = resp.PrimaryContainer.Image
		}
		if resp.PrimaryContainer.ImageConfig != nil {
			f6f3 := &svcapitypes.ImageConfig{}
			if resp.PrimaryContainer.ImageConfig.RepositoryAccessMode != nil {
				f6f3.RepositoryAccessMode = resp.PrimaryContainer.ImageConfig.RepositoryAccessMode
			}
			f6.ImageConfig = f6f3
		}
		if resp.PrimaryContainer.Mode != nil {
			f6.Mode = resp.PrimaryContainer.Mode
		}
		if resp.PrimaryContainer.ModelDataUrl != nil {
			f6.ModelDataURL = resp.PrimaryContainer.ModelDataUrl
		}
		if resp.PrimaryContainer.ModelPackageName != nil {
			f6.ModelPackageName = resp.PrimaryContainer.ModelPackageName
		}
		ko.Spec.PrimaryContainer = f6
	}
	if resp.VpcConfig != nil {
		f7 := &svcapitypes.VPCConfig{}
		if resp.VpcConfig.SecurityGroupIds != nil {
			f7f0 := []*string{}
			for _, f7f0iter := range resp.VpcConfig.SecurityGroupIds {
				var f7f0elem string
				f7f0elem = *f7f0iter
				f7f0 = append(f7f0, &f7f0elem)
			}
			f7.SecurityGroupIDs = f7f0
		}
		if resp.VpcConfig.Subnets != nil {
			f7f1 := []*string{}
			for _, f7f1iter := range resp.VpcConfig.Subnets {
				var f7f1elem string
				f7f1elem = *f7f1iter
				f7f1 = append(f7f1, &f7f1elem)
			}
			f7.Subnets = f7f1
		}
		ko.Spec.VPCConfig = f7
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.ModelName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeModelInput, error) {
	res := &svcsdk.DescribeModelInput{}

	if r.ko.Spec.ModelName != nil {
		res.SetModelName(*r.ko.Spec.ModelName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateModelWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateModel", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.ModelArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.ModelArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateModelInput, error) {
	res := &svcsdk.CreateModelInput{}

	if r.ko.Spec.Containers != nil {
		f0 := []*svcsdk.ContainerDefinition{}
		for _, f0iter := range r.ko.Spec.Containers {
			f0elem := &svcsdk.ContainerDefinition{}
			if f0iter.ContainerHostname != nil {
				f0elem.SetContainerHostname(*f0iter.ContainerHostname)
			}
			if f0iter.Environment != nil {
				f0elemf1 := map[string]*string{}
				for f0elemf1key, f0elemf1valiter := range f0iter.Environment {
					var f0elemf1val string
					f0elemf1val = *f0elemf1valiter
					f0elemf1[f0elemf1key] = &f0elemf1val
				}
				f0elem.SetEnvironment(f0elemf1)
			}
			if f0iter.Image != nil {
				f0elem.SetImage(*f0iter.Image)
			}
			if f0iter.ImageConfig != nil {
				f0elemf3 := &svcsdk.ImageConfig{}
				if f0iter.ImageConfig.RepositoryAccessMode != nil {
					f0elemf3.SetRepositoryAccessMode(*f0iter.ImageConfig.RepositoryAccessMode)
				}
				f0elem.SetImageConfig(f0elemf3)
			}
			if f0iter.Mode != nil {
				f0elem.SetMode(*f0iter.Mode)
			}
			if f0iter.ModelDataURL != nil {
				f0elem.SetModelDataUrl(*f0iter.ModelDataURL)
			}
			if f0iter.ModelPackageName != nil {
				f0elem.SetModelPackageName(*f0iter.ModelPackageName)
			}
			f0 = append(f0, f0elem)
		}
		res.SetContainers(f0)
	}
	if r.ko.Spec.EnableNetworkIsolation != nil {
		res.SetEnableNetworkIsolation(*r.ko.Spec.EnableNetworkIsolation)
	}
	if r.ko.Spec.ExecutionRoleARN != nil {
		res.SetExecutionRoleArn(*r.ko.Spec.ExecutionRoleARN)
	}
	if r.ko.Spec.ModelName != nil {
		res.SetModelName(*r.ko.Spec.ModelName)
	}
	if r.ko.Spec.PrimaryContainer != nil {
		f4 := &svcsdk.ContainerDefinition{}
		if r.ko.Spec.PrimaryContainer.ContainerHostname != nil {
			f4.SetContainerHostname(*r.ko.Spec.PrimaryContainer.ContainerHostname)
		}
		if r.ko.Spec.PrimaryContainer.Environment != nil {
			f4f1 := map[string]*string{}
			for f4f1key, f4f1valiter := range r.ko.Spec.PrimaryContainer.Environment {
				var f4f1val string
				f4f1val = *f4f1valiter
				f4f1[f4f1key] = &f4f1val
			}
			f4.SetEnvironment(f4f1)
		}
		if r.ko.Spec.PrimaryContainer.Image != nil {
			f4.SetImage(*r.ko.Spec.PrimaryContainer.Image)
		}
		if r.ko.Spec.PrimaryContainer.ImageConfig != nil {
			f4f3 := &svcsdk.ImageConfig{}
			if r.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode != nil {
				f4f3.SetRepositoryAccessMode(*r.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode)
			}
			f4.SetImageConfig(f4f3)
		}
		if r.ko.Spec.PrimaryContainer.Mode != nil {
			f4.SetMode(*r.ko.Spec.PrimaryContainer.Mode)
		}
		if r.ko.Spec.PrimaryContainer.ModelDataURL != nil {
			f4.SetModelDataUrl(*r.ko.Spec.PrimaryContainer.ModelDataURL)
		}
		if r.ko.Spec.PrimaryContainer.ModelPackageName != nil {
			f4.SetModelPackageName(*r.ko.Spec.PrimaryContainer.ModelPackageName)
		}
		res.SetPrimaryContainer(f4)
	}
	if r.ko.Spec.Tags != nil {
		f5 := []*svcsdk.Tag{}
		for _, f5iter := range r.ko.Spec.Tags {
			f5elem := &svcsdk.Tag{}
			if f5iter.Key != nil {
				f5elem.SetKey(*f5iter.Key)
			}
			if f5iter.Value != nil {
				f5elem.SetValue(*f5iter.Value)
			}
			f5 = append(f5, f5elem)
		}
		res.SetTags(f5)
	}
	if r.ko.Spec.VPCConfig != nil {
		f6 := &svcsdk.VpcConfig{}
		if r.ko.Spec.VPCConfig.SecurityGroupIDs != nil {
			f6f0 := []*string{}
			for _, f6f0iter := range r.ko.Spec.VPCConfig.SecurityGroupIDs {
				var f6f0elem string
				f6f0elem = *f6f0iter
				f6f0 = append(f6f0, &f6f0elem)
			}
			f6.SetSecurityGroupIds(f6f0)
		}
		if r.ko.Spec.VPCConfig.Subnets != nil {
			f6f1 := []*string{}
			for _, f6f1iter := range r.ko.Spec.VPCConfig.Subnets {
				var f6f1elem string
				f6f1elem = *f6f1iter
				f6f1 = append(f6f1, &f6f1elem)
			}
			f6.SetSubnets(f6f1)
		}
		res.SetVpcConfig(f6)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteModelWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteModel", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteModelInput, error) {
	res := &svcsdk.DeleteModelInput{}

	if r.ko.Spec.ModelName != nil {
		res.SetModelName(*r.ko.Spec.ModelName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Model,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
			break
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else if terminalCondition != nil {
		terminalCondition.Status = corev1.ConditionFalse
		terminalCondition.Message = nil
	}
	if terminalCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
