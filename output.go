package aci

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_USERNAME", nil),
				Description: "Username of the ACI user",
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_PASSWORD", nil),
				Description: "Password of the ACI user",
			},

			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_URL", nil),
				Description: "URL of the ACI server",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"aci_tenant":              resourceAciTenant(),
			"aci_application_profile": resourceAciApplicationProfile(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"aci_tenant":              dataSourceAciTenant(),
			"aci_application_profile": dataSourceAciApplicationProfile(),
		},

		ConfigureFunc: configureClient,
	}
}

func configureClient(d *schema.ResourceData) (interface{}, error) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	url := d.Get("url").(string)
	return nil, nil
}
