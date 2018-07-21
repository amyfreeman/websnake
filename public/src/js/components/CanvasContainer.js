import React, { Component } from "react";

class CanvasContainer extends Component {
  constructor() {
    super();
  }
  render() {
    var style = {
        minHeight: "100vh",
        width: "100%",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "green"
      };
    return (
        <div id="canvas-container" style={style}>
            <canvas id="maincanvas"></canvas>
        </div>
    );
  }
}
export default CanvasContainer;