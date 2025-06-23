// Read file
let fs = require('fs');
const { text } = require('stream/consumers');
fs.readFile('./input.txt', 'utf-8', (err, data) => {
    if (err) {
        console.error('Error reading file: ' + err);
        return;
    }
    console.log(data);
    let content = `Data read from input.txt at ${new Date()}\n${data}`;
    fs.writeFile(
        'async.txt', content, () => {
            console.log('File written SuccessFully');
    });
});
console.log("Reading File....");

// Writing File

