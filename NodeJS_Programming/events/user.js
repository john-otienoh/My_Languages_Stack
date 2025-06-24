const events = require('node:events');

module.exports = class extends events.EventEmitter{
    constructor(){
        super();
    }
}