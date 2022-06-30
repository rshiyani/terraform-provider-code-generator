from ast import If, arguments
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

def urlsplit(path,ln):
    pathSplit = path 
    res = pathSplit.split("/")[:ln]
    return '/'.join(res)

def rmlaststr(s):
    str = s 
    res = str[:-1]
    return res


def urlpassvar(path, *args):
    arguments = path[1:].split("/")
    listArgs = []
    for i in args[0]:
        listArgs.append(i["name"])
    finalurl = ""
    for i in range(len(arguments)):
        if arguments[i] in listArgs:
            finalurl += "%s"
        else:
            finalurl += "/"+arguments[i] + "/"  
    finalurl = '"'+finalurl+'"'
    for i in range(len(listArgs)):
        check =  any(item in listArgs[i] for item in arguments)
        if check is True:
            finalurl += ","+ listArgs[i]
    return finalurl 



    