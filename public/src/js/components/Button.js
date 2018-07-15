import React, { Component } from "react";
import ReactDOM from "react-dom";

function Button(props){
    return (
        <button type="button" onClick={props.onClick}>{props.label}</button>
    );
}

export default Button;