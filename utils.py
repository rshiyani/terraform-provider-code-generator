from re import sub


def camelize(s):
    s = sub(r"(_|-)+", " ", s).title().replace(" ", "")
    return ''.join([s[0].lower(), s[1:]])


def pascalize(s):
    s = sub(r"(_|-)+", " ", s).title().replace(" ", "")
    return ''.join([s[0].upper(), s[1:]])


def snakify(name):
    s1 = sub('(.)([A-Z][a-z]+)', r'\1_\2', name)
    return sub('([a-z0-9])([A-Z])', r'\1_\2', s1).lower()


def is_list(value):
    return isinstance(value, list)

def quote(s):
    if str(s)[:5] == 'json(':
        return '`'+str(s)+'`'
    return '"'+str(s)+'"'

def make_dot_string(value, *args):
    lst = value + list(args)
    s = ".".join(lst)
    return s

def eliminate_zeroes(value):
    return '.'.join(value.split(".0."))

def eliminate_zeroes_and_capitalize(value):
    return pascalize('_'.join(value.split(".0.")))