package aci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciContract() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAciContractRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "name",
			},

			"ipv4": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP-address v4",
			},

			"ipv6": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP-address v6",
			},

			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"cidr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"time": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"url_https": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"url_http": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"base_64": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"json": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"reg_exp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"gender": &schema.Schema{
				Type: schema.TypeString,
			},

			"port_number": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"port_with_zero": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"nuclear_code": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Nuclear code",
			},

			"test_score": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Computed:    true,
				Description: "range",
			},

			"percentage": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
				Computed: true,
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
									"id_list": &schema.Schema{
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

									"ipv6": &schema.Schema{
										Type:        schema.TypeString,
										Description: "ipv6",
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
		Name: d.Get("name"),

		Ipv4: d.Get("ipv4"),

		Ipv6: d.Get("ipv6"),

		Mac: d.Get("mac"),

		Cidr: d.Get("cidr"),

		Time: d.Get("time"),

		UrlHttps: d.Get("urlHttps"),

		UrlHttp: d.Get("urlHttp"),

		Uuid: d.Get("uuid"),

		Base64: d.Get("base64"),

		Json: d.Get("json"),

		RegExp: d.Get("regExp"),

		PortNumber: d.Get("portNumber"),

		PortWithZero: d.Get("portWithZero"),

		NuclearCode: d.Get("nuclearCode"),

		TestScore: d.Get("testScore"),

		Percentage: d.Get("percentage"),
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
