import React from "react";
import Button from "./Button.jsx";

class ModalContent extends React.Component {
  constructor() {
    super();
    this.state = {
      message: "Welcome",
      buttonVisible: true
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
    this.setState({
      message: "Waiting...",
      buttonVisible: false
    })
  }
}

export default ModalContent;