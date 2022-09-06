package identity

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("oci_identity_compartment", func(r *config.Resource) {
       // r.ShortGroup = "compartment"
        r.Version = "v1alpha1"
        // Identifier for this resource is assigned by the provider. In other
        // words it is not simply the name of the resource.
        r.ExternalName = config.IdentifierFromProvider

    })
    
}