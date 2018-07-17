import {glue} from './glue.js';
import {Root} from './components/Root.js'

class IO{
    constructor(){
        this._socket = glue();

        this._socket.onMessage((data)=>{
            // use channels
            this.onMessage(data);
        });
    
        this._socket.on("connected", function() {
            console.log("socket connected");
        });
    
        this._socket.on("connecting", function() {
            console.log("socket connecting");
        });
    
        this._socket.on("disconnected", function() {
            console.log("socket disconnected");
        });
    
        this._socket.on("reconnecting", function() {
            console.log("socket reconnecting");
        });
    
        this._socket.on("error", function(e, msg) {
            console.log("socket error: " + msg);
        });
    
        this._socket.on("connect_timeout", function() {
            console.log("socket connect_timeout");
        });
    
        this._socket.on("timeout", function() {
            console.log("socket timeout");
        });
    
        this._socket.on("discard_send_buffer", function() {
            console.log("socket discard_send_buffer");
        });

        window.addEventListener('keydown', (event) => {
            switch (event.keyCode) {
                case 37:
                    this.leftPress();
                    break;
                case 38:
                    this.upPress();
                    break;
                case 39:
                    this.rightPress();
                    break;
                case 40:
                    this.downPress();
            }
        });
    }

    leftPress(){
        console.log("left key pressed");
        this._socket.send("left");
    }

    rightPress(){
        console.log("right key pressed");
        this._socket.send("right");
    }

    upPress(){
        console.log("up key pressed");
        this._socket.send("up");
    }

    downPress(){
        console.log("down key pressed");
        this._socket.send("down");
    }

    startButtonPress(){
        console.log("start button pressed");
        this._socket.send("start");
    }

    gameStarting(){
        console.log("the game is starting");
    }
    
    onMessage(data){
        console.log("in onMessage");
    }
}

var io = new IO();
export {io};