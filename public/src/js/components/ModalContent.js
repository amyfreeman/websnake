import React, { Component } from "react";
import Button from "./Button.js"
import {io} from "../IO.js";

class ModalContent extends Component {
  constructor() {
    super();
    this.state = {
      message: "Welcome",
      buttonVisible: "false"
    }
    this.startButtonPress = this.startButtonPress.bind(this)
  }
  render() {
    var style = {
      color:"#FFFFFF",
      marginRight:"200px",
      marginBottom:"200px"
    };
    return (
      <div id="content" style={style}>
        <h1>{this.state.message}</h1>
        {
          this.state.buttonVisible? 
          <Button label={"Start a Game"} onClick={this.startButtonPress} /> :
          null
        }
      </div>
    );
  }
  startButtonPress(){
    this.setState({message: "Waiting...", buttonVisible: false})
    io.startButtonPress();
  }
}

export default ModalContent;