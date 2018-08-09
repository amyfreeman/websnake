import React from "react";
import Button from "./Button.jsx";
import css from '../w3.css';

class ModalContent extends React.Component {
    constructor(props) {
        super();
        this.state = {
            message: "WEBSNAKE",
            buttonVisible: true
        };

        this.startButtonPress = this.startButtonPress.bind(this);
        this.onSTATUS = this.onSTATUS.bind(this);
        props.registerHandler("STATUS", this.onSTATUS);
    }
    render() {
        var divStyle = {
            display: "flex",
            justifyContent: "center", alignItems: "center",
            flexDirection: "column",
            fontFamily: "Courier New, Courier, monospace",
        }
        var h1Style = {
            color:"#FFFFFF",
            fontSize: "5em",
        };
        var inputStyle = {
            marginBottom: "10px",
            fontSize: "1em",
            lineHeight: "1.5em",
        };
        return (
            <div style={divStyle} id="ModalContent">
                <h1 style={h1Style}>{this.state.message}</h1>
                {
                    this.state.buttonVisible?
                    <div style={{width: "100%"}}>
                        <input className="w3-input" style={inputStyle} type="text" placeholder="Nickname"></input>
                        <Button label={"Start a Game"} onClick={this.startButtonPress} />
                    </div> :
                    null
                }
            </div>
        );
    }

    startButtonPress(){
        this.setState({
            message: "WAITING...",
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