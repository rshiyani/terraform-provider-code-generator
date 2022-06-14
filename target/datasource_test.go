package acctest

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func TestAccAciContractDataSource_Basic(t *testing.T) {
	resourceName := "aci_contract.test"
	dataSourceName := "data.aci_contract.test"
	rName := makeTestVariable(acctest.RandString(5))
	randomParameter := acctest.RandStringFromCharSet(5, "abcdefghijklmnopqrstuvwxyz")
	randomValue := acctest.RandString(5)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccCheckAciContractDestroy,
		Steps: []resource.TestStep{
			{
				Config:      CreateAccContractDSWithoutRequiredName("name"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredIpv4("ipv4"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredIpv6("ipv6"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredMac("mac"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredCidr("cidr"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredTime("time"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredUrlHttps("url_https"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredUrlHttp("url_http"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredUuid("uuid"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredBase64("base_64"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredJson("json"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredRegExp("reg_exp"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredPortNumber("port_number"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredPortWithZero("port_with_zero"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredNuclearCode("nuclear_code"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredTestScore("test_score"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config:      CreateAccContractDSWithoutRequiredPercentage("percentage"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config: CreateAccContractConfigDataSource(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),

					resource.TestCheckResourceAttrPair(dataSourceName, "ipv4", resourceName, "ipv4"),

					resource.TestCheckResourceAttrPair(dataSourceName, "ipv6", resourceName, "ipv6"),

					resource.TestCheckResourceAttrPair(dataSourceName, "mac", resourceName, "mac"),

					resource.TestCheckResourceAttrPair(dataSourceName, "cidr", resourceName, "cidr"),

					resource.TestCheckResourceAttrPair(dataSourceName, "time", resourceName, "time"),

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

					resource.TestCheckResourceAttrPair(dataSourceName, "percentage", resourceName, "percentage"),

					resource.TestCheckResourceAttrPair(dataSourceName, "filter.#", resourceName, "filter.#"),
				),
			},
			{
				Config:      CreateAccContractDataSourceUpdate(rName, randomParameter, randomValue),
				ExpectError: regexp.MustCompile(`An argument named (.)+ is not expected here.`),
			},

			{
				Config:      CreateAccContractWithInvalidParentDn(rName),
				ExpectError: regexp.MustCompile(`(.)+ Object may not exists`),
			},

			{
				Config: CreateAccContractDataSourceUpdatedResource(rName, "annotation", "orchestrator:terraform-testacc"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "annotation", resourceName, "annotation"),
				),
			},
		},
	})
}

func CreateAccContractDSWithoutRequiredName(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Name")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		name = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredIpv4(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Ipv4")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		ipv4 = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredIpv6(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Ipv6")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		ipv6 = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredMac(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Mac")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		mac = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredCidr(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Cidr")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		cidr = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredTime(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Time")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		time = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredUrlHttps(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required UrlHttps")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		url_https = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredUrlHttp(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required UrlHttp")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		url_http = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredUuid(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Uuid")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		uuid = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredBase64(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Base64")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		base_64 = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredJson(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Json")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		json = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredRegExp(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required RegExp")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		reg_exp = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredGender(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Gender")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		gender = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredPortNumber(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required PortNumber")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		port_number = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredPortWithZero(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required PortWithZero")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		port_with_zero = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredNuclearCode(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required NuclearCode")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		nuclear_code = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredTestScore(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required TestScore")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		test_score = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredPercentage(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Percentage")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		percentage = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequiredFilter(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Filter")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		filter = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractDSWithoutRequired(rName string) string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required")
	resource := fmt.Sprintf(`
	data "aci_contract" "test" {
		name = "%s"
	}
	`, rName)
	return resource
}

func CreateAccContractConfigDataSource(rName string) string {
	fmt.Println("=== STEP  Basic: testing Contract data source creation with required arguments only")
	resource := fmt.Sprintf(`
	resource "aci_resource" "test"{
		name = "%s"
	}
	resource "aci_contract" "test" {
		tenant_dn = aci_tenant.test.id
		name = "%s"
	}
	data "aci_contract" "test" {
		tenant_dn = aci_contract.test.tenant_dn
		name = aci_contract.test.name
	}
	`, rName, rName)
	return resource
}
