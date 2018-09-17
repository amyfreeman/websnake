import React from "react";

function Button(props){
    return (
        <button className="btn btn-default" type="button" onClick={props.onClick}>
            {props.label}
        </button>
    );
}

export default Button;