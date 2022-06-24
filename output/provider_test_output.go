package aci

import (
    "os"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			"valid": []interface{}{ "YmM1NjJkNWd0eg==", "OTRlNXdkcjM0ZA==", "Nnc1NzlhN203eQ==", "a3Bia3d2aWdhMA==", "ang1eXJjb3JmaA==", "Z25mczZ1ZHFyaA==", "OTBrNjVzaHlhZA==", "b3Z6dmh2b2NyZg==", "NXBneTl1dm41NA==", "MHl6ZXA0b3lmNQ==", "MnRwYXg0Zzh6dw==", "cXp0bXcwN2FrOA==", "aXF0cGZkZmFwYw==", "aThweHNzdzV3OQ==", "ZXk4bjBobHlpaw==" },
			"invalid": []interface{}{ "a3+J1b%mFs//" },
		},
		"cidr": map[string]interface{}{
			"valid": []interface{}{ "176.196.32.0/20", "176.196.192.0/20", "176.196.0.0/20", "176.196.16.0/20", "176.196.224.0/20", "176.196.160.0/20", "176.196.128.0/20", "176.196.208.0/20", "176.196.112.0/20", "176.196.80.0/20", "176.196.176.0/20", "176.196.64.0/20", "176.196.144.0/20", "176.196.240.0/20", "176.196.48.0/20" },
			"invalid": []interface{}{ "277.278.289.290/6" },
		},
		"ipv4": map[string]interface{}{
			"valid": []interface{}{ "176.196.46.88", "176.196.28.211", "176.196.250.129", "176.196.38.82", "176.196.220.70", "176.196.184.219", "176.196.180.220", "176.196.123.21", "176.196.133.184", "176.196.249.171", "176.196.176.160", "176.196.128.87", "176.196.179.139", "176.196.63.153", "176.196.199.182" },
			"invalid": []interface{}{ "284.294.285.267" },
		},
		"ipv6": map[string]interface{}{
			"valid": []interface{}{ "2001:db8::34f4:0:0:f372", "2001:db8::34f4:0:0:f3b2", "2001:db8::34f4:0:0:f3a1", "2001:db8::34f4:0:0:f388", "2001:db8::34f4:0:0:f352", "2001:db8::34f4:0:0:f3f4", "2001:db8::34f4:0:0:f3eb", "2001:db8::34f4:0:0:f33d", "2001:db8::34f4:0:0:f3db", "2001:db8::34f4:0:0:f390", "2001:db8::34f4:0:0:f333", "2001:db8::34f4:0:0:f3fa", "2001:db8::34f4:0:0:f3c9", "2001:db8::34f4:0:0:f3aa", "2001:db8::34f4:0:0:f38c" },
			"invalid": []interface{}{ "invalidIPv6" },
		},
		"json": map[string]interface{}{
			"valid": []interface{}{ "json({ \"attribute\" : \"value0\" })", "json({ \"attribute\" : \"value1\" })", "json({ \"attribute\" : \"value2\" })", "json({ \"attribute\" : \"value3\" })", "json({ \"attribute\" : \"value4\" })", "json({ \"attribute\" : \"value5\" })", "json({ \"attribute\" : \"value6\" })", "json({ \"attribute\" : \"value7\" })", "json({ \"attribute\" : \"value8\" })", "json({ \"attribute\" : \"value9\" })", "json({ \"attribute\" : \"value10\" })", "json({ \"attribute\" : \"value11\" })", "json({ \"attribute\" : \"value12\" })", "json({ \"attribute\" : \"value13\" })", "json({ \"attribute\" : \"value14\" })" },
			"invalid": []interface{}{ "json({ name : val)" },
		},
		"mac": map[string]interface{}{
			"valid": []interface{}{ "b1:8f:da:db:9e:60", "69:3f:be:2c:9a:82", "1d:c7:95:b8:2e:aa", "f8:74:b6:5f:df:e9", "f7:3a:ba:cc:72:82", "ed:b5:24:c5:65:83", "2e:d3:f6:b4:9c:84", "97:56:75:73:7c:e7", "a0:17:c5:4d:c8:d3", "62:c6:45:c7:bc:c4", "6e:ba:d6:e2:a4:63", "2f:bc:7d:8f:74:65", "33:fc:98:92:53:c7", "a9:9f:57:58:01:39", "85:63:5c:0a:4b:c2" },
			"invalid": []interface{}{ "invalidMAC" },
		},
		"regex": map[string]interface{}{
			"valid": []interface{}{ "(?m)^[0-9]{2}$", "^(\\$)(\\d)+" },
			"invalid": []interface{}{ "[0-9)++" },
		},
		"string": map[string]interface{}{
			"valid": []interface{}{ "bd8m2eeccg", "vly3eh6lx4", "t0ud66eyv8", "3axr02msbc", "qn6n8ci2io", "o0r99uzcu7", "d63xpem5bl", "z7ejkf6dr0", "wmjlemcmwb", "78j08g8yh2", "qkboqrlri0", "f8qlp2btsk", "a1ecjffivj", "gqalpcn2mu", "n1i8932zx5" },
			"invalid": []interface{}{ 12345 },
		},
		"time": map[string]interface{}{
			"valid": []interface{}{ "2022-06-24T12:49:44.879349+00:00", "2022-07-17T12:49:44.879349+00:00", "2022-08-09T12:49:44.879349+00:00", "2022-09-01T12:49:44.879349+00:00", "2022-09-24T12:49:44.879349+00:00", "2022-10-17T12:49:44.879349+00:00", "2022-11-09T12:49:44.879349+00:00", "2022-12-02T12:49:44.879349+00:00", "2022-12-25T12:49:44.879349+00:00", "2023-01-17T12:49:44.879349+00:00", "2023-02-09T12:49:44.879349+00:00", "2023-03-04T12:49:44.879349+00:00", "2023-03-27T12:49:44.879349+00:00", "2023-04-19T12:49:44.879349+00:00", "2023-05-12T12:49:44.879349+00:00" },
			"invalid": []interface{}{ "2022-06-24 18:19:44.879349" },
		},
		"url-http": map[string]interface{}{
			"valid": []interface{}{ "http://b46dhjeba3sxf5n.com", "http://z296u4zxlybrwuq.com", "http://qehx17kjklta77q.com", "http://xdsbyhnu83yptw7.com", "http://oc5hylj92pa7n9g.com", "http://9tmjlpev6r7fi9h.com", "http://yxcmgfhmv6ck2kc.com", "http://fc1rze57ruf7vvk.com", "http://koqm0nvkrxlji8f.com", "http://xbjwpf1b9l28x33.com", "http://f0339t7rcnein3h.com", "http://2nzycwquehk3tt3.com", "http://0umhvu0j0qvih14.com", "http://fnejufowwmt52hs.com", "http://5clj99737nm7ghf.com" },
			"invalid": []interface{}{ "ht:/ymrxxfuye2pzhyw.com" },
		},
		"url-https": map[string]interface{}{
			"valid": []interface{}{ "https://rbfojf3ja1if19n.com", "https://09pv6vuzusqrr0s.com", "https://9hw4ks2hiua9en7.com", "https://ay7ku2ne66wlh0s.com", "https://o0dyttc9dqspkis.com", "https://jgkyc1nmqjy5rqq.com", "https://eevuky50q1xaw7a.com", "https://80m4nss4z0m9fbo.com", "https://m0og8p5xi7wvk4g.com", "https://lyvpo64az54m0wb.com", "https://8x5gqz4z3f5sp29.com", "https://492l0yitwgc9yy7.com", "https://4y8m5sf3553qqfe.com", "https://ldnkwv6xpmlbtuz.com", "https://wu7gt9n1g83n8qd.com" },
			"invalid": []interface{}{ "hts:/6ux6ehtmf2qq6ja.com" },
		},
		"uuid": map[string]interface{}{
			"valid": []interface{}{ "1f74d192-f3bc-11ec-98ce-7c8ae196ee3b", "1f74d193-f3bc-11ec-aa16-7c8ae196ee3b", "1f74d194-f3bc-11ec-93ef-7c8ae196ee3b", "1f74d195-f3bc-11ec-b5e8-7c8ae196ee3b", "1f74d196-f3bc-11ec-8030-7c8ae196ee3b", "1f74d197-f3bc-11ec-8dac-7c8ae196ee3b", "1f74d198-f3bc-11ec-8811-7c8ae196ee3b", "1f74d199-f3bc-11ec-954a-7c8ae196ee3b", "1f74d19a-f3bc-11ec-8c92-7c8ae196ee3b", "1f74d19b-f3bc-11ec-97c4-7c8ae196ee3b", "1f74d19c-f3bc-11ec-8f0c-7c8ae196ee3b", "1f74d19d-f3bc-11ec-a164-7c8ae196ee3b", "1f74d19e-f3bc-11ec-be63-7c8ae196ee3b", "1f74d19f-f3bc-11ec-97b5-7c8ae196ee3b", "1f74d1a0-f3bc-11ec-9145-7c8ae196ee3b" },
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