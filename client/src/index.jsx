import React, { Component } from 'react';
import ReactDOM from 'react-dom';

class App extends Component {
  render() {
    return (
      <h1>Hello {this.props.name}</h1>
    );
  }
}
ReactDOM.render(<App name="world" />, document.getElementById('root'));
