package bucket

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("oci_objectstorage_bucket", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
        // we need to override the default group that terrajet generated for
        // this resource, which would be "github"  
        r.ShortGroup = "bucket"
    })
}