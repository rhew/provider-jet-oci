//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Object) DeepCopyInto(out *Object) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Object.
func (in *Object) DeepCopy() *Object {
	if in == nil {
		return nil
	}
	out := new(Object)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Object) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectList) DeepCopyInto(out *ObjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Object, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectList.
func (in *ObjectList) DeepCopy() *ObjectList {
	if in == nil {
		return nil
	}
	out := new(ObjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectObservation) DeepCopyInto(out *ObjectObservation) {
	*out = *in
	if in.ContentLength != nil {
		in, out := &in.ContentLength, &out.ContentLength
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.VersionID != nil {
		in, out := &in.VersionID, &out.VersionID
		*out = new(string)
		**out = **in
	}
	if in.WorkRequestID != nil {
		in, out := &in.WorkRequestID, &out.WorkRequestID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectObservation.
func (in *ObjectObservation) DeepCopy() *ObjectObservation {
	if in == nil {
		return nil
	}
	out := new(ObjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectParameters) DeepCopyInto(out *ObjectParameters) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.CacheControl != nil {
		in, out := &in.CacheControl, &out.CacheControl
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.ContentDisposition != nil {
		in, out := &in.ContentDisposition, &out.ContentDisposition
		*out = new(string)
		**out = **in
	}
	if in.ContentEncoding != nil {
		in, out := &in.ContentEncoding, &out.ContentEncoding
		*out = new(string)
		**out = **in
	}
	if in.ContentLanguage != nil {
		in, out := &in.ContentLanguage, &out.ContentLanguage
		*out = new(string)
		**out = **in
	}
	if in.ContentMd5 != nil {
		in, out := &in.ContentMd5, &out.ContentMd5
		*out = new(string)
		**out = **in
	}
	if in.ContentType != nil {
		in, out := &in.ContentType, &out.ContentType
		*out = new(string)
		**out = **in
	}
	if in.DeleteAllObjectVersions != nil {
		in, out := &in.DeleteAllObjectVersions, &out.DeleteAllObjectVersions
		*out = new(bool)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(string)
		**out = **in
	}
	if in.OpcSseKMSKeyID != nil {
		in, out := &in.OpcSseKMSKeyID, &out.OpcSseKMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.SourceURIDetails != nil {
		in, out := &in.SourceURIDetails, &out.SourceURIDetails
		*out = make([]SourceURIDetailsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StorageTier != nil {
		in, out := &in.StorageTier, &out.StorageTier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectParameters.
func (in *ObjectParameters) DeepCopy() *ObjectParameters {
	if in == nil {
		return nil
	}
	out := new(ObjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSpec) DeepCopyInto(out *ObjectSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSpec.
func (in *ObjectSpec) DeepCopy() *ObjectSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStatus) DeepCopyInto(out *ObjectStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStatus.
func (in *ObjectStatus) DeepCopy() *ObjectStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceURIDetailsObservation) DeepCopyInto(out *SourceURIDetailsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceURIDetailsObservation.
func (in *SourceURIDetailsObservation) DeepCopy() *SourceURIDetailsObservation {
	if in == nil {
		return nil
	}
	out := new(SourceURIDetailsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceURIDetailsParameters) DeepCopyInto(out *SourceURIDetailsParameters) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.DestinationObjectIfMatchEtag != nil {
		in, out := &in.DestinationObjectIfMatchEtag, &out.DestinationObjectIfMatchEtag
		*out = new(string)
		**out = **in
	}
	if in.DestinationObjectIfNoneMatchEtag != nil {
		in, out := &in.DestinationObjectIfNoneMatchEtag, &out.DestinationObjectIfNoneMatchEtag
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SourceObjectIfMatchEtag != nil {
		in, out := &in.SourceObjectIfMatchEtag, &out.SourceObjectIfMatchEtag
		*out = new(string)
		**out = **in
	}
	if in.SourceVersionID != nil {
		in, out := &in.SourceVersionID, &out.SourceVersionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceURIDetailsParameters.
func (in *SourceURIDetailsParameters) DeepCopy() *SourceURIDetailsParameters {
	if in == nil {
		return nil
	}
	out := new(SourceURIDetailsParameters)
	in.DeepCopyInto(out)
	return out
}
