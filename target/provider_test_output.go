package aci

import (
    "os"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var Test = map[string]interface{}{
		"base64": map[string]interface{}{
			"valid": []interface{}{ "aWdrcXNhYTMwMA==" },
			"invalid": []interface{}{ "a3+J1b%mFs//" },
		},
		"cidr": map[string]interface{}{
			"valid": []interface{}{ "191.223.38.116/16" },
			"invalid": []interface{}{ "290.272.261.290/18" },
		},
		"ipv4": map[string]interface{}{
			"valid": []interface{}{ "26.224.23.179" },
			"invalid": []interface{}{ "263.266.257.296" },
		},
		"ipv6": map[string]interface{}{
			"valid": []interface{}{ "c2ba:8207:1c3e:5f9f:97e9:d921" },
			"invalid": []interface{}{ "invalidIPv6" },
		},
		"json": map[string]interface{}{
			"valid": []interface{}{ "json({ \"attribute\" : \"value\" })" },
			"invalid": []interface{}{ "json({ name : val)" },
		},
		"mac": map[string]interface{}{
			"valid": []interface{}{ "6d:9f:78:69:c2:27" },
			"invalid": []interface{}{ "invalidMAC" },
		},
		"regex": map[string]interface{}{
			"valid": []interface{}{ "(?m)^[0-9]{2}$" },
			"invalid": []interface{}{ "[0-9)++" },
		},
		"string": map[string]interface{}{
			"valid": []interface{}{ "lbs1g9r4cy" },
			"invalid": []interface{}{ 12345 },
		},
		"time": map[string]interface{}{
			"valid": []interface{}{ "2022-06-14T13:22:46.949170+00:00" },
			"invalid": []interface{}{ "2022-06-14 18:52:46.949170" },
		},
		"url-http": map[string]interface{}{
			"valid": []interface{}{ "http://nyll6x8izll0pmy.com" },
			"invalid": []interface{}{ "ht:/ucjj3eauy2djxix.com" },
		},
		"url-https": map[string]interface{}{
			"valid": []interface{}{ "https://ols3bp4h8mw9zur.com" },
			"invalid": []interface{}{ "hts:/fndlnh16wlwou33.com" },
		},
		"uuid": map[string]interface{}{
			"valid": []interface{}{ "14bb2bf8-ebe5-11ec-afe4-7c8ae1943087" },
			"invalid": []interface{}{ "invalid323Uuid12" },
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