from jinja2 import Environment, FileSystemLoader
import yaml
import random
from utils import camelize, pascalize, snakify, is_list, quote
import string
import datetime
from netaddr import *
import uuid
import base64


def generate_random_values(data):
    data["types"] = {}
    cidrRandom = IPNetwork('.'.join('%s' % random.randint(0,255) for i in range(2))+'.0.0/16')
    cidrList = [str(i) for i in cidrRandom.subnet(20)]
    random.shuffle(cidrList)
    cidrList = cidrList[:15]
    ipList = [str(i) for i in cidrRandom]
    random.shuffle(ipList)
    ipList = ipList[:15]
    ipv6Random = IPNetwork('2001:0db8:0000:0000:34f4:0000:0000:f3dd/120')
    ipv6List = [str(i) for i in ipv6Random]
    random.shuffle(ipv6List)
    ipv6List = ipv6List[:15]
    data["types"]["cidr"] = {
        "valid": cidrList,
        "invalid": ['.'.join('%s' % random.randint(256, 300) for i in range(4)) + "/" + str(random.randint(0, 31))]
    }
    data["types"]["ipv4"] = {
        "valid": ipList,
        "invalid": ['.'.join('%s' % random.randint(256, 300) for i in range(4))]
    }
    data["types"]["ipv6"] = {
        "valid": ipv6List,
        "invalid": ['invalidIPv6']
    }
    data["types"]["mac"] = {
        "valid": [':'.join('%02x' % random.randint(0, 255) for i in range(6)) for i in range(15)],
        "invalid": ['invalidMAC']
    }
    # data["port"] = {
    #     "valid": random.randint(1, 65535),
    #     "invalid": random.randint(65536, 66000)
    # }
    # data["port0"] = {
    #     "valid": random.randint(0, 65535),
    #     "invalid": random.randint(65536, 66000)
    # }
    data["types"]["time"] = {
        "valid": [str((datetime.datetime.now(datetime.timezone.utc) + datetime.timedelta(days=i*23)).isoformat()) for i in range(15)],
        "invalid": [str(datetime.datetime.now())]
    }
    data["types"]["url-http"] = {
        "valid": ["http://{}.com".format(''.join(random.choices(string.ascii_lowercase + string.digits, k=15))) for i in range(15)],
        "invalid": ["ht:/{}.com".format(''.join(random.choices(string.ascii_lowercase + string.digits, k=15)))]
    }
    data["types"]["url-https"] = {
        "valid": ["https://{}.com".format(''.join(random.choices(string.ascii_lowercase + string.digits, k=15))) for i in range(15)],
        "invalid": ["hts:/{}.com".format(''.join(random.choices(string.ascii_lowercase + string.digits, k=15)))]
    }
    data["types"]["uuid"] = {
        "valid": [str(uuid.uuid1()) for i in range(15)],
        "invalid": ["invalid323Uuid12"]
    }
    data["types"]["string"] = {
        "valid": [''.join(random.choices(string.ascii_lowercase + string.digits, k=10)) for i in range(15)],
        "invalid": [12345]
    }
    data["types"]["json"] = {
        "valid": ["json({ \"attribute\" : \"value"+str(i)+"\" })" for i in range(15)],
        "invalid": ["json({ name : val)"]
    }
    data["types"]["regex"] = {
        "valid": [r'(?m)^[0-9]{2}$', r'^(\$)(\d)+'],
        "invalid": [r'[0-9)++']
    }
    data["types"]["base64"] = {
        "valid": [(base64.b64encode((''.join(random.choices(string.ascii_lowercase + string.digits, k=10))).encode("ascii"))).decode("ascii") for i in range(15)],
        "invalid": ["a3+J1b%mFs//"]
    }
    return data

def handleListSetMap(data, schema):
    if schema["element"]["type"] == "schema":
        if schema["element"]["schema"]["type"] == "string":
            schema["test_params"] = {
                "valid": [''.join(random.choices(string.ascii_lowercase + string.digits, k=10)) for i in range(5)],
                "invalid": [10, 12.43]
            }
        elif schema["element"]["schema"]["type"] == "int":
            schema["test_params"] = {
                "valid": [random.randint(-1000,1000) for i in range(5)],
                "invalid": ["random",10.023]
            }
        elif schema["element"]["schema"]["type"] == "float":
            schema["test_params"] = {
                "valid": [random.randint(-1000,1000)*random.random() for i in range(5)],
                "invalid": ["random",10]
            }            
    else:    
        for schema in schema["element"]["schema"]:
            if schema["type"] == "string":
                if "validation" in schema:
                    if schema["validation"]["func_name"] in typeMap:
                        schema["subtype"] = typeMap[schema["validation"]["func_name"]]
                        schema["test_params"] = {
                            "valid": data["types"][schema["subtype"]]["valid"],
                            "invalid": data["types"][schema["subtype"]]["invalid"]
                        }
                    elif schema["validation"]["func_name"] == "StringInSlice" or schema["validation"]["func_name"] == "StringNotInSlice":
                        schema["subtype"] = "string"
                        schema["test_params"] = {
                            "valid": [i for i in schema["validation"]["params"]],
                            "invalid": [''.join(random.choices(string.ascii_lowercase + string.digits, k=10))],
                        }
                    elif schema["validation"]["func_name"] == "IsCIDRNetwork":
                        schema["subtype"] = "string"
                        x = schema["validation"]["params"][0]
                        y = schema["validation"]["params"][1]
                        schema["test_params"] = {
                            "valid": [x, y, (x+y)//2] + [random.randint(x+1,y-1) for i in range(12) if i != (x+y)//2],
                            "invalid": [x-1, y+1]
                        }
                    else:
                        schema["subtype"] = "string"
                        schema["test_params"] = {
                            "valid": [''.join(random.choices(string.ascii_lowercase + string.digits, k=10)) for i in range(5)],
                            "invalid": [10, 12.43]
                        }
                else:
                    schema["subtype"] = "string"
                    schema["test_params"] = {
                        "valid": [''.join(random.choices(string.ascii_lowercase + string.digits, k=10)) for i in range(5)],
                        "invalid": [10, 12.43]
                    }
            elif schema["type"] == "int":
                if "validation" in schema:
                    if schema["validation"]["func_name"] == "IntBetween":
                        schema["subtype"] = "range"
                        x = schema["validation"]["params"][0]
                        y = schema["validation"]["params"][1]
                        schema["test_params"] = {
                            "valid": [x, y, (x+y)//2] + [random.randint(x+1,y-1) for i in range(12) if i != (x+y)//2],
                            "invalid": [x-1, y+1]
                        }
                    elif schema["validation"]["func_name"] == "IsPortNumber":
                        schema["subtype"] = "port"
                        schema["test_params"] = {
                            "valid": [1, 65535] + [random.randint(2,65534) for i in range(13)],
                            "invalid": [0, 65536]
                        }
                    elif schema["validation"]["func_name"] == "IsPortNumberOrZero":
                        schema["subtype"] = "port0"
                        schema["test_params"] = {
                            "valid": [0, 65535] + [random.randint(1,65534) for i in range(13)],
                            "invalid": [-1, 65536]
                        }
                    else:
                        schema["subtype"] = "int"
                        schema["test_params"] = {
                            "valid": [random.randint(-1000,1000) for i in range(5)],
                            "invalid": ["random",10.023]
                        }
                else:
                    schema["subtype"] = "int"
                    schema["test_params"] = {
                        "valid": [random.randint(-1000,1000) for i in range(5)],
                        "invalid": ["random",10.023]
                    }
            elif schema["type"] == "float":
                if "validation" in schema:
                    if schema["validation"]["func_name"] == "FloatBetween":
                        schema["subtype"] = "range"
                        x = schema["validation"]["params"][0]
                        y = schema["validation"]["params"][1]
                        schema["test_params"] = {
                            "valid": [x, y, (x+y)/2] + [random.randint(x,y)*random.random() for i in range(13)],
                            "invalid": [x-1, y+1]
                        }
                    else:
                        schema["subtype"] = "float"
                        schema["test_params"] = {
                            "valid": [random.randint(-1000,1000)*random.random() for i in range(5)],
                            "invalid": ["random",10]
                        }
                else:
                    schema["subtype"] = "float"
                    schema["test_params"] = {
                        "valid": [random.randint(-1000,1000)*random.random() for i in range(5)],
                        "invalid": ["random",10]
                    }
            elif schema["type"] == "bool":
                schema["subtype"] = "bool"
            elif schema["type"] == "map":
                schema["subtype"] = "map"
                if schema["element"]["schema"]["type"] == "string":
                    schema["test_params"] = {
                        "valid": [''.join(random.choices(string.ascii_lowercase + string.digits, k=10)) for i in range(5)],
                        "invalid": [10, 12.43]
                    }
                elif schema["element"]["schema"]["type"] == "int":
                    schema["test_params"] = {
                        "valid": [random.randint(-1000,1000) for i in range(5)],
                        "invalid": ["random",10.023]
                    }
                elif schema["element"]["schema"]["type"] == "float":
                    schema["test_params"] = {
                        "valid": [random.randint(-1000,1000)*random.random() for i in range(5)],
                        "invalid": ["random",10]
                    } 
                elif schema["element"]["schema"]["type"] == "bool":
                    schema["test_params"] = {
                        "valid": ["true", "false"],
                        "invalid": ["random", 10]
                    }
            elif schema["type"] in ["list","set"]:
                handleListSetMap(data,schema)

typeMap = {
    "IsCIDR": "cidr",
    "IsIPAddress": "ipv4",
    "IsIPv4Address": "ipv4",
    "IsIPv6Address": "ipv6",
    "IsMACAddress": "mac",
    "IsRFC3339Time": "time",
    "IsURLWithHTTPS": "url-https",
    "IsURLWithHTTPorHTTPS": "url-http",
    "IsUUID": "uuid",
    "StringIsBase64": "base64",
    "StringIsJSON": "json",
    "StringIsValidRegExp": "regex",
}

def pre_process():
    with open("./config/resources/contract.yml", 'r') as stream:
        data = yaml.safe_load(stream)
    data = generate_random_values(data)

    for schema in data['schemas']:
        if schema["type"] == "string":
            if "validation" in schema:
                if schema["validation"]["func_name"] in typeMap:
                    schema["subtype"] = typeMap[schema["validation"]["func_name"]]
                    schema["test_params"] = {
                        "valid": data["types"][schema["subtype"]]["valid"],
                        "invalid": data["types"][schema["subtype"]]["invalid"]
                    }
                elif schema["validation"]["func_name"] == "StringInSlice" or schema["validation"]["func_name"] == "StringNotInSlice":
                    schema["subtype"] = "string"
                    schema["test_params"] = {
                        "valid": [i for i in schema["validation"]["params"]],
                        "invalid": [''.join(random.choices(string.ascii_lowercase + string.digits, k=10))],
                    }
                elif schema["validation"]["func_name"] == "IsCIDRNetwork":
                    schema["subtype"] = "string"
                    x = schema["validation"]["params"][0]
                    y = schema["validation"]["params"][1]
                    schema["test_params"] = {
                        "valid": [x, y, (x+y)//2] + [random.randint(x+1,y-1) for i in range(12) if i != (x+y)//2],
                        "invalid": [x-1, y+1]
                    }
                else:
                    schema["subtype"] = "string"
                    schema["test_params"] = {
                        "valid": [''.join(random.choices(string.ascii_lowercase + string.digits, k=10)) for i in range(5)],
                        "invalid": [10, 12.43]
                    }
            else:
                schema["subtype"] = "string"
                schema["test_params"] = {
                    "valid": [''.join(random.choices(string.ascii_lowercase + string.digits, k=10)) for i in range(5)],
                    "invalid": [10, 12.43]
                }
        elif schema["type"] == "int":
            if "validation" in schema:
                if schema["validation"]["func_name"] == "IntBetween":
                    schema["subtype"] = "range"
                    x = schema["validation"]["params"][0]
                    y = schema["validation"]["params"][1]
                    schema["test_params"] = {
                        "valid": [x, y, (x+y)//2] + [random.randint(x+1,y-1) for i in range(12) if i != (x+y)//2],
                        "invalid": [x-1, y+1]
                    }
                elif schema["validation"]["func_name"] == "IsPortNumber":
                    schema["subtype"] = "port"
                    schema["test_params"] = {
                        "valid": [1, 65535] + [random.randint(2,65534) for i in range(13)],
                        "invalid": [0, 65536]
                    }
                elif schema["validation"]["func_name"] == "IsPortNumberOrZero":
                    schema["subtype"] = "port0"
                    schema["test_params"] = {
                        "valid": [0, 65535] + [random.randint(1,65534) for i in range(13)],
                        "invalid": [-1, 65536]
                    }
                else:
                    schema["subtype"] = "int"
                    schema["test_params"] = {
                        "valid": [random.randint(-1000,1000) for i in range(5)],
                        "invalid": ["random",10.023]
                    }
            else:
                schema["subtype"] = "int"
                schema["test_params"] = {
                        "valid": [random.randint(-1000,1000) for i in range(5)],
                        "invalid": ["random",10.023]
                }
        elif schema["type"] == "float":
            if "validation" in schema:
                if schema["validation"]["func_name"] == "FloatBetween":
                    schema["subtype"] = "range"
                    x = schema["validation"]["params"][0]
                    y = schema["validation"]["params"][1]
                    schema["test_params"] = {
                        "valid": [x, y, (x+y)/2] + [random.randint(x,y)*random.random() for i in range(13)],
                        "invalid": [x-1, y+1]
                    }
                else:
                    schema["subtype"] = "float"
                    schema["test_params"] = {
                        "valid": [random.randint(-1000,1000)*random.random() for i in range(5)],
                        "invalid": ["random",10]
                    }
            else:
                schema["subtype"] = "float"
                schema["test_params"] = {
                    "valid": [random.randint(-1000,1000)*random.random() for i in range(5)],
                    "invalid": ["random",10]
                }
        elif schema["type"] == "bool":
            schema["subtype"] = "bool"
        elif schema["type"] == "map":
            schema["subtype"] = "map"
            if schema["element"]["schema"]["type"] == "string":
                schema["test_params"] = {
                    "valid": [''.join(random.choices(string.ascii_lowercase + string.digits, k=10)) for i in range(5)],
                    "invalid": [10, 12.43]
                }
            elif schema["element"]["schema"]["type"] == "int":
                schema["test_params"] = {
                    "valid": [random.randint(-1000,1000) for i in range(5)],
                    "invalid": ["random",10.023]
                }
            elif schema["element"]["schema"]["type"] == "float":
                schema["test_params"] = {
                    "valid": [random.randint(-1000,1000)*random.random() for i in range(5)],
                    "invalid": ["random",10]
                } 
            elif schema["element"]["schema"]["type"] == "bool":
                schema["test_params"] = {
                    "valid": ["true", "false"],
                    "invalid": ["random", 10]
                }
        elif schema["type"] in ["list","set"]:
            handleListSetMap(data,schema)
    with open('./config/resources/contract_generated.yml', 'w') as outfile:
        yaml.dump(data, outfile, default_flow_style=False)

def pre_process_for_provider():
    with open("./config/provider.yml", 'r') as stream:
        data = yaml.safe_load(stream)
    data = generate_random_values(data)
    new_data = {}
    for key, _ in data.items():
        if key == "types":
            for k, v in data[key].items():
                new_data[k] = {
                    "valid": data["types"][str(k)]["valid"],
                    "invalid": data["types"][str(k)]["invalid"]
                }
    data["types"] = new_data
    with open('./config/provider_generated.yml', 'w') as outfile:
        yaml.dump(data, outfile, default_flow_style=False)

pre_process()
pre_process_for_provider()

config = yaml.full_load(open('./config/provider_generated.yml'))
env = Environment(loader=FileSystemLoader('./templates'),
                  trim_blocks=True, lstrip_blocks=True)


env.filters["camelize"] = camelize
env.filters["pascalize"] = pascalize
env.filters["snakify"] = snakify
env.filters["is_list"] = is_list
env.filters["quote"] = quote


template = env.get_template('provider_test.j2')

# to save the results
with open("output/provider_test_output.go", "w") as fh:
    fh.write(template.render(config))


config = yaml.full_load(open('./config/resources/contract_generated.yml'))
env = Environment(loader=FileSystemLoader('./templates'),
                  trim_blocks=True, lstrip_blocks=True)


env.filters["camelize"] = camelize
env.filters["pascalize"] = pascalize
env.filters["snakify"] = snakify
env.filters["is_list"] = is_list
env.filters["quote"] = quote


template = env.get_template('resource_test_new.j2')

# to save the results
with open("output/resource_test_output.go", "w") as fh:
    fh.write(template.render(config))
