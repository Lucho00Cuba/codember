const URL = 'https://codember.dev/data/message_02.txt'

fetch(URL)
  .then(response => response.text())
  .then(data => {
    let VALUE = 0
    for (let i = 0; i < data.length; i++) {
      if ( data[i] === '#') {
        VALUE += 1
      } else if (data[i] === '@') {
        VALUE -= 1
      } else if (data[i] === '*') {
        VALUE *= VALUE
      } else if (data[i] === '&') {
        process.stdout.write(`${VALUE}`);
      }
    }
  })
  .catch(err => console.log('Error:', err))