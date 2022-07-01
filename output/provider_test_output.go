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
			"valid": []interface{}{ "MWM4cjYxNXBoZw==", "ZTZnNnlxeGd2cg==", "bjM0YmlmOWpkOQ==", "Mjg4dzcyeDd3aQ==" },
			"invalid": []interface{}{ "a3+J1b%mFs//" },
		},
		"cidr": map[string]interface{}{
			"valid": []interface{}{ "130.108.0.0/20", "130.108.192.0/20", "130.108.144.0/20", "130.108.80.0/20" },
			"invalid": []interface{}{ "274.266.283.262/30" },
		},
		"ipv4": map[string]interface{}{
			"valid": []interface{}{ "130.108.231.110", "130.108.202.160", "130.108.43.127", "130.108.135.181" },
			"invalid": []interface{}{ "294.259.295.299" },
		},
		"ipv6": map[string]interface{}{
			"valid": []interface{}{ "2001:db8::34f4:0:0:f379", "2001:db8::34f4:0:0:f34b", "2001:db8::34f4:0:0:f312", "2001:db8::34f4:0:0:f3e8" },
			"invalid": []interface{}{ "invalidIPv6" },
		},
		"json": map[string]interface{}{
			"valid": []interface{}{ "json({ \"attribute\" : \"value0\" })", "json({ \"attribute\" : \"value1\" })", "json({ \"attribute\" : \"value2\" })", "json({ \"attribute\" : \"value3\" })" },
			"invalid": []interface{}{ "json({ name : val)" },
		},
		"mac": map[string]interface{}{
			"valid": []interface{}{ "be:e8:ba:fb:3c:72", "3c:85:03:83:98:8f", "39:a8:d1:99:2d:97", "ef:1b:64:b0:fa:97" },
			"invalid": []interface{}{ "invalidMAC" },
		},
		"regex": map[string]interface{}{
			"valid": []interface{}{ "(?m)^[0-9]{2}$", "^(\\$)(\\d)+" },
			"invalid": []interface{}{ "[0-9)++" },
		},
		"string": map[string]interface{}{
			"valid": []interface{}{ "tu59znwjey", "9j60mvxaey", "davprq8qtm", "nwfgq0uxli" },
			"invalid": []interface{}{ 12345 },
		},
		"time": map[string]interface{}{
			"valid": []interface{}{ "2022-07-01T13:00:28.611928+00:00", "2022-07-24T13:00:28.611928+00:00", "2022-08-16T13:00:28.611928+00:00", "2022-09-08T13:00:28.611928+00:00" },
			"invalid": []interface{}{ "2022-07-01 18:30:28.611928" },
		},
		"url-http": map[string]interface{}{
			"valid": []interface{}{ "http://sw834mnli9ho2ts.com", "http://xz2henjjhbhln63.com", "http://ld3sfb6vf8hu1xf.com", "http://xiddmwyqlcoqjsh.com" },
			"invalid": []interface{}{ "ht:/x9b3p9e88idrfu5.com" },
		},
		"url-https": map[string]interface{}{
			"valid": []interface{}{ "https://rswhqj3l65617zg.com", "https://sn2nr7f1af8kd7m.com", "https://nzkqawgjuo20d76.com", "https://3keuv9rnej0djtf.com" },
			"invalid": []interface{}{ "hts:/e48el2dczs72z7x.com" },
		},
		"uuid": map[string]interface{}{
			"valid": []interface{}{ "c80add71-f93d-11ec-92f6-7c8ae1979a87", "c80add72-f93d-11ec-a425-7c8ae1979a87", "c80add73-f93d-11ec-93fd-7c8ae1979a87", "c80add74-f93d-11ec-9373-7c8ae1979a87" },
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