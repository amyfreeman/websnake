import React, { Component } from "react";
import ReactDOM from "react-dom";
import Modal from "./Modal.js";
import {startGame3} from "../controller.js";

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
          ? <Modal startGame={this.startGame}/>
          : null
        }
      </div>
    );
  }
  startGame() {
    this.setState(prevState => ({ modalVisible: false }));
    console.log(1);
    startGame3();
    console.log(2);
  }
}
export default Root;

const wrapper = document.getElementById("react-root");
wrapper ? ReactDOM.render(<Root />, wrapper) : false;