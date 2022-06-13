package aci

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var Test = map[string]interface{}{
	"base64": map[string]interface{}{
		"valid":   "anFsOXlpM3hqdA==",
		"invalid": "a3+J1b%mFs//",
	},
	"cidr": map[string]interface{}{
		"valid":   "166.180.39.7/3",
		"invalid": "297.273.273.279/11",
	},
	"ipv4": map[string]interface{}{
		"valid":   "139.231.130.228",
		"invalid": "264.284.273.258",
	},
	"ipv6": map[string]interface{}{
		"valid":   "f62e:7c46:412:2b24:8ccd:3c7c",
		"invalid": "invalidIPv6",
	},
	"json": map[string]interface{}{
		"valid":   `json({ "attribute" : "value" })`,
		"invalid": `json({ name : val)`,
	},
	"mac": map[string]interface{}{
		"valid":   "cf:8b:85:36:b5:08",
		"invalid": "invalidMAC",
	},
	"regex": map[string]interface{}{
		"valid":   "(?m)^[0-9]{2}$",
		"invalid": "[0-9)++",
	},
	"string": map[string]interface{}{
		"valid":   "tvwuc7y4se",
		"invalid": "12345",
	},
	"time": map[string]interface{}{
		"valid":   "2022-06-13T11:44:30.398002+00:00",
		"invalid": "2022-06-13 17:14:30.398002",
	},
	"url-http": map[string]interface{}{
		"valid":   "http://mtfgvi1dml20xaj.com",
		"invalid": "ht:/e831z5ijewcfm7b.com",
	},
	"url-https": map[string]interface{}{
		"valid":   "https://skjp9b56jtd5fav.com",
		"invalid": "hts:/kh4it35gw0ccb88.com",
	},
	"uuid": map[string]interface{}{
		"valid":   "2fb31dfc-eb0e-11ec-ad4c-7c8ae1943087",
		"invalid": "invalid323Uuid12",
	},
}

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"aci": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("ACI_USERNAME"); v == "" {
		t.Fatal("ACI_USERNAME env variable must be set for acceptance tests")
	}
	if v := os.Getenv("ACI_PASSWORD"); v == "" {
		t.Fatal("ACI_PASSWORD env variable must be set for acceptance tests")
	}
	if v := os.Getenv("ACI_URL"); v == "" {
		t.Fatal("ACI_URL env variable must be set for acceptance tests")
	}
}

var providerFactories = map[string]func() (*schema.Provider, error){
	"aci": func() (*schema.Provider, error) {
		return testAccProvider, nil
	},
}
