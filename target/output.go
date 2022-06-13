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
	"is_good_student": map[string]interface{}{
		"valid":   []bool{true, false},
		"invalid": []interface{}{235, "World"},
	},
	"ipv4": map[string]interface{}{
		"valid":   []string{"101.235.83.204"},
		"invalid": []interface{}{"296.281.285.274"},
	},
	"ipv6": map[string]interface{}{
		"valid":   []string{"1cc9:d6b9:7b99:4f83:eb3b:d2aa"},
		"invalid": []interface{}{"invalidIPv6"},
	},
	"mac": map[string]interface{}{
		"valid":   []string{"d9:88:f1:b2:2c:66"},
		"invalid": []interface{}{"invalidMAC"},
	},
	"cidr": map[string]interface{}{
		"valid":   []string{"92.170.124.56/7"},
		"invalid": []interface{}{"280.283.272.286/11"},
	},
	"time": map[string]interface{}{
		"valid":   []string{"2022-06-13T12:51:15.953419+00:00"},
		"invalid": []interface{}{"2022-06-13 18:21:15.953419"},
	},
	"url_https": map[string]interface{}{
		"valid":   []string{"https://vue6ep1bbldkfa9.com"},
		"invalid": []interface{}{"hts:/wmw6xh4gq4xevbc.com"},
	},
	"url_http": map[string]interface{}{
		"valid":   []string{"http://bxgacoq7epli06c.com"},
		"invalid": []interface{}{"ht:/5f6kjcinv9vqo6c.com"},
	},
	"uuid": map[string]interface{}{
		"valid":   []string{"83325e73-eb17-11ec-9b4b-7c8ae1943087"},
		"invalid": []interface{}{"invalid323Uuid12"},
	},
	"base_64": map[string]interface{}{
		"valid":   []string{"OGc2Nm9rOW54eQ=="},
		"invalid": []interface{}{"a3+J1b%mFs//"},
	},
	"json": map[string]interface{}{
		"valid":   []string{"json({ \"attribute\" : \"value\" })"},
		"invalid": []interface{}{"json({ name : val)"},
	},
	"reg_exp": map[string]interface{}{
		"valid":   []string{"(?m)^[0-9]{2}$"},
		"invalid": []interface{}{"[0-9)++"},
	},
	"gender": map[string]interface{}{
		"valid":   []string{"701vogs508"},
		"invalid": []interface{}{12345},
	},
	"port_number": map[string]interface{}{
		"valid":   []int{1, 53, 65535},
		"invalid": []interface{}{0, 65536},
	},
	"port_with_zero": map[string]interface{}{
		"valid":   []int{0, 1, 53, 65535},
		"invalid": []interface{}{-1, 65536},
	},
	"nuclear_code": map[string]interface{}{
		"valid":   []string{"701vogs508"},
		"invalid": []interface{}{12345},
	},
	"test_score": map[string]interface{}{
		"valid":   []int{1, 100, 50},
		"invalid": []interface{}{0, 101},
	},
	"percentage": map[string]interface{}{
		"valid":   []float64{0, 100, 50},
		"invalid": []interface{}{-1, 101},
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
				Config:      CreateAccContractWithoutId(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutWeight(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutIsGoodStudent(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutIpv6(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutMac(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutTime(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutUrlHttp(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutRegExp(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config: CreateAccContractConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractExists(resourceName, &contract_default),
					resource.TestCheckResourceAttr(resourceName, "name", ""),
					resource.TestCheckResourceAttr(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "weight"),
					resource.TestCheckResourceAttr(resourceName, "is_good_student"),
					resource.TestCheckResourceAttr(resourceName, "ipv4", ""),
					resource.TestCheckResourceAttr(resourceName, "ipv6", ""),
					resource.TestCheckResourceAttr(resourceName, "mac", ""),
					resource.TestCheckResourceAttr(resourceName, "cidr", ""),
					resource.TestCheckResourceAttr(resourceName, "time", ""),
					resource.TestCheckResourceAttr(resourceName, "url_https", ""),
					resource.TestCheckResourceAttr(resourceName, "url_http", ""),
					resource.TestCheckResourceAttr(resourceName, "uuid", ""),
					resource.TestCheckResourceAttr(resourceName, "base_64", ""),
					resource.TestCheckResourceAttr(resourceName, "json", ""),
					resource.TestCheckResourceAttr(resourceName, "reg_exp", ""),
					resource.TestCheckResourceAttr(resourceName, "gender", ""),
					resource.TestCheckResourceAttr(resourceName, "port_number"),
					resource.TestCheckResourceAttr(resourceName, "port_with_zero"),
					resource.TestCheckResourceAttr(resourceName, "nuclear_code", ""),
					resource.TestCheckResourceAttr(resourceName, "test_score"),
					resource.TestCheckResourceAttr(resourceName, "percentage"),
				),
			},
		},
	})
}

func CreateAccContractWithoutId(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					weight = "%v"
							
					is_good_student = "%v"
							
					ipv6 = "%v"
							
					mac = "%v"
							
					time = "%v"
							
					url_http = "%v"
							
					reg_exp = "%v"
							
		}
	`, resourceContractTest["weight"].(map[string]interface{})["valid"].([]float64)[0], resourceContractTest["is_good_student"].(map[string]interface{})["valid"].([]bool)[0], resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0])
}
func CreateAccContractWithoutWeight(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					id = "%v"
							
					is_good_student = "%v"
							
					ipv6 = "%v"
							
					mac = "%v"
							
					time = "%v"
							
					url_http = "%v"
							
					reg_exp = "%v"
							
		}
	`, resourceContractTest["id"].(map[string]interface{})["valid"].([]int)[0], resourceContractTest["is_good_student"].(map[string]interface{})["valid"].([]bool)[0], resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0])
}
func CreateAccContractWithoutIsGoodStudent(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					id = "%v"
							
					weight = "%v"
							
					ipv6 = "%v"
							
					mac = "%v"
							
					time = "%v"
							
					url_http = "%v"
							
					reg_exp = "%v"
							
		}
	`, resourceContractTest["id"].(map[string]interface{})["valid"].([]int)[0], resourceContractTest["weight"].(map[string]interface{})["valid"].([]float64)[0], resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0])
}
func CreateAccContractWithoutIpv6(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					id = "%v"
							
					weight = "%v"
							
					is_good_student = "%v"
							
					mac = "%v"
							
					time = "%v"
							
					url_http = "%v"
							
					reg_exp = "%v"
							
		}
	`, resourceContractTest["id"].(map[string]interface{})["valid"].([]int)[0], resourceContractTest["weight"].(map[string]interface{})["valid"].([]float64)[0], resourceContractTest["is_good_student"].(map[string]interface{})["valid"].([]bool)[0], resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0])
}
func CreateAccContractWithoutMac(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					id = "%v"
							
					weight = "%v"
							
					is_good_student = "%v"
							
					ipv6 = "%v"
							
					time = "%v"
							
					url_http = "%v"
							
					reg_exp = "%v"
							
		}
	`, resourceContractTest["id"].(map[string]interface{})["valid"].([]int)[0], resourceContractTest["weight"].(map[string]interface{})["valid"].([]float64)[0], resourceContractTest["is_good_student"].(map[string]interface{})["valid"].([]bool)[0], resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0])
}
func CreateAccContractWithoutTime(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					id = "%v"
							
					weight = "%v"
							
					is_good_student = "%v"
							
					ipv6 = "%v"
							
					mac = "%v"
							
					url_http = "%v"
							
					reg_exp = "%v"
							
		}
	`, resourceContractTest["id"].(map[string]interface{})["valid"].([]int)[0], resourceContractTest["weight"].(map[string]interface{})["valid"].([]float64)[0], resourceContractTest["is_good_student"].(map[string]interface{})["valid"].([]bool)[0], resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0])
}
func CreateAccContractWithoutUrlHttp(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					id = "%v"
							
					weight = "%v"
							
					is_good_student = "%v"
							
					ipv6 = "%v"
							
					mac = "%v"
							
					time = "%v"
							
					reg_exp = "%v"
							
		}
	`, resourceContractTest["id"].(map[string]interface{})["valid"].([]int)[0], resourceContractTest["weight"].(map[string]interface{})["valid"].([]float64)[0], resourceContractTest["is_good_student"].(map[string]interface{})["valid"].([]bool)[0], resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0])
}
func CreateAccContractWithoutRegExp(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
					id = "%v"
							
					weight = "%v"
							
					is_good_student = "%v"
							
					ipv6 = "%v"
							
					mac = "%v"
							
					time = "%v"
							
					url_http = "%v"
							
		}
	`, resourceContractTest["id"].(map[string]interface{})["valid"].([]int)[0], resourceContractTest["weight"].(map[string]interface{})["valid"].([]float64)[0], resourceContractTest["is_good_student"].(map[string]interface{})["valid"].([]bool)[0], resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0], resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0])
}

func CreateAccContractConfig(rName string) string {
	return fmt.Sprintf(`
		resource  "aci_contract" "test" {
					id = %v
					weight = %v
					is_good_student = %v
					ipv6 = "%v"
					mac = "%v"
					time = "%v"
					url_http = "%v"
					reg_exp = "%v"
		}
	`)
}
