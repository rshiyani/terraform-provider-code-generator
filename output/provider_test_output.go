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
			"valid": []interface{}{ "aHg5Nmw2dHJnZw==", "aW94d2dkZHM1Mg==", "NGNkbWJmYWl5MQ==", "ODQwNzNjMDY1dQ==" },
			"invalid": []interface{}{ "a3+J1b%mFs//" },
			"multiple_valids": []interface{}{ "bDVtOXkxdmNtag==", "c3ExMjMxMjB6ZQ==", "Z3F2MnRrMzl5eA==", "NDV2MnY0OXJhdA==", "ZGI5MTE3anRpbA==", "YTJtaDY5bjV3Nw==", "aXQxaHU1dzF4dg==", "dTM4b21uYml3Ng==", "OGFzNnR2eGlwbQ==", "ajB3anFuemFhNg==", "b25lbzhlMnBpdg==", "c3E2bWVkeTdpYQ==", "eXV0bXpxb2Vxcw==", "NTRtdmdndm93YQ==", "dnR1M2hzaW0yZw==" },
		},
		"cidr": map[string]interface{}{
			"valid": []interface{}{ "137.27.16.0/20", "137.27.80.0/20", "137.27.192.0/20", "137.27.112.0/20" },
			"invalid": []interface{}{ "284.287.293.271/12" },
			"multiple_valids": []interface{}{ "137.27.16.0/20", "137.27.80.0/20", "137.27.192.0/20", "137.27.112.0/20", "137.27.0.0/20", "137.27.176.0/20", "137.27.128.0/20", "137.27.240.0/20", "137.27.48.0/20", "137.27.144.0/20", "137.27.64.0/20", "137.27.96.0/20", "137.27.160.0/20", "137.27.224.0/20", "137.27.208.0/20" },
		},
		"ipv4": map[string]interface{}{
			"valid": []interface{}{ "137.27.246.44", "137.27.212.144", "137.27.21.184", "137.27.191.191" },
			"invalid": []interface{}{ "273.277.267.259" },
			"multiple_valids": []interface{}{ "137.27.246.44", "137.27.212.144", "137.27.21.184", "137.27.191.191", "137.27.79.194", "137.27.132.7", "137.27.204.201", "137.27.32.212", "137.27.251.120", "137.27.147.173", "137.27.46.42", "137.27.13.28", "137.27.77.24", "137.27.116.69", "137.27.54.101" },
		},
		"ipv6": map[string]interface{}{
			"valid": []interface{}{ "2001:db8::34f4:0:0:f3e8", "2001:db8::34f4:0:0:f339", "2001:db8::34f4:0:0:f30a", "2001:db8::34f4:0:0:f36c" },
			"invalid": []interface{}{ "invalidIPv6" },
			"multiple_valids": []interface{}{ "2001:db8::34f4:0:0:f3e8", "2001:db8::34f4:0:0:f339", "2001:db8::34f4:0:0:f30a", "2001:db8::34f4:0:0:f36c", "2001:db8::34f4:0:0:f33d", "2001:db8::34f4:0:0:f338", "2001:db8::34f4:0:0:f35d", "2001:db8::34f4:0:0:f304", "2001:db8::34f4:0:0:f350", "2001:db8::34f4:0:0:f39f", "2001:db8::34f4:0:0:f375", "2001:db8::34f4:0:0:f333", "2001:db8::34f4:0:0:f33e", "2001:db8::34f4:0:0:f3cf", "2001:db8::34f4:0:0:f305" },
		},
		"json": map[string]interface{}{
			"valid": []interface{}{ "json({ \"attribute\" : \"value0\" })", "json({ \"attribute\" : \"value1\" })", "json({ \"attribute\" : \"value2\" })", "json({ \"attribute\" : \"value3\" })" },
			"invalid": []interface{}{ "json({ name : val)" },
			"multiple_valids": []interface{}{ "json({ \"attribute\" : \"value0\" })", "json({ \"attribute\" : \"value1\" })", "json({ \"attribute\" : \"value2\" })", "json({ \"attribute\" : \"value3\" })", "json({ \"attribute\" : \"value4\" })", "json({ \"attribute\" : \"value5\" })", "json({ \"attribute\" : \"value6\" })", "json({ \"attribute\" : \"value7\" })", "json({ \"attribute\" : \"value8\" })", "json({ \"attribute\" : \"value9\" })", "json({ \"attribute\" : \"value10\" })", "json({ \"attribute\" : \"value11\" })", "json({ \"attribute\" : \"value12\" })", "json({ \"attribute\" : \"value13\" })", "json({ \"attribute\" : \"value14\" })" },
		},
		"mac": map[string]interface{}{
			"valid": []interface{}{ "03:f6:67:f7:e5:2c", "c7:29:2e:ee:a2:41", "26:dd:14:04:34:d2", "67:3b:ef:79:fc:ba" },
			"invalid": []interface{}{ "invalidMAC" },
			"multiple_valids": []interface{}{ "18:4d:d3:2f:97:00", "41:54:05:f4:74:bf", "2c:d1:0f:c8:da:83", "6a:d3:c7:50:1e:61", "e1:5a:08:11:50:be", "e2:c8:6d:89:6a:c9", "2a:55:51:d3:d9:e9", "dc:3a:c5:40:98:27", "e8:4f:38:81:30:93", "5c:1a:b3:20:6b:1c", "da:25:af:1a:a1:03", "eb:60:43:3c:e5:a0", "d3:ce:f1:08:1b:4d", "98:e4:d9:34:2f:d3", "94:fc:14:76:e9:34" },
		},
		"regex": map[string]interface{}{
			"valid": []interface{}{ "(?m)^[0-9]{2}$", "^(\\$)(\\d)+" },
			"invalid": []interface{}{ "[0-9)++" },
			"multiple_valids": []interface{}{ "(?m)^[0-9]{2}$", "^(\\$)(\\d)+" },
		},
		"string": map[string]interface{}{
			"valid": []interface{}{ "oxladho2p8", "1difghle0r", "jesow8jsro", "5da89s5gts" },
			"invalid": []interface{}{ 12345 },
			"multiple_valids": []interface{}{ "hrbwtzg9gr", "3ktqpp7awt", "z4wxv9f2wr", "y0wqcxo2l0", "m2kxlp46yt", "1fmhvzm37u", "eznsx5gfbj", "55o0p6ums8", "apgkj5cuka", "yhevvglayn", "n4fksc2f11", "b3ed7tx0pb", "70en7g2q0f", "kjgnfapdzk", "6yjx7j2py6" },
		},
		"time": map[string]interface{}{
			"valid": []interface{}{ "2022-07-11T05:17:51.153537+00:00", "2022-08-03T05:17:51.153537+00:00", "2022-08-26T05:17:51.153537+00:00", "2022-09-18T05:17:51.153537+00:00" },
			"invalid": []interface{}{ "2022-07-11 10:47:51.153537" },
			"multiple_valids": []interface{}{ "2022-07-11T05:17:51.153537+00:00", "2022-08-03T05:17:51.153537+00:00", "2022-08-26T05:17:51.153537+00:00", "2022-09-18T05:17:51.153537+00:00", "2022-10-11T05:17:51.153537+00:00", "2022-11-03T05:17:51.153537+00:00", "2022-11-26T05:17:51.153537+00:00", "2022-12-19T05:17:51.153537+00:00", "2023-01-11T05:17:51.153537+00:00", "2023-02-03T05:17:51.153537+00:00", "2023-02-26T05:17:51.153537+00:00", "2023-03-21T05:17:51.153537+00:00", "2023-04-13T05:17:51.153537+00:00", "2023-05-06T05:17:51.153537+00:00", "2023-05-29T05:17:51.153537+00:00" },
		},
		"url-http": map[string]interface{}{
			"valid": []interface{}{ "http://8qmsbgbvdbzd9x8.com", "http://uqas8imjrjzevdr.com", "http://hj2r0zmfzt8y53a.com", "http://i2kr1z8w457pklu.com" },
			"invalid": []interface{}{ "ht:/dssyu4oppwtes4z.com" },
			"multiple_valids": []interface{}{ "http://pfpyhl5x8pws9tg.com", "http://vm3cdidyga9ghf4.com", "http://qcnzl4fpz2vckvv.com", "http://jspapnmslsjsoe4.com", "http://mo34506s0l2vzau.com", "http://ae2ybt7viavkf4c.com", "http://rwsq6w19oyxx1s4.com", "http://3rkjkb1gkrzic8y.com", "http://9adk94dyclfyijm.com", "http://t5e8yqaojb67t7r.com", "http://f8aoop1124yv24v.com", "http://gyidbxcw7aqz2my.com", "http://ou760ed9o81aaft.com", "http://rsvj0y8msx6abxh.com", "http://ukarenevhb0vvhy.com" },
		},
		"url-https": map[string]interface{}{
			"valid": []interface{}{ "https://lh0fv4pbzmbb727.com", "https://lz3p8bipd3g5zeb.com", "https://0o8tbcn7ui25hvl.com", "https://pnwjcxznx4aa1t6.com" },
			"invalid": []interface{}{ "hts:/xv4a2dsbpoypmjz.com" },
			"multiple_valids": []interface{}{ "https://246sj6rzemplvsh.com", "https://lidengxq8927r6k.com", "https://57p4x5964542hc0.com", "https://fk0py5k5l842thx.com", "https://wxjgcarsu64g61a.com", "https://kau8h35zyx8aqdq.com", "https://35w2obq218uw009.com", "https://6xrllbx2rci55jx.com", "https://1avlw6acjpv7hzo.com", "https://nwh8ie92c9ktar5.com", "https://cigywfrv3gx7wis.com", "https://uml6aapc63060z5.com", "https://bi89869zaopjjxa.com", "https://wuruhinuhxxy1bw.com", "https://lzphyl3jc254b04.com" },
		},
		"uuid": map[string]interface{}{
			"valid": []interface{}{ "cf705711-00d8-11ed-b171-7c8ae1979a87", "cf705712-00d8-11ed-9f97-7c8ae1979a87", "cf705713-00d8-11ed-885b-7c8ae1979a87", "cf705714-00d8-11ed-bd6d-7c8ae1979a87" },
			"invalid": []interface{}{ "invalid323Uuid12" },
			"multiple_valids": []interface{}{ "cf705715-00d8-11ed-a839-7c8ae1979a87", "cf705716-00d8-11ed-9c8b-7c8ae1979a87", "cf705717-00d8-11ed-b221-7c8ae1979a87", "cf705718-00d8-11ed-a146-7c8ae1979a87", "cf705719-00d8-11ed-aeb5-7c8ae1979a87", "cf70571a-00d8-11ed-a3ff-7c8ae1979a87", "cf70571b-00d8-11ed-8e21-7c8ae1979a87", "cf70571c-00d8-11ed-b34e-7c8ae1979a87", "cf70571d-00d8-11ed-9091-7c8ae1979a87", "cf70571e-00d8-11ed-b536-7c8ae1979a87", "cf70571f-00d8-11ed-9f31-7c8ae1979a87", "cf705720-00d8-11ed-8d1b-7c8ae1979a87", "cf705721-00d8-11ed-81f7-7c8ae1979a87", "cf705722-00d8-11ed-b5e5-7c8ae1979a87", "cf705723-00d8-11ed-86bb-7c8ae1979a87" },
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