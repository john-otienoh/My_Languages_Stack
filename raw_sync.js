const fs = require('fs');
const { text } = require('stream/consumers');
let text_data = fs.readFileSync(
    './input.txt', 'utf-8'
);
console.log(text_data);

// writing file
let content = `Data read from input.txt at ${new Date()}\n${text_data}`;
fs.writeFileSync(
    './sync.txt', content
);