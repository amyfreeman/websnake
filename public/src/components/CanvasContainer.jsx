import React from "react";

const NUM_ROWS = 10;
const NUM_COLS = 10;
const BOARD_SIZE = 0.8 * Math.min(window.innerWidth, window.innerHeight);
const CELL_WIDTH = BOARD_SIZE / NUM_COLS;
const CELL_HEIGHT = BOARD_SIZE / NUM_ROWS;
var canvas;
var ctx;

class CanvasContainer extends React.Component {
  constructor(props) {
    super();

    this.drawCell = this.drawCell.bind(this);
    this.onGAMESTATE = this.onGAMESTATE.bind(this);
    props.registerHandler("GAMESTATE", this.onGAMESTATE);
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
    canvas = document.getElementById("maincanvas");
    ctx = canvas.getContext("2d");
    ctx.canvas.width  = BOARD_SIZE;
    ctx.canvas.height = BOARD_SIZE;
      this.drawBoard();
  }
  drawBoard(){
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

  drawCell(x, y, color){
    ctx.fillStyle = color;
    ctx.fillRect(x * CELL_WIDTH, y * CELL_HEIGHT, CELL_WIDTH, CELL_HEIGHT, color);
  }

  onGAMESTATE(data){
    console.log(data);
    data = data.split("");
    for (var i = 0; i < NUM_COLS; i++){
        for (var j = 0; j < NUM_ROWS; j++){
            var c = data[i * NUM_COLS + j];
            if (c == "."){
                this.drawCell(i, j, "black");
            }
            else if (c == "F"){
                this.drawCell(i, j, "red");
            }
            else if (c == "0"){
                this.drawCell(i, j, "green");
            }
            else if (c == "1"){
                this.drawCell(i, j, "blue");
            }
            else{
                this.drawCell(i, j, "yellow");
            }
        }
    }
    this.drawBoard();
  }
}
export default CanvasContainer;