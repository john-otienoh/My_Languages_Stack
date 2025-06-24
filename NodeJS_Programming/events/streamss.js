const fs = require('fs');
const { error } = require('node:console');
const { createServer } = require('node:http');
const { request } = require('node:https');

const hostname = '127.0.0.1';
const port = 8080;
const server = createServer()

// Efficient for small files
server.on('request', (req, res) => {
    fs.readFile('../data/data.json', 'utf-8', (err, data) => {
        fs.writeFile('../txt/stream.txt', data, () => {
            if(err){
                res.end("Something went wrong");
                return;
            }
            // console.log('File written from source to destination');
            res.end(data)
        })
    })

})

// Using streams 
server.on('request', (req, res) => {
    let rs = fs.createReadStream('../data/data.json');
    rs.on('data', (chunk) => {
        res.write(chunk);
    })
    rs.on('end', () => {
        res.end();
    })

    rs.on('error', (error) => {
        res.end(error.message);
    })
})

// Using Pipe method 
server.on('request', (req, res) => {
    let rs = fs.createReadStream('../data/data.json');
    rs.pipe(res);
})
server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});
