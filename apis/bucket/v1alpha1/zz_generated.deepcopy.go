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
func (in *Bucket) DeepCopyInto(out *Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket.
func (in *Bucket) DeepCopy() *Bucket {
	if in == nil {
		return nil
	}
	out := new(Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketList) DeepCopyInto(out *BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketList.
func (in *BucketList) DeepCopy() *BucketList {
	if in == nil {
		return nil
	}
	out := new(BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketObservation) DeepCopyInto(out *BucketObservation) {
	*out = *in
	if in.ApproximateCount != nil {
		in, out := &in.ApproximateCount, &out.ApproximateCount
		*out = new(string)
		**out = **in
	}
	if in.ApproximateSize != nil {
		in, out := &in.ApproximateSize, &out.ApproximateSize
		*out = new(string)
		**out = **in
	}
	if in.BucketID != nil {
		in, out := &in.BucketID, &out.BucketID
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsReadOnly != nil {
		in, out := &in.IsReadOnly, &out.IsReadOnly
		*out = new(bool)
		**out = **in
	}
	if in.ObjectLifecyclePolicyEtag != nil {
		in, out := &in.ObjectLifecyclePolicyEtag, &out.ObjectLifecyclePolicyEtag
		*out = new(string)
		**out = **in
	}
	if in.ReplicationEnabled != nil {
		in, out := &in.ReplicationEnabled, &out.ReplicationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketObservation.
func (in *BucketObservation) DeepCopy() *BucketObservation {
	if in == nil {
		return nil
	}
	out := new(BucketObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketParameters) DeepCopyInto(out *BucketParameters) {
	*out = *in
	if in.AccessType != nil {
		in, out := &in.AccessType, &out.AccessType
		*out = new(string)
		**out = **in
	}
	if in.AutoTiering != nil {
		in, out := &in.AutoTiering, &out.AutoTiering
		*out = new(string)
		**out = **in
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.DefinedTags != nil {
		in, out := &in.DefinedTags, &out.DefinedTags
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
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
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
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
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
	if in.ObjectEventsEnabled != nil {
		in, out := &in.ObjectEventsEnabled, &out.ObjectEventsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RetentionRules != nil {
		in, out := &in.RetentionRules, &out.RetentionRules
		*out = make([]RetentionRulesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StorageTier != nil {
		in, out := &in.StorageTier, &out.StorageTier
		*out = new(string)
		**out = **in
	}
	if in.Versioning != nil {
		in, out := &in.Versioning, &out.Versioning
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketParameters.
func (in *BucketParameters) DeepCopy() *BucketParameters {
	if in == nil {
		return nil
	}
	out := new(BucketParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketSpec) DeepCopyInto(out *BucketSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketSpec.
func (in *BucketSpec) DeepCopy() *BucketSpec {
	if in == nil {
		return nil
	}
	out := new(BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketStatus) DeepCopyInto(out *BucketStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketStatus.
func (in *BucketStatus) DeepCopy() *BucketStatus {
	if in == nil {
		return nil
	}
	out := new(BucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DurationObservation) DeepCopyInto(out *DurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DurationObservation.
func (in *DurationObservation) DeepCopy() *DurationObservation {
	if in == nil {
		return nil
	}
	out := new(DurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DurationParameters) DeepCopyInto(out *DurationParameters) {
	*out = *in
	if in.TimeAmount != nil {
		in, out := &in.TimeAmount, &out.TimeAmount
		*out = new(string)
		**out = **in
	}
	if in.TimeUnit != nil {
		in, out := &in.TimeUnit, &out.TimeUnit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DurationParameters.
func (in *DurationParameters) DeepCopy() *DurationParameters {
	if in == nil {
		return nil
	}
	out := new(DurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionRulesObservation) DeepCopyInto(out *RetentionRulesObservation) {
	*out = *in
	if in.RetentionRuleID != nil {
		in, out := &in.RetentionRuleID, &out.RetentionRuleID
		*out = new(string)
		**out = **in
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
	if in.TimeModified != nil {
		in, out := &in.TimeModified, &out.TimeModified
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionRulesObservation.
func (in *RetentionRulesObservation) DeepCopy() *RetentionRulesObservation {
	if in == nil {
		return nil
	}
	out := new(RetentionRulesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionRulesParameters) DeepCopyInto(out *RetentionRulesParameters) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = make([]DurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimeRuleLocked != nil {
		in, out := &in.TimeRuleLocked, &out.TimeRuleLocked
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionRulesParameters.
func (in *RetentionRulesParameters) DeepCopy() *RetentionRulesParameters {
	if in == nil {
		return nil
	}
	out := new(RetentionRulesParameters)
	in.DeepCopyInto(out)
	return out
}
