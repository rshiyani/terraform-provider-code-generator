package acctest

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

//remove after resource_test
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
	"valid_cidr": map[string]interface{}{
		"valid":   []int{0, 32},
		"invalid": []int{349, 57},
	},
}

func TestAccAciContractDataSource_Basic(t *testing.T) {
	resourceName := "aci_contract.test"
	dataSourceName := "data.aci_contract.test"
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
				Config: CreateAccContractConfigDataSource(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),

					resource.TestCheckResourceAttrPair(dataSourceName, "annotation", resourceName, "annotation"),

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
				Config:      CreateAccContractUpdatedConfigDataSourceRandomAttr(randomParameter, randomValue),
				ExpectError: regexp.MustCompile(`An argument named (.)+ is not expected here.`),
			},

			{
				Config:      CreateAccContractDataSourceWithInvalidName(),
				ExpectError: regexp.MustCompile(`(.)+ Object may not exists`),
			},

			{
				Config: CreateAccContractUpdateConfigDataSource("annotation", randomValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "annotation", resourceName, "annotation"),
				),
			},
		},
	})
}

func CreateAccContractDSWithoutRequiredName() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Name")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredAnnotation() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Annotation")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredIpv4() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Ipv4")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredIpv6() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Ipv6")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredMac() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Mac")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredCidr() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Cidr")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredTime() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Time")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredUrlHttps() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required UrlHttps")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredUrlHttp() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required UrlHttp")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredUuid() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Uuid")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredBase64() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Base64")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredJson() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Json")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredRegExp() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required RegExp")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredGender() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Gender")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredPortNumber() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required PortNumber")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredPortWithZero() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required PortWithZero")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredNuclearCode() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required NuclearCode")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredTestScore() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required TestScore")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredPercentage() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Percentage")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0])
	return resource
}
func CreateAccContractDSWithoutRequiredFilter() string {
	fmt.Println("=== STEP  Basic: Testing Contract data source creation without required Filter")

	resource := CreateAccContractConfig()
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
        name = "%v"
        ipv4 = "%v"
        ipv6 = "%v"
        mac = "%v"
        cidr = "%v"
        time = "%v"
        url_https = "%v"
        url_http = "%v"
        uuid = "%v"
        base_64 = "%v"
        json = "%v"
        reg_exp = "%v"
        port_number = "%v"
        port_with_zero = "%v"
        nuclear_code = "%v"
        test_score = "%v"
        percentage = "%v"
	}
	`, resourceContractTest["name"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv4"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["ipv6"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["mac"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["cidr"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["time"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_https"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["url_http"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["uuid"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["base_64"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["json"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["reg_exp"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["port_number"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["port_with_zero"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["nuclear_code"].(map[string]interface{})["valid"].([]string)[0],
		resourceContractTest["test_score"].(map[string]interface{})["valid"].([]int)[0],
		resourceContractTest["percentage"].(map[string]interface{})["valid"].([]float64)[0])
	return resource
}

func CreateAccContractConfigDataSource() string {
	fmt.Println("=== STEP  Basic: testing Contract data source creation with required arguments only")
	resource := CreateAccContractConfig() //from resource_test
	resource += fmt.Sprintf(`

	data "aci_contract" "test" {
			name = aci_contract.test.name
			ipv4 = aci_contract.test.ipv4
			ipv6 = aci_contract.test.ipv6
			mac = aci_contract.test.mac
			cidr = aci_contract.test.cidr
			time = aci_contract.test.time
			url_https = aci_contract.test.url_https
			url_http = aci_contract.test.url_http
			uuid = aci_contract.test.uuid
			base_64 = aci_contract.test.base_64
			json = aci_contract.test.json
			reg_exp = aci_contract.test.reg_exp
			port_number = aci_contract.test.port_number
			port_with_zero = aci_contract.test.port_with_zero
			nuclear_code = aci_contract.test.nuclear_code
			test_score = aci_contract.test.test_score
			percentage = aci_contract.test.percentage
	}
	`)
	return resource
}

func CreateAccContractUpdatedConfigDataSourceRandomAttr(key, value string) string {
	fmt.Println("=== STEP  Basic: testing Contract data source creation with random attributes")
	resource := CreateAccContractConfig() //from resource_test
	resource += fmt.Sprintf(`

	data "aci_contract" "test" {
			name = aci_contract.test.name
			ipv4 = aci_contract.test.ipv4
			ipv6 = aci_contract.test.ipv6
			mac = aci_contract.test.mac
			cidr = aci_contract.test.cidr
			time = aci_contract.test.time
			url_https = aci_contract.test.url_https
			url_http = aci_contract.test.url_http
			uuid = aci_contract.test.uuid
			base_64 = aci_contract.test.base_64
			json = aci_contract.test.json
			reg_exp = aci_contract.test.reg_exp
			port_number = aci_contract.test.port_number
			port_with_zero = aci_contract.test.port_with_zero
			nuclear_code = aci_contract.test.nuclear_code
			test_score = aci_contract.test.test_score
			percentage = aci_contract.test.percentage
			%s = "%s"
	}
	`, key, value)
	return resource
}

func CreateAccContractDataSourceWithInvalidName() string {
	fmt.Println("=== STEP  Basic: testing Contract data source with invalid name")
	resource := CreateAccContractConfig() //from resource_test
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
			ipv4 = aci_contract.test.ipv4
			ipv6 = aci_contract.test.ipv6
			mac = aci_contract.test.mac
			cidr = aci_contract.test.cidr
			time = aci_contract.test.time
			url_https = aci_contract.test.url_https
			url_http = aci_contract.test.url_http
			uuid = aci_contract.test.uuid
			base_64 = aci_contract.test.base_64
			json = aci_contract.test.json
			reg_exp = aci_contract.test.reg_exp
			port_number = aci_contract.test.port_number
			port_with_zero = aci_contract.test.port_with_zero
			nuclear_code = aci_contract.test.nuclear_code
			test_score = aci_contract.test.test_score
			percentage = aci_contract.test.percentage
			name = "${ aci_contract.test.name}abc"
	}
	`)
	return resource
}

func CreateAccContractUpdatedConfigDataSource(key, value string) string {
	fmt.Println("=== STEP  Basic: testing Contract data source with updated resource")
	resource := CreateAccContractUpdatedAttr(key, value) //from resource_test
	resource += fmt.Sprintf(`
	data "aci_contract" "test" {
			name = aci_contract.test.name
			ipv4 = aci_contract.test.ipv4
			ipv6 = aci_contract.test.ipv6
			mac = aci_contract.test.mac
			cidr = aci_contract.test.cidr
			time = aci_contract.test.time
			url_https = aci_contract.test.url_https
			url_http = aci_contract.test.url_http
			uuid = aci_contract.test.uuid
			base_64 = aci_contract.test.base_64
			json = aci_contract.test.json
			reg_exp = aci_contract.test.reg_exp
			port_number = aci_contract.test.port_number
			port_with_zero = aci_contract.test.port_with_zero
			nuclear_code = aci_contract.test.nuclear_code
			test_score = aci_contract.test.test_score
			percentage = aci_contract.test.percentage
	}
	`)
	return resource
}

//remove after resource_test
func CreateAccContractConfig() string {
	resource := fmt.Sprintf(`
	resource "aci_tenant" "test"{
		name = "abcd"
	}
	resource "aci_contract" "test" {
		tenant_dn = aci_tenant.test.id
		name = "xyzw"
	}
	`)
	return resource
}

//remove after resource_test
func CreateAccContractUpdatedAttr(attribute, value string) string {
	resource := fmt.Sprintf(`
	resource "aci_tenant" "test"{
		name = "abcd"
	}
	resource "aci_contract" "test"{
		tenant_dn = aci_tenant.test.id
		name = "xyzw"
		%s = "%s"
	}
	`, attribute, value)
	return resource
}
