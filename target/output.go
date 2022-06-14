package aci

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAciContract() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAciContractCreate,
		UpdateContext: resourceAciContractUpdate,
		ReadContext:   resourceAciContractRead,
		DeleteContext: resourceAciContractDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "name",
			},

			"ipv4": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				Description:      "IP-address v4",
				ValidateDiagFunc: validation.ToDiagFunc(validation.IsIPv4Address),
			},

			"ipv6": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				Description:      "IP-address v6",
				ValidateDiagFunc: validation.ToDiagFunc(validation.IsIPv6Address),
			},

			"mac": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IsMACAddress),
			},

			"cidr": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IsCIDR),
			},

			"time": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IsRFC3339Time),
			},

			"url_https": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IsURLWithHTTPS),
			},

			"url_http": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IsURLWithHTTPorHTTPS),
			},

			"uuid": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
			},

			"base_64": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsBase64),
			},

			"json": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsJSON),
			},

			"reg_exp": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsValidRegExp),
			},

			"gender": &schema.Schema{
				Type: schema.TypeString,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"male",
					"female",
					"other",
				}, true),
				),
			},

			"port_number": &schema.Schema{
				Type:             schema.TypeInt,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IsPortNumber),
			},

			"port_with_zero": &schema.Schema{
				Type:             schema.TypeInt,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IsPortNumberOrZero),
			},

			"nuclear_code": &schema.Schema{
				Type:        schema.TypeString,
				Sensitive:   true,
				Required:    true,
				Description: "Nuclear code",
				// [ERROR]: NotInYourList may be a Typo or not in AutoGen List. Please refer docs once.

			},

			"test_score": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Computed:    true,
				Description: "range",
				// [ERROR]: IntBetween may be a Typo or not in AutoGen List. Please refer docs once.

			},

			"percentage": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
				Computed: true,
				// [ERROR]: FloatBetween may be a Typo or not in AutoGen List. Please refer docs once.

			},

			"filter": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "filter list",
				MaxItems:    10,
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
							DefaultFunc: schema.EnvDefaultFunc("FILTER_ID", nil),
							Description: "id of filter",
						},

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							DefaultFunc: func() (interface{}, error) {
								// [TODO]: Write your code here
								return nil, nil
							},
							Description: "description of filter",
						},

						"filter_entry": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Optional:    true,
							Description: "list of filter_entry",
							MaxItems:    4,
							MinItems:    1,
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
										DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
											// [TODO]: Write your code here
											return false
										},
									},

									"ipv6": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										Description:      "ipv6",
										ValidateDiagFunc: validation.ToDiagFunc(validation.IsIPv6Address),
									},

									"apply_to_frag": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Optional:    true,
										Description: "apply to fragment",
										ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
											"yes",
											"no",
										}, false),
										),
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
