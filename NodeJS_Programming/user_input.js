// Getting User Input
const { stdin } = require("process");
const readline = require("readline");
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout

});
rl.question("What is your name: ", (name) => {
    console.log(`You name is ${name}`);
    rl.close();
});

rl.on('close', () => {
    console.log("Interface Closed");
    process.exit(0);
});
