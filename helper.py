import os
from utils import *
from preprocess import *

# Setting the filter in the environment
def set_filter(env):
    env.filters["camelize"] = camelize
    env.filters["pascalize"] = pascalize
    env.filters["snakify"] = snakify
    env.filters["is_list"] = is_list
    env.filters["quote"] = quote
    env.filters["make_dot_string"] = make_dot_string
    env.filters["eliminate_zeroes"] = eliminate_zeroes
    env.filters["eliminate_zeroes_and_capitalize"] = eliminate_zeroes_and_capitalize
    env.filters["eliminate_dots_and_capitalize"] = eliminate_dots_and_capitalize
    env.filters["get_first"] = get_first
    env.filters["eliminate_first"] = eliminate_first
  
# Generate the Terraform Resource File
def generate_resource_file(file, provider_name):
    print(f"=== Started Creation of Resource File")
    try:
        config = yaml.full_load(open(f'./config/resources/{file}.yml'))
    except Exception as e:
        print("Error: ", e)
        print(f"=== Error in Creation of Resource File")
        return False
    env = Environment(loader=FileSystemLoader('./templates'),
                    trim_blocks=True, lstrip_blocks=True)
    set_filter(env)

    template = env.get_template('resource.j2')

    # to save the results
    try:
        with open(f"output/resources/resource_{provider_name}_{file}.go", "w") as fh:
            fh.write(template.render(config))
    except Exception as e:
        print("Error: ", e)
        print(f"=== Error in Creation of Resource File")
        return False

    print(f"=== Completed Creation of Resource File")
    return True


# Function will be called by cli, which is responsible for generation of Resource File
def generate_resource(files, provider_name):
    isDir = os.path.isdir('./output/resources')
    if not isDir:
        os.mkdir('./output/resources')
    isDirPreprocess = os.path.isdir('./config/resources/preprocess')
    if not isDirPreprocess:
        os.mkdir('./config/resources/preprocess')

    for file in files:
        if file == "preprocess":
            continue
        if file[-4:] == ".yml":
            file = file[:-4]
        status = pre_process(file)
        if not status:
           exit()
        
        status = generate_resource_file(file, provider_name)
        if not status:
           exit()


# Function will be called by cli, which is responsible for generation of Resource Test File
def generate_resource_test(files,provider_name):

    for file in files:
        if file == "preprocess":
            continue
        if file[-4:] == ".yml":
            file = file[:-4]
        status = generate_resource_test_file(file,provider_name)
        if not status:
           exit()

# Generate the Terraform Provider File
def generate_provider_file():   
    print(f"=== Started Creation of Provider File")
    try:
        config = yaml.full_load(open(f'./config/provider.yml'))
    except Exception as e:
        print("Error: ", e)
        print(f"=== Error in Creation of Provider File")
        return False
    env = Environment(loader=FileSystemLoader('./templates'),
                    trim_blocks=True, lstrip_blocks=True)
    set_filter(env)

    template = env.get_template('provider.j2')

    # to save the results
    try:
        with open(f"output/provider.go", "w") as fh:
            fh.write(template.render(config))
    except Exception as e:
        print("Error: ", e)
        print(f"=== Error in Creation of Provider File")
        return False

    print(f"=== Completed Creation of Provider File")
    return True


# Function will be called by cli, which is responsible for generation of Provider File
def generate_provider():
    status = generate_provider_file()
    if not status:
        exit()

# Function will be called by cli, which is responsible for generation of Provider Test File
def generate_provider_test():
    status = pre_process_for_provider()
    if not status:
       exit()
    status = generate_provider_test_file()
    if not status:
       exit()

# Generate the Terraform Datasource Test File
def generate_datasource_test_file(file, provider_name):
    print(f"=== Started Creation of Datasource Test File named {file}")
    try:
        config = yaml.full_load(open(f'./config/datasources/{file}.yml'))
    except Exception as e:
        print("Error: ", e)
        print(f"=== Error in Creation of Datasource Test File named {file}")
        return False
    env = Environment(loader=FileSystemLoader('./templates'),
                    trim_blocks=True, lstrip_blocks=True)
    
    set_filter(env)
    
    template = env.get_template('datasource_test.j2')

    # to save the results
    try:
        with open(f"output/datasources/datasource_{provider_name}_{file}_test.go", "w") as fh:
            fh.write(template.render(config))
    except Exception as e:
        print("Error: ", e)
        print(f"=== Error in Creation of Datasource Test File named {file}")
        return False

    print(f"=== Completed Creation of Datasource Test File named {file}")
    return True
        
# Function will be called by cli, which is responsible for generation of Datasource Test File
def generate_datasource_test(files,provider_name):
    isDir = os.path.isdir('./output/datasources')
    if not isDir:
        os.mkdir('./output/datasources')
        
    for file in files:
        if file[-4:] == ".yml":
            file = file[:-4]
        status = generate_datasource_test_file(file,provider_name)
        if not status:
            exit()

# Generate the Terraform DataSource File
def generate_datasource_file(file,provider_name):
    print(f"=== Started Creation of Datasource File named {file}")
    try:
        config = yaml.full_load(open(f'./config/datasources/{file}.yml'))
    except Exception as e:
        print("Error: ", e)
        print(f"=== Error in Creation of Datasource File named {file}")
        return False
    env = Environment(loader=FileSystemLoader('./templates'),
                    trim_blocks=True, lstrip_blocks=True)

    set_filter(env)

    template = env.get_template('datasource.j2')

    # to save the results
    try:
        with open(f"output/datasources/datasource_{provider_name}_{file}.go", "w") as fh:
            fh.write(template.render(config))
    except Exception as e:
        print("Error: ", e)
        print(f"=== Error in Creation of Datasource File named {file}")
        return False

    print(f"=== Completed Creation of Datasource File named {file}")
    return True


# Function will be called by cli, which is responsible for generation of Datasource File
def generate_datasource(files,provider_name):
    isDir = os.path.isdir('./output/datasources')
    if not isDir:
        os.mkdir('./output/datasources')
    for file in files:
        if file[-4:] == ".yml":
            file = file[:-4]
        status = generate_datasource_file(file,provider_name)
        if not status:
            exit()

# Generate the Terraform Model File
def generate_model_file(file, provider_name):
    print(f"=== Started Creation of Model File for Resource named {file}")
    try:
        config = yaml.full_load(open(f'./config/resources/{file}.yml'))
    except Exception as e:
        print("Error: ", e)
        print(f"=== Error in Creation of Model File for Resource named {file}")
        return False
    env = Environment(loader=FileSystemLoader('./templates'),
                    trim_blocks=True, lstrip_blocks=True)

    set_filter(env)

    template = env.get_template('model.j2')

    # to save the results
    try:
        with open(f"output/models/{file}.go", "w") as fh:
            fh.write(template.render(config))
    except Exception as e:
        print("Error: ", e)
        print(f"=== Error in Creation of Model File for Resource named {file}")
        return False

    print(f"=== Completed Creation of Model File for Resource named {file}")
    return True



# Function will be called by cli, which is responsible for generation of Model File
def generate_model(files,provider_name):
    isDir = os.path.isdir('./output/models')
    if not isDir:
        os.mkdir('./output/models')
    for file in files:
        if file == "preprocess":
            continue
        if file[-4:] == ".yml":
            file = file[:-4]
        status = generate_model_file(file,provider_name)
        if not status:
            exit()
    
def generate_client_test(inputs):
    generate_client(inputs, "client_test.j2", "_test.go", " test")

def generate_client(inputs, template="client.j2", postfix=".go", test = ""):
    for file in inputs:
        # check that file exists
        file = file.split(".")[0] + ".yml"
        fileExists = os.path.isfile(f"./config/client/{file}")
        if not fileExists:
            print(f"=== WARNING: ./config/client/{file} does not exist, skipping.")
            continue
        
        # parse yaml
        print(f"=== INFO: Beginning creation of client{test} for {file}")
        try:
            config = yaml.full_load(open(f'./config/client/{file}'))
        except Exception as e:
            print(f"=== ERROR: Error parsing {file}. skipping.")
            print(e)
            continue
        
        # load jinja environment
        env = Environment(loader=FileSystemLoader('./templates'),
                    trim_blocks=True, lstrip_blocks=True)
        template = env.get_template(f'{template}')

        # render the template
        filename = file.split(".")[0]
        try:
            with open(f"output/client/{filename}{postfix}", "w") as fh:
                fh.write(template.render(config))
        except Exception as e:
            print(f"=== ERROR: This error may have happened due to insufficient write permissions or due to some syntatical error in the jinja template. Refer the error below for more info.")
            print(e)
            continue

        print(f"=== INFO: Created - client{test} for {file}")
        return