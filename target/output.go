package aci

import (
	"context"

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
			"tenant_dn": &schema.Schema{
				Type:        schema.TypeString,
				Sensitive:   true,
				Required:    true,
				ForceNew:    true,
				Description: "tenant DN",
			},

			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "contract name",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"1",
					"2",
					"3",
				}, false),
				),
			},

			"my_map": &schema.Schema{
				Type:        schema.TypeMap,
				Sensitive:   true,
				Required:    true,
				ForceNew:    true,
				Description: "My map for testing",
				// [ERROR]: StringinSLice may be a Typo or not in AutoGen List. Please refer docs once.
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"prio": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "prio",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.IsCIDRNetwork(1, 2),
				),
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
							DefaultFunc: schema.EnvDefaultFunc("ID", nil),
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
										DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
											// [TODO]: Write your code here
											return false
										},
									},

									"id": &schema.Schema{
										Type:             schema.TypeString,
										Computed:         true,
										Description:      "id of filter entry",
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

func resourceAciContractCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	aciClient := m.(*client.Client)
	contractAttr = models.ContractAttributes{}

	contractAttr.TenantDn = TenantDn.(string)
	contractAttr.Name = Name.(string)
	if Prio, ok := d.GetOk("prio"); ok {
		contractAttr.Prio = Prio.(string)
	}

	contract := models.NewContract(contractAttr)

	err := aciClient.Save(contract)
	if err != nil {
		return diag.FromErr(err)
	}
	if Filters, ok := d.GetOk("filter"); ok {
		filters := Filters.([]interface{})
		for _, val := range filters {
			filterAttr := models.FilterAttributes{}
			filter := val.(map[string]interface{})
			filterAttr.FilterName = filter["filter_name"].(string)
			if Id, ok := d.Get("id"); ok {
				filterAttr.Id = filter["id"].(string)
			}
			if Description, ok := d.Get("description"); ok {
				filterAttr.Description = filter["description"].(string)
			}
			if filter["filter_entry"] != nil {
				filter_entrys := filter["filter_entry"].([]interface{})

				for _, val := range filter_entrys {
					filter_entryAttr := models.FilterEntryAttributes{}
					filter_entry := val.(map[string]interface{})
					filter_entryAttr.FilterEntryName = filter_entry["filter_entry_name"].(string)
					if ApplyToFrag, ok := d.GetOk("apply_to_frag"); ok {
						filter_entryAttr.ApplyToFrag = filter_entry["apply_to_frag"].(string)
					}
				}
				filter_entry := models.NewFilterEntry(filter_entryAttr)
				err := aciClient.Save(filter_entrymodel)
				if err != nil {
					return diag.FromErr(err)
				}
			}

		}
		filtermodel := models.NewFilter(filterAttr)

		err := Client.Save(filtermodel)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceAciContractRead(ctx, d, m)
}
