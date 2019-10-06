import React from 'react';
import logo from './logo.svg';
import './App.scss';

function App() {
  return (
    <div className="App section">
      <figure className="image is-128x128">
        <img className="image is-rounded" src={logo} alt="logo" />
      </figure>
      <h1>Hello CodeSandbox</h1>
      <h2>Start editing to see some magic happen!</h2>

      <button className="button is-danger is-outlined">
        Hello
      </button>
    </div>
  );
}

export default App;
