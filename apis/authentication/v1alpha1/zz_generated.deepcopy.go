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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedServiceAccount) DeepCopyInto(out *ManagedServiceAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedServiceAccount.
func (in *ManagedServiceAccount) DeepCopy() *ManagedServiceAccount {
	if in == nil {
		return nil
	}
	out := new(ManagedServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedServiceAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedServiceAccountList) DeepCopyInto(out *ManagedServiceAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedServiceAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedServiceAccountList.
func (in *ManagedServiceAccountList) DeepCopy() *ManagedServiceAccountList {
	if in == nil {
		return nil
	}
	out := new(ManagedServiceAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedServiceAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedServiceAccountRotation) DeepCopyInto(out *ManagedServiceAccountRotation) {
	*out = *in
	out.Validity = in.Validity
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedServiceAccountRotation.
func (in *ManagedServiceAccountRotation) DeepCopy() *ManagedServiceAccountRotation {
	if in == nil {
		return nil
	}
	out := new(ManagedServiceAccountRotation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedServiceAccountSpec) DeepCopyInto(out *ManagedServiceAccountSpec) {
	*out = *in
	out.Rotation = in.Rotation
	if in.TTLSecondsAfterCreation != nil {
		in, out := &in.TTLSecondsAfterCreation, &out.TTLSecondsAfterCreation
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedServiceAccountSpec.
func (in *ManagedServiceAccountSpec) DeepCopy() *ManagedServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedServiceAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedServiceAccountStatus) DeepCopyInto(out *ManagedServiceAccountStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExpirationTimestamp != nil {
		in, out := &in.ExpirationTimestamp, &out.ExpirationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.TokenSecretRef != nil {
		in, out := &in.TokenSecretRef, &out.TokenSecretRef
		*out = new(SecretRef)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedServiceAccountStatus.
func (in *ManagedServiceAccountStatus) DeepCopy() *ManagedServiceAccountStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedServiceAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
	in.LastRefreshTimestamp.DeepCopyInto(&out.LastRefreshTimestamp)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}