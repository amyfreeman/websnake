import React from "react";

function Button(props){
    return (
        <button className="w3-button w3-white" type="button" onClick={props.onClick}>
            {props.label}
        </button>
    );
}

export default Button;