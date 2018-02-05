import React, { Component } from "react";
import ReactDOM from "react-dom";
import Button from "./Button.js"

class ModalContent extends Component {
  constructor() {
    super();
  }
  render() {
    var style = {
      color:"#FFFFFF",
      marginRight:"200px",
      marginBottom:"200px"
    };
    return (
      <div id="content" style={style}>
        <h1>Welcome to WebSnake</h1>
        <p>The best game ever</p>
        <Button label={"words"} clickFunction={tempFunction} />
        <Button label={"words3"} clickFunction={tempFunction}/>
      </div>
    );
  }
}

function tempFunction(){
  console.log(10);
}

export default ModalContent;