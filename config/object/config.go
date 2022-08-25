package object

import "github.com/crossplane/terrajet/pkg/config"

func Configure(p *config.Provider) {
    p.AddResourceConfigurator("oci_objectstorage_object", func(r *config.Resource) {

        // we need to override the default group that terrajet generated for
        // this resource, which would be "github" 
        r.ShortGroup = "object"

        // Identifier for this resource is assigned by the provider. In other
        // words it is not simply the name of the resource.
        r.ExternalName = config.NameAsIdentifier

    })
}