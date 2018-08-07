import React from "react";
import Button from "./Button.jsx";

class ModalContent extends React.Component {
    constructor(props) {
        super();
        this.state = {
            message: "websnake.",
            buttonVisible: true
        };

        this.startButtonPress = this.startButtonPress.bind(this);
        this.onSTATUS = this.onSTATUS.bind(this);
        props.registerHandler("STATUS", this.onSTATUS);
    }
    render() {
        var style = {
            color:"#FFFFFF",
            fontFamily: "Courier New, Courier, monospace"
        };
        return (
            <div id="content" style={style}>
                <h1>{this.state.message}</h1>
                {
                    this.state.buttonVisible? 
                    <Button label={"Start a Game"} onClick={this.startButtonPress} /> :
                    null
                }
            </div>
        );
    }

    startButtonPress(){
        this.setState({
            message: "Waiting...",
            buttonVisible: false
        });
        this.props.send("STATUS", "READY");
    }

    onSTATUS(data){
        switch(data) {
        case "OPPONENT_FOUND":
            this.setState({
                message: "Game beginning in 3 seconds.",
                buttonVisible: false
            })
            setTimeout(()=>{
                this.setState({
                message: "Game beginning in 2 seconds."
                })
            }, 1000);
            setTimeout(()=>{
                this.setState({
                message: "Game beginning in 1 seconds."
                })
            }, 2000);
            break;
        case "GAMEOVER":
            this.setState({
                message: "Game over."
            })
        } 
    }
}

export default ModalContent;