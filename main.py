from jinja2 import Environment, FileSystemLoader
import yaml

config = yaml.full_load(open('./config/provider.yml'))
env = Environment(loader=FileSystemLoader('./templates'),
                  trim_blocks=True, lstrip_blocks=True)

template = env.get_template('provider.j2')

# to save the results
with open("output.go", "w") as fh:
    fh.write(template.render(config))
