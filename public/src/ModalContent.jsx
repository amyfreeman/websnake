import React from "react";

class ModalContent extends React.Component {
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
        <h1>Hello</h1>
      </div>
    );
  }
}

export default ModalContent;