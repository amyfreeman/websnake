import React, { Component } from "react";
import ReactDOM from "react-dom";
import Modal from "./Modal.js"

class Root extends Component {
  constructor() {
    super();
  }
  render() {
    return (
      <div id="root">
        <Modal startGame={this.startGame}/>
      </div>
    );
  }
  startGame() {
    console.log("booyeah");
  }
}
export default Root;

const wrapper = document.getElementById("react-root");
wrapper ? ReactDOM.render(<Root />, wrapper) : false;