import React, { Component } from "react";
import ReactDOM from "react-dom";

import Modal from "./Modal.js"
import CanvasContainer from "./CanvasContainer.js";

class Root extends Component {
  constructor() {
    super();
  }
  render() {
    return (
      <div id="root">
      <CanvasContainer />
      <Modal />
      </div>
    );
  }
}
export default Root;

const wrapper = document.getElementById("react-root");
wrapper ? ReactDOM.render(<Root />, wrapper) : false;