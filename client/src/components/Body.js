import React, { Component } from "react";
import '../App.css';
import Button from './Button';
import './Body.css';

class Body extends Component {
  constructor(props){
    super(props);
    this.state = {
        name: ""
      }
    }

  render(){
    return (
      <div className='body-container'>
        <h1>Order</h1>
        <div className='input-areas'>
          <form>
            <input
              className='body-input'
              name='code'
              type='text'
              placeholder='Your code'
            />
            <Button buttonStyle='btn--outline'>Create / Enter an order</Button>
          </form>
        </div>
      </div>
    );
  }
}

export default Body;
