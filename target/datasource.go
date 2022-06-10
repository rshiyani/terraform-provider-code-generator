package aci

import (
	"context"

	"github.com/RutvikS-crest/movies-go-client/client"
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

			"cast": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"filter": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
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
							Required:    true,
							Computed:    true,
							Description: "list of filter_entry",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entry_next": &schema.Schema{
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"entry_next_name": &schema.Schema{
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"cast": &schema.Schema{
										Type:     schema.TypeSet,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cast2": &schema.Schema{
													Type:     schema.TypeSet,
													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
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

	aciClient := m.(*client.Client)

	contractMap := models.Contract{
		TenantDn: d.Get("tenant_dn"),

		Name: d.Get("name"),
	}

	contractId, err := getIdFromContractModel(&contractMap)
	if err != nil {
		return diag.FromErr(err)
	}

	contract, err := getContractAttributes(aciClient, contractId)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = setContractAttributes(contract, d)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
