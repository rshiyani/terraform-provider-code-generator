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


isDir = os.path.isdir('./output')
if not isDir:
    os.mkdir('./output')
isDirResources = os.path.isdir('./output/resources')
if not isDirResources:
    os.mkdir('./output/resources')
isDirDatasources = os.path.isdir('./output/datasources')
if not isDirDatasources:
    os.mkdir('./output/datasources')
isDirModels = os.path.isdir('./output/models')
if not isDirModels:
    os.mkdir('./output/models')


if "all" in inputs:
    if "all" in generate:
        inputs = os.listdir("./config/resources")
        generate_resource(inputs, pname)
        generate_resource_test(inputs, pname)
        generate_provider_test()
        generate_provider()  
        generate_model(inputs,pname)
        inputs = os.listdir("./config/datasources")
        generate_datasource(inputs,pname)
        generate_datasource_test(inputs,pname)    
    else:
        if "resource" in generate:
            inputs = os.listdir("./config/resources")
            generate_resource(inputs, pname)
        if "resource_test" in  generate:
            inputs = os.listdir("./config/resources")
            generate_resource_test(inputs, pname)
        if "provider_test" in generate:
            generate_provider_test()
        if "provider" in generate:
            generate_provider()    
        if "datasource" in generate:
            inputs = os.listdir("./config/datasources")
            generate_datasource(inputs,pname)
        if "datasource_test" in generate:
            inputs = os.listdir("./config/datasources")
            generate_datasource_test(inputs,pname)    
        if "model" in generate:
            inputs = os.listdir("./config/resources")
            generate_model(inputs,pname)
            
else:
    if "all" in generate:
        generate_resource(inputs, pname)
        generate_resource_test(inputs, pname)
        generate_provider_test()
        generate_provider()  
        generate_datasource(inputs,pname)
        generate_datasource_test(inputs,pname)    
        generate_model(inputs,pname)
    else:
        if "resource" in generate:
            generate_resource(inputs, pname)
        if "resource_test" in  generate:
            generate_resource_test(inputs, pname)
        if "provider_test" in generate:
            generate_provider_test()
        if "provider" in generate:
            generate_provider()    
        if "datasource" in generate:
            generate_datasource(inputs,pname)
        if "datasource_test" in generate:
            generate_datasource_test(inputs,pname)
        if "model" in generate:
            generate_model(inputs,pname)

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
