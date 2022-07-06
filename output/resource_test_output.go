package aci

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RutvikS-crest/movies-go-client/client"
)

const contractSelfRequiredCount = 5

var resourceContractTest = map[string]interface{}{
	"name": map[string]interface{}{
		"valid":           []interface{}{"b6epjiajkc", "c7o6sg433u", "y04f9l7p5s", "j6wcreq2do"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"3bqkrd25nz", "ps33hh05jw", "atdj7y9tgh", "p8a8i1kth0", "x043jtvts2", "s0daqkly3e", "nzkj5jib3a", "iayfzajfix", "mqvlijttxf", "zuij23gl4m", "o37n5xulux", "6sq3bul6r6", "tnv0bvsm8b", "glgpm6c99h", "jz9b0on36l"},
	},

	"temp": map[string]interface{}{
		"valid":           []interface{}{180, 45, 650, 388},
		"invalid":         []interface{}{"random", 10.023},
		"multiple_valids": []interface{}{-72, 395, -455, 84, 283, -265, -804, 36, -925, -32, 469, 730, 425, -416, -564},
	},

	"weight": map[string]interface{}{
		"valid":           []interface{}{-968.2083562524816, 319.83828987522924, 914.3886433874081, 533.7449091697114},
		"invalid":         []interface{}{"random", 10},
		"multiple_valids": []interface{}{39.85083893763871, -631.6408667714682, -944.267139291758, 213.58423554456127, 373.9668403322991, 552.8699640142963, -621.7213035453773, 971.9117263000096, -656.5891299950213, 108.12052311373095, 993.2418126639959, -935.9610603231043, -353.30039837746784, 628.8930792119718, -657.903676337481},
	},

	"ipv4_for": map[string]interface{}{
		"valid":           searchInObject(Test, "ipv4_for.valid"),
		"invalid":         searchInObject(Test, "ipv4_for.invalid"),
		"multiple_valids": searchInObject(Test, "ipv4_for.multiple_valids"),
	},

	"port_number": map[string]interface{}{
		"valid":           []interface{}{1, 65535, 37225, 53880},
		"invalid":         []interface{}{0, 65536},
		"multiple_valids": []interface{}{1, 65535, 14311, 26358, 41876, 11422, 47173, 21233, 34776, 11221, 47279, 10056, 751, 46741, 57119},
	},

	"temp_schema_list": map[string]interface{}{
		"valid":           []interface{}{"wvwkl1j7qm", "svu5byk2os", "o9cquhcauv", "ie9zo3bglk"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"1cwl6sg3ko", "jqww9eauck", "z2dcvjr1d1", "5bys2v04dj", "o4o4lw4rx3", "li376ugct2", "7zbmbfe29g", "lcjs78lrdm", "h00prlm13w", "9sz8zpzwr0", "2ttf0qre1g", "v57d9fkf7x", "dzec0df43b", "nk78kitujg", "h6g5p0osrw"},
	},

	"test_score": map[string]interface{}{
		"valid":           []interface{}{1, 100, 50, 88},
		"invalid":         []interface{}{0, 101},
		"multiple_valids": []interface{}{1, 100, 50, 5, 83, 2, 61, 65, 17, 84, 15, 86, 67, 35, 27},
	},

	"string_in_some_names": map[string]interface{}{
		"valid":           []interface{}{"parth", "aarsh", "arjun", "alfatah", "krunal"},
		"invalid":         []interface{}{"6t4mt1im0l"},
		"multiple_valids": []interface{}{"parth", "aarsh", "arjun", "alfatah", "krunal"},
	},

	"valid_cidr": map[string]interface{}{
		"valid":           []interface{}{0, 32, 16, 23},
		"invalid":         []interface{}{-1, 33},
		"multiple_valids": []interface{}{0, 32, 16, 29, 6, 9, 10, 24, 11, 25, 13, 28, 7, 23, 13},
	},

	"percentage": map[string]interface{}{
		"valid":           []interface{}{0, 100, 50.0, 8.901470198524594},
		"invalid":         []interface{}{-1, 101},
		"multiple_valids": []interface{}{0, 100, 50.0, 69.23049722298951, 45.187452338007105, 35.251861221822146, 59.00225935528169, 54.564866948679914, 35.90301308432537, 98.18484397088255, 87.3436783369663, 9.696505845874814, 98.95790156011914, 53.173833837510124, 0.7119919193560287},
	},

	"testingmap": map[string]interface{}{
		"valid":           []interface{}{-471, 586, -778, -176},
		"invalid":         []interface{}{"random", 10.023},
		"multiple_valids": []interface{}{392, -58, -662, 981, 939, -926, -574, 715, -99, 283, -41, -534, -317, 324, -437},
	},

	"filter": map[string]interface{}{
		"filter_name": map[string]interface{}{
			"valid":           []interface{}{50, 100, 75.0, 52.03287158785944},
			"invalid":         []interface{}{49, 101},
			"multiple_valids": []interface{}{50, 100, 75.0, 76.86387871635571, 89.11103592316053, 85.6898488194801, 83.7025704049203, 92.43865170222696, 63.04113698293864, 73.51116852350565, 62.462066017526006, 64.66893937964593, 84.41348388396842, 82.67839221075506, 56.51072346271268},
		},

		"id": map[string]interface{}{
			"valid":           []interface{}{"fast21p525", "of946zaloh", "6dnxsknkbj", "fdai9wx99u"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"pzpe3ogcbx", "zuyayqaw21", "em1eausher", "j0hxifgp0i", "5wol7aigt0", "g5yd82ho8s", "8emu78z6hd", "4115j9xbtu", "qdgt8nvxes", "v2pmdff59x", "5dq2dumfzd", "op99crepbg", "p0rcp62wrd", "1a5d8dx4xv", "me492u99t5"},
		},

		"description": map[string]interface{}{
			"valid":           []interface{}{"82attosmwc", "a2rs7no28o", "it6e3grd7s", "prith0q4t7"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"yl0v44j8qg", "vmcpc973zn", "6nod5wt0wd", "fgzgu9wt04", "dc8nll0tqh", "p6iy35npp9", "pp4fgmxz0a", "ap2vwgx3m5", "yolvvkjk1o", "12qq6e75di", "rtv0uhujlu", "cth118uwnd", "3tql8m34p9", "fp58g43a3c", "6o8dzm3oej"},
		},

		"filter_entry": map[string]interface{}{
			"id_list": map[string]interface{}{
				"valid":           []interface{}{"cq7fzzofpi", "a7trdzcdrd", "7gpnn3gb6q", "gm21lda2jj"},
				"invalid":         []interface{}{10, 12.43},
				"multiple_valids": []interface{}{"zkoeuwgcki", "wileekqdi2", "9f122iqcym", "7ceuwkiu9w", "l6839w4vf1", "7n53gqncqq", "3lgji2c3tg", "u4hzmrzrn5", "z3fl1398rh", "ysrvjhzexl", "xzkipoh4od", "t9eark9dsj", "xi1sjnfdyo", "6emwdqz320", "s9p2zlju2t"},
			},

			"filter_entry_name": map[string]interface{}{
				"valid":           []interface{}{"60jtv6ol7v", "aatca1c5xm", "a7q43k503l", "2ubim52ixk"},
				"invalid":         []interface{}{10, 12.43},
				"multiple_valids": []interface{}{"f2frj2twzt", "pn2zbozfir", "pe0txf3l2v", "cbpcuj9d1g", "o5gu806orb", "lx8s6tfhbs", "lha0x53dud", "otxn48oh7b", "idd5hcoqq3", "n2j891fgml", "eotyqmsld8", "0l2fskm89b", "7498kf59m2", "trlbh6qmjc", "ootzzctspc"},
			},

			"ipv6": map[string]interface{}{
				"valid":           searchInObject(Test, "ipv6.valid"),
				"invalid":         searchInObject(Test, "ipv6.invalid"),
				"multiple_valids": searchInObject(Test, "ipv6.multiple_valids"),
			},

			"apply_to_frag": map[string]interface{}{
				"valid":           []interface{}{"u91jaxum1e", "5og5tw9xl5", "ktub750mch", "8uhh3stlms"},
				"invalid":         []interface{}{"yes", "no"},
				"multiple_valids": []interface{}{"nnjp87lzy5", "gx6kh8t99b", "1daquyuznb", "vwx46ikrur", "rkordixx7f", "1pwosqviza", "j9pf9fphwm", "tm9hscvy99", "rzk4xlcigv", "x7zkvzbbln", "h6tx35h7al", "6f0lw83c18", "9dhpacrob5", "rztkm4lnk7", "t2w535jh8r"},
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
		Steps: append([]resource.TestStep{
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
				Config:      CreateAccContractWithoutTempSchemaList(),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutFilter(),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config: CreateAccContractConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractExists(resourceName, &contract_default),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%v", searchInObject(resourceContractTest, "name.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "temp", fmt.Sprintf("%v", searchInObject(resourceContractTest, "temp.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "weight", fmt.Sprintf("%v", searchInObject(resourceContractTest, "weight.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "ipv4_for", ""),

					resource.TestCheckResourceAttr(resourceName, "port_number", "0"),

					resource.TestCheckResourceAttr(resourceName, "temp_schema_list.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "temp_schema_list.0", fmt.Sprintf("%v", searchInObject(resourceContractTest, "temp_schema_list.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "temp_schema_list.1", fmt.Sprintf("%v", searchInObject(resourceContractTest, "temp_schema_list.valid.1"))),

					resource.TestCheckResourceAttr(resourceName, "test_score", "0"),

					resource.TestCheckResourceAttr(resourceName, "string_in_some_names", "parth"),

					resource.TestCheckResourceAttr(resourceName, "valid_cidr", ""),

					resource.TestCheckResourceAttr(resourceName, "percentage", "0.0"),

					resource.TestCheckResourceAttr(resourceName, "filter.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_name", "0.0"),
					resource.TestCheckResourceAttr(resourceName, "filter.0.id", ""),
					resource.TestCheckResourceAttr(resourceName, "filter.0.description", fmt.Sprintf("%v", searchInObject(resourceContractTest, "filter.description.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.id_list.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.id_list.0", fmt.Sprintf("%v", searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.id_list.1", fmt.Sprintf("%v", searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.1"))),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.filter_entry_name", fmt.Sprintf("%v", searchInObject(resourceContractTest, "filter.filter_entry.filter_entry_name.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.ipv6", ""),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.apply_to_frag", ""),
					resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.apply_to_frag_liist_schema.#", "0"),
				),
			},
			{
				Config: CreateAccContractConfigWithOptional(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractExists(resourceName, &contract_updated),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%v", searchInObject(resourceContractTest, "name.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "temp", fmt.Sprintf("%v", searchInObject(resourceContractTest, "temp.valid.0"))),
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
		}, generateStepForUpdatedRequiredAttr(resourceName, &contract_default, &contract_updated)...),
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

func TestAccAciContract_NegativeCases(t *testing.T) {
	resourceName := "aci_contract.test"

	// [TODO]: Add makeTestVariable() to utils.go file
	// rName := makeTestVariable(acctest.RandString(5))
	// rOther := makeTestVariable(acctest.RandString(5))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAciContractDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccContractConfig(),
			},
		}, generateNegativeSteps(resourceName)...),
	})
}

func TestAccAciContract_MultipleCreateDelete(t *testing.T) {
	resourceName := "aci_contract.test"

	// [TODO]: Add makeTestVariable() to utils.go file
	// rName := makeTestVariable(acctest.RandString(5))
	// rOther := makeTestVariable(acctest.RandString(5))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAciContractDestroy,
		Steps: []resource.TestStep{
			{
				Config: CreateAccContractMultipleConfig(),
			},
		},
	})
}

func CreateAccContractWithoutName() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "aci_contract" "test" {

									temp = %v

									weight = %v

									ipv4_for = "%v"

									port_number = %v

									temp_schema_list = ["%v","%v"]

									test_score = %v

									string_in_some_names = "%v"

									valid_cidr = "%v"

									percentage = %v

									filter {
    
									                        
                                        filter_name = %v
                        
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
			`, searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
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

									name = "%v"

									weight = %v

									ipv4_for = "%v"

									port_number = %v

									temp_schema_list = ["%v","%v"]

									test_score = %v

									string_in_some_names = "%v"

									valid_cidr = "%v"

									percentage = %v

									filter {
    
									                        
                                        filter_name = %v
                        
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
			`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
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

									name = "%v"

									temp = %v

									ipv4_for = "%v"

									port_number = %v

									temp_schema_list = ["%v","%v"]

									test_score = %v

									string_in_some_names = "%v"

									valid_cidr = "%v"

									percentage = %v

									filter {
    
									                        
                                        filter_name = %v
                        
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
			`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
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
func CreateAccContractWithoutTempSchemaList() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "aci_contract" "test" {

									name = "%v"

									temp = %v

									weight = %v

									ipv4_for = "%v"

									port_number = %v

									test_score = %v

									string_in_some_names = "%v"

									valid_cidr = "%v"

									percentage = %v

									filter {
    
									                        
                                        filter_name = %v
                        
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
			`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
		searchInObject(resourceContractTest, "ipv4_for.valid.0"),
		searchInObject(resourceContractTest, "port_number.valid.0"),
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
func CreateAccContractWithoutFilter() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "aci_contract" "test" {

									name = "%v"

									temp = %v

									weight = %v

									ipv4_for = "%v"

									port_number = %v

									temp_schema_list = ["%v","%v"]

									test_score = %v

									string_in_some_names = "%v"

									valid_cidr = "%v"

									percentage = %v
				}
			`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
		searchInObject(resourceContractTest, "ipv4_for.valid.0"),
		searchInObject(resourceContractTest, "port_number.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.1"),
		searchInObject(resourceContractTest, "test_score.valid.0"),
		searchInObject(resourceContractTest, "string_in_some_names.valid.0"),
		searchInObject(resourceContractTest, "valid_cidr.valid.0"),
		searchInObject(resourceContractTest, "percentage.valid.0"))
	return resource
}

func CreateAccContractConfig() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
		resource  "aci_contract" "test" {

							name = "%v"

							temp = %v

							weight = %v

							temp_schema_list = ["%v","%v"]

							filter {
    
							 

						          description = "%v"

						          filter_entry {
							
						              id_list = ["%v","%v"]
 

						              filter_entry_name = "%v"

						            }

							}
		}
	`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.1"),
		searchInObject(resourceContractTest, "filter.description.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.1"),
		searchInObject(resourceContractTest, "filter.filter_entry.filter_entry_name.valid.0"))
	return resource
}

func CreateAccContractConfigWithOptional() string {
	resource := createContractConfig(getParentContract())
	return resource
}

func generateStepForUpdatedRequiredAttr(resourceName string, contract_default, contract_updated *models.Contract) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var value interface{}
	value = searchInObject(resourceContractTest, "name.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccContractUpdateRequiredName(),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAciContractExists(resourceName, contract_updated),
			resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%v", value)),
			testAccCheckAciContractIdNotEqual(contract_default, contract_updated),
		),
	})
	value = searchInObject(resourceContractTest, "temp.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccContractUpdateRequiredTemp(),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAciContractExists(resourceName, contract_updated),
			resource.TestCheckResourceAttr(resourceName, "temp", fmt.Sprintf("%v", value)),
			testAccCheckAciContractIdNotEqual(contract_default, contract_updated),
		),
	})
	value = searchInObject(resourceContractTest, "weight.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccContractUpdateRequiredWeight(),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAciContractExists(resourceName, contract_updated),
			resource.TestCheckResourceAttr(resourceName, "weight", fmt.Sprintf("%v", value)),
			testAccCheckAciContractIdNotEqual(contract_default, contract_updated),
		),
	})
	value = searchInObject(resourceContractTest, "temp_schema_list.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccContractUpdateRequiredTempSchemaList(),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAciContractExists(resourceName, contract_updated),
			resource.TestCheckResourceAttr(resourceName, "temp_schema_list.0", fmt.Sprintf("%v", value)),
			testAccCheckAciContractIdNotEqual(contract_default, contract_updated),
		),
	})
	value = searchInObject(resourceContractTest, "filter.description.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccContractUpdateRequiredFilterDescription(),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAciContractExists(resourceName, contract_updated),
			resource.TestCheckResourceAttr(resourceName, "filter.0.description", fmt.Sprintf("%v", value)),
			testAccCheckAciContractIdNotEqual(contract_default, contract_updated),
		),
	})
	value = searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccContractUpdateRequiredFilterFilterEntryIdList(),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAciContractExists(resourceName, contract_updated),
			resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.id_list.0", fmt.Sprintf("%v", value)),
			testAccCheckAciContractIdNotEqual(contract_default, contract_updated),
		),
	})
	value = searchInObject(resourceContractTest, "filter.filter_entry.filter_entry_name.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccContractUpdateRequiredFilterFilterEntryFilterEntryName(),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAciContractExists(resourceName, contract_updated),
			resource.TestCheckResourceAttr(resourceName, "filter.0.filter_entry.0.filter_entry_name", fmt.Sprintf("%v", value)),
			testAccCheckAciContractIdNotEqual(contract_default, contract_updated),
		),
	})

	return testSteps
}
func CreateAccContractUpdateRequiredName() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	value := searchInObject(resourceContractTest, "name.valid.1")
	resource += fmt.Sprintf(`
			resource "aci_contract" "test" {
							
							name = "%v"

							temp = %v

							weight = %v

							ipv4_for = "%v"

							port_number = %v

							temp_schema_list = ["%v", "%v"]

							test_score = %v

							string_in_some_names = "%v"

							valid_cidr = "%v"

							percentage = %v

							filter {
    
							                        
                            filter_name = %v
                        
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
		`, value,
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
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
func CreateAccContractUpdateRequiredTemp() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	value := searchInObject(resourceContractTest, "temp.valid.1")
	resource += fmt.Sprintf(`
			resource "aci_contract" "test" {

							name = "%v"
							
							temp = %v

							weight = %v

							ipv4_for = "%v"

							port_number = %v

							temp_schema_list = ["%v", "%v"]

							test_score = %v

							string_in_some_names = "%v"

							valid_cidr = "%v"

							percentage = %v

							filter {
    
							                        
                            filter_name = %v
                        
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
		`, searchInObject(resourceContractTest, "name.valid.0"),
		value,
		searchInObject(resourceContractTest, "weight.valid.0"),
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
func CreateAccContractUpdateRequiredWeight() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	value := searchInObject(resourceContractTest, "weight.valid.1")
	resource += fmt.Sprintf(`
			resource "aci_contract" "test" {

							name = "%v"

							temp = %v
							
							weight = %v

							ipv4_for = "%v"

							port_number = %v

							temp_schema_list = ["%v", "%v"]

							test_score = %v

							string_in_some_names = "%v"

							valid_cidr = "%v"

							percentage = %v

							filter {
    
							                        
                            filter_name = %v
                        
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
		`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		value,
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
func CreateAccContractUpdateRequiredTempSchemaList() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	value := searchInObject(resourceContractTest, "temp_schema_list.valid.1")
	resource += fmt.Sprintf(`
			resource "aci_contract" "test" {

							name = "%v"

							temp = %v

							weight = %v

							ipv4_for = "%v"

							port_number = %v

							temp_schema_list = ["%v"]

							test_score = %v

							string_in_some_names = "%v"

							valid_cidr = "%v"

							percentage = %v

							filter {
    
							                        
                            filter_name = %v
                        
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
		`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
		searchInObject(resourceContractTest, "ipv4_for.valid.0"),
		searchInObject(resourceContractTest, "port_number.valid.0"),
		value,
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
func CreateAccContractUpdateRequiredFilterDescription() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	value := searchInObject(resourceContractTest, "filter.description.valid.1")
	resource += fmt.Sprintf(`
			resource "aci_contract" "test" {

							name = "%v"

							temp = %v

							weight = %v

							ipv4_for = "%v"

							port_number = %v

							temp_schema_list = ["%v", "%v"]

							test_score = %v

							string_in_some_names = "%v"

							valid_cidr = "%v"

							percentage = %v

							filter {
    
							                        
                            filter_name = %v
                        
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
		`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
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
		value,
		searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.1"),
		searchInObject(resourceContractTest, "filter.filter_entry.filter_entry_name.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.ipv6.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid.0"),
		searchInObject(resourceContractTest, "filter.filter_entry.apply_to_frag_liist_schema.valid.1"))
	return resource
}
func CreateAccContractUpdateRequiredFilterFilterEntryIdList() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	value := searchInObject(resourceContractTest, "filter.filter_entry.id_list.valid.1")
	resource += fmt.Sprintf(`
			resource "aci_contract" "test" {

							name = "%v"

							temp = %v

							weight = %v

							ipv4_for = "%v"

							port_number = %v

							temp_schema_list = ["%v", "%v"]

							test_score = %v

							string_in_some_names = "%v"

							valid_cidr = "%v"

							percentage = %v

							filter {
    
							                        
                            filter_name = %v
                        
                            id = "%v"
                        
                            description = "%v"

						    filter_entry {
						
						        id_list = ["%v"]
  
                        
                                filter_entry_name = "%v"
                        
                                ipv6 = "%v"
                        
                                apply_to_frag = "%v"

						        apply_to_frag_liist_schema = ["%v", "%v"]
  

						      }

							}
			}
		`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
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
		value,
		searchInObject(resourceContractTest, "filter_entry.filter_entry_name.valid.0"),
		searchInObject(resourceContractTest, "filter_entry.ipv6.valid.0"),
		searchInObject(resourceContractTest, "filter_entry.apply_to_frag.valid.0"),
		searchInObject(resourceContractTest, "apply_to_frag_liist_schema.valid.0"),
		searchInObject(resourceContractTest, "apply_to_frag_liist_schema.valid.1"))
	return resource
}
func CreateAccContractUpdateRequiredFilterFilterEntryFilterEntryName() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	value := searchInObject(resourceContractTest, "filter.filter_entry.filter_entry_name.valid.1")
	resource += fmt.Sprintf(`
			resource "aci_contract" "test" {

							name = "%v"

							temp = %v

							weight = %v

							ipv4_for = "%v"

							port_number = %v

							temp_schema_list = ["%v", "%v"]

							test_score = %v

							string_in_some_names = "%v"

							valid_cidr = "%v"

							percentage = %v

							filter {
    
							                        
                            filter_name = %v
                        
                            id = "%v"
                        
                            description = "%v"

						    filter_entry {
						
						        id_list = ["%v", "%v"]
  
					
						        filter_entry_name = "%v"
                        
                                ipv6 = "%v"
                        
                                apply_to_frag = "%v"

						        apply_to_frag_liist_schema = ["%v", "%v"]
  

						      }

							}
			}
		`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
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
		searchInObject(resourceContractTest, "id_list.valid.0"),
		searchInObject(resourceContractTest, "id_list.valid.1"),
		value,
		searchInObject(resourceContractTest, "filter_entry.ipv6.valid.0"),
		searchInObject(resourceContractTest, "filter_entry.apply_to_frag.valid.0"),
		searchInObject(resourceContractTest, "apply_to_frag_liist_schema.valid.0"),
		searchInObject(resourceContractTest, "apply_to_frag_liist_schema.valid.1"))
	return resource
}

func CreateAccContractUpdatedAttrIpv4For(value interface{}) string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "aci_contract" "test" {

							name = "%v"

							temp = %v

							weight = %v
							
							ipv4_for = "%v"

							port_number = %v

							temp_schema_list = ["%v", "%v"]

							test_score = %v

							string_in_some_names = "%v"

							valid_cidr = "%v"

							percentage = %v

							filter {
    
							                        
                            filter_name = %v
                        
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
	`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
		value,
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
func CreateAccContractUpdatedAttrPortNumber(value interface{}) string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "aci_contract" "test" {

							name = "%v"

							temp = %v

							weight = %v

							ipv4_for = "%v"
							
							port_number = %v

							temp_schema_list = ["%v", "%v"]

							test_score = %v

							string_in_some_names = "%v"

							valid_cidr = "%v"

							percentage = %v

							filter {
    
							                        
                            filter_name = %v
                        
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
	`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
		searchInObject(resourceContractTest, "ipv4_for.valid.0"),
		value,
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
func CreateAccContractUpdatedAttrTestScore(value interface{}) string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "aci_contract" "test" {

							name = "%v"

							temp = %v

							weight = %v

							ipv4_for = "%v"

							port_number = %v

							temp_schema_list = ["%v", "%v"]
							
							test_score = %v

							string_in_some_names = "%v"

							valid_cidr = "%v"

							percentage = %v

							filter {
    
							                        
                            filter_name = %v
                        
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
	`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
		searchInObject(resourceContractTest, "ipv4_for.valid.0"),
		searchInObject(resourceContractTest, "port_number.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.1"),
		value,
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
func CreateAccContractUpdatedAttrStringInSomeNames(value interface{}) string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "aci_contract" "test" {

							name = "%v"

							temp = %v

							weight = %v

							ipv4_for = "%v"

							port_number = %v

							temp_schema_list = ["%v", "%v"]

							test_score = %v
							
							string_in_some_names = "%v"

							valid_cidr = "%v"

							percentage = %v

							filter {
    
							                        
                            filter_name = %v
                        
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
	`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
		searchInObject(resourceContractTest, "ipv4_for.valid.0"),
		searchInObject(resourceContractTest, "port_number.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.1"),
		searchInObject(resourceContractTest, "test_score.valid.0"),
		value,
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
func CreateAccContractUpdatedAttrValidCidr(value interface{}) string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "aci_contract" "test" {

							name = "%v"

							temp = %v

							weight = %v

							ipv4_for = "%v"

							port_number = %v

							temp_schema_list = ["%v", "%v"]

							test_score = %v

							string_in_some_names = "%v"
							
							valid_cidr = "%v"

							percentage = %v

							filter {
    
							                        
                            filter_name = %v
                        
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
	`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
		searchInObject(resourceContractTest, "ipv4_for.valid.0"),
		searchInObject(resourceContractTest, "port_number.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.1"),
		searchInObject(resourceContractTest, "test_score.valid.0"),
		searchInObject(resourceContractTest, "string_in_some_names.valid.0"),
		value,
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
func CreateAccContractUpdatedAttrPercentage(value interface{}) string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "aci_contract" "test" {

							name = "%v"

							temp = %v

							weight = %v

							ipv4_for = "%v"

							port_number = %v

							temp_schema_list = ["%v", "%v"]

							test_score = %v

							string_in_some_names = "%v"

							valid_cidr = "%v"
							
							percentage = %v

							filter {
    
							                        
                            filter_name = %v
                        
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
	`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
		searchInObject(resourceContractTest, "ipv4_for.valid.0"),
		searchInObject(resourceContractTest, "port_number.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.0"),
		searchInObject(resourceContractTest, "temp_schema_list.valid.1"),
		searchInObject(resourceContractTest, "test_score.valid.0"),
		searchInObject(resourceContractTest, "string_in_some_names.valid.0"),
		searchInObject(resourceContractTest, "valid_cidr.valid.0"),
		value,
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

func generateStepForUpdatedAttr(resourceName string, contract_default, contract_updated *models.Contract) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var valid []interface{}
	valid = searchInObject(resourceContractTest, "ipv4_for.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccContractUpdatedAttrIpv4For(value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAciContractExists(resourceName, contract_updated),
				resource.TestCheckResourceAttr(resourceName, "ipv4_for", v),
				testAccCheckAciContractIdEqual(contract_default, contract_updated),
			),
		})
	}
	valid = searchInObject(resourceContractTest, "port_number.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccContractUpdatedAttrPortNumber(value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAciContractExists(resourceName, contract_updated),
				resource.TestCheckResourceAttr(resourceName, "port_number", v),
				testAccCheckAciContractIdEqual(contract_default, contract_updated),
			),
		})
	}
	valid = searchInObject(resourceContractTest, "test_score.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccContractUpdatedAttrTestScore(value),
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
			Config: CreateAccContractUpdatedAttrStringInSomeNames(value),
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
			Config: CreateAccContractUpdatedAttrValidCidr(value),
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
			Config: CreateAccContractUpdatedAttrPercentage(value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAciContractExists(resourceName, contract_updated),
				resource.TestCheckResourceAttr(resourceName, "percentage", v),
				testAccCheckAciContractIdEqual(contract_default, contract_updated),
			),
		})
	}
	return testSteps
}

func generateNegativeSteps(resourceName string) []resource.TestStep {
	//Use Update Config Function with false value
	testSteps := make([]resource.TestStep, 0, 1)
	var invalid []interface{}
	invalid = searchInObject(resourceContractTest, "ipv4_for.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccContractUpdatedAttrIpv4For(value),
			ExpectError: regexp.MustCompile(expectErrorMap["IsIPv4Address"]),
		})
	}
	invalid = searchInObject(resourceContractTest, "port_number.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccContractUpdatedAttrPortNumber(value),
			ExpectError: regexp.MustCompile(expectErrorMap["IsPortNumber"]),
		})
	}
	invalid = searchInObject(resourceContractTest, "test_score.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccContractUpdatedAttrTestScore(value),
			ExpectError: regexp.MustCompile(expectErrorMap["IntBetween"]),
		})
	}
	invalid = searchInObject(resourceContractTest, "string_in_some_names.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccContractUpdatedAttrStringInSomeNames(value),
			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
		})
	}
	invalid = searchInObject(resourceContractTest, "valid_cidr.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccContractUpdatedAttrValidCidr(value),
			ExpectError: regexp.MustCompile(expectErrorMap["IsCIDRNetwork"]),
		})
	}
	invalid = searchInObject(resourceContractTest, "percentage.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccContractUpdatedAttrPercentage(value),
			ExpectError: regexp.MustCompile(expectErrorMap["FloatBetween"]),
		})
	}
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccMovieConfig(),
	})
	return testSteps
}

func CreateAccContractMultipleConfig() string {
	var resource string
	parentResources := getParentContract()
	parentResources = parentResources[:len(parentResources)-1]
	resource += createContractConfig(parentResources)
	multipleValues := searchInObject(resourceContractTest, "name.multiple_valids").([]interface{})
	for i, val := range multipleValues {
		resource += fmt.Sprintf(`
			resource "aci_contract" "test%d" {
							
							name = "%v"

							temp = %v

							weight = %v

							ipv4_for = "%v"

							port_number = %v

							temp_schema_list = ["%v", "%v"]

							test_score = %v

							string_in_some_names = "%v"

							valid_cidr = "%v"

							percentage = %v

							filter {
    
							                        
                            filter_name = %v
                        
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
		`, i, val,
			searchInObject(resourceContractTest, "temp.valid.0"),
			searchInObject(resourceContractTest, "weight.valid.0"),
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
	return resource
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

func testAccCheckAciContractIdNotEqual(contract1, contract2 *models.Contract) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		Id1, err := getIdFromContractModel(contract1)
		if err != nil {
			return err
		}
		Id2, err := getIdFromContractModel(contract2)
		if err != nil {
			return err
		}
		if Id1 == Id2 {
			return fmt.Errorf("Contract IDs are equal")
		}
		return nil
	}
}

func getParentContract() []string {
	t := []string{}
	t = append(t, contractBlock())
	return t
}

func contractBlock() string {
	return fmt.Sprintf(`
		resource  "aci_contract" "test" {

						name = "%v"

						temp = %v

						weight = %v

						ipv4_for = "%v"

						port_number = %v

				        temp_schema_list = ["%v","%v"]

						test_score = %v

						string_in_some_names = "%v"

						valid_cidr = "%v"

						percentage = %v

                        filter {
    
                                                
                            filter_name = %v
                        
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
	`, searchInObject(resourceContractTest, "name.valid.0"),
		searchInObject(resourceContractTest, "temp.valid.0"),
		searchInObject(resourceContractTest, "weight.valid.0"),
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
