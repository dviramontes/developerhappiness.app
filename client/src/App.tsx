import React, { useState, useCallback, useEffect } from "react";
import { format } from "timeago.js";
import { useInterval } from "./hooks/useInterval";
import "./App.css";

interface Row {
  user: string;
  active: boolean;
  bot: boolean;
  email: string;
  imgUrl: string;
  timezone: string;
  admin: boolean;
  owner: boolean;
}

const firstRender = Date.now();

const UserRow = ({
  bot,
  user,
  owner,
  email,
  admin,
  imgUrl,
  active,
  timezone,
}: Row) => (
  <div className="row">
    <p>{user}</p>
    {active ? <CheckboxEmoji /> : <XboxEmoji />}
    {bot ? <BotEmoji /> : <XboxEmoji />}
    <EmailButton email={email} />
    <p>{timezone}</p>
    <p>
      <img src={imgUrl} alt="profile" />
    </p>
    {admin ? <CheckboxEmoji /> : <XboxEmoji />}
    {owner ? <CheckboxEmoji /> : <XboxEmoji />}
  </div>
);

export default function App() {
  let baseEndpoint: string;

  const refreshInterval: number = 1000;
  const [users, setUsers] = useState([]);
  const [refresh, setRefresh] = useState(firstRender);

  if (process.env.NODE_ENV === "production") {
    baseEndpoint = "";
  } else {
    baseEndpoint = "http://localhost:3000";
  }

  const fetchUsers = useCallback(async () => {
    const res = await fetch(`${baseEndpoint}/ping`);
    const text = await res.text();
    console.log(text);
  }, [baseEndpoint]);

  useEffect(() => {
    fetchUsers();
  });

  useInterval(async () => {
    // TODO: fix time go
    // const diff: number = Date.now() - firstRender;
    // setRefresh(Date.now() - diff);
    fetchUsers();
  }, refreshInterval);

  return (
    <div className="App">
      <header className="App-header">
        <h1>slack user list</h1>
        <p className="App-link">Last refreshed: {format(refresh)}</p>
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
        <UserRow
          user={"david"}
          active={true}
          bot={false}
          email={"me@mail.com"}
          timezone={"America/Denver"}
          imgUrl={
            "https://secure.gravatar.com/avatar/fe5373af89a931ab1660970a9b25ff2c.jpg?s=32&d=https%3A%2F%2Fa.slack-edge.com%2Fdf10d%2Fimg%2Favatars%2Fava_0010-32.png"
          }
          admin={true}
          owner={true}
        />
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

const CheckboxEmoji = () => (
  <p>
    <span role="img" aria-label="check">
      ✅
    </span>
  </p>
);

const XboxEmoji = () => (
  <p>
    <span role="img" aria-label="x">
      ❌
    </span>
  </p>
);

const BotEmoji = () => (
  <p>
    <span role="img" aria-label="x">
      🤖
    </span>
  </p>
);

const EmailButton = ({ email }: any) => (
  <p>
    <a href={`mailto:${email}`} target="_blank" rel="noopener noreferrer">
      <span role="img" aria-label="email">
        ✉️
      </span>
    </a>
  </p>
);
