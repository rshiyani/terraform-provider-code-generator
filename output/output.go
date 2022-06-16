package aci

import (
	"fmt"
	"regexp"
	"testing"
)

var resourceContractTest = map[string]interface{}{
	"name": map[string]interface{}{
		"valid":   []interface{}{"Hello", "World"},
		"invalid": []interface{}{234, 987},
	},
	"id": map[string]interface{}{
		"valid":   []interface{}{234, 987},
		"invalid": []interface{}{"Hello", "World"},
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
		"invalid": []interface{}{"j3d9d0yum6"},
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
	rName := makeTestVariable(acctest.RandString(5))
	rOther := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAciContractDestroy,
		Steps: []resource.TestStep{
			{
				Config:      CreateAccContractWithoutName(),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutId(),
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
					resource.TestCheckResourceAttr(resourceName, "name", resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0]),

					resource.TestCheckResourceAttr(resourceName, "id", resourceContractTest["id"].(map[string]interface{})["valid"].([]int)[0]),

					resource.TestCheckResourceAttr(resourceName, "weight", resourceContractTest["weight"].(map[string]interface{})["valid"].([]float64)[0]),

					resource.TestCheckResourceAttr(resourceName, "ipv4_for", resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]string)[0]),

					resource.TestCheckResourceAttr(resourceName, "port_number", 0),

					resource.TestCheckResourceAttr(resourceName, "test_score", 0),

					resource.TestCheckResourceAttr(resourceName, "string_in_some_names", "parth"),

					resource.TestCheckResourceAttr(resourceName, "valid_cidr", ""),
				),
			},
			{
				Config: CreateAccContractConfigWithOptional(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractExists(resourceName, &contract_updated),
					resource.TestCheckResourceAttr(resourceName, "name", resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0]),
					resource.TestCheckResourceAttr(resourceName, "id", resourceContractTest["id"].(map[string]interface{})["valid"].([]int)[0]),
					resource.TestCheckResourceAttr(resourceName, "weight", resourceContractTest["weight"].(map[string]interface{})["valid"].([]float64)[0]),
					resource.TestCheckResourceAttr(resourceName, "ipv4_for", resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]string)[0]),
					resource.TestCheckResourceAttr(resourceName, "port_number", resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0]),
					resource.TestCheckResourceAttr(resourceName, "test_score", resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0]),
					resource.TestCheckResourceAttr(resourceName, "string_in_some_names", resourceContractTest["string_in_some_names"].(map[string]interface{})["valid"].([]string)[0]),
					resource.TestCheckResourceAttr(resourceName, "valid_cidr", resourceContractTest["valid_cidr"].(map[string]interface{})["valid"].([]string)[0]),
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

func generateStepForUpdatedAttr(attr string) string {
	valid := resourceContractTest[attr].(map[string]interface{})["valid"].([]interface{})
	str := ""
	for _, value := range valid {
		str += fmt.Sprintf(`{
			Config: CreateAccContractUpdatedAttr(),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAciContractExists(resourceName, &contract_updated),
				resource.TestCheckResourceAttr(resourceName, "%v", "%v"),
				testAccCheckAciContractdEqual(&contract_default, &contract_updated),
			), 
		},
		`, attr, value)
	}
	return str
}

func TestAccAciContract_Update(t *testing.T) {
	var contract_default models.Contract
	var contract_updated models.Contract
	resourceName := "aci_contract.test"

	// [TODO]: Add makeTestVariable() to utils.go file
	rName := makeTestVariable(acctest.RandString(5))
	rOther := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAciContractDestroy,
		Steps: []resource.TestStep{
			{
				Config: CreateAccContractConfig(),
				Check:  testAccCheckAciContractExists(resourceName, &contract_default),
			},
			generateStepForUpdatedAttr("port_number"),
			generateStepForUpdatedAttr("test_score"),
			generateStepForUpdatedAttr("string_in_some_names"),
			generateStepForUpdatedAttr("valid_cidr"),
		},
	})
}

func CreateAccContractWithoutName() string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					id = "%v"
					weight = "%v"
					ipv4_for = "%v"
		}
	`, resourceContractTest["id"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["weight"].(map[string]interface{})["valid"].([]float64)[0],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]string)[0])
}
func CreateAccContractWithoutId() string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					name = "%v"
					weight = "%v"
					ipv4_for = "%v"
		}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["weight"].(map[string]interface{})["valid"].([]float64)[0],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]string)[0])
}
func CreateAccContractWithoutWeight() string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					name = "%v"
					id = "%v"
					ipv4_for = "%v"
		}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["id"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]string)[0])
}
func CreateAccContractWithoutIpv4For() string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					name = "%v"
					id = "%v"
					weight = "%v"
		}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["id"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["weight"].(map[string]interface{})["valid"].([]float64)[0])
}

func CreateAccContractConfig() string {
	var resource string
	resource += fmt.Sprintf(`
		resource  "aci_contract" "test" {
						name = "%v"
						id = "%v"
						weight = "%v"
						ipv4_for = "%v"
		}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["id"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["weight"].(map[string]interface{})["valid"].([]float64)[0],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]string)[0])
	return resource
}

func CreateAccContractConfigWithOptional() string {
	var resource string
	resource += fmt.Sprintf(`
		resource  "aci_contract" "test" {
						name = "%v"
						id = "%v"
						weight = "%v"
						ipv4_for = "%v"
						port_number = "%v"
						test_score = "%v"
						string_in_some_names = "%v"
						valid_cidr = "%v"
		}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["id"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["weight"].(map[string]interface{})["valid"].([]float64)[0],
		resourceContractTest["ipv4_for"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["string_in_some_names"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["valid_cidr"].(map[string]interface{})["valid"].([]string)[0])
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
