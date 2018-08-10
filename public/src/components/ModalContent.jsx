import React from "react";
import Button from "./Button.jsx";
import css from '../w3.css';
import css2 from '../custom.css';

class ModalContent extends React.Component {
    constructor(props) {
        super();
        this.state = {
            message: "WEBSNAKE",
            buttonVisible: true,
            opacity: 1,
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
            opacity: this.state.opacity,
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
            <div className="ModalContent" style={divStyle} id="ModalContent">
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
        setTimeout(()=>{
            this.setState({
                message: "WAITING.\u00A0\u00A0",
                buttonVisible: false,
                opacity: 1
            });
            setInterval(()=>{
                this.updateMessage()
            }, 1000);
        }, 1000);
        this.setState({
            opacity: 0
        });

        this.props.send("STATUS", "READY");
    }

    onSTATUS(data){
        switch(data) {
        case "OPPONENT_FOUND":
            this.setState({
                opacity: 0
            });
            setTimeout(()=>{
                this.setState({
                    message: "Game beginning in 3 seconds.",
                    buttonVisible: false,
                    opacity: 1,
                });
            }, 1000);
            setTimeout(()=>{
                this.setState({
                    message: "Game beginning in 2 seconds."
                })
            }, 2000);
            setTimeout(()=>{
                this.setState({
                    message: "Game beginning in 1 seconds."
                })
            }, 3000);
            break;
        case "GAMEOVER":
            this.setState({
                message: "Game over."
            })
        } 
    }
    updateMessage(){
        switch(this.state.message){
        case "WAITING...":
            this.setState({message: "WAITING.\u00A0\u00A0"});
            return;
        case "WAITING.\u00A0\u00A0":
            this.setState({message: "WAITING..\u00A0"});
            return;
        case "WAITING..\u00A0":
            this.setState({message: "WAITING..."});
        }
    }
}

export default ModalContent;