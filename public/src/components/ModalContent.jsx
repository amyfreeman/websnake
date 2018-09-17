import React from "react";
import Button from "./Button.jsx";
import css from './ModalContent.css';

class ModalContent extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            message: "WEBSNAKE",
            inputsVisible: true,
            nicknameInputVisible: true,
            opacity: 1,
            startButtonEnabled: true,
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
                    this.state.inputsVisible?
                    <div style={{width: "100%"}}>
                        {this.state.nicknameInputVisible?
                        <input id="nickname-input" className="nickname-input" style={inputStyle} type="text" placeholder="Nickname"></input> :
                        null}
                        <Button label={"Start a Game"} onClick={this.startButtonPress} />
                    </div> :
                    null
                }
            </div>
        );
    }

    componentDidMount() {
        this.registerInputKeyupListener();
        this.setInputFocused();
    }

    componentDidUpdate() {
        if (this.state.inputsVisible && this.state.nicknameInputVisible){
            this.registerInputKeyupListener();
            this.setInputFocused();
        }
    }

    registerInputKeyupListener() {
        document.getElementById("nickname-input").addEventListener("keyup", (event) => {
            event.preventDefault();
            if (event.keyCode === 13) {
                this.startButtonPress();
            }
        });
    }

    setInputFocused() {
        document.getElementById("nickname-input").focus();
    }

    startButtonPress(){
        console.log(this.state.startButtonEnabled);
        if (this.state.startButtonEnabled){
            this.setState({
                opacity: 0,
                startButtonEnabled: false,
            });
            setTimeout(()=>{
                var updateMessageIntervalId = setInterval(()=>{
                    this.updateMessage()
                }, 1000);
                this.setState({
                    message: "WAITING.\u00A0\u00A0",
                    inputsVisible: false,
                    opacity: 1,
                    updateMessageIntervalId: updateMessageIntervalId,
                    startButtonEnabled: true,
                });
            }, 1000);
            if (this.state.nicknameInputVisible){
                this.props.send("NICKNAME", document.getElementById("nickname-input").value);
            }
            this.props.send("STATUS", "READY");
        }
    }

    onSTATUS(data){
        switch(data) {
            case "OPPONENT_FOUND":
                clearInterval(this.state.updateMessageIntervalId);
                this.setState({
                    opacity: 0
                });
                setTimeout(()=>{
                    this.setState({
                        message: "Game starting in 3",
                        inputsVisible: false,
                        opacity: 1,
                    });
                }, 1000);
                setTimeout(()=>{
                    this.setState({
                        message: "Game starting in 2"
                    })
                }, 2000);
                setTimeout(()=>{
                    this.setState({
                        message: "Game starting in 1"
                    })
                }, 3000);
                break;
            case "GAMEOVER":
                this.setState({
                    message: "Game over.",
                    inputsVisible: true,
                    nicknameInputVisible: false,
                });
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