package aci

import (
	"fmt"
	"regexp"
	"testing"
)

var resourceContractTest = map[string]interface{}{
	"name": map[string]interface{}{
		"valid":           searchInObject(resourceTenantTest, "filter.filter_entry.name.valid"),
		"invalid":         searchInObject(resourceTenantTest, "filter.filter_entry.name.invalid"),
		"multiple_valids": searchInObject(resourceTenantTest, "filter.filter_entry.name.multiple_valids"),
	},

	"temp": map[string]interface{}{
		"valid":           searchInObject(resourceApplicationProfileTest, "application_dn.valid"),
		"invalid":         searchInObject(resourceApplicationProfileTest, "application_dn.invalid"),
		"multiple_valids": searchInObject(resourceApplicationProfileTest, "application_dn.multiple_valids"),
	},

	"weight": map[string]interface{}{
		"valid":           []interface{}{-232.95313868437975, -96.95418868143406, -100.53403174895872, -153.46931907789437},
		"invalid":         []interface{}{"random", 10},
		"multiple_valids": []interface{}{-0.42853307801350426, 188.67339266292572, 39.462590117803934, -639.7643873528164, -145.8702448802587, -306.2923820171308, -336.4600235502115, 643.6038105074292, 152.81087148216577, -230.13598768146088, -273.86151467246873, -123.17694977239557, -467.7534722258367, 175.6709677092868, 103.03241124207479},
	},

	"ipv4_for": map[string]interface{}{
		"valid":           []interface{}{"208.41.80.53", "208.41.0.131", "208.41.32.130", "208.41.58.15"},
		"invalid":         []interface{}{"267.277.281.286"},
		"multiple_valids": []interface{}{"208.41.80.53", "208.41.0.131", "208.41.32.130", "208.41.58.15", "208.41.181.61", "208.41.151.152", "208.41.213.168", "208.41.229.158", "208.41.18.205", "208.41.208.140", "208.41.137.0", "208.41.205.210", "208.41.210.121", "208.41.41.108", "208.41.104.207"},
	},

	"port_number": map[string]interface{}{
		"valid":           []interface{}{1, 65535, 39319, 62959},
		"invalid":         []interface{}{0, 65536},
		"multiple_valids": []interface{}{1, 65535, 2092, 27356, 5698, 21721, 8200, 3654, 4679, 63781, 52123, 39407, 22047, 60818, 45608},
	},

	"temp_schema_list": map[string]interface{}{
		"valid":           []interface{}{"l0ce6bicqr", "t6tqfxypan", "jn79uvss2k", "9axw2vh42x"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"asnyiktdiu", "yx8vp6ia82", "wru6clk9fy", "d4pwld83iw", "77xo07fv5l", "emybd0cxzz", "u375ju5b9i", "0lhwadcoij", "k6j0iu02ys", "8habasmzem", "zhcos7tk1u", "crs09ihmnx", "rdown2p67y", "okb4ngk7hh", "hfzkc3s3jm"},
	},

	"test_score": map[string]interface{}{
		"valid":           []interface{}{1, 100, 50, 48},
		"invalid":         []interface{}{0, 101},
		"multiple_valids": []interface{}{1, 100, 50, 30, 37, 36, 66, 53, 37, 77, 46, 3, 24, 29, 35},
	},

	"string_in_some_names": map[string]interface{}{
		"valid":           []interface{}{"parth", "aarsh", "arjun", "alfatah", "krunal"},
		"invalid":         []interface{}{"wnoe5v1ltx"},
		"multiple_valids": []interface{}{"parth", "aarsh", "arjun", "alfatah", "krunal"},
	},

	"valid_cidr": map[string]interface{}{
		"valid":           []interface{}{0, 32, 16, 21},
		"invalid":         []interface{}{-1, 33},
		"multiple_valids": []interface{}{0, 32, 16, 30, 10, 5, 12, 29, 21, 28, 15, 31, 5, 20, 4},
	},

	"percentage": map[string]interface{}{
		"valid":           []interface{}{0, 100, 50.0, 79.25895836369423},
		"invalid":         []interface{}{-1, 101},
		"multiple_valids": []interface{}{0, 100, 50.0, 31.39012924699733, 14.552797639146688, 7.683031514319582, 53.339113566464306, 39.74839659807847, 23.870415037284626, 55.48580085505294, 11.85264443892742, 0.4145263169945834, 7.358523616848771, 13.811146447944788, 0.5287296784194194},
	},

	"testingmap": map[string]interface{}{
		"valid":           []interface{}{-319, 717, 845, -48},
		"invalid":         []interface{}{"random", 10.023},
		"multiple_valids": []interface{}{656, 943, 387, 905, 344, -10, 505, -279, 198, -398, 419, -913, -752, 762, -761},
	},

	"filter": map[string]interface{}{
		"filter_name": map[string]interface{}{
			"valid":           []interface{}{"b3m2rks0pa", "bqw0fnv2dv", "yufofii005", "ha8b0w39ad"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"1h01rgdbbx", "v58fxfrfm8", "q6rni9yuym", "1x1v7dxk42", "hhaj7zxaw2", "32hfmysi97", "lozdc6ypd2", "xd6ppmyhkv", "3mnd78x97k", "33adg17a8f", "nyoeb5k53z", "ah905i61rq", "kwxvs3hxcs", "fng6v8h2kp", "c0u70qto06"},
		},

		"id": map[string]interface{}{
			"valid":           []interface{}{"cwtl9jl5dl", "3tbpki57j1", "d1wzbikz34", "w0np8ipbbu"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"4uvcqkruvk", "b5o5spxgpn", "ie91zbt1cw", "osv52f820q", "g2vzz1p6eu", "o7s2k423gg", "eriri4lsyd", "hkit2zsq8p", "dhayvod9q2", "4uxwe0x7ew", "kmer14xvrm", "jza9r5g1ob", "ngbvi7prne", "l4x8cwvluo", "zqh27hs388"},
		},

		"description": map[string]interface{}{
			"valid":           []interface{}{"grfar7akjm", "eokmc4yxw6", "qxh9uctivj", "p4b065xhj6"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"qi53fweqwu", "ts9zdkcbyo", "40t9y82ayz", "zwrf9hk7hy", "ko6w9t6f6u", "v331lkxda6", "kosn4ija6z", "o4ocssfhms", "igzxcvrw63", "axxbablnz8", "pdu9vvsdmv", "wm6eodb6ao", "8cjbquxov8", "hsm9ictzxj", "30irtvxhya"},
		},

		"filter_entry": map[string]interface{}{
			"id_list": map[string]interface{}{
				"valid":           []interface{}{"zgcpqsjbf8", "wieiuvwpva", "2f856ubhuu", "xu2p664hzr"},
				"invalid":         []interface{}{10, 12.43},
				"multiple_valids": []interface{}{"opdbinmxur", "2tkx0a814s", "3fs5uz5we7", "ehnt8cb4au", "ua4w2wvr13", "7fyfg0elw4", "d7vkf7r0q4", "xdn60l052r", "tzr6xjb1jz", "xggvme6hxk", "v06m4ycnf9", "yeddvm3tsp", "j6ix7djs2x", "nnwkzr16rl", "8bz45hazhp"},
			},

			"filter_entry_name": map[string]interface{}{
				"valid":           []interface{}{"dj1fbp5hek", "v7jdigiz0x", "dnarfr2byc", "soppt6a4nc"},
				"invalid":         []interface{}{10, 12.43},
				"multiple_valids": []interface{}{"445slsb139", "e7eu9kgm66", "rfw3w2tm9k", "p90iv52bcd", "jdzu2qy85w", "4ln6f8jk8j", "0pj0wqp06y", "nchmd5nfvl", "a5a5eawook", "mybpoib260", "hkgz532dcn", "1u0vu7cqlw", "v6mhd9wwtl", "cfdewt4rfp", "lw3wydup37"},
			},

			"ipv6": map[string]interface{}{
				"valid":           []interface{}{"2001:db8::34f4:0:0:f3d4", "2001:db8::34f4:0:0:f363", "2001:db8::34f4:0:0:f317", "2001:db8::34f4:0:0:f304"},
				"invalid":         []interface{}{"invalidIPv6"},
				"multiple_valids": []interface{}{"2001:db8::34f4:0:0:f3d4", "2001:db8::34f4:0:0:f363", "2001:db8::34f4:0:0:f317", "2001:db8::34f4:0:0:f304", "2001:db8::34f4:0:0:f329", "2001:db8::34f4:0:0:f30b", "2001:db8::34f4:0:0:f3e4", "2001:db8::34f4:0:0:f35d", "2001:db8::34f4:0:0:f362", "2001:db8::34f4:0:0:f335", "2001:db8::34f4:0:0:f3aa", "2001:db8::34f4:0:0:f3a9", "2001:db8::34f4:0:0:f3d7", "2001:db8::34f4:0:0:f35b", "2001:db8::34f4:0:0:f349"},
			},

			"apply_to_frag": map[string]interface{}{
				"valid":           []interface{}{"gqdrd01uu3", "qkbx6thckg", "qz6bk4akah", "26fa61kfo6"},
				"invalid":         []interface{}{"yes", "no"},
				"multiple_valids": []interface{}{"57yjrpnc2o", "865i9yrqjb", "2p6q5rit52", "zshbeqoh3h", "tc07o5w9it", "kfbu25w9or", "cvdblfr6vb", "gmbo567a5l", "2sarz43gry", "8dk1qnq3cq", "xhe12vre6h", "fd0wuk6l0v", "mucg7teqse", "yjq8wrizns", "y475ody7ir"},
			},

			"apply_to_frag_liist_schema": map[string]interface{}{
				"valid":           []interface{}{true, false},
				"invalid":         []interface{}{"random", 10},
				"multiple_valids": []interface{}{true, false},
			},
		},
	},
}

func TestAccAciContract_Basic(t *testing.T) {
	var contract_default models.Contract
	var contract_updated models.Contract
	resourceName := "aci_contract.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAciContractDestroy,
		Steps: []resource.TestStep{
			{
				Config:      CreateAccContractWithoutName(),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutTemp(),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutWeight(),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutIpv4For(),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config: CreateAccContractConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractExists(resourceName, &contract_default),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%v", searchInObject(resourceTenantTest, "filter.filter_entry.name.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "temp", fmt.Sprintf("%v", searchInObject(resourceApplicationProfileTest, "application_dn.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "weight", fmt.Sprintf("%v", searchInObject(resourceContractTest, "weight.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "ipv4_for", fmt.Sprintf("%v", searchInObject(resourceContractTest, "ipv4_for.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "port_number", "0"),

					resource.TestCheckResourceAttr(resourceName, "temp_schema_list.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "temp_schema_list.0", "Hello"),
					resource.TestCheckResourceAttr(resourceName, "temp_schema_list.1", "10.3442"),
					resource.TestCheckResourceAttr(resourceName, "temp_schema_list.2", "2"),

					resource.TestCheckResourceAttr(resourceName, "test_score", "0"),

					resource.TestCheckResourceAttr(resourceName, "string_in_some_names", "parth"),

					resource.TestCheckResourceAttr(resourceName, "valid_cidr", ""),

					resource.TestCheckResourceAttr(resourceName, "percentage", "0.0"),

					resource.TestCheckResourceAttr(resourceName, "filter.#", "0"),
				),
			},
			{
				Config: CreateAccContractConfigWithOptional(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractExists(resourceName, &contract_updated),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%v", searchInObject(resourceTenantTest, "filter.filter_entry.name.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "temp", fmt.Sprintf("%v", searchInObject(resourceApplicationProfileTest, "application_dn.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "weight", fmt.Sprintf("%v", searchInObject(resourceContractTest, "weight.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "ipv4_for", fmt.Sprintf("%v", searchInObject(resourceContractTest, "ipv4_for.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "port_number", fmt.Sprintf("%v", searchInObject(resourceContractTest, "port_number.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "temp_schema_list.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "temp_schema_list.0", fmt.Sprintf("%v", searchInObject(resourceContractTest, "temp_schema_list.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "temp_schema_list.1", fmt.Sprintf("%v", searchInObject(resourceContractTest, "temp_schema_list.valid.1"))),
					resource.TestCheckResourceAttr(resourceName, "test_score", fmt.Sprintf("%v", searchInObject(resourceContractTest, "test_score.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "string_in_some_names", fmt.Sprintf("%v", searchInObject(resourceContractTest, "string_in_some_names.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "valid_cidr", fmt.Sprintf("%v", searchInObject(resourceContractTest, "valid_cidr.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "percentage", fmt.Sprintf("%v", searchInObject(resourceContractTest, "percentage.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "filter.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_name", fmt.Sprintf("%v", searchInObject(resourceContractTest, "filter.filter_name.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "filter.0.id", fmt.Sprintf("%v", searchInObject(resourceContractTest, "filter.id.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "filter.0.description", fmt.Sprintf("%v", searchInObject(resourceContractTest, "filter.description.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.id_list.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.id_list.0", fmt.Sprintf("%v", searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.id_list.1", fmt.Sprintf("%v", searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.1"))),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.filter_entry_name", fmt.Sprintf("%v", searchInObject(resourceContractTest, "filter.filter_entry.filter_entry_name.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.ipv6", fmt.Sprintf("%v", searchInObject(resourceContractTest, "filter.filter_entry.ipv6.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.apply_to_frag", fmt.Sprintf("%v", searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.apply_to_frag_liist_schema.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.apply_to_frag_liist_schema.0", fmt.Sprintf("%v", searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.apply_to_frag_liist_schema.1", fmt.Sprintf("%v", searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid.1"))),

					testAccCheckAciContractIdEqual(&contract_default, &contract_updated),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: CreateAccContractConfig(),
			},
		},
	})
}

func TestAccAciContract_Update(t *testing.T) {
	var contract_default models.Contract
	var contract_updated models.Contract
	resourceName := "aci_contract.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAciContractDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccContractConfig(),
				Check:  testAccCheckAciContractExists(resourceName, &contract_default),
			},
		}, generateStepForUpdatedAttr(resourceName, &contract_default, &contract_updated)...),
	})
}

func CreateAccContractWithoutName() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "aci_contract" "test" {
									temp = aci_application_profile.test.application_dn
									weight = "%v"
									ipv4_for = "%v"
									port_number = "%v"
									temp_schema_list = ["%v","%v"]
									test_score = "%v"
									string_in_some_names = "%v"
									valid_cidr = "%v"
									percentage = "%v"
									filter {
    
									                        
                                        filter_name = "%v"
                        
                                        id = "%v"
                        
                                        description = "%v"
                                        filter_entry {
                                                    
                                            id_list = ["%v","%v"]
                        
                                            filter_entry_name = "%v"
                        
                                            ipv6 = "%v"
                        
                                            apply_to_frag = "%v"
                        
                                            apply_to_frag_liist_schema = ["%v","%v"]

                                          }

									}
				}
			`, searchInObject(resourceContractTest, "weight.valid.0"),
		searchInObject(resourceContractTest, "ipv4_for.valid.0"),
		searchInObject(resourceContractTest, "port_number.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.1"),
		searchInObject(resourceContractTest, "test_score.valid.0"),
		searchInObject(resourceContractTest, "string_in_some_names.valid.0"),
		searchInObject(resourceContractTest, "valid_cidr.valid.0"),
		searchInObject(resourceContractTest, "percentage.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_name.valid.0"),
		searchInObject(resourceContractTest, "filter.id.valid.0"),
		searchInObject(resourceContractTest, "filter.description.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.1"),
		searchInObject(resourceContractTest, "filter.filter_entry.filter_entry_name.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.ipv6.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid.1"))
	return resource
}
func CreateAccContractWithoutTemp() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "aci_contract" "test" {
									name = aci_tenant.test.filter.filter_entry.name
									weight = "%v"
									ipv4_for = "%v"
									port_number = "%v"
									temp_schema_list = ["%v","%v"]
									test_score = "%v"
									string_in_some_names = "%v"
									valid_cidr = "%v"
									percentage = "%v"
									filter {
    
									                        
                                        filter_name = "%v"
                        
                                        id = "%v"
                        
                                        description = "%v"
                                        filter_entry {
                                                    
                                            id_list = ["%v","%v"]
                        
                                            filter_entry_name = "%v"
                        
                                            ipv6 = "%v"
                        
                                            apply_to_frag = "%v"
                        
                                            apply_to_frag_liist_schema = ["%v","%v"]

                                          }

									}
				}
			`, searchInObject(resourceContractTest, "weight.valid.0"),
		searchInObject(resourceContractTest, "ipv4_for.valid.0"),
		searchInObject(resourceContractTest, "port_number.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.1"),
		searchInObject(resourceContractTest, "test_score.valid.0"),
		searchInObject(resourceContractTest, "string_in_some_names.valid.0"),
		searchInObject(resourceContractTest, "valid_cidr.valid.0"),
		searchInObject(resourceContractTest, "percentage.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_name.valid.0"),
		searchInObject(resourceContractTest, "filter.id.valid.0"),
		searchInObject(resourceContractTest, "filter.description.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.1"),
		searchInObject(resourceContractTest, "filter.filter_entry.filter_entry_name.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.ipv6.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid.1"))
	return resource
}
func CreateAccContractWithoutWeight() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "aci_contract" "test" {
									name = aci_tenant.test.filter.filter_entry.name
									temp = aci_application_profile.test.application_dn
									ipv4_for = "%v"
									port_number = "%v"
									temp_schema_list = ["%v","%v"]
									test_score = "%v"
									string_in_some_names = "%v"
									valid_cidr = "%v"
									percentage = "%v"
									filter {
    
									                        
                                        filter_name = "%v"
                        
                                        id = "%v"
                        
                                        description = "%v"
                                        filter_entry {
                                                    
                                            id_list = ["%v","%v"]
                        
                                            filter_entry_name = "%v"
                        
                                            ipv6 = "%v"
                        
                                            apply_to_frag = "%v"
                        
                                            apply_to_frag_liist_schema = ["%v","%v"]

                                          }

									}
				}
			`, searchInObject(resourceContractTest, "ipv4_for.valid.0"),
		searchInObject(resourceContractTest, "port_number.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.1"),
		searchInObject(resourceContractTest, "test_score.valid.0"),
		searchInObject(resourceContractTest, "string_in_some_names.valid.0"),
		searchInObject(resourceContractTest, "valid_cidr.valid.0"),
		searchInObject(resourceContractTest, "percentage.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_name.valid.0"),
		searchInObject(resourceContractTest, "filter.id.valid.0"),
		searchInObject(resourceContractTest, "filter.description.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.1"),
		searchInObject(resourceContractTest, "filter.filter_entry.filter_entry_name.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.ipv6.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid.1"))
	return resource
}
func CreateAccContractWithoutIpv4For() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "aci_contract" "test" {
									name = aci_tenant.test.filter.filter_entry.name
									temp = aci_application_profile.test.application_dn
									weight = "%v"
									port_number = "%v"
									temp_schema_list = ["%v","%v"]
									test_score = "%v"
									string_in_some_names = "%v"
									valid_cidr = "%v"
									percentage = "%v"
									filter {
    
									                        
                                        filter_name = "%v"
                        
                                        id = "%v"
                        
                                        description = "%v"
                                        filter_entry {
                                                    
                                            id_list = ["%v","%v"]
                        
                                            filter_entry_name = "%v"
                        
                                            ipv6 = "%v"
                        
                                            apply_to_frag = "%v"
                        
                                            apply_to_frag_liist_schema = ["%v","%v"]

                                          }

									}
				}
			`, searchInObject(resourceContractTest, "weight.valid.0"),
		searchInObject(resourceContractTest, "port_number.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.1"),
		searchInObject(resourceContractTest, "test_score.valid.0"),
		searchInObject(resourceContractTest, "string_in_some_names.valid.0"),
		searchInObject(resourceContractTest, "valid_cidr.valid.0"),
		searchInObject(resourceContractTest, "percentage.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_name.valid.0"),
		searchInObject(resourceContractTest, "filter.id.valid.0"),
		searchInObject(resourceContractTest, "filter.description.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.1"),
		searchInObject(resourceContractTest, "filter.filter_entry.filter_entry_name.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.ipv6.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid.1"))
	return resource
}

func CreateAccContractConfig() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
		resource  "aci_contract" "test" {
							name = aci_tenant.test.filter.filter_entry.name
							temp = aci_application_profile.test.application_dn
							weight = "%v"
							ipv4_for = "%v"
		}
	`, searchInObject(resourceContractTest, "weight.valid.0"),
		searchInObject(resourceContractTest, "ipv4_for.valid.0"))
	return resource
}

func CreateAccContractConfigWithOptional() string {
	resource := createContractConfig(getParentContract())
	return resource
}

func generateStepForUpdatedAttr(resourceName string, contract_default, contract_updated *models.Contract) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var valid []interface{}
	valid = searchInObject(resourceContractTest, "port_number.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccContractUpdatedAttr("port_number", value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAciContractExists(resourceName, contract_updated),
				resource.TestCheckResourceAttr(resourceName, "port_number", v),
				testAccCheckAciContractIdEqual(contract_default, contract_updated),
			),
		})
	}
	valid = searchInObject(resourceContractTest, "temp_schema_list.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("[%v]", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccContractUpdatedAttrTempSchemaList("temp_schema_list", value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAciContractExists(resourceName, contract_updated),
				resource.TestCheckResourceAttr(resourceName, "temp_schema_list", v),
				testAccCheckAciContractIdEqual(contract_default, contract_updated),
			),
		})
	}
	valid = searchInObject(resourceContractTest, "test_score.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccContractUpdatedAttr("test_score", value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAciContractExists(resourceName, contract_updated),
				resource.TestCheckResourceAttr(resourceName, "test_score", v),
				testAccCheckAciContractIdEqual(contract_default, contract_updated),
			),
		})
	}
	valid = searchInObject(resourceContractTest, "string_in_some_names.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccContractUpdatedAttr("string_in_some_names", value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAciContractExists(resourceName, contract_updated),
				resource.TestCheckResourceAttr(resourceName, "string_in_some_names", v),
				testAccCheckAciContractIdEqual(contract_default, contract_updated),
			),
		})
	}
	valid = searchInObject(resourceContractTest, "valid_cidr.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccContractUpdatedAttr("valid_cidr", value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAciContractExists(resourceName, contract_updated),
				resource.TestCheckResourceAttr(resourceName, "valid_cidr", v),
				testAccCheckAciContractIdEqual(contract_default, contract_updated),
			),
		})
	}
	valid = searchInObject(resourceContractTest, "percentage.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccContractUpdatedAttr("percentage", value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAciContractExists(resourceName, contract_updated),
				resource.TestCheckResourceAttr(resourceName, "percentage", v),
				testAccCheckAciContractIdEqual(contract_default, contract_updated),
			),
		})
	}
	valid = searchInObject(resourceContractTest, "filter.description.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccContractUpdatedAttrFilter("description", value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAciContractExists(resourceName, contract_updated),
				resource.TestCheckResourceAttr(resourceName, "filter.0.description", v),
				testAccCheckAciContractIdEqual(contract_default, contract_updated),
			),
		})
	}
	valid = searchInObject(resourceContractTest, "filter.filter_entry.ipv6.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccContractUpdatedAttrFilterFilterEntry("ipv6", value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAciContractExists(resourceName, contract_updated),
				resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.ipv6", v),
				testAccCheckAciContractIdEqual(contract_default, contract_updated),
			),
		})
	}
	valid = searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccContractUpdatedAttrFilterFilterEntry("apply_to_frag", value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAciContractExists(resourceName, contract_updated),
				resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.apply_to_frag", v),
				testAccCheckAciContractIdEqual(contract_default, contract_updated),
			),
		})
	}
	valid = searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("[%v]", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccContractUpdatedAttrFilterFilterEntry("apply_to_frag_liist_schema", value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAciContractExists(resourceName, contract_updated),
				resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.apply_to_frag_liist_schema", v),
				testAccCheckAciContractIdEqual(contract_default, contract_updated),
			),
		})
	}

	return testSteps
}

func testAccCheckAciContractExists(name string, contract *models.Contract) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// [TODO]: Write your code here
	}
}

func testAccCheckMoviesMovieDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_contract" {
			// [TODO]: Write your code here
		}
	}
	return nil
}

func testAccCheckAciContractIdEqual(contract1, contract2 *models.Contract) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		Id1, err := getIdFromContractModel(contract1)
		if err != nil {
			return err
		}
		Id2, err := getIdFromContractModel(contract2)
		if err != nil {
			return err
		}
		if Id1 != Id2 {
			return fmt.Errorf("Contract IDs are not equal")
		}
		return nil
	}
}

func getParentContract() []string {
	t := []string{}
	t = append(t, getParentTenant()...)
	t = append(t, getParentApplicationProfile()...)
	t = append(t, contractBlock())
	return t
}

func contractBlock() string {
	return fmt.Sprintf(`
		resource  "aci_contract" "test" {
						name = aci_tenant.test.filter.filter_entry.name
						temp = aci_application_profile.test.application_dn
						weight = "%v"
						ipv4_for = "%v"
						port_number = "%v"
				        temp_schema_list = ["%v","%v"]
						test_score = "%v"
						string_in_some_names = "%v"
						valid_cidr = "%v"
						percentage = "%v"
                        filter {
    
                                                
                            filter_name = "%v"
                        
                            id = "%v"
                        
                            description = "%v"
                            filter_entry {
                                                    
                                id_list = ["%v","%v"]
                        
                                filter_entry_name = "%v"
                        
                                ipv6 = "%v"
                        
                                apply_to_frag = "%v"
                        
                                apply_to_frag_liist_schema = ["%v","%v"]

                              }

                        }
		}
	`, searchInObject(resourceContractTest, "weight.valid.0"),
		searchInObject(resourceContractTest, "ipv4_for.valid.0"),
		searchInObject(resourceContractTest, "port_number.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.1"),
		searchInObject(resourceContractTest, "test_score.valid.0"),
		searchInObject(resourceContractTest, "string_in_some_names.valid.0"),
		searchInObject(resourceContractTest, "valid_cidr.valid.0"),
		searchInObject(resourceContractTest, "percentage.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_name.valid.0"),
		searchInObject(resourceContractTest, "filter.id.valid.0"),
		searchInObject(resourceContractTest, "filter.description.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.1"),
		searchInObject(resourceContractTest, "filter.filter_entry.filter_entry_name.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.ipv6.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid.1"))
}

// To eliminate duplicate resource block from slice of resource blocks
func createContractConfig(configSlice []string) string {
	keys := make(map[string]bool)
	str := ""

	for _, entry := range configSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			str += entry
		}
	}

	return str
}
