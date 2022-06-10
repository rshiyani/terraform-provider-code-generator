package aci

import (
	"fmt"
	"regexp"
	"testing"
)

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
				Config:      CreateAccContractWithoutName(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutIpv4(rName),
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
				Config:      CreateAccContractWithoutCidr(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutTime(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutUrlHttps(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutUrlHttp(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutUuid(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutBase64(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutJson(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutRegExp(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutPortNumber(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutPortWithZero(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutNuclearCode(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutTestScore(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccContractWithoutPercentage(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
		},
	})
}

func CreateAccContractWithoutName(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutIpv4(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutIpv6(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutMac(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutCidr(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutTime(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutUrlHttps(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutUrlHttp(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutUuid(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutBase64(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutJson(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutRegExp(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			port_number = 1
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutPortNumber(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutPortWithZero(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutNuclearCode(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			port_with_zero = 0
			test_score = 1
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutTestScore(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			percentage = 0
		}
	`)
}
func CreateAccContractWithoutPercentage(rName string) string {
	return fmt.Sprintf(`
		resource "aci_contract" "test" {
			name = "ufl4ow8f3q"
			ipv4 = "71.230.128.82"
			ipv6 = "fb2d:2b35:b822:be17:b361:6691"
			mac = "52:af:5f:a6:77:24"
			cidr = "198.45.35.129/9"
			time = "2022-06-10T05:05:34.280848+00:00"
			url_https = "https://er388ujn1sa2dn2.com"
			url_http = "http://clb80ylsgjrrf6d.com"
			uuid = "f56c3da5-e87a-11ec-a5f9-7c8ae19799de"
			base_64 = "b2tjYThwc25pcg=="
			json = '{ "attribute" : "value" }'
			reg_exp = "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b"
			port_number = 1
			port_with_zero = 0
			nuclear_code = "ufl4ow8f3q"
			test_score = 1
		}
	`)
}
