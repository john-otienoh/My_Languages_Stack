const fs = require('fs');
const url = require('url');
const { text } = require('stream/consumers');
const replaceHtml = require('./modules/replaceHtml');

let html = fs.readFileSync('./template/index.html', 'utf-8');
let data = JSON.parse(fs.readFileSync('./data/data.json', 'utf-8'));
let productListHtml = fs.readFileSync('./template/products.html', 'utf-8');
let productDetailHtml = fs.readFileSync('./template/product_detail.html', 'utf-8');


// function replaceHtml(template, product){ 
//     let output = template.replace('{{%IMAGE%}}', product.images);
//     output = output.replace('{{%DESCRIPTION%}}', product.description);
//     output = output.replace('{{%NAME%}}', product.title);
//     output = output.replace('{{%STOCK%}}', product.stock);
//     output = output.replace('{{%RATING%}}', product.rating);
//     output = output.replace('{{%PRICE%}}', product.price);
//     output = output.replace('{{%ID%}}', product.id);
//     output = output.replace('{{%DISC%}}', product.discountPercentage)

//     return output;

// }

const { createServer } = require('node:http');
const { request } = require('node:https');

const hostname = '127.0.0.1';
const port = 8080;

// Craete Routes Home and about
const server = createServer((req, res) => {
//   res.statusCode = 200;
//   res.setHeader();
  let {query, pathname: path} = url.parse(req.url, true);
  console.log(query);
//   let path = req.url;
  if(path === '/' || path.toLocaleLowerCase() === '/home'){
    res.writeHead(200, {
        'Content-Type': 'text/html',
        'header-mine': 'hello-john'
    });
    res.end(html.replace('{{%CONTENT%}}', "HomePage"));
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
  }else if(path.toLocaleLowerCase() === '/products'){
    if(!query.id){
        let dataArray = data.map((d) => {
            return replaceHtml(productListHtml, d)
        })
        let productsHtmlResponse = html.replace('{{%CONTENT%}}', dataArray.join(','));
        res.writeHead(200, {'Content-Type': 'text/html','header-mine': 'hello-john'});
        res.end(productsHtmlResponse);
    }else {
        let productDetailResponse = replaceHtml(productDetailHtml, data[query.id]);
        res.end( html.replace('{{%CONTENT%}}', productDetailResponse));
    }
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
