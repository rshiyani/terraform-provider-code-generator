from jinja2 import Environment, FileSystemLoader
import yaml
from utils import camelize, pascalize, snakify, is_list, quote, make_dot_string, eliminate_zeroes, eliminate_zeroes_and_capitalize, eliminate_dots_and_capitalize,get_first,eliminate_first


config = yaml.full_load(open('./config/datasources/movie.yml'))
env = Environment(loader=FileSystemLoader('./templates'),
                  trim_blocks=True, lstrip_blocks=True)

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

template = env.get_template('datasource_test.j2')

# to save the results
with open("output/datasource_test.go", "w") as fh:
    fh.write(template.render(config))
