import React, { Component } from "react";
import ReactDOM from "react-dom";
import ModalContent from "./ModalContent.js"

class Modal extends Component {
  constructor() {
    super();
  }
  render() {
    var style = {
      position:"fixed",
      zIndex:1,
      left:"0",
      top:"0",
      width:"100%",
      height:"100%",
      overflow:"auto",
      backgroundColor:"rgba(0, 0, 0, 0.8)",
      display: "flex",
      justifyContent: "center", alignItems: "center"
    };
    return (
      <div id="modal" style={style}>
        <ModalContent startGame={this.startGame} />
      </div>
    );
  }
  startGame() {
    console.log("booyeah");
  }
}
export default Modal;

const wrapper = document.getElementById("react-root");
wrapper ? ReactDOM.render(<Modal />, wrapper) : false;