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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CompartmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InactiveState *string `json:"inactiveState,omitempty" tf:"inactive_state,omitempty"`

	IsAccessible *bool `json:"isAccessible,omitempty" tf:"is_accessible,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`
}

type CompartmentParameters struct {

	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	EnableDelete *bool `json:"enableDelete,omitempty" tf:"enable_delete,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// CompartmentSpec defines the desired state of Compartment
type CompartmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CompartmentParameters `json:"forProvider"`
}

// CompartmentStatus defines the observed state of Compartment.
type CompartmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CompartmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Compartment is the Schema for the Compartments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type Compartment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CompartmentSpec   `json:"spec"`
	Status            CompartmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CompartmentList contains a list of Compartments
type CompartmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Compartment `json:"items"`
}

// Repository type metadata.
var (
	Compartment_Kind             = "Compartment"
	Compartment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Compartment_Kind}.String()
	Compartment_KindAPIVersion   = Compartment_Kind + "." + CRDGroupVersion.String()
	Compartment_GroupVersionKind = CRDGroupVersion.WithKind(Compartment_Kind)
)

func init() {
	SchemeBuilder.Register(&Compartment{}, &CompartmentList{})
}
