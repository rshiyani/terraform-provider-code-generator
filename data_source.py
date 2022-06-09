from jinja2 import Environment, FileSystemLoader
import yaml
from utils import camelize, pascalize, snakify

config = yaml.full_load(open('./config/resources/contract.yml'))
env = Environment(loader=FileSystemLoader('./templates'),
                  trim_blocks=True, lstrip_blocks=True)

env.filters["camelize"] = camelize
env.filters["pascalize"] = pascalize
env.filters["snakify"] = snakify

template = env.get_template('datasource.j2')

# to save the results
with open("target/datasource.go", "w") as fh:
    fh.write(template.render(config))
