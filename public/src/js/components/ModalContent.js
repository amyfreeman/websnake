import React, { Component } from "react";
import ReactDOM from "react-dom";

class ModalContent extends Component {
  constructor() {
    super();
  }
  render() {
    return (
      <div id="content" style={{color:"#FFFFFF", marginRight:"200px", marginBottom:"200px"}}>
        <h1>Welcome to WebSnake</h1>
        <p>The best game ever</p>
      </div>
    );
  }
}
export default ModalContent;