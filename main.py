from jinja2 import Environment, FileSystemLoader
import yaml
import random
from utils import camelize, pascalize, snakify, is_list
import string
import datetime
import uuid
import base64


def generateRandomValues(data):
    data["cidr"] = {
        "valid": '.'.join('%s' % random.randint(0, 255) for i in range(4)) + "/" + str(random.randint(0, 31)),
        "invalid": '.'.join('%s' % random.randint(256, 300) for i in range(4)) + "/" + str(random.randint(0, 31))
    }
    data["ipv4"] = {
        "valid": '.'.join('%s' % random.randint(0, 255) for i in range(4)),
        "invalid": '.'.join('%s' % random.randint(256, 300) for i in range(4))
    }
    data["ipv6"] = {
        "valid": ':'.join('%x' % random.randint(0, 16**4) for i in range(6)),
        "invalid": 'invalidIPv6'
    }
    data["mac"] = {
        "valid": ':'.join('%02x' % random.randint(0, 255) for i in range(6)),
        "invalid": 'invalidMAC'
    }
    # data["port"] = {
    #     "valid": random.randint(1, 65535),
    #     "invalid": random.randint(65536, 66000)
    # }
    # data["port0"] = {
    #     "valid": random.randint(0, 65535),
    #     "invalid": random.randint(65536, 66000)
    # }
    data["time"] = {
        "valid": datetime.datetime.now(datetime.timezone.utc).isoformat(),
        "invalid": datetime.datetime.now()
    }
    data["url-http"] = {
        "valid": 'http://{}.com'.format(''.join(random.choices(string.ascii_lowercase + string.digits, k=15))),
        "invalid": 'ht:/{}.com'.format(''.join(random.choices(string.ascii_lowercase + string.digits, k=15)))
    }
    data["url-https"] = {
        "valid": 'https://{}.com'.format(''.join(random.choices(string.ascii_lowercase + string.digits, k=15))),
        "invalid": 'hts:/{}.com'.format(''.join(random.choices(string.ascii_lowercase + string.digits, k=15)))
    }
    data["uuid"] = {
        "valid": str(uuid.uuid1()),
        "invalid": "invalid323Uuid12"
    }
    data["string"] = {
        "valid": ''.join(random.choices(string.ascii_lowercase + string.digits, k=10)),
        "invalid": 12345
    }
    data["json"] = {
        "valid": "{ \"attribute\" : \"value\" }",
        "invalid": "{ name : val"
    }
    data["regex"] = {
        "valid": r'\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b',
        "invalid": r'[0-9)++'
    }
    data["base64"] = {
        "valid": (base64.b64encode((''.join(random.choices(string.ascii_lowercase + string.digits, k=10))).encode("ascii"))).decode("ascii"),
        "invalid": "a3+J1b%mFs//"
    }
    return data


def preProcess():
    with open("./config/resources/pre-contract.yml", 'r') as stream:
        data = yaml.safe_load(stream)
    data = generateRandomValues(data)
    # print(data)

    typeMap = {
        "IsCIDR": "cidr",
        "IsIPAddress": "ipv4",
        "IsIPv4Address": "ipv4",
        # "IsIPv4Range": "",
        "IsIPv6Address": "ipv6",
        "IsMACAddress": "mac",
        "IsPortNumber": "port",
        "IsPortNumberOrZero": "port0",
        "IsRFC3339Time": "time",
        "IsURLWithHTTPS": "url-https",
        "IsURLWithHTTPorHTTPS": "url-http",
        "IsUUID": "uuid",
        # "ListOfUniqueStrings": "",
        # "NoZeroValues": "",
        "StringIsBase64": "base64",
        # "StringIsEmpty": "",
        "StringIsJSON": "json",
        # "StringIsNotEmpty": "",
        # "StringIsNotWhiteSpace": "",
        "StringIsValidRegExp": "regex",
        # "StringIsWhiteSpace": ""
    }
    for schema in data['schemas']:
        if schema["type"] == "string":
            if "validation" in schema:
                if schema["validation"]["func_name"] in typeMap:
                    schema["subtype"] = typeMap[schema["validation"]["func_name"]]
                else:
                    schema["subtype"] = "string"
            else:
                schema["subtype"] = "string"
            # set random value as per subtype
            schema["test_params"] = {
                "valid": data[schema["subtype"]]["valid"],
                "invalid": data[schema["subtype"]]["invalid"]
            }
        elif schema["type"] == "int":
            if "validation" in schema:
                if schema["validation"]["func_name"] == "IntBetween":
                    schema["subtype"] = "range"
                    x = schema["validation"]["params"][0]
                    y = schema["validation"]["params"][1]
                    schema["test_params"] = {
                        "valid": [x, y, (x+y)//2],
                        "invalid": [x-1, y+1]
                    }
                elif schema["validation"]["func_name"] == "IsPortNumber":
                    schema["subtype"] = "port"
                    schema["test_params"] = {
                        "valid": [1, 53, 65535],
                        "invalid": [0, 65536]
                    }
                elif schema["validation"]["func_name"] == "IsPortNumberOrZero":
                    schema["subtype"] = "port0"
                    schema["test_params"] = {
                        "valid": [0, 1, 53, 65535],
                        "invalid": [-1, 65536]
                    }
                else:
                    schema["subtype"] = "int"
            else:
                schema["subtype"] = "int"
        elif schema["type"] == "float":
            if "validation" in schema:
                if schema["validation"]["func_name"] == "FloatBetween":
                    schema["subtype"] = "range"
                    x = schema["validation"]["params"][0]
                    y = schema["validation"]["params"][1]
                    schema["test_params"] = {
                        "valid": [x, y, (x+y)//2],
                        "invalid": [x-1, y+1]
                    }
                else:
                    schema["subtype"] = "float"
            else:
                schema["subtype"] = "float"
        elif schema["type"] == "bool":
            schema["test_params"] = {
                "valid": "true",
                "invalid": "truee"
            }
    with open('./config/resources/contract.yml', 'w') as outfile:
        yaml.dump(data, outfile, default_flow_style=False)


preProcess()

config = yaml.full_load(open('./config/resources/contract.yml'))
env = Environment(loader=FileSystemLoader('./templates'),
                  trim_blocks=True, lstrip_blocks=True)


env.filters["camelize"] = camelize
env.filters["pascalize"] = pascalize
env.filters["snakify"] = snakify
env.filters["is_list"] = is_list


template = env.get_template('resource_test.j2')

# to save the results
with open("target/output.go", "w") as fh:
    fh.write(template.render(config))
