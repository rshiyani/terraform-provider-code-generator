# terraform-provider-code-generator

# string
- ipv4
- ipv6
- cidr
- mac
- enum (string-in-slice)
- url
- file (content)

# int
- range (x, y)
- notInRange (not in range (x, y))
- port
-- test cases to check for range(x, y):
    - <x
    - =x
    - (x+y)/2
    - =y
    - >y

# float
- range (x, y)
- notInRange (not in range (x, y))

# bool

# 