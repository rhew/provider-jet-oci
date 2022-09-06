package core

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("oci_core_instance", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
        // we need to override the default group that terrajet generated for
        r.Version = "v1alpha1"
        // r.ShortGroup = "compute"

        // This resource reside inside a compartment. By defining it as a reference to Compartment
        // object, we can build cross resource referencing.
        r.References["compartment_id"] = config.Reference{
            Type: "github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1.Compartment",
        }

        r.References["create_vnic_details.subnet_id"] = config.Reference{
            Type: "Subnet",
        }

        r.References["create_vnic_details.vlan_id"] = config.Reference{
            Type: "Vlan",
        }


        r.References["dedicated_vm_host_id"] = config.Reference{
            Type: "DedicatedVMHost",
        }

        r.References["source_details.source_id"] = config.Reference{
            Type: "Image",
        }

        // r.LateInitializer = config.LateInitializer{
		// 	// NOTE(): These are ignored because they conflict with each other.
		// 	// See the following for more details: https://github.com/crossplane/terrajet/issues/107
		// 	IgnoredFields: []string{
		// 		"compartment_id",
        //         "dedicated_vm_host_id"
	
		// 	},

    })

    p.AddResourceConfigurator("oci_core_image", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
        // we need to override the default group that terrajet generated for
        r.Version = "v1alpha1"
        // r.ShortGroup = "compute"

    })

    p.AddResourceConfigurator("oci_core_vcn", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
        // we need to override the default group that terrajet generated for
        r.Version = "v1alpha1"
        // r.ShortGroup = "compute"

        r.References["compartment_id"] = config.Reference{
            Type: "github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1.Compartment",
        }

    })

    p.AddResourceConfigurator("oci_core_subnet", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
        // we need to override the default group that terrajet generated for
        r.Version = "v1alpha1"

        r.References["vcn_id"] = config.Reference{
            Type: "Vcn",
        }

        r.References["compartment_id"] = config.Reference{
            Type: "github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1.Compartment",
        }

    })

    p.AddResourceConfigurator("oci_core_dedicated_vm_host", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
        // we need to override the default group that terrajet generated for
        r.Version = "v1alpha1"

    })
    
}