package aci

import (
    "os"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var Test = map[string]interface{}{
		"base64": map[string]interface{}{
			"valid": []interface{}{ "cHJtNWIwaHVlaw==" },
			"invalid": []interface{}{ "a3+J1b%mFs//" },
		},
		"cidr": map[string]interface{}{
			"valid": []interface{}{ "19.152.40.31/20" },
			"invalid": []interface{}{ "277.268.279.265/29" },
		},
		"ipv4": map[string]interface{}{
			"valid": []interface{}{ "238.52.197.160" },
			"invalid": []interface{}{ "275.297.295.284" },
		},
		"ipv6": map[string]interface{}{
			"valid": []interface{}{ "9a50:82d5:e85f:e33c:da27:ce4d" },
			"invalid": []interface{}{ "invalidIPv6" },
		},
		"json": map[string]interface{}{
			"valid": []interface{}{ "json({ \"attribute\" : \"value\" })" },
			"invalid": []interface{}{ "json({ name : val)" },
		},
		"mac": map[string]interface{}{
			"valid": []interface{}{ "b1:3b:93:74:3c:79" },
			"invalid": []interface{}{ "invalidMAC" },
		},
		"regex": map[string]interface{}{
			"valid": []interface{}{ "(?m)^[0-9]{2}$" },
			"invalid": []interface{}{ "[0-9)++" },
		},
		"string": map[string]interface{}{
			"valid": []interface{}{ "n8udrsczga" },
			"invalid": []interface{}{ 12345 },
		},
		"time": map[string]interface{}{
			"valid": []interface{}{ "2022-06-15T05:03:17.275769+00:00" },
			"invalid": []interface{}{ "2022-06-15 10:33:17.275769" },
		},
		"url-http": map[string]interface{}{
			"valid": []interface{}{ "http://87aterv456pm5fx.com" },
			"invalid": []interface{}{ "ht:/5uwx9x0grkim642.com" },
		},
		"url-https": map[string]interface{}{
			"valid": []interface{}{ "https://mw360o1q8h8282j.com" },
			"invalid": []interface{}{ "hts:/uu08asw6d10w6ie.com" },
		},
		"uuid": map[string]interface{}{
			"valid": []interface{}{ "77d3acbe-ec68-11ec-9279-7c8ae1943087" },
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