import React from "react";
import css from './HomeModal.css';
import Button from '@material-ui/core/Button';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import { withStyles } from '@material-ui/core/styles';

const styles = {
    typography: {
        color:"#FFFFFF",
        fontSize: "5em",
        fontFamily: "Courier New, Courier, monospace",
    },
    div: {
        display: "flex",
        justifyContent: "center", alignItems: "center",
        flexDirection: "column",
    },
    paper: {
        width: "100%",
        display: "flex",
        justifyContent: "center", alignItems: "center",
    },
    textField: {
        marginRight: "15px",
    },
};

class HomeModal extends React.Component {
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
        return (
            <div className={this.props.classes.div}>
                <Typography className={this.props.classes.typography} variant="h5" component="h3">
                    WEBSNAKE
                </Typography>
                <Paper className={this.props.classes.paper}>
                    {
                        this.state.nicknameInputVisible ?
                        <TextField margin="normal" label="Nickname" className={this.props.classes.textField}/> :
                        null
                    }
                    <Button size="large" variant="contained" color="primary">
                        BEGIN
                    </Button>
                </Paper>
            </div>
        );
    }

    componentDidMount() {
        //this.registerInputKeyupListener();
        //this.setInputFocused();
    }

    componentDidUpdate() {
        if (this.state.inputsVisible && this.state.nicknameInputVisible){
            //this.registerInputKeyupListener();
            //this.setInputFocused();
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

export default withStyles(styles)(HomeModal);