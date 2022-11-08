//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Nephio Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	if in.NetworkInstance != nil {
		in, out := &in.NetworkInstance, &out.NetworkInstance
		*out = new(string)
		**out = **in
	}
	if in.NetworkName != nil {
		in, out := &in.NetworkName, &out.NetworkName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FiveGCoreTopology) DeepCopyInto(out *FiveGCoreTopology) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FiveGCoreTopology.
func (in *FiveGCoreTopology) DeepCopy() *FiveGCoreTopology {
	if in == nil {
		return nil
	}
	out := new(FiveGCoreTopology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FiveGCoreTopology) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FiveGCoreTopologyList) DeepCopyInto(out *FiveGCoreTopologyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FiveGCoreTopology, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FiveGCoreTopologyList.
func (in *FiveGCoreTopologyList) DeepCopy() *FiveGCoreTopologyList {
	if in == nil {
		return nil
	}
	out := new(FiveGCoreTopologyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FiveGCoreTopologyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FiveGCoreTopologySpec) DeepCopyInto(out *FiveGCoreTopologySpec) {
	*out = *in
	if in.UPFs != nil {
		in, out := &in.UPFs, &out.UPFs
		*out = make([]UPFClusterSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FiveGCoreTopologySpec.
func (in *FiveGCoreTopologySpec) DeepCopy() *FiveGCoreTopologySpec {
	if in == nil {
		return nil
	}
	out := new(FiveGCoreTopologySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FiveGCoreTopologyStatus) DeepCopyInto(out *FiveGCoreTopologyStatus) {
	*out = *in
	if in.UPFStatuses != nil {
		in, out := &in.UPFStatuses, &out.UPFStatuses
		*out = make([]UPFStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FiveGCoreTopologyStatus.
func (in *FiveGCoreTopologyStatus) DeepCopy() *FiveGCoreTopologyStatus {
	if in == nil {
		return nil
	}
	out := new(FiveGCoreTopologyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceConfig) DeepCopyInto(out *InterfaceConfig) {
	*out = *in
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GatewayIPs != nil {
		in, out := &in.GatewayIPs, &out.GatewayIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceConfig.
func (in *InterfaceConfig) DeepCopy() *InterfaceConfig {
	if in == nil {
		return nil
	}
	out := new(InterfaceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *N6Endpoint) DeepCopyInto(out *N6Endpoint) {
	*out = *in
	in.Endpoint.DeepCopyInto(&out.Endpoint)
	in.UEPool.DeepCopyInto(&out.UEPool)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new N6Endpoint.
func (in *N6Endpoint) DeepCopy() *N6Endpoint {
	if in == nil {
		return nil
	}
	out := new(N6Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *N6InterfaceConfig) DeepCopyInto(out *N6InterfaceConfig) {
	*out = *in
	in.Interface.DeepCopyInto(&out.Interface)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new N6InterfaceConfig.
func (in *N6InterfaceConfig) DeepCopy() *N6InterfaceConfig {
	if in == nil {
		return nil
	}
	out := new(N6InterfaceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pool) DeepCopyInto(out *Pool) {
	*out = *in
	if in.NetworkInstance != nil {
		in, out := &in.NetworkInstance, &out.NetworkInstance
		*out = new(string)
		**out = **in
	}
	if in.NetworkName != nil {
		in, out := &in.NetworkName, &out.NetworkName
		*out = new(string)
		**out = **in
	}
	if in.PrefixSize != nil {
		in, out := &in.PrefixSize, &out.PrefixSize
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pool.
func (in *Pool) DeepCopy() *Pool {
	if in == nil {
		return nil
	}
	out := new(Pool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UPF) DeepCopyInto(out *UPF) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UPF.
func (in *UPF) DeepCopy() *UPF {
	if in == nil {
		return nil
	}
	out := new(UPF)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UPF) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UPFCapacity) DeepCopyInto(out *UPFCapacity) {
	*out = *in
	out.UplinkThroughput = in.UplinkThroughput.DeepCopy()
	out.DownlinkThroughput = in.DownlinkThroughput.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UPFCapacity.
func (in *UPFCapacity) DeepCopy() *UPFCapacity {
	if in == nil {
		return nil
	}
	out := new(UPFCapacity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UPFClass) DeepCopyInto(out *UPFClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UPFClass.
func (in *UPFClass) DeepCopy() *UPFClass {
	if in == nil {
		return nil
	}
	out := new(UPFClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UPFClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UPFClassList) DeepCopyInto(out *UPFClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UPFClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UPFClassList.
func (in *UPFClassList) DeepCopy() *UPFClassList {
	if in == nil {
		return nil
	}
	out := new(UPFClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UPFClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UPFClassSpec) DeepCopyInto(out *UPFClassSpec) {
	*out = *in
	out.PackageRef = in.PackageRef
	if in.DNNs != nil {
		in, out := &in.DNNs, &out.DNNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UPFClassSpec.
func (in *UPFClassSpec) DeepCopy() *UPFClassSpec {
	if in == nil {
		return nil
	}
	out := new(UPFClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UPFClassStatus) DeepCopyInto(out *UPFClassStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UPFClassStatus.
func (in *UPFClassStatus) DeepCopy() *UPFClassStatus {
	if in == nil {
		return nil
	}
	out := new(UPFClassStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UPFClusterSet) DeepCopyInto(out *UPFClusterSet) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.UPF.DeepCopyInto(&out.UPF)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UPFClusterSet.
func (in *UPFClusterSet) DeepCopy() *UPFClusterSet {
	if in == nil {
		return nil
	}
	out := new(UPFClusterSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UPFDeployment) DeepCopyInto(out *UPFDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UPFDeployment.
func (in *UPFDeployment) DeepCopy() *UPFDeployment {
	if in == nil {
		return nil
	}
	out := new(UPFDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UPFDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UPFDeploymentList) DeepCopyInto(out *UPFDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UPFDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UPFDeploymentList.
func (in *UPFDeploymentList) DeepCopy() *UPFDeploymentList {
	if in == nil {
		return nil
	}
	out := new(UPFDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UPFDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UPFDeploymentSpec) DeepCopyInto(out *UPFDeploymentSpec) {
	*out = *in
	in.Capacity.DeepCopyInto(&out.Capacity)
	if in.N3Interfaces != nil {
		in, out := &in.N3Interfaces, &out.N3Interfaces
		*out = make([]InterfaceConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.N4Interfaces != nil {
		in, out := &in.N4Interfaces, &out.N4Interfaces
		*out = make([]InterfaceConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.N6Interfaces != nil {
		in, out := &in.N6Interfaces, &out.N6Interfaces
		*out = make([]N6InterfaceConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.N9Interfaces != nil {
		in, out := &in.N9Interfaces, &out.N9Interfaces
		*out = make([]InterfaceConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UPFDeploymentSpec.
func (in *UPFDeploymentSpec) DeepCopy() *UPFDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(UPFDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UPFDeploymentStatus) DeepCopyInto(out *UPFDeploymentStatus) {
	*out = *in
	in.ComputeUpTime.DeepCopyInto(&out.ComputeUpTime)
	in.OperationUpTime.DeepCopyInto(&out.OperationUpTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UPFDeploymentStatus.
func (in *UPFDeploymentStatus) DeepCopy() *UPFDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(UPFDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UPFList) DeepCopyInto(out *UPFList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UPF, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UPFList.
func (in *UPFList) DeepCopy() *UPFList {
	if in == nil {
		return nil
	}
	out := new(UPFList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UPFList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UPFSpec) DeepCopyInto(out *UPFSpec) {
	*out = *in
	in.Capacity.DeepCopyInto(&out.Capacity)
	if in.N3 != nil {
		in, out := &in.N3, &out.N3
		*out = make([]Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.N4 != nil {
		in, out := &in.N4, &out.N4
		*out = make([]Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.N6 != nil {
		in, out := &in.N6, &out.N6
		*out = make([]N6Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.N9 != nil {
		in, out := &in.N9, &out.N9
		*out = make([]Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UPFSpec.
func (in *UPFSpec) DeepCopy() *UPFSpec {
	if in == nil {
		return nil
	}
	out := new(UPFSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UPFStatus) DeepCopyInto(out *UPFStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UPFStatus.
func (in *UPFStatus) DeepCopy() *UPFStatus {
	if in == nil {
		return nil
	}
	out := new(UPFStatus)
	in.DeepCopyInto(out)
	return out
}
