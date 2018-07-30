import React from "react";
import ReactDOM from 'react-dom';
import Modal from "./components/Modal.jsx";
import CanvasContainer from "./components/CanvasContainer.jsx";
import {glue} from './glue.js';

var handlers = {};

class App extends React.Component {
  constructor() {
    super();
    this.state = {
      modalPresent: true
    }
    this.initiateSockets();
    this.sendREADY = this.sendREADY.bind(this);
    this.registerHandler = this.registerHandler.bind(this);
  }
  render() {
    return (
      <div id="root">
        <CanvasContainer registerHandler={this.registerHandler}/>
        {this.state.modalPresent? <Modal on_sendREADY={this.sendREADY}/> : null}
      </div>
    );
  }
  initiateSockets(){
    this._socket = glue();
      this._socket.onMessage((data)=>{
        if (handlers[data]){
            handlers[data]();
        }
        if (data == "BEGIN"){
            this.setState({modalPresent: false});
        }
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
  registerHandler(message, handler){
    handlers[message] = handler;
  }
  sendREADY(){
      this._socket.send("READY");
  }
}

ReactDOM.render(<App/>, document.getElementById('app'));