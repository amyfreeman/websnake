import {glue} from './glue.js';

class SocketConnection{
    constructor(){
        this.socket = glue();

        this.socket.onMessage(function(data) {
            // use channels
            if (data == "game beginning"){
                updateStatus(data)
            }
            else if (data == "confirmed. waiting for stranger"){
                updateStatus(data);
            }
            else{
                drawFromMsg(data);
            }
        });
    
        this.socket.on("connected", function() {
            console.log("connected");
        });
    
        this.socket.on("connecting", function() {
            console.log("connecting");
        });
    
        this.socket.on("disconnected", function() {
            console.log("disconnected");
        });
    
        this.socket.on("reconnecting", function() {
            console.log("reconnecting");
        });
    
        this.socket.on("error", function(e, msg) {
            console.log("error: " + msg);
        });
    
        this.socket.on("connect_timeout", function() {
            console.log("connect_timeout");
        });
    
        this.socket.on("timeout", function() {
            console.log("timeout");
        });
    
        this.socket.on("discard_send_buffer", function() {
            console.log("some data could not be sent and was discarded.");
        });
    }
    send(message){
        this.socket.send(message);
    }
}
  
export {SocketConnection};