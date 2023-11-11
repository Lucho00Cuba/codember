const URL = 'https://codember.dev/data/message_01.txt';

fetch(URL)
  .then(response => response.text())
  .then(data => {
    const items = data.split(' ');
    const maps = {};

    for (const item of items) {
      // serialization
      const serializedItem = item.toLowerCase().trim();

      if (!(serializedItem in maps)) {
        maps[serializedItem] = 1;
      } else {
        maps[serializedItem] += 1;
      }
    }

    for (const [key, value] of Object.entries(maps)) {
      //console.log(`${key}${value}`);
      process.stdout.write(`${key}${value}`);
    }
  })
  .catch(error => console.error('Error:', error));