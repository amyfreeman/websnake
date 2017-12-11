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
    //assert(msg.length == GAME_WIDTH * GAME_HEIGHT);
    for (var i = 0; i < GAME_WIDTH; i++){
        for (var j = 0; j < GAME_HEIGHT; j++){
            ctx.beginPath();
            var c = msg[i * GAME_WIDTH + j]
            if (c == "."){
                ctx.fillStyle = "black";
            }
            else if (c == "F"){
                ctx.fillStyle = "red";
            }
            else if (c == "0"){
                ctx.fillStyle = "green";
            }
            else if (c == "1"){
                ctx.fillStyle = "blue";
            }
            else{
                ctx.fillStyle = "yellow";
            }
            ctx.rect(i * COLUMN_WIDTH, ROW_HEIGHT * GAME_HEIGHT - (j + 1)* ROW_HEIGHT, COLUMN_WIDTH, ROW_HEIGHT);
            ctx.fill();
        }
    }
}