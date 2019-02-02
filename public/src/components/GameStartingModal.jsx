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

class GameStartingModal extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            message: "Game starting in 4",
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
    }
}

export default withStyles(styles)(GameStartingModal);