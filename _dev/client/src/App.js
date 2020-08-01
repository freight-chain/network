import React from 'react';
import logo from './logo.svg';
import './App.css';

class App extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      value: '',
    }

    this.ws = new WebSocket(process.env.REACT_APP_WS_URL)

    this.ws.onopen = e => {
      this.ws.send(`{"action": "subscribe", "id": "foo"}`)
    }

    this.ws.onmessage = e => {
      this.handleWSEvent(e)
    }
  }

  updateInput = (e) => {
    const value = e.target.value;
    this.setState({value});

    this.ws.send(`{"action": "update", "id": "${e.target.id}", "value": "${value}"`)
  }

  handleWSEvent = (e) => {
    window.console.log('received message on ws', e.data);
    this.setState({
      value: e.data.value,
    })
  }

  render () {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Open two browsers and make changes to the input, below.
          </p>
          <input
            id="foo"
            type="text"
            value={this.state.value}
            onChange={this.updateInput}
            placeholder='enter text, here...'
          >
          </input>
        </header>
      </div>
    );
  }
}

export default App;
