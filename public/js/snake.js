//game variables
var GAME_WIDTH = 10;
var GAME_HEIGHT = 10;
var NUM_SNAKES = 2;
var NUM_FOODS = 1;

//gui variables
var TOTAL_WIDTH = 200;
var TOTAL_HEIGHT = 200;

var COLUMN_WIDTH = TOTAL_WIDTH / GAME_WIDTH;
var ROW_HEIGHT = TOTAL_HEIGHT / GAME_HEIGHT

function drawField(){
    ctx.moveTo(0, 0);
    ctx.lineTo(100, 100);
    ctx.stroke();
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
            else{
                ctx.fillStyle = "green";
            }
            ctx.rect(i * COLUMN_WIDTH, ROW_HEIGHT * GAME_HEIGHT - (j + 1)* ROW_HEIGHT, COLUMN_WIDTH, ROW_HEIGHT);
            ctx.fill();
        }
    }
}