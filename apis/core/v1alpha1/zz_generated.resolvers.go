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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.CreateVnicDetails); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CreateVnicDetails[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CreateVnicDetails[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.CreateVnicDetails[i3].SubnetIDSelector,
			To: reference.To{
				List:    &SubnetList{},
				Managed: &Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CreateVnicDetails[i3].SubnetID")
		}
		mg.Spec.ForProvider.CreateVnicDetails[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CreateVnicDetails[i3].SubnetIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.CreateVnicDetails); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CreateVnicDetails[i3].VlanID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CreateVnicDetails[i3].VlanIDRef,
			Selector:     mg.Spec.ForProvider.CreateVnicDetails[i3].VlanIDSelector,
			To: reference.To{
				List:    &VlanList{},
				Managed: &Vlan{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CreateVnicDetails[i3].VlanID")
		}
		mg.Spec.ForProvider.CreateVnicDetails[i3].VlanID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CreateVnicDetails[i3].VlanIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DedicatedVMHostID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DedicatedVMHostIDRef,
		Selector:     mg.Spec.ForProvider.DedicatedVMHostIDSelector,
		To: reference.To{
			List:    &DedicatedVMHostList{},
			Managed: &DedicatedVMHost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DedicatedVMHostID")
	}
	mg.Spec.ForProvider.DedicatedVMHostID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DedicatedVMHostIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SourceDetails); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceDetails[i3].SourceID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.SourceDetails[i3].SourceIDRef,
			Selector:     mg.Spec.ForProvider.SourceDetails[i3].SourceIDSelector,
			To: reference.To{
				List:    &ImageList{},
				Managed: &Image{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SourceDetails[i3].SourceID")
		}
		mg.Spec.ForProvider.SourceDetails[i3].SourceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SourceDetails[i3].SourceIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Subnet.
func (mg *Subnet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VcnID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VcnIDRef,
		Selector:     mg.Spec.ForProvider.VcnIDSelector,
		To: reference.To{
			List:    &VcnList{},
			Managed: &Vcn{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VcnID")
	}
	mg.Spec.ForProvider.VcnID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VcnIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Vcn.
func (mg *Vcn) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	return nil
}
