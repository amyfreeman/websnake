import React from "react";

function Button(props){
    var style = {
        fontFamily: "Courier New, Courier, monospace"
    };
    return (
        <button type="button" onClick={props.onClick} style={style}>
            {props.label}
        </button>
    );
}

export default Button;