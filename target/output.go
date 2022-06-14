package aci

import (
	"fmt"
	"regexp"
	"testing"
)

var resourceContractTest = map[string]interface{}{
	"name": map[string]interface{}{
		"valid":   []string{"Hello", "World"},
		"invalid": []interface{}{234, 987},
	},
	"id": map[string]interface{}{
		"valid":   []int{234, 987},
		"invalid": []interface{}{"Hello", "World"},
	},
	"weight": map[string]interface{}{
		"valid":   []float64{23.4, 987},
		"invalid": []interface{}{"Hello", "World"},
	},
	"ipv4_for": map[string]interface{}{
		"valid":   Test["ipv4"].(map[string]interface{})["valid"].([]string),
		"invalid": Test["ipv4"].(map[string]interface{})["invalid"].([]interface{}),
	},
	"port_number": map[string]interface{}{
		"valid":   []int{1, 53, 65535},
		"invalid": []interface{}{0, 65536},
	},
	"test_score": map[string]interface{}{
		"valid":   []int{1, 100, 50},
		"invalid": []interface{}{0, 101},
	},
	"string_in_some_names": map[string]interface{}{
		"valid":   []string{"parth", "aarsh", "arjun", "alfatah", "krunal"},
		"invalid": []string{"nqa0pvxzrh"},
	},
	"valid_cidr": map[string]interface{}{
		"valid":   []int{0, 32, 16},
		"invalid": []int{-1, 33},
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
		ProviderFactories: testAccProviders,
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
	return fmt.Sprintf(`
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
}

func CreateAccContractConfigWithOptional() string {
	return fmt.Sprintf(`
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
}
