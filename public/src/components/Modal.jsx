import React from "react";
import ModalContent from "./ModalContent.jsx";

class Modal extends React.Component {
    constructor(props) {
        super();
        this.state = {
            opacity: 1
        };
        this.onSTATUS = this.onSTATUS.bind(this);
        props.registerHandler("STATUS", this.onSTATUS);
    }

    render() {
        var style = {
            position:"fixed",
            zIndex:1,
            left:"0",
            top:"0",
            width:"100%",
            height:"100%",
            overflow:"auto",
            backgroundColor:"rgba(0, 0, 0, 0.8)",
            opacity: this.state.opacity,
            display: "flex",
            justifyContent: "center", alignItems: "center",
        };
        return (
            <div id="Modal" style={style}>
                <ModalContent registerHandler={this.props.registerHandler} send={this.props.send}/>
            </div>
        );
    }

    onSTATUS(data){
        switch(data) {
            case "BEGIN":
                this.setState({opacity: 0});
                break;
            case "GAMEOVER":
                this.setState({opacity: 1});
        }
    }
}
export default Modal;