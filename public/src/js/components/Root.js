import React, { Component } from "react";
import ReactDOM from "react-dom";
import Modal from "./Modal.js";

class Root extends Component {
  constructor() {
    super();
    this.state = {
      modalVisible: true
    }
    this.startGame = this.startGame.bind(this)
  }
  render() {
    return (
      <div id="root">
        {
          this.state.modalVisible
          ? <Modal />
          : null
        }
      </div>
    );
  }
}
export default Root;

const wrapper = document.getElementById("react-root");
wrapper ? ReactDOM.render(<Root />, wrapper) : false;