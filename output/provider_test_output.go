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
			"valid": []interface{}{ "aDdmejFqZnh4cQ==", "cGk0ems1bmZ0ZQ==", "b2ViNWE4bjFsMA==", "N3NnbjcwazQ2aA==" },
			"invalid": []interface{}{ "a3+J1b%mFs//" },
			"multiple_valids": []interface{}{ "MWY2ZGg0bnJzeg==", "bG95NG9zemhsYg==", "dzliY2swdGh4Yg==", "cDd2MHBlMng4Ng==", "NW1xMzJzaWZrMQ==", "azk5aTFndXF4eA==", "c2h3MXYzNzBucg==", "ODVxMnZ2OWllcQ==", "MngxNXd6d3BjYg==", "OXB1cTkzeTEwbA==", "NGd3ODhjNWl5ag==", "cGx5YTViaGJmMg==", "eW9uY2M5N3Vycw==", "a3Y5d3B6NGt4MQ==", "eXRueGxpa2tlYw==" },
		},
		"cidr": map[string]interface{}{
			"valid": []interface{}{ "35.13.208.0/20", "35.13.0.0/20", "35.13.160.0/20", "35.13.240.0/20" },
			"invalid": []interface{}{ "292.300.298.274/17" },
			"multiple_valids": []interface{}{ "35.13.208.0/20", "35.13.0.0/20", "35.13.160.0/20", "35.13.240.0/20", "35.13.192.0/20", "35.13.16.0/20", "35.13.32.0/20", "35.13.112.0/20", "35.13.48.0/20", "35.13.80.0/20", "35.13.144.0/20", "35.13.64.0/20", "35.13.128.0/20", "35.13.224.0/20", "35.13.176.0/20" },
		},
		"ipv4": map[string]interface{}{
			"valid": []interface{}{ "35.13.74.251", "35.13.71.94", "35.13.235.31", "35.13.196.58" },
			"invalid": []interface{}{ "256.292.300.282" },
			"multiple_valids": []interface{}{ "35.13.74.251", "35.13.71.94", "35.13.235.31", "35.13.196.58", "35.13.215.213", "35.13.150.226", "35.13.196.231", "35.13.171.163", "35.13.51.117", "35.13.84.42", "35.13.1.48", "35.13.20.103", "35.13.205.110", "35.13.210.174", "35.13.97.71" },
		},
		"ipv6": map[string]interface{}{
			"valid": []interface{}{ "2001:db8::34f4:0:0:f331", "2001:db8::34f4:0:0:f3f6", "2001:db8::34f4:0:0:f3a9", "2001:db8::34f4:0:0:f3e0" },
			"invalid": []interface{}{ "invalidIPv6" },
			"multiple_valids": []interface{}{ "2001:db8::34f4:0:0:f331", "2001:db8::34f4:0:0:f3f6", "2001:db8::34f4:0:0:f3a9", "2001:db8::34f4:0:0:f3e0", "2001:db8::34f4:0:0:f333", "2001:db8::34f4:0:0:f31b", "2001:db8::34f4:0:0:f34f", "2001:db8::34f4:0:0:f372", "2001:db8::34f4:0:0:f39e", "2001:db8::34f4:0:0:f303", "2001:db8::34f4:0:0:f32d", "2001:db8::34f4:0:0:f37e", "2001:db8::34f4:0:0:f33a", "2001:db8::34f4:0:0:f306", "2001:db8::34f4:0:0:f3a8" },
		},
		"json": map[string]interface{}{
			"valid": []interface{}{ "json({ \"attribute\" : \"value0\" })", "json({ \"attribute\" : \"value1\" })", "json({ \"attribute\" : \"value2\" })", "json({ \"attribute\" : \"value3\" })" },
			"invalid": []interface{}{ "json({ name : val)" },
			"multiple_valids": []interface{}{ "json({ \"attribute\" : \"value0\" })", "json({ \"attribute\" : \"value1\" })", "json({ \"attribute\" : \"value2\" })", "json({ \"attribute\" : \"value3\" })", "json({ \"attribute\" : \"value4\" })", "json({ \"attribute\" : \"value5\" })", "json({ \"attribute\" : \"value6\" })", "json({ \"attribute\" : \"value7\" })", "json({ \"attribute\" : \"value8\" })", "json({ \"attribute\" : \"value9\" })", "json({ \"attribute\" : \"value10\" })", "json({ \"attribute\" : \"value11\" })", "json({ \"attribute\" : \"value12\" })", "json({ \"attribute\" : \"value13\" })", "json({ \"attribute\" : \"value14\" })" },
		},
		"mac": map[string]interface{}{
			"valid": []interface{}{ "bb:e4:6b:71:af:97", "a6:b6:94:33:7b:b8", "41:69:00:ea:37:98", "5a:46:a2:79:b0:c2" },
			"invalid": []interface{}{ "invalidMAC" },
			"multiple_valids": []interface{}{ "c9:f0:0c:22:3c:29", "3b:c9:99:65:03:99", "d1:47:d8:b9:c1:f2", "df:ab:20:cc:8b:3d", "8b:35:36:05:75:9e", "91:23:b7:7d:da:2b", "67:91:20:11:d6:ec", "b6:37:70:22:76:5f", "3a:90:6d:71:90:1b", "ee:40:46:c8:2b:95", "50:38:03:72:2e:ba", "85:61:fa:98:c6:14", "24:80:55:f9:fe:8d", "f6:24:b4:20:ec:d9", "ef:a4:cb:8b:16:10" },
		},
		"regex": map[string]interface{}{
			"valid": []interface{}{ "(?m)^[0-9]{2}$", "^(\\$)(\\d)+" },
			"invalid": []interface{}{ "[0-9)++" },
			"multiple_valids": []interface{}{ "(?m)^[0-9]{2}$", "^(\\$)(\\d)+" },
		},
		"string": map[string]interface{}{
			"valid": []interface{}{ "q7pbh7tzl7", "slop6h5avg", "nwj6jy0bup", "iwfgcxyfsx" },
			"invalid": []interface{}{ 12345 },
			"multiple_valids": []interface{}{ "x4kknuphp6", "5gq84pdgi0", "jn7sfzjj1g", "1vx5qdvdgk", "62c48i8b2p", "3efsatt5pg", "2mb1gqubhs", "guljgqfesf", "z13n8xys00", "c93uevyu4q", "9j1ip9o8sh", "wt6g602spc", "q8lm3ek22f", "wlxbui0ons", "ordfmvqbzg" },
		},
		"time": map[string]interface{}{
			"valid": []interface{}{ "2022-07-06T05:50:10.242739+00:00", "2022-07-29T05:50:10.242739+00:00", "2022-08-21T05:50:10.242739+00:00", "2022-09-13T05:50:10.242739+00:00" },
			"invalid": []interface{}{ "2022-07-06 11:20:10.242739" },
			"multiple_valids": []interface{}{ "2022-07-06T05:50:10.242739+00:00", "2022-07-29T05:50:10.242739+00:00", "2022-08-21T05:50:10.242739+00:00", "2022-09-13T05:50:10.242739+00:00", "2022-10-06T05:50:10.242739+00:00", "2022-10-29T05:50:10.242739+00:00", "2022-11-21T05:50:10.242739+00:00", "2022-12-14T05:50:10.242739+00:00", "2023-01-06T05:50:10.242739+00:00", "2023-01-29T05:50:10.242739+00:00", "2023-02-21T05:50:10.242739+00:00", "2023-03-16T05:50:10.242739+00:00", "2023-04-08T05:50:10.242739+00:00", "2023-05-01T05:50:10.242739+00:00", "2023-05-24T05:50:10.242739+00:00" },
		},
		"url-http": map[string]interface{}{
			"valid": []interface{}{ "http://1y07wmabtapurfg.com", "http://l9hfgwcuo8juxdg.com", "http://qq3j81f5jcb6i0b.com", "http://k0ek07bm9zbcbg9.com" },
			"invalid": []interface{}{ "ht:/rwkg8z0r6iowg7i.com" },
			"multiple_valids": []interface{}{ "http://efzhbalfozkjm5u.com", "http://5qe645zfch2c8x8.com", "http://w25ryvfbdu1iehh.com", "http://f4i5jmgm169q9yj.com", "http://r0qv0nqyed3kseu.com", "http://lovxtpuy14b2ixn.com", "http://zpxj79twnjea0dl.com", "http://8or8vqxtmlkhznj.com", "http://fdf8o4qf8alczuk.com", "http://kizlnf5cyqm27wn.com", "http://idhkaacoxg7gnbg.com", "http://blfplg9rqnzvnk6.com", "http://mx8wpu1k34b41a3.com", "http://y11dnmxomy2jm0i.com", "http://35pqrb4mqwdh7t5.com" },
		},
		"url-https": map[string]interface{}{
			"valid": []interface{}{ "https://sj5f3ms2q3zsycx.com", "https://h1ebt4bltq9s94z.com", "https://9pm0otddvc3fhge.com", "https://bn52gyh7nl4uyar.com" },
			"invalid": []interface{}{ "hts:/i7oh4kxgdoedcbh.com" },
			"multiple_valids": []interface{}{ "https://32sav2331b9vqjl.com", "https://kl5efof5ymhypmk.com", "https://ih2t3fdi7lisdsf.com", "https://cf6ne16ft5vvw6s.com", "https://49yhsvmmr97t96b.com", "https://40cnfpeadhkk3uv.com", "https://3jiycg73eivqpme.com", "https://ciri7gb4jjpy2fe.com", "https://h0l4v50hbop85ko.com", "https://vju0tgwo071s2qx.com", "https://x0hajpjt1xln3r7.com", "https://9xw6mbezcih7mnj.com", "https://jejqym88s4ylub0.com", "https://vj3zg4wgnpi39wr.com", "https://w4pu195784hubvc.com" },
		},
		"uuid": map[string]interface{}{
			"valid": []interface{}{ "7f291f00-fcef-11ec-b244-7c8ae196ee3b", "7f291f01-fcef-11ec-a383-7c8ae196ee3b", "7f291f02-fcef-11ec-9d19-7c8ae196ee3b", "7f291f03-fcef-11ec-85cc-7c8ae196ee3b" },
			"invalid": []interface{}{ "invalid323Uuid12" },
			"multiple_valids": []interface{}{ "7f291f04-fcef-11ec-a6d7-7c8ae196ee3b", "7f291f05-fcef-11ec-b4c1-7c8ae196ee3b", "7f291f06-fcef-11ec-a4e5-7c8ae196ee3b", "7f291f07-fcef-11ec-ad5c-7c8ae196ee3b", "7f291f08-fcef-11ec-acdb-7c8ae196ee3b", "7f291f09-fcef-11ec-91d0-7c8ae196ee3b", "7f291f0a-fcef-11ec-b4a1-7c8ae196ee3b", "7f291f0b-fcef-11ec-8005-7c8ae196ee3b", "7f291f0c-fcef-11ec-8432-7c8ae196ee3b", "7f291f0d-fcef-11ec-93ba-7c8ae196ee3b", "7f291f0e-fcef-11ec-93df-7c8ae196ee3b", "7f291f0f-fcef-11ec-a66b-7c8ae196ee3b", "7f291f10-fcef-11ec-b99d-7c8ae196ee3b", "7f291f11-fcef-11ec-af5d-7c8ae196ee3b", "7f291f12-fcef-11ec-b3c7-7c8ae196ee3b" },
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