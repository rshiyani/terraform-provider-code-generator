#!/usr/bin/env python3
import os
import argparse
from helper import *
from loading_animation import *
 
parser = argparse.ArgumentParser(formatter_class=argparse.RawDescriptionHelpFormatter,description='''\
TerraJinja Tool that generates files for Terraform Provider.
     
Made By:  █▀▀ █▀█ █▀▀ █▀ ▀█▀   █▀▄ ▄▀█ ▀█▀ ▄▀█   █▀ █▄█ █▀ ▀█▀ █▀▀ █▀▄▀█ █▀
          █▄▄ █▀▄ ██▄ ▄█ ░█░   █▄▀ █▀█ ░█░ █▀█   ▄█ ░█░ ▄█ ░█░ ██▄ █░▀░█ ▄█
''')


parser.add_argument("--generate", type=str,default='all', help='Enter the type you want to generate Eg. "resource,resource_test,datasource" ')
parser.add_argument("--input", type=str,default='all', help='Enter the resource file name you want to generate for type Eg. "contract,subnet" ')
parser.add_argument("--pname", type=str,required=True, help='Enter the provider name Eg. "aws" ')
args = parser.parse_args()

generate = args.generate.split(",")
inputs = args.input.split(",")
pname=args.pname

dirs = ['./output', './output/resources', './output/datasources', './output/models', './output/client','./output/terraformer']
for dir in dirs:
    if not os.path.isdir(dir):
        os.mkdir(dir)

isConfigResources = os.path.isdir('./config/resources')
isConfigProvider = os.path.isfile("./config/provider.yml")
isConfigDatasources = os.path.isdir('./config/datasources')
isConfigClient = os.path.isdir('./config/client')
isConfigTerraformer = os.path.isdir('./config/terraformer')


resourceInputs = None
datasourceInput = None
clientInputs = None
terraformerInputs=None

if "all" in inputs:
    resourceInputs = os.listdir("./config/resources")
    datasourceInput = os.listdir("./config/datasources")
    clientInputs = os.listdir('./config/client')
    terraformerInputs = os.listdir('./config/terraformer')
else:
    resourceInputs = datasourceInput = clientInputs = terraformerInputs = inputs
   
if "all" in generate:
    if isConfigResources:
        generate_resource(resourceInputs, pname)
        generate_resource_test(resourceInputs, pname)
        generate_model(resourceInputs, pname)
    if isConfigProvider:
        generate_provider_test()
        generate_provider()
    if isConfigDatasources:
        generate_datasource(datasourceInput, pname)
        generate_datasource_test(datasourceInput, pname)
    if isConfigClient:
        generate_client(clientInputs)
        generate_client_test(clientInputs)
    if isConfigTerraformer:
        generate_terraformer(terraformerInputs, pname)
else:
    if "resource" in generate and isConfigResources:
        generate_resource(resourceInputs, pname)
    if "resource_test" in generate and isConfigResources:
        generate_resource_test(resourceInputs, pname)
    if "provider_test" in generate and isConfigProvider:
        generate_provider_test()
    if "provider" in generate and isConfigProvider:
        generate_provider()    
    if "datasource" in generate and isConfigDatasources:
        generate_datasource(datasourceInput, pname)
    if "datasource_test" in generate and isConfigDatasources:
        generate_datasource_test(datasourceInput, pname)    
    if "model" in generate and isConfigResources:
        generate_model(resourceInputs, pname)
    if "client" in generate and isConfigClient:
        generate_client(clientInputs)
    if "client_test" in generate and isConfigClient:
        generate_client_test(clientInputs)
    if "terraformer" in generate and isConfigTerraformer:
        generate_terraformer(terraformerInputs, pname)    

# Formatting the go files created
format_all_files()

# print(types[1])
# for filename in os.listdir("./config"):
#     print(filename)

# --all generate all res, data,provider
# -data "contract,subnet" : generate datasource file -datasourcetest
#    -data -all 
# -resource "contract,subnet" : generate resource file 
# -markdown "contract,subnet" : generate markdown file 

# -make="datasource,resource,datasourcetest" -files="contrat,subnet" -files="all"

# -data -file=contract generate datasource file 
