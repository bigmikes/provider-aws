/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package mounttarget

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/efs"

	svcapitypes "github.com/crossplane/provider-aws/apis/efs/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeMountTargetsInput returns input for read
// operation.
func GenerateDescribeMountTargetsInput(cr *svcapitypes.MountTarget) *svcsdk.DescribeMountTargetsInput {
	res := &svcsdk.DescribeMountTargetsInput{}

	if cr.Status.AtProvider.FileSystemID != nil {
		res.SetFileSystemId(*cr.Status.AtProvider.FileSystemID)
	}
	if cr.Status.AtProvider.MountTargetID != nil {
		res.SetMountTargetId(*cr.Status.AtProvider.MountTargetID)
	}

	return res
}

// GenerateMountTarget returns the current state in the form of *svcapitypes.MountTarget.
func GenerateMountTarget(resp *svcsdk.DescribeMountTargetsOutput) *svcapitypes.MountTarget {
	cr := &svcapitypes.MountTarget{}

	found := false
	for _, elem := range resp.MountTargets {
		if elem.AvailabilityZoneId != nil {
			cr.Status.AtProvider.AvailabilityZoneID = elem.AvailabilityZoneId
		} else {
			cr.Status.AtProvider.AvailabilityZoneID = nil
		}
		if elem.AvailabilityZoneName != nil {
			cr.Status.AtProvider.AvailabilityZoneName = elem.AvailabilityZoneName
		} else {
			cr.Status.AtProvider.AvailabilityZoneName = nil
		}
		if elem.FileSystemId != nil {
			cr.Status.AtProvider.FileSystemID = elem.FileSystemId
		} else {
			cr.Status.AtProvider.FileSystemID = nil
		}
		if elem.IpAddress != nil {
			cr.Spec.ForProvider.IPAddress = elem.IpAddress
		} else {
			cr.Spec.ForProvider.IPAddress = nil
		}
		if elem.LifeCycleState != nil {
			cr.Status.AtProvider.LifeCycleState = elem.LifeCycleState
		} else {
			cr.Status.AtProvider.LifeCycleState = nil
		}
		if elem.MountTargetId != nil {
			cr.Status.AtProvider.MountTargetID = elem.MountTargetId
		} else {
			cr.Status.AtProvider.MountTargetID = nil
		}
		if elem.NetworkInterfaceId != nil {
			cr.Status.AtProvider.NetworkInterfaceID = elem.NetworkInterfaceId
		} else {
			cr.Status.AtProvider.NetworkInterfaceID = nil
		}
		if elem.OwnerId != nil {
			cr.Status.AtProvider.OwnerID = elem.OwnerId
		} else {
			cr.Status.AtProvider.OwnerID = nil
		}
		if elem.SubnetId != nil {
			cr.Status.AtProvider.SubnetID = elem.SubnetId
		} else {
			cr.Status.AtProvider.SubnetID = nil
		}
		if elem.VpcId != nil {
			cr.Status.AtProvider.VPCID = elem.VpcId
		} else {
			cr.Status.AtProvider.VPCID = nil
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateMountTargetInput returns a create input.
func GenerateCreateMountTargetInput(cr *svcapitypes.MountTarget) *svcsdk.CreateMountTargetInput {
	res := &svcsdk.CreateMountTargetInput{}

	if cr.Spec.ForProvider.IPAddress != nil {
		res.SetIpAddress(*cr.Spec.ForProvider.IPAddress)
	}

	return res
}

// GenerateDeleteMountTargetInput returns a deletion input.
func GenerateDeleteMountTargetInput(cr *svcapitypes.MountTarget) *svcsdk.DeleteMountTargetInput {
	res := &svcsdk.DeleteMountTargetInput{}

	if cr.Status.AtProvider.MountTargetID != nil {
		res.SetMountTargetId(*cr.Status.AtProvider.MountTargetID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "MountTargetNotFound"
}