import requests

URL = 'https://codember.dev/data/message_02.txt'

req = requests.get(URL)

SECRET = req.text
VALUE = 0

for item in SECRET:
  if item == '#':
    VALUE = VALUE + 1
  elif item == '@':
    VALUE = VALUE - 1
  elif item == '*':
    VALUE = VALUE * VALUE
  elif item == '&':
    #print(f"current value: {VALUE}")
    print(VALUE, end='')
  else:
    pass
  
#print(f"\nfinal value {VALUE}")