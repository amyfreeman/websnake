import React from "react";
import css from './HomeModal.css';
import Typography from '@material-ui/core/Typography';
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
};

class WaitingModal extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            message: "WAITING...",
            updateMessageIntervalId: null,
        };
    }

    render() {
        return (
            <div className={this.props.classes.div}>
                <Typography className={this.props.classes.typography} variant="h5" component="h3">
                    {this.state.message}
                </Typography>
            </div>
        );
    }

    componentDidMount() {
        var updateMessageIntervalId = setInterval(()=>{
            this.updateMessage()
        }, 1000);
        this.setState({
            message: "WAITING.\u00A0\u00A0",
            updateMessageIntervalId: updateMessageIntervalId,
        });
    }

    componentWillUnmount() {
        clearInterval(this.state.updateMessageIntervalId);
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
                return;
        }
    }
}

export default withStyles(styles)(WaitingModal);