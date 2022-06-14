package acctest

import (
	"fmt"
	"regexp"
	"testing"
)

func TestAccAciTenantDataSource_Basic(t *testing.T) {
	resourceName := "aci_tenant.test"
	dataSourceName := "data.aci_tenant.test"
	rName := makeTestVariable(acctest.RandString(5))
	randomParameter := acctest.RandStringFromCharSet(5, "abcdefghijklmnopqrstuvwxyz")
	randomValue := acctest.RandString(5)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccCheckAciTenantDestroy,
		Steps: []resource.TestStep{
			{
				Config:      CreateAccTenantDSWithoutRequired("name"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccTenantDSWithoutRequired("url_https"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccTenantDSWithoutRequired("url_http"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccTenantDSWithoutRequired("uuid"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccTenantDSWithoutRequired("base_64"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccTenantDSWithoutRequired("json"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccTenantDSWithoutRequired("reg_exp"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccTenantDSWithoutRequired("port_number"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccTenantDSWithoutRequired("port_with_zero"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccTenantDSWithoutRequired("nuclear_code"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccTenantDSWithoutRequired("test_score"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config: CreateAccTenantConfigDataSource(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),

					resource.TestCheckResourceAttrPair(dataSourceName, "annotation", resourceName, "annotation"),

					resource.TestCheckResourceAttrPair(dataSourceName, "name_alias", resourceName, "name_alias"),

					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),

					resource.TestCheckResourceAttrPair(dataSourceName, "url_https", resourceName, "url_https"),

					resource.TestCheckResourceAttrPair(dataSourceName, "url_http", resourceName, "url_http"),

					resource.TestCheckResourceAttrPair(dataSourceName, "uuid", resourceName, "uuid"),

					resource.TestCheckResourceAttrPair(dataSourceName, "base_64", resourceName, "base_64"),

					resource.TestCheckResourceAttrPair(dataSourceName, "json", resourceName, "json"),

					resource.TestCheckResourceAttrPair(dataSourceName, "reg_exp", resourceName, "reg_exp"),

					resource.TestCheckResourceAttrPair(dataSourceName, "gender", resourceName, "gender"),

					resource.TestCheckResourceAttrPair(dataSourceName, "port_number", resourceName, "port_number"),

					resource.TestCheckResourceAttrPair(dataSourceName, "port_with_zero", resourceName, "port_with_zero"),

					resource.TestCheckResourceAttrPair(dataSourceName, "nuclear_code", resourceName, "nuclear_code"),

					resource.TestCheckResourceAttrPair(dataSourceName, "test_score", resourceName, "test_score"),
				),
			},
			{
				Config:      CreateAccTenantDataSourceUpdate(rName, randomParameter, randomValue),
				ExpectError: regexp.MustCompile(`An argument named (.)+ is not expected here.`),
			},

			{
				Config:      CreateAccTenantWithInvalidParentDn(rName),
				ExpectError: regexp.MustCompile(`(.)+ Object may not exists`),
			},

			{
				Config: CreateAccTenantDataSourceUpdatedResource(rName, "annotation", "orchestrator:terraform-testacc"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "annotation", resourceName, "annotation"),
				),
			},
		},
	})
}
func CreateAccTenantDSWithoutRequired(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Tenant data source creation without required")
	resource := fmt.Sprintf(`
	data "aci_tenant" "test" {
		name = "%s"
	}
	`, rName)
	return resource
}

func CreateAccTenantConfigDataSource(rName string) string {
	fmt.Println("=== STEP  Basic: testing Tenant data source creation with required arguments only")
	resource := fmt.Sprintf(`
	resource "aci_tenant" "test"{
		name = "%s"
	}
	resource "aci_tenant" "test" {
		tenant_dn = aci_tenant.test.id
		name = "%s"
	}
	data "aci_tenant" "test" {
		tenant_dn = aci_tenant.test.tenant_dn
		name = aci_tenant.test.name
	}
	`, rName, rName)
	return resource
}
