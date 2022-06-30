from jinja2 import Environment, FileSystemLoader
import yaml
from utils import urlsplit, rmlaststr, camelize, pascalize, snakify,urlpassvar


config = yaml.full_load(open('./config/servicesNew.yml'))
env = Environment(loader=FileSystemLoader('./templates'),
                  trim_blocks=True, lstrip_blocks=True)

env.filters["urlsplit"] = urlsplit
env.filters["rmlaststr"] = rmlaststr
env.filters["camelize"] = camelize
env.filters["pascalize"] = pascalize
env.filters["snakify"] = snakify
env.filters["urlpassvar"] = urlpassvar

template = env.get_template('services.j2')

# with open('./config/services.yml') as f:
#    dataMap = yaml.safe_load(f)
# print(proxy = Cut(dataMap))
# to save the results
with open("./target/services.go", "w") as fh:
    fh.write(template.render(config))
