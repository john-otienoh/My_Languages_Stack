const fs = require('fs');
const { text } = require('stream/consumers');
let html = fs.readFileSync(
    './index.html', 'utf-8'
);
const { createServer } = require('node:http');

const hostname = '127.0.0.1';
const port = 8080;

// Craete Routes Home and about
const server = createServer((req, res) => {
//   res.statusCode = 200;
//   res.setHeader();
  let path = req.url;
  if(path === '/' || path.toLocaleLowerCase() === '/home'){
    res.writeHead(200, {
        'Content-Type': 'text/html',
        'header-mine': 'hello-john'
    });
    res.end(html.replace('{{%CONTENT%}}', 'You are in Homepage'));
  }else if(path.toLocaleLowerCase() === '/contact'){
    res.writeHead(200, {
        'Content-Type': 'text/html',
        'header-mine': 'hello-john'
    });
    res.end(html.replace("{{%CONTENT%}}", "You are in the Contact Us page."));
  }else if(path.toLocaleLowerCase() === '/about'){
    res.writeHead(200, {
        'Content-Type': 'text/html',
        'header-mine': 'hello-john'
    });
    res.end(html.replace("{{%CONTENT%}}", "You are in the About Page."));
  }else{
    res.writeHead(404, {
        'Content-Type': 'text/html',
        'header-mine': 'hello-john'
    });
    res.end(html.replace("{{%CONTENT%}}", 'Error 404!: Page not found</h1>'));
  }
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});
