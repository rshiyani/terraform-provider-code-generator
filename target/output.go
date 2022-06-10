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
							Required:    true,
							Computed:    true,
							Description: "list of filter_entry",
							MaxItems:    4,
							MinItems:    1,
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
	contract = models.Contract{
		TenantDn: d.Get("tenant_dn").(string),
		Name:     d.Get("name").(string),
		MyMap:    d.Get("my_map"),
	}

	if Prio, ok := d.GetOk("prio"); ok {
		contract.Prio = Prio.(string)
	}

	contract.Casts = d.Get("cast").([]string)

	filters := d.Get("filter").([]interface{})

	for _, val := range filters {
		filterMap := val.(map[string]interface{})
		filter := models.Filter{
			FilterName: filterMap["filter_name"].(string),
		}
		if filterMap["description"] != nil {
			filter.Description = filterMap["description"].(string)
		}
		if filterMap["filter_entry"] != nil {
			filterEntrys := filterMap["filter_entry"].([]interface{})

			for _, val := range filterEntrys {
				filterEntryMap := val.(map[string]interface{})
				filterEntry := models.FilterEntry{
					FilterEntryName: filterEntryMap["filter_entry_name"].(string),
				}
				if filterEntryMap["entry_next"] != nil {
					entryNexts := filterEntryMap["entry_next"].([]interface{})

					for _, val := range entryNexts {
						entryNextMap := val.(map[string]interface{})
						entryNext := models.EntryNext{
							EntryNextName: entryNextMap["entry_next_name"].(string),
						}

						filterEntry.EntryNexts = append(filter_entry.EntryNexts, entryNext)
					}

				}

				if filterEntryMap["cast"] != nil {
					casts := filterEntryMap["cast"].(*schema.Set).List()

					for _, val := range casts {
						castMap := val.(map[string]interface{})
						cast := models.Cast{}
						cast.Cast2s = castMap["cast2"].(string)

						filterEntry.Casts = append(filter_entry.Casts, cast)
					}

				}

				if filterEntryMap["applyToFrag"] != nil {
					filterEntry.ApplyToFrag = filterEntryMap["apply_to_frag"].(string)
				}

				filter.FilterEntrys = append(filter.FilterEntrys, filterEntry)
			}

		}

		contract.Filters = append(Contract.Filters, filter)
	}

	err := aciClient.CreateContract(contract)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceAciContractRead(ctx, d, m)
}

func resourceAciContractUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	aciClient := m.(*client.Client)
	contract = models.Contract{
		TenantDn: d.Get("tenant_dn").(string),
		Name:     d.Get("name").(string),
		MyMap:    d.Get("my_map"),
	}

	if Prio, ok := d.GetOk("prio"); ok {
		contract.Prio = Prio.(string)
	}

	contract.Casts = d.Get("cast").([]string)

	filters := d.Get("filter").([]interface{})

	for _, val := range filters {
		filterMap := val.(map[string]interface{})
		filter := models.Filter{
			FilterName: filterMap["filter_name"].(string),
		}
		if filterMap["description"] != nil {
			filter.Description = filterMap["description"].(string)
		}
		if filterMap["filter_entry"] != nil {
			filterEntrys := filterMap["filter_entry"].([]interface{})

			for _, val := range filterEntrys {
				filterEntryMap := val.(map[string]interface{})
				filterEntry := models.FilterEntry{
					FilterEntryName: filterEntryMap["filter_entry_name"].(string),
				}
				if filterEntryMap["entry_next"] != nil {
					entryNexts := filterEntryMap["entry_next"].([]interface{})

					for _, val := range entryNexts {
						entryNextMap := val.(map[string]interface{})
						entryNext := models.EntryNext{
							EntryNextName: entryNextMap["entry_next_name"].(string),
						}

						filterEntry.EntryNexts = append(filter_entry.EntryNexts, entryNext)
					}

				}

				if filterEntryMap["cast"] != nil {
					casts := filterEntryMap["cast"].(*schema.Set).List()

					for _, val := range casts {
						castMap := val.(map[string]interface{})
						cast := models.Cast{}
						cast.Cast2s = castMap["cast2"].(string)

						filterEntry.Casts = append(filter_entry.Casts, cast)
					}

				}

				if filterEntryMap["applyToFrag"] != nil {
					filterEntry.ApplyToFrag = filterEntryMap["apply_to_frag"].(string)
				}

				filter.FilterEntrys = append(filter.FilterEntrys, filterEntry)
			}

		}

		contract.Filters = append(Contract.Filters, filter)
	}

	err := aciClient.UpdateContract(contract)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceAciContractRead(ctx, d, m)
}

func resourceAciContractDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	aciClient := m.(*client.Client)

	err := aciClient.DeleteContract(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diag.FromErr(err)
}
