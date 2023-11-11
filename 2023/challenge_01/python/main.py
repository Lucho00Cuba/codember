import requests

URL = 'https://codember.dev/data/message_01.txt'

req = requests.get(URL)
items = req.text.split(' ')

MAPS = {}

for item in items:
  # serialization
  item = item.lower().strip()

  if item not in MAPS:
    MAPS[item] = 1
  else:
    MAPS[item] = MAPS[item] + 1

for key, value in MAPS.items():
  print(f"{key}{value}", end='')