from argparse import ArgumentParser
from jinja2 import FileSystemLoader, Environment
import pathlib

environment = Environment(loader=FileSystemLoader("."))

def render_templates(variables, filename):
    template = environment.get_template(filename)
    rendered_template = template.render(variables)
    print(rendered_template)
    
def main():
    parser = ArgumentParser()
    parser.add_argument('filename')
    args, variables = parser.parse_known_args()
    environment_variables = {}

    p = pathlib.Path(args.filename)
    if not pathlib.Path.exists(p):
        raise FileNotFoundError(f"File {p} doesn't exists")

    try:
        for var in  [dict([variable.lstrip("--").split("=",1)]) for variable in variables]:
            environment_variables.update(var)

    except ValueError:
        print("Bad argument")
        return
    
    render_templates(environment_variables, args.filename)

if __name__ == "__main__":
    main()





    
