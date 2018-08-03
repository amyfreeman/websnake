import React from "react";
import ReactDOM from 'react-dom';
import Modal from "./components/Modal.jsx";
import CanvasContainer from "./components/CanvasContainer.jsx";
import {glue} from './glue.js';

class App extends React.Component {
    constructor() {
        super();
        this.state = {
            modalPresent: true
        }
        this.initiateSockets();
        this.send = this.send.bind(this);
        this.registerHandler = this.registerHandler.bind(this);
        this.onSTATUS = this.onSTATUS.bind(this);
        this.registerHandler("STATUS", this.onSTATUS);
    }
    render() {
        return (
            <div id="root">
            <CanvasContainer registerHandler={this.registerHandler} send={this.send}/>
            {this.state.modalPresent? <Modal send={this.send}/> : null}
            </div>
        );
    }
    initiateSockets(){
    this._socket = glue();
        this._socket.onMessage((data)=>{
            console.log("hey so...we've received a socket message without a channel...dat bad: " + data);
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
    }
    registerHandler(channel, handler){
        this._socket.channel(channel).onMessage((data)=>{
            handler(data);
        });
    }
    send(channel, message){
        this._socket.channel(channel).send(message);
    }
    onSTATUS(data){
        console.log(data);
        switch(data) {
            case "OPPONENT_FOUND":
                break;
            case "BEGIN":
                this.setState({
                    modalPresent: false
                });
                break;
            case "GAMEOVER":
                break;
            default:
                console.log("Unkown 'STATUS' command detected: " + data);
        } 
    }
}

ReactDOM.render(<App/>, document.getElementById('app'));