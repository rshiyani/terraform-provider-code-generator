package aci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciContract() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAciContractRead,

		Schema: map[string]*schema.Schema{
			"tenant_dn": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "tenant DN",
			},

			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "contract name",
			},

			"my_map": &schema.Schema{
				Type:        schema.TypeMap,
				Required:    true,
				Description: "My map for testing",
			},

			"prio": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "prio",
			},

			"filter": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "filter list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "name of filter",
						},

						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "id of filter",
						},

						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "description of filter",
						},

						"filter_entry": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "list of filter_entry",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cast": &schema.Schema{
										Type:     schema.TypeSet,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"filter_entry_name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "name of filter entry",
									},

									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "id of filter entry",
									},

									"apply_to_frag": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "apply to fragment",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceAciContractRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	AciClient := m.(*client.Client)

	ContractMap := models.Contract{
		TenantDn: d.Get("tenantDn"),

		Name: d.Get("name"),

		MyMap: d.Get("myMap"),
	}

	ContractId, err := getIdFromContractModel(&ContractMap)
	if err != nil {
		return diag.FromErr(err)
	}

	Contract, err := getContractAttributes(AciClient, ContractId)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = setContractAttributes(Contract, d)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
