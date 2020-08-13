import React from 'react';
import './App.css';

export default function App() {
  let baseEndpoint = null;

  if (process.env.NODE_ENV === 'production') {
    baseEndpoint = ""
  } else {
    baseEndpoint = "http://localhost:3000"
  }

  fetch(`${baseEndpoint}/ping`)
    .then(res => res.text())
    .catch(err => console.error(err))
    .then(res => console.log(res))

  return (
    <div className="App">
      <header className="App-header">
        <h1>slack user list</h1>
        <h2></h2>
        <div className="row">
          <p className="col">user</p>
          <p className="col">active</p>
          <p className="col">bot</p>
          <p className="col">email</p>
          <p className="col">timezone</p>
          <p className="col">image</p>
          <p className="col">admin</p>
          <p className="col">owner</p>
        </div>
        <div className="row">
          <p>david</p>
          <p>✅</p>
          <p>🤖</p>
          <p>✉️</p>
          <p>America/Denver</p>
          <p><img src="https://secure.gravatar.com/avatar/fe5373af89a931ab1660970a9b25ff2c.jpg?s=32&d=https%3A%2F%2Fa.slack-edge.com%2Fdf10d%2Fimg%2Favatars%2Fava_0010-32.png" alt="profile-picture"/></p>
          <p>✅</p>
          <p>✅</p>
        </div>
        <div className="row">
          <p>david</p>
          <p>❌</p>
          <p>🤖</p>
          <p>✉️</p>
          <p>America/Denver</p>
          <p><img src="https://secure.gravatar.com/avatar/fe5373af89a931ab1660970a9b25ff2c.jpg?s=32&d=https%3A%2F%2Fa.slack-edge.com%2Fdf10d%2Fimg%2Favatars%2Fava_0010-32.png" alt="profile-picture"/></p>
          <p>✅</p>
          <p>✅</p>
        </div>
        <div className="row">
          <p>bar</p>
        </div>
        <div className="row">
          <p>bar</p>
        </div>
      </header>
    </div>
  );
}
