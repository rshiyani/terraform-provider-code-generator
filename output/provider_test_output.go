package aci

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var expectErrorMap = map[string]string{
	"IsCIDR":               "expected (.)+ to be a valid IPv4 Value, got (.)+: (.)+",
	"IsIPAddress":          "expected (.)+ to contain a valid IP, got: (.)+",
	"IsIPv4Address":        "expected (.)+ to contain a valid IPv4 address, got: (.)+",
	"IsIPv6Address":        "expected (.)+ to contain a valid IPv6 address, got: (.)+",
	"IsMACAddress":         "expected (.)+ to be a valid MAC address, got (.)+: (.)+",
	"IsRFC3339Time":        "expected (.)+ to be a valid RFC3339 date, got (.)+: (.)+",
	"IsURLWithHTTPS":       "",
	"IsURLWithHTTPorHTTPS": "",
	"IsUUID":               "expected (.)+ to be a valid UUID, got (.)+",
	"StringIsBase64":       "expected (.)+ to be a base64 string, got (.)+",
	"StringIsJSON":         "(.)+ contains an invalid JSON: (.)+",
	"StringIsValidRegExp":  "(.)+: (.)+",
	"StringInSlice":        "expected (.)+ to be one of (.)+, got (.)+",
	"StringNotInSlice":     "expected (.)+ to not be any of (.)+, got (.)+",
	"IsCIDRNetwork":        "expected (.)+ to contain a network Value with between (.)+ and (.)+ significant bits, got: (.)+",
	"IntBetween":           "expected (.)+ to be in the range ((.)+ - (.)+), got (.)+",
	"IsPortNumber":         "expected (.)+ to be a valid port number, got: (.)+",
	"IsPortNumberOrZero":   "expected (.)+ to be a valid port number or 0, got: (.)+",
	"FloatBetween":         "expected (.)+ to be in the range ((.)+ - (.)+), got (.)+",
}

var Test = map[string]interface{}{
	"base64": map[string]interface{}{
		"valid":   []interface{}{"cDRmcXdncHkzbQ==", "c3d5cWFoMDRtYg==", "OXVpam1jZ2loag==", "Z3N6NmpzM2Jyag==", "ZjV4cHpxZWI2Yw==", "MGJ0cjVsOXZpZQ==", "aDlvemRyd2xvZw==", "ZGYwYnhwODhvMA==", "MzVpazBxZjdiNg==", "ejg0aW9jMjhkMA==", "ejhqdXJ3bDhpOQ==", "Yjh1cGJvbmRmYQ==", "ZmEzMTI4ODhoaA==", "YTNjbW12cDZmZQ==", "MTFwZjJvdG84dg=="},
		"invalid": []interface{}{"a3+J1b%mFs//"},
	},
	"cidr": map[string]interface{}{
		"valid":   []interface{}{"157.38.240.0/20", "157.38.144.0/20", "157.38.160.0/20", "157.38.0.0/20", "157.38.80.0/20", "157.38.32.0/20", "157.38.48.0/20", "157.38.96.0/20", "157.38.224.0/20", "157.38.176.0/20", "157.38.16.0/20", "157.38.64.0/20", "157.38.128.0/20", "157.38.192.0/20", "157.38.112.0/20"},
		"invalid": []interface{}{"278.287.266.294/30"},
	},
	"ipv4": map[string]interface{}{
		"valid":   []interface{}{"157.38.57.106", "157.38.84.107", "157.38.33.202", "157.38.52.74", "157.38.134.230", "157.38.130.146", "157.38.240.232", "157.38.68.158", "157.38.60.112", "157.38.101.40", "157.38.159.238", "157.38.179.120", "157.38.3.232", "157.38.250.173", "157.38.210.0"},
		"invalid": []interface{}{"260.295.286.277"},
	},
	"ipv6": map[string]interface{}{
		"valid":   []interface{}{"2001:db8::34f4:0:0:f373", "2001:db8::34f4:0:0:f33d", "2001:db8::34f4:0:0:f34c", "2001:db8::34f4:0:0:f38c", "2001:db8::34f4:0:0:f3a0", "2001:db8::34f4:0:0:f3a1", "2001:db8::34f4:0:0:f3a6", "2001:db8::34f4:0:0:f3ae", "2001:db8::34f4:0:0:f3c7", "2001:db8::34f4:0:0:f358", "2001:db8::34f4:0:0:f36d", "2001:db8::34f4:0:0:f306", "2001:db8::34f4:0:0:f32c", "2001:db8::34f4:0:0:f3fa", "2001:db8::34f4:0:0:f3fc"},
		"invalid": []interface{}{"invalidIPv6"},
	},
	"json": map[string]interface{}{
		"valid":   []interface{}{"json({ \"attribute\" : \"value0\" })", "json({ \"attribute\" : \"value1\" })", "json({ \"attribute\" : \"value2\" })", "json({ \"attribute\" : \"value3\" })", "json({ \"attribute\" : \"value4\" })", "json({ \"attribute\" : \"value5\" })", "json({ \"attribute\" : \"value6\" })", "json({ \"attribute\" : \"value7\" })", "json({ \"attribute\" : \"value8\" })", "json({ \"attribute\" : \"value9\" })", "json({ \"attribute\" : \"value10\" })", "json({ \"attribute\" : \"value11\" })", "json({ \"attribute\" : \"value12\" })", "json({ \"attribute\" : \"value13\" })", "json({ \"attribute\" : \"value14\" })"},
		"invalid": []interface{}{"json({ name : val)"},
	},
	"mac": map[string]interface{}{
		"valid":   []interface{}{"44:e2:f4:4a:ac:a6", "79:3b:bb:d4:62:73", "a8:53:9a:80:e1:24", "dd:84:7c:2f:8d:0a", "e9:c2:8f:86:57:67", "3a:c0:26:25:4d:b0", "27:9a:ef:f4:aa:c4", "43:a5:f2:3f:a8:6f", "d9:79:e6:da:87:1e", "ea:5e:6a:c6:7a:d8", "58:04:ca:67:6b:08", "c6:63:49:66:b7:b9", "a7:7d:a7:e6:93:0f", "47:a1:9e:c9:b0:05", "70:d3:9b:8b:72:31"},
		"invalid": []interface{}{"invalidMAC"},
	},
	"regex": map[string]interface{}{
		"valid":   []interface{}{"(?m)^[0-9]{2}$", "^(\\$)(\\d)+"},
		"invalid": []interface{}{"[0-9)++"},
	},
	"string": map[string]interface{}{
		"valid":   []interface{}{"75d3ocj0ho", "hiocthp8ih", "wb19ih9z6e", "7r215bn203", "u6yj9m3bgu", "d3b3x9rc69", "hiijyjngte", "xx7ruzayuc", "c2su21hwpz", "2zn7cosd1g", "t4lpb3ym0l", "afugdzmbqp", "o4h1mpnb5i", "z4aqcns5g5", "w1mixgxf3t"},
		"invalid": []interface{}{12345},
	},
	"time": map[string]interface{}{
		"valid":   []interface{}{"2022-06-24T06:27:04.777900+00:00", "2022-07-17T06:27:04.777900+00:00", "2022-08-09T06:27:04.777900+00:00", "2022-09-01T06:27:04.777900+00:00", "2022-09-24T06:27:04.777900+00:00", "2022-10-17T06:27:04.777900+00:00", "2022-11-09T06:27:04.777900+00:00", "2022-12-02T06:27:04.777900+00:00", "2022-12-25T06:27:04.777900+00:00", "2023-01-17T06:27:04.777900+00:00", "2023-02-09T06:27:04.777900+00:00", "2023-03-04T06:27:04.777900+00:00", "2023-03-27T06:27:04.777900+00:00", "2023-04-19T06:27:04.777900+00:00", "2023-05-12T06:27:04.777900+00:00"},
		"invalid": []interface{}{"2022-06-24 11:57:04.777900"},
	},
	"url-http": map[string]interface{}{
		"valid":   []interface{}{"http://r4xnw9x00zjg855.com", "http://w7pks92qq8e6c4f.com", "http://v2ncibkv6qnb1kc.com", "http://vjjtzsiypb0u6e3.com", "http://3tlxdyqgfb33cs1.com", "http://28rh3khtrrfk1km.com", "http://qmo7ew47c1c7ghd.com", "http://mxb85z2mxmpl1j2.com", "http://y8oo1do3g8x8par.com", "http://chc17rhszjq8rrn.com", "http://o74ch02eavu8vqe.com", "http://zvgoc2jq58r44ky.com", "http://3qfiyk1ifqte6gh.com", "http://bx438btpksehrmj.com", "http://bz8r2yq4n8q2ay5.com"},
		"invalid": []interface{}{"ht:/n2yg1zw44ucqzco.com"},
	},
	"url-https": map[string]interface{}{
		"valid":   []interface{}{"https://8vcbkkkm6ns4h4a.com", "https://sszn0jir7m1gfas.com", "https://3m55t71bya3p9z3.com", "https://94bppi30v3iu3ov.com", "https://aq44qe5s9anesjz.com", "https://elwsbesggusrcee.com", "https://qcbh7rc90fr1oi8.com", "https://hjq6mymjklj3ihz.com", "https://f4d6vf9jlfx5tqc.com", "https://hah0ziw7nhxslta.com", "https://ejism8ngfz095sk.com", "https://q8ffyozkx377f5r.com", "https://f8cplzaizk0i747.com", "https://k4jbh2adlhjwr0a.com", "https://x2ddcaseivt35o3.com"},
		"invalid": []interface{}{"hts:/k80z2aayt4ztr32.com"},
	},
	"uuid": map[string]interface{}{
		"valid":   []interface{}{"aa2b5ec0-f386-11ec-b8d7-7c8ae1943087", "aa2b5ec1-f386-11ec-b6ef-7c8ae1943087", "aa2b5ec2-f386-11ec-9f84-7c8ae1943087", "aa2b5ec3-f386-11ec-af9c-7c8ae1943087", "aa2b5ec4-f386-11ec-9d3b-7c8ae1943087", "aa2b5ec5-f386-11ec-9cbb-7c8ae1943087", "aa2b5ec6-f386-11ec-9d89-7c8ae1943087", "aa2b5ec7-f386-11ec-928c-7c8ae1943087", "aa2b5ec8-f386-11ec-9169-7c8ae1943087", "aa2b5ec9-f386-11ec-ab30-7c8ae1943087", "aa2b5eca-f386-11ec-a34a-7c8ae1943087", "aa2b5ecb-f386-11ec-8e50-7c8ae1943087", "aa2b5ecc-f386-11ec-b0df-7c8ae1943087", "aa2b5ecd-f386-11ec-93c0-7c8ae1943087", "aa2b5ece-f386-11ec-9919-7c8ae1943087"},
		"invalid": []interface{}{"invalid323Uuid12"},
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
