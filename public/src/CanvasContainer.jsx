import React from "react";

const NUM_ROWS = 10;
const NUM_COLS = 10;
const BOARD_SIZE = 0.8 * Math.min(window.innerWidth, window.innerHeight);
const CELL_WIDTH = BOARD_SIZE / NUM_COLS;
const CELL_HEIGHT = BOARD_SIZE / NUM_ROWS;

var ctx;

class CanvasContainer extends React.Component {
  constructor() {
    super();

    this.props.registerHandler("startGame", this.startGame);
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
      this.drawBoard();
  }
  drawBoard(){
    const canvas = this.refs.canvas;
    ctx = canvas.getContext("2d");
    ctx.canvas.width  = BOARD_SIZE;
    ctx.canvas.height = BOARD_SIZE;

    ctx.fillStyle = "#000000";
    ctx.fillRect(0, 0, BOARD_SIZE, BOARD_SIZE);
    
    ctx.strokeStyle = "#00FF00";
    ctx.lineWidth = 10;
    ctx.beginPath();
    for (var i = 0; i <= NUM_ROWS; i++){
        ctx.moveTo(0, i * CELL_HEIGHT);
        ctx.lineTo(BOARD_SIZE, i * CELL_HEIGHT);
    }
    for (var i = 0; i <= NUM_COLS; i++){
        ctx.moveTo(i * CELL_WIDTH, 0);
        ctx.lineTo(i * CELL_WIDTH, BOARD_SIZE); 
    }
    ctx.stroke();
  }

  startGame(){
    console.log("game is starting, for real");
  }
}
export default CanvasContainer;