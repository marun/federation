// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package certificates

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateSigningRequest).DeepCopyInto(out.(*CertificateSigningRequest))
			return nil
		}, InType: reflect.TypeOf(&CertificateSigningRequest{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateSigningRequestCondition).DeepCopyInto(out.(*CertificateSigningRequestCondition))
			return nil
		}, InType: reflect.TypeOf(&CertificateSigningRequestCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateSigningRequestList).DeepCopyInto(out.(*CertificateSigningRequestList))
			return nil
		}, InType: reflect.TypeOf(&CertificateSigningRequestList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateSigningRequestSpec).DeepCopyInto(out.(*CertificateSigningRequestSpec))
			return nil
		}, InType: reflect.TypeOf(&CertificateSigningRequestSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateSigningRequestStatus).DeepCopyInto(out.(*CertificateSigningRequestStatus))
			return nil
		}, InType: reflect.TypeOf(&CertificateSigningRequestStatus{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSigningRequest) DeepCopyInto(out *CertificateSigningRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSigningRequest.
func (in *CertificateSigningRequest) DeepCopy() *CertificateSigningRequest {
	if in == nil {
		return nil
	}
	out := new(CertificateSigningRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateSigningRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSigningRequestCondition) DeepCopyInto(out *CertificateSigningRequestCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSigningRequestCondition.
func (in *CertificateSigningRequestCondition) DeepCopy() *CertificateSigningRequestCondition {
	if in == nil {
		return nil
	}
	out := new(CertificateSigningRequestCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSigningRequestList) DeepCopyInto(out *CertificateSigningRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificateSigningRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSigningRequestList.
func (in *CertificateSigningRequestList) DeepCopy() *CertificateSigningRequestList {
	if in == nil {
		return nil
	}
	out := new(CertificateSigningRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateSigningRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSigningRequestSpec) DeepCopyInto(out *CertificateSigningRequestSpec) {
	*out = *in
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Usages != nil {
		in, out := &in.Usages, &out.Usages
		*out = make([]KeyUsage, len(*in))
		copy(*out, *in)
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Extra != nil {
		in, out := &in.Extra, &out.Extra
		*out = make(map[string]ExtraValue, len(*in))
		for key, val := range *in {
			(*out)[key] = make(ExtraValue, len(val))
			copy((*out)[key], val)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSigningRequestSpec.
func (in *CertificateSigningRequestSpec) DeepCopy() *CertificateSigningRequestSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateSigningRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSigningRequestStatus) DeepCopyInto(out *CertificateSigningRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CertificateSigningRequestCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSigningRequestStatus.
func (in *CertificateSigningRequestStatus) DeepCopy() *CertificateSigningRequestStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateSigningRequestStatus)
	in.DeepCopyInto(out)
	return out
}
