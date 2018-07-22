import React from "react";

class CanvasContainer extends React.Component {
  constructor() {
    super();
  }
  render() {
    var style = {
        minHeight: "100vh",
        width: "100%",
        display: "flex",
        justifyContent: "center",
        alignItems: "center"
      };
    return (
        <div id="canvas-container" style={style}>
            <canvas id="maincanvas" ref="canvas"></canvas>
        </div>
    );
  }
  componentDidMount() {
      console.log("hmidfmmsd");
      this.refs.canvas.style.backgroundColor = "red";
  }
}
export default CanvasContainer;