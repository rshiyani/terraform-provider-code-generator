package aci

import (
    "os"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/Jeffail/gabs/v2"
)

var expectErrorMap = map[string]string{
	"IsCIDR": "expected (.)+ to be a valid IPv4 Value, got (.)+: (.)+",
	"IsIPAddress":"expected (.)+ to contain a valid IP, got: (.)+",
	"IsIPv4Address": "expected (.)+ to contain a valid IPv4 address, got: (.)+",
	"IsIPv6Address": "expected (.)+ to contain a valid IPv6 address, got: (.)+",
	"IsMACAddress":"expected (.)+ to be a valid MAC address, got (.)+: (.)+",
	"IsRFC3339Time": "expected (.)+ to be a valid RFC3339 date, got (.)+: (.)+",
	"IsURLWithHTTPS": "",
	"IsURLWithHTTPorHTTPS": "",
	"IsUUID": "expected (.)+ to be a valid UUID, got (.)+",
	"StringIsBase64": "expected (.)+ to be a base64 string, got (.)+",
	"StringIsJSON": "(.)+ contains an invalid JSON: (.)+",
	"StringIsValidRegExp":"(.)+: (.)+" ,
	"StringInSlice": "expected (.)+ to be one of (.)+, got (.)+",
	"StringNotInSlice": "expected (.)+ to not be any of (.)+, got (.)+",
	"IsCIDRNetwork": "expected (.)+ to contain a network Value with between (.)+ and (.)+ significant bits, got: (.)+",
	"IntBetween": "expected (.)+ to be in the range ((.)+ - (.)+), got (.)+",
	"IsPortNumber": "expected (.)+ to be a valid port number, got: (.)+",
	"IsPortNumberOrZero": "expected (.)+ to be a valid port number or 0, got: (.)+",
	"FloatBetween": "expected (.)+ to be in the range ((.)+ - (.)+), got (.)+",
}

var Test = map[string]interface{}{
		"base64": map[string]interface{}{
			"valid": []interface{}{ "ejRkMXBvanFodA==", "YnQ2bXJiZTZ5eA==", "bnNmaHV1ZGR5Yg==", "Z2RodjI2Y3RnaQ==" },
			"invalid": []interface{}{ "a3+J1b%mFs//" },
		},
		"cidr": map[string]interface{}{
			"valid": []interface{}{ "244.119.144.0/20", "244.119.240.0/20", "244.119.80.0/20", "244.119.208.0/20" },
			"invalid": []interface{}{ "270.271.281.277/20" },
		},
		"ipv4": map[string]interface{}{
			"valid": []interface{}{ "244.119.54.218", "244.119.36.244", "244.119.97.114", "244.119.4.88" },
			"invalid": []interface{}{ "268.268.257.289" },
		},
		"ipv6": map[string]interface{}{
			"valid": []interface{}{ "2001:db8::34f4:0:0:f304", "2001:db8::34f4:0:0:f389", "2001:db8::34f4:0:0:f3f0", "2001:db8::34f4:0:0:f3c1" },
			"invalid": []interface{}{ "invalidIPv6" },
		},
		"json": map[string]interface{}{
			"valid": []interface{}{ "json({ \"attribute\" : \"value0\" })", "json({ \"attribute\" : \"value1\" })", "json({ \"attribute\" : \"value2\" })", "json({ \"attribute\" : \"value3\" })" },
			"invalid": []interface{}{ "json({ name : val)" },
		},
		"mac": map[string]interface{}{
			"valid": []interface{}{ "3b:ab:04:78:14:eb", "ad:75:fe:a7:fe:e7", "77:cc:0a:3f:95:c1", "5a:03:7f:36:b4:ec" },
			"invalid": []interface{}{ "invalidMAC" },
		},
		"regex": map[string]interface{}{
			"valid": []interface{}{ "(?m)^[0-9]{2}$", "^(\\$)(\\d)+" },
			"invalid": []interface{}{ "[0-9)++" },
		},
		"string": map[string]interface{}{
			"valid": []interface{}{ "26s3fh68yv", "ejkg2wzka0", "oft2shujf9", "bk7aieggc3" },
			"invalid": []interface{}{ 12345 },
		},
		"time": map[string]interface{}{
			"valid": []interface{}{ "2022-06-29T12:39:34.897900+00:00", "2022-07-22T12:39:34.897900+00:00", "2022-08-14T12:39:34.897900+00:00", "2022-09-06T12:39:34.897900+00:00" },
			"invalid": []interface{}{ "2022-06-29 18:09:34.897900" },
		},
		"url-http": map[string]interface{}{
			"valid": []interface{}{ "http://w1fvaum3qgod67o.com", "http://hmnv4vtt4l2blif.com", "http://sdsjlwmbgblzi72.com", "http://ewpd2xilcst7a12.com" },
			"invalid": []interface{}{ "ht:/gme0w5dtwmxwmyj.com" },
		},
		"url-https": map[string]interface{}{
			"valid": []interface{}{ "https://uwrixu4kvb6sjo9.com", "https://9idsu4ajojkaj3x.com", "https://la39uzl85cxnpck.com", "https://2h5dthrfrfct0m5.com" },
			"invalid": []interface{}{ "hts:/0s41ej8edgtn82w.com" },
		},
		"uuid": map[string]interface{}{
			"valid": []interface{}{ "87f1c93b-f7a8-11ec-a40b-7c8ae1979a87", "87f1c93c-f7a8-11ec-a617-7c8ae1979a87", "87f1c93d-f7a8-11ec-a2c0-7c8ae1979a87", "87f1c93e-f7a8-11ec-b1a6-7c8ae1979a87" },
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


func searchInObject(testMap map[string]interface{}, attr string) interface{} {
	jsonStr, _ := json.Marshal(testMap)
	jsonParsed, _ := gabs.ParseJSON([]byte(jsonStr))
	return jsonParsed.Path(attr).Data()
}