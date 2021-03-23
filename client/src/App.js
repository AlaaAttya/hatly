import React, { Component } from "react";
import PropTypes from "prop-types";
import Home from './components/pages/Home';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';

import NavigationBar from './components/NavigationBar';
import './App.css';
import tree from "./tree";
import initClient from "./actions/appActions"

class App extends Component {
  static propTypes = {
    dispatch:     PropTypes.func,
    history:      PropTypes.object,
    translations: PropTypes.object,
    loading:      PropTypes.bool
}

constructor(props) {
    super(props);

    initClient();

    this.state = {
      language:     tree.get("language"),
      translations:   tree.get("translations")
  };
}


  render() {
    console.log(this.state);

    return (
      <div>
        <Router>
          <NavigationBar />
          <Switch>
            <Route path='/' exact component={Home} />
          </Switch>
        </Router>
      </div>
    );
}
}
export default App;
