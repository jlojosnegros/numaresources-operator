//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	machineconfiguration_openshift_iov1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigPool) DeepCopyInto(out *MachineConfigPool) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]machineconfiguration_openshift_iov1.MachineConfigPoolCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigPool.
func (in *MachineConfigPool) DeepCopy() *MachineConfigPool {
	if in == nil {
		return nil
	}
	out := new(MachineConfigPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NUMAResourcesOperator) DeepCopyInto(out *NUMAResourcesOperator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NUMAResourcesOperator.
func (in *NUMAResourcesOperator) DeepCopy() *NUMAResourcesOperator {
	if in == nil {
		return nil
	}
	out := new(NUMAResourcesOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NUMAResourcesOperator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NUMAResourcesOperatorList) DeepCopyInto(out *NUMAResourcesOperatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NUMAResourcesOperator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NUMAResourcesOperatorList.
func (in *NUMAResourcesOperatorList) DeepCopy() *NUMAResourcesOperatorList {
	if in == nil {
		return nil
	}
	out := new(NUMAResourcesOperatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NUMAResourcesOperatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NUMAResourcesOperatorSpec) DeepCopyInto(out *NUMAResourcesOperatorSpec) {
	*out = *in
	if in.NodeGroups != nil {
		in, out := &in.NodeGroups, &out.NodeGroups
		*out = make([]NodeGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NUMAResourcesOperatorSpec.
func (in *NUMAResourcesOperatorSpec) DeepCopy() *NUMAResourcesOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(NUMAResourcesOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NUMAResourcesOperatorStatus) DeepCopyInto(out *NUMAResourcesOperatorStatus) {
	*out = *in
	if in.DaemonSets != nil {
		in, out := &in.DaemonSets, &out.DaemonSets
		*out = make([]NamespacedName, len(*in))
		copy(*out, *in)
	}
	if in.MachineConfigPools != nil {
		in, out := &in.MachineConfigPools, &out.MachineConfigPools
		*out = make([]MachineConfigPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NUMAResourcesOperatorStatus.
func (in *NUMAResourcesOperatorStatus) DeepCopy() *NUMAResourcesOperatorStatus {
	if in == nil {
		return nil
	}
	out := new(NUMAResourcesOperatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedName) DeepCopyInto(out *NamespacedName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedName.
func (in *NamespacedName) DeepCopy() *NamespacedName {
	if in == nil {
		return nil
	}
	out := new(NamespacedName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroup) DeepCopyInto(out *NodeGroup) {
	*out = *in
	if in.MachineConfigPoolSelector != nil {
		in, out := &in.MachineConfigPoolSelector, &out.MachineConfigPoolSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroup.
func (in *NodeGroup) DeepCopy() *NodeGroup {
	if in == nil {
		return nil
	}
	out := new(NodeGroup)
	in.DeepCopyInto(out)
	return out
}
