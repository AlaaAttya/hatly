import React, { Component } from "react";
import "./Button.css"
class Button extends Component {

    onClick(e) {
      return alert("Hello");
    }

    render() {
        return (
            <button className= "button"
            onClick={this.onClick}
            >
            <span>Order</span>
            </button>
        );
    }
}

export default Button;