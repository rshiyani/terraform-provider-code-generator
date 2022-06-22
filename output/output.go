package aci

import (
	"fmt"
	"regexp"
	"testing"
)

const contractSelfRequiredCount = 2

var resourceContractTest = map[string]interface{}{
	"temp": map[string]interface{}{
		"valid":   resourceApplicationProfileTest["application_dn"].(map[string]interface{})["valid"].([]interface{}),
		"invalid": resourceApplicationProfileTest["application_dn"].(map[string]interface{})["invalid"].([]interface{}),
	},
	"weight": map[string]interface{}{
		"valid":   []interface{}{23.4, 987},
		"invalid": []interface{}{"Hello", "World"},
	},

	"ipv4_for": map[string]interface{}{
		"valid":   Test["ipv4"].(map[string]interface{})["valid"].([]interface{}),
		"invalid": Test["ipv4"].(map[string]interface{})["invalid"].([]interface{}),
	},
	"port_number": map[string]interface{}{
		"valid":   []interface{}{1, 53, 65535},
		"invalid": []interface{}{0, 65536},
	},
	"test_score": map[string]interface{}{
		"valid":   []interface{}{1, 100, 50},
		"invalid": []interface{}{0, 101},
	},
	"string_in_some_names": map[string]interface{}{
		"valid":   []interface{}{"parth", "aarsh", "arjun", "alfatah", "krunal"},
		"invalid": []interface{}{"6xxqoku5qi"},
	},
	"valid_cidr": map[string]interface{}{
		"valid":   []interface{}{0, 32, 16},
		"invalid": []interface{}{-1, 33},
	},
}

func TestAccAciContract_Basic(t *testing.T) {
	var contract_default models.Contract
	var contract_updated models.Contract
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
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%v",
						func() string {
							id, err := getIdFromContractModel(&contract_default)
							if err != nil {
								return ""
							}
							return id
						}())),

					resource.TestCheckResourceAttr(resourceName, "temp", fmt.Sprintf("%v", resourceContractTest["temp"].(map[string]interface{})["valid"].([]interface{})[0])),

					resource.TestCheckResourceAttr(resourceName, "weight", fmt.Sprintf("%v", resourceContractTest["weight"].(map[string]interface{})["valid"].([]interface{})[0])),

					resource.TestCheckResourceAttr(resourceName, "ipv4_for", fmt.Sprintf("%v", resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]interface{})[0])),

					resource.TestCheckResourceAttr(resourceName, "port_number", "0"),
					resource.TestCheckResourceAttr(resourceName, "test_score", "0"),
					resource.TestCheckResourceAttr(resourceName, "string_in_some_names", "parth"),
					resource.TestCheckResourceAttr(resourceName, "valid_cidr", ""),
				),
			},
			{
				Config: CreateAccContractConfigWithOptional(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractExists(resourceName, &contract_updated),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%v",
						func() string {
							id, err := getIdFromContractModel(&contract_updated)
							if err != nil {
								return ""
							}
							return id
						}())), // Function to get ID based on the model provided
					resource.TestCheckResourceAttr(resourceName, "temp", fmt.Sprintf("%v", resourceContractTest["temp"].(map[string]interface{})["valid"].([]interface{})[0])),
					resource.TestCheckResourceAttr(resourceName, "weight", fmt.Sprintf("%v", resourceContractTest["weight"].(map[string]interface{})["valid"].([]interface{})[0])),
					resource.TestCheckResourceAttr(resourceName, "ipv4_for", fmt.Sprintf("%v", resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]interface{})[0])),
					resource.TestCheckResourceAttr(resourceName, "port_number", fmt.Sprintf("%v", resourceContractTest["port_number"].(map[string]interface{})["valid"].([]interface{})[0])),
					resource.TestCheckResourceAttr(resourceName, "test_score", fmt.Sprintf("%v", resourceContractTest["test_score"].(map[string]interface{})["valid"].([]interface{})[0])),
					resource.TestCheckResourceAttr(resourceName, "string_in_some_names", fmt.Sprintf("%v", resourceContractTest["string_in_some_names"].(map[string]interface{})["valid"].([]interface{})[0])),
					resource.TestCheckResourceAttr(resourceName, "valid_cidr", fmt.Sprintf("%v", resourceContractTest["valid_cidr"].(map[string]interface{})["valid"].([]interface{})[0])),
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
			{
				Config: CreateAccContractUpdateParentRequiredArgumentName(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractExists(resourceName, &contract_updated),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%v",
						func() string {
							id, err := getIdFromContractModel(&contract_updated)
							if err != nil {
								return ""
							}
							return id
						}())), // Function to get ID based on the model provided
					func(model1, model2 *models.Contract) resource.TestCheckFunc {
						// Check if Tenant have some independent required field
						if tenantSelfRequiredCount > 0 {
							return testAccCheckAciContractIdNotEqual(model1, model2)
						}
						return testAccCheckAciContractIdEqual(model1, model2)
					}(contract_default, contract_updated),
				),
			},
			{
				Config: CreateAccContractUpdateParentRequiredArgumentTemp(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractExists(resourceName, &contract_updated),
					resource.TestCheckResourceAttr(resourceName, "temp", fmt.Sprintf("%v", resourceContractTest["temp"].(map[string]interface{})["valid"].([]interface{})[1])),
					func(model1, model2 *models.Contract) resource.TestCheckFunc {
						// Check if ApplicationProfile have some independent required field
						if applicationProfileSelfRequiredCount > 0 {
							return testAccCheckAciContractIdNotEqual(model1, model2)
						}
						return testAccCheckAciContractIdEqual(model1, model2)
					}(contract_default, contract_updated),
				),
			},
			{
				Config: CreateAccContractUpdateRequiredArgumentWeight(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractExists(resourceName, &contract_updated),
					resource.TestCheckResourceAttr(resourceName, "weight", fmt.Sprintf("%v", resourceContractTest["weight"].(map[string]interface{})["valid"].([]interface{})[1])),
					testAccCheckAciContractIdNotEqual(&contract_default, &contract_updated),
				),
			},
			{
				Config: CreateAccContractUpdateRequiredArgumentIpv4For(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractExists(resourceName, &contract_updated),
					resource.TestCheckResourceAttr(resourceName, "ipv4_for", fmt.Sprintf("%v", resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]interface{})[1])),
					testAccCheckAciContractIdNotEqual(&contract_default, &contract_updated),
				),
			},
		},
	})
}

// Returns the []TestSteps consisiting of Updation of optional attributes
func generateStepForUpdatedAttr(resourceName string, contract_default, contract_updated *models.Contract) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var valid []interface{}
	valid = resourceContractTest["port_number"].(map[string]interface{})["valid"].([]interface{})
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
	valid = resourceContractTest["test_score"].(map[string]interface{})["valid"].([]interface{})
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
	valid = resourceContractTest["string_in_some_names"].(map[string]interface{})["valid"].([]interface{})
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
	valid = resourceContractTest["valid_cidr"].(map[string]interface{})["valid"].([]interface{})
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
	return testSteps
}

func TestAccAciContract_Update(t *testing.T) {
	var contract_default models.Contract
	var contract_updated models.Contract
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
				Check:  testAccCheckAciContractExists(resourceName, &contract_default),
			},
		}, generateStepForUpdatedAttr(resourceName, &contract_default, &contract_updated)...),
	})
}

func CreateAccContractWithoutName() string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					temp = "%v"
					weight = "%v"
					ipv4_for = "%v"
		}
	`, resourceContractTest["temp"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["weight"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]interface{})[0])
}
func CreateAccContractWithoutTemp() string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					name = "%v"
					weight = "%v"
					ipv4_for = "%v"
		}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["weight"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]interface{})[0])
}
func CreateAccContractWithoutWeight() string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					name = "%v"
					temp = "%v"
					ipv4_for = "%v"
		}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["temp"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]interface{})[0])
}
func CreateAccContractWithoutIpv4For() string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					name = "%v"
					temp = "%v"
					weight = "%v"
		}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["temp"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["weight"].(map[string]interface{})["valid"].([]interface{})[0])
}

func CreateAccContractConfig() string {
	resource := createContractConfig(getParentContract())
	return resource
}

func CreateAccContractConfigWithOptional() string {
	var resource string
	resource = CreateAccTenantConfig()
	resource = CreateAccApplicationProfileConfig()
	resource += fmt.Sprintf(`
		resource  "aci_contract" "test" {
						name = aci_tenant.test.id
						temp = aci_application_profile.test.application_dn
						weight = "%v"
						ipv4_for = "%v"
						port_number = "%v"
						test_score = "%v"
						string_in_some_names = "%v"
						valid_cidr = "%v"
		}
	`, resourceContractTest["weight"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["string_in_some_names"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["valid_cidr"].(map[string]interface{})["valid"].([]interface{})[0])
	return resource
}

func CreateAccContractUpdatedAttr(attr string, value interface{}) string {
	var resource string
	resource = CreateAccTenantConfig()
	resource = CreateAccApplicationProfileConfig()
	resource += fmt.Sprintf(`
		resource  "aci_contract" "test" {
						name = aci_tenant.test.id
						temp = aci_application_profile.test.application_dn
						weight = "%v"
						ipv4_for = "%v"
						%v = "%v"
		}
	`, resourceContractTest["weight"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]interface{})[0], attr, value)
	return resource
}

func CreateAccContractUpdateRequiredArgumentWeight() string {
	t := []string{}
	t = append(t, getParentTenant()...)
	t = append(t, getParentApplicationProfile()...)
	t = append(t, fmt.Sprintf(`
					resource  "aci_contract" "test" {
									name = aci_tenant.test.id
									temp = aci_application_profile.test.application_dn
										weight = "%v"
										
										ipv4_for = "%v"
										
					}
				`, resourceContractTest["weight"].(map[string]interface{})["valid"].([]interface{})[1],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]interface{})[0]))
	resource := createContractConfig(t)
	return resource
}
func CreateAccContractUpdateRequiredArgumentIpv4For() string {
	t := []string{}
	t = append(t, getParentTenant()...)
	t = append(t, getParentApplicationProfile()...)
	t = append(t, fmt.Sprintf(`
					resource  "aci_contract" "test" {
									name = aci_tenant.test.id
									temp = aci_application_profile.test.application_dn
										weight = "%v"
										
										ipv4_for = "%v"
										
					}
				`, resourceContractTest["weight"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]interface{})[1]))
	resource := createContractConfig(t)
	return resource
}

func CreateAccContractUpdateParentRequiredArgumentName() string {
	t := []string{}
	t = append(t, getUpdatedParentTenant()...)
	t = append(t, getParentApplicationProfile()...)
	t = append(t, fmt.Sprintf(`
					resource  "aci_contract" "test" {
									name = aci_tenant.test.id
									temp = aci_application_profile.test.application_dn
										weight = "%v"
										
										ipv4_for = "%v"
										
					}
				`, resourceContractTest["weight"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]interface{})[0]))
	resource := createContractConfig(t)
	return resource
}
func CreateAccContractUpdateParentRequiredArgumentTemp() string {
	t := []string{}
	t = append(t, getParentTenant()...)
	t = append(t, getUpdatedParentApplicationProfile()...)
	t = append(t, fmt.Sprintf(`
					resource  "aci_contract" "test" {
									name = aci_tenant.test.id
									temp = aci_application_profile.test.application_dn
										weight = "%v"
										
										ipv4_for = "%v"
										
					}
				`, resourceContractTest["weight"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]interface{})[0]))
	resource := createContractConfig(t)
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

func getUpdatedParentContract() []string {
	t := []string{}
	t = append(t, getParentTenant()...)
	t = append(t, getParentApplicationProfile()...)
	t = append(t, contractUpdatedParentBlock())
	return t
}

func contractUpdatedParentBlock() string {
	return fmt.Sprintf(`
	resource  "aci_contract" "test" {
						name = aci_tenant.test.id
						temp = aci_application_profile.test.application_dn
						weight = "%v"
							
						ipv4_for = "%v"
							
	}
`, resourceContractTest["weight"].(map[string]interface{})["valid"].([]interface{})[1],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]interface{})[0])
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
					name = aci_tenant.test.id
					temp = aci_application_profile.test.application_dn
					weight = "%v"
					
					ipv4_for = "%v"
					
	}
`, resourceContractTest["weight"].(map[string]interface{})["valid"].([]interface{})[0],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]interface{})[0])
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
