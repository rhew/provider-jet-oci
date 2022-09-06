package bucket

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("oci_objectstorage_bucket", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
        // we need to override the default group that terrajet generated for
    
        r.ShortGroup = "bucket"

        // This resource reside inside a compartment. By defining it as a reference to Compartment
        // object, we can build cross resource referencing.
        r.References["compartment_id"] = config.Reference{
            Type: "github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1.Compartment",
        }
        // r.LateInitializer = config.LateInitializer{
		// 	// NOTE(muvaf): These are ignored because they conflict with each other.
		// 	// See the following for more details: https://github.com/crossplane/terrajet/issues/107
		// 	IgnoredFields: []string{
		// 		"compartment_id",
	
		// 	},
		// }
    })
}