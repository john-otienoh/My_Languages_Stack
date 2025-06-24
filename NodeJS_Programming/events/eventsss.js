// Emitting and handling custom events
const events = require('events');
const user = require('./user');
// let emmitter = new events.EventEmitter();

// emmitter.on('userCreated', () => {
//     console.log('A new user is created')
// })

// emmitter.emit('userCreated');

let myEmmitter = new user()

myEmmitter.on('userCreated', () => {
    console.log('A new user is created')
})

myEmmitter.emit('userCreated');