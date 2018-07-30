import React from "react";

const NUM_ROWS = 10;
const NUM_COLS = 10;
const BOARD_SIZE = 0.8 * Math.min(window.innerWidth, window.innerHeight);
const CELL_WIDTH = BOARD_SIZE / NUM_COLS;
const CELL_HEIGHT = BOARD_SIZE / NUM_ROWS;

var ctx;

class CanvasContainer extends React.Component {
  constructor(props) {
    super();

    props.registerHandler("BEGIN", this.startGame);
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
    this._ctx = canvas.getContext("2d");
    this._ctx.canvas.width  = BOARD_SIZE;
    this._ctx.canvas.height = BOARD_SIZE;

    this._ctx.fillStyle = "#000000";
    this._ctx.fillRect(0, 0, BOARD_SIZE, BOARD_SIZE);
    
    this._ctx.strokeStyle = "#00FF00";
    this._ctx.lineWidth = 10;
    this._ctx.beginPath();
    for (var i = 0; i <= NUM_ROWS; i++){
        this._ctx.moveTo(0, i * CELL_HEIGHT);
        this._ctx.lineTo(BOARD_SIZE, i * CELL_HEIGHT);
    }
    for (var i = 0; i <= NUM_COLS; i++){
        this._ctx.moveTo(i * CELL_WIDTH, 0);
        this._ctx.lineTo(i * CELL_WIDTH, BOARD_SIZE); 
    }
    this._ctx.stroke();
  }

  drawCell(x, y, color){
    this._ctx.fillStyle = color;
    this._ctx.fillRect(x * CELL_WIDTH, y * CELL_HEIGHT, CELL_WIDTH, CELL_HEIGHT, color);
  }

  startGame(){
    console.log("game is starting, for real");
  }
}
export default CanvasContainer;