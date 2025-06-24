const fs = require('fs');
const { text } = require('stream/consumers');
let html = fs.readFileSync(
    './index.html', 'utf-8'
);
// console.log(text_data);
const { createServer } = require('node:http');

const hostname = '127.0.0.1';
const port = 8080;

const server = createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/html');
  res.end(html);
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});
