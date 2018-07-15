class IO{
    constructor(socketConnection){
        this.socketConnection = socketConnection;

        window.addEventListener('keydown', function(event) {
            switch (event.keyCode) {
                case 37:
                    leftPress();
                    break;
                case 38:
                    upPress();
                    break;
                case 39:
                    rightPress();
                    break;
                case 40:
                    downPress();
            }
        }, false);
    }

    leftPress(){
        Console.log("left press registered");
        this.socketConnections.send("left");
    }

    rightPress(){
        this.socketConnections.send("right");
    }

    upPress(){
        this.socketConnections.send("up");
    }

    downPress(){
        this.socketConnections.send("down");
    }
}
  
export {IO};