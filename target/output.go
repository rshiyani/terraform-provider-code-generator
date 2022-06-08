package aci

import (
	"context"

	"github.com/RutvikS-crest/movies-go-client/client"
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

									"entry_next": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Optional:    true,
										Description: "list of filter_entry",
										MaxItems:    5,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"test": &schema.Schema{
													Type:     schema.TypeSet,
													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
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
	contract = models.Contract{
		TenantDn: d.Get("tenant_dn").(string),
		Name:     d.Get("name").(string),
	}

	if Prio, ok := d.GetOk("prio"); ok {
		contract.Prio = Prio.(string)
	}

	if Filters, ok := d.GetOk("filter"); ok {
		filters := Filters.([]interface{})

		for _, val := range filters {
			filter := models.Filter{}
			filterMap := val.(map[string]interface{})
			filter.FilterName = filterMap["filter_name"].(string)
			if filterMap["description"] != nil {
				filter.Description = filterMap["description"].(string)
			}
			if filterMap["filter_entry"] != nil {
				filter_entrys := filterMap["filter_entry"].([]interface{})

				for _, val := range filter_entrys {
					filter_entry := models.FilterEntry{}
					filterEntryMap := val.(map[string]interface{})
					filter_entry.Cast = filterEntryMap["cast"].(*schema.Set).List()
					filter_entry.FilterEntryName = filterEntryMap["filter_entry_name"].(string)
					if filterEntryMap["entry_next"] != nil {
						entry_nexts := filterEntryMap["entry_next"].([]interface{})

						for _, val := range entry_nexts {
							entry_next := models.EntryNext{}
							entryNextMap := val.(map[string]interface{})
							entry_next.Test = entryNextMap["test"].(*schema.Set).List()

							filter_entry.EntryNexts = append(filter_entry.EntryNexts, entry_next)
						}
					}

					if filterEntryMap["apply_to_frag"] != nil {
						filter_entry.ApplyToFrag = filterEntryMap["apply_to_frag"].(string)
					}

					filter.FilterEntrys = append(filter.FilterEntrys, filter_entry)
				}
			}

			contract.Filters = append(Contract.Filters, filter)
		}
	}

	err := aciClient.CreateContract(contract)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceAciContractRead(ctx, d, m)
}
