//game variables
var GAME_WIDTH = 10;
var GAME_HEIGHT = 10;
var NUM_SNAKES = 2;
var NUM_FOODS = 1;

//gui variables
var CANVAS_WIDTH = 200;
var CANVAS_HEIGHT = 200;

var COLUMN_WIDTH = CANVAS_WIDTH / GAME_WIDTH;
var ROW_HEIGHT = CANVAS_HEIGHT / GAME_HEIGHT

function initializeCanvas(){
    canvas = document.querySelector("#back");
    ctx = canvas.getContext("2d");
    ctx.canvas.width  = CANVAS_WIDTH;
    ctx.canvas.height = CANVAS_HEIGHT;
}

function drawFromMsg(msg){
    console.log(msg);
    msg = msg.split("");
    for (var i = 0; i < GAME_WIDTH; i++){
        for (var j = 0; j < GAME_HEIGHT; j++){
            var c = msg[i * GAME_WIDTH + j]
            switch (c) {
                case ".":
                    ctx.fillStyle = "black";
                    break;
                case "F":
                    ctx.fillStyle = "red";
                    break;
                case "0":
                    ctx.fillStyle = "green";
                    break;
                case "1":
                    ctx.fillStyle = "blue";
                    break;
                default:
                    ctx.fillStyle = "yellow";
            }
            ctx.fillRect(i * COLUMN_WIDTH, ROW_HEIGHT * GAME_HEIGHT - (j + 1)* ROW_HEIGHT, COLUMN_WIDTH, ROW_HEIGHT);
        }
    }
}