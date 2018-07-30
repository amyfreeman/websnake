import React from "react";
import ModalContent from "./ModalContent.jsx";

class Modal extends React.Component {
  constructor() {
    super();
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
      display: "flex",
      justifyContent: "center", alignItems: "center"
    };
    return (
      <div id="modal" style={style}>
        <ModalContent />
      </div>
    );
  }
}
export default Modal;