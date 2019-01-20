import React from "react";

const NUM_ROWS = 10;
const NUM_COLS = 10;
const BOARD_SIZE = 0.8 * Math.min(window.innerWidth, window.innerHeight);
const CELL_WIDTH = BOARD_SIZE / NUM_COLS;
const CELL_HEIGHT = BOARD_SIZE / NUM_ROWS;
const GRID_COLOR = "#00FF00";
const MY_COLOR = "white";
const OPPONENT_COLOR = "blue";
const FOOD_COLOR = "red";
const BG_COLOR = "black";
const UNKNOWN_COLOR = "yellow";


var canvas;
var ctx;

class CanvasLayer extends React.Component {
    constructor(props) {
        super(props);

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
            <div id="canvas-layer" style={style}>
                <canvas id="main-canvas" ref="canvas"></canvas>
            </div>
        );
    }

    componentDidMount() {
        canvas = document.getElementById("main-canvas");
        ctx = canvas.getContext("2d");
        ctx.canvas.width  = BOARD_SIZE;
        ctx.canvas.height = BOARD_SIZE;
        this.drawGrid();

        window.addEventListener('keydown', (event) => {
            switch (event.keyCode) {
                case 37: // Left
                    this.props.send("GAMEPLAY", "LEFT");
                    break;
                case 38: // Up
                    this.props.send("GAMEPLAY", "UP");
                    break;
                case 39: // Right
                    this.props.send("GAMEPLAY", "RIGHT");
                    break;
                case 40: // Down
                    this.props.send("GAMEPLAY", "DOWN");
                    break;
            }
        }, false);
        this.drawStartingPositions();
    }

    drawGrid(){
        ctx.strokeStyle = GRID_COLOR;
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
        // Remember to account for canvas origin at upper-left
        ctx.fillRect(x * CELL_WIDTH, (NUM_ROWS - y - 1) * CELL_HEIGHT, CELL_WIDTH, CELL_HEIGHT, color);
    }

    onGAMESTATE(data){
        console.log(data);
        data = data.split("");
        for (var i = 0; i < NUM_COLS; i++){
            for (var j = 0; j < NUM_ROWS; j++){
                var c = data[i * NUM_COLS + j];
                if (c == "."){
                    this.drawCell(i, j, BG_COLOR);
                }
                else if (c == "F"){
                    this.drawCell(i, j, FOOD_COLOR);
                }
                else if (c == "0"){
                    this.drawCell(i, j, MY_COLOR);
                }
                else if (c == "1"){
                    this.drawCell(i, j, OPPONENT_COLOR);
                }
                else{
                    this.drawCell(i, j, UNKNOWN_COLOR);
                }
            }
        }
        this.drawGrid();
    }

    drawStartingPositions(){
        this.drawCell(0, 0, MY_COLOR);
        this.drawCell(NUM_COLS - 1, NUM_ROWS - 1, OPPONENT_COLOR);
        this.drawGrid();
    }
}
export default CanvasLayer;