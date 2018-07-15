import React, { Component } from "react";
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
        <Button label={"Start a Game"} onClick={this.props.startGame} />
      </div>
    );
  }
}

export default ModalContent;