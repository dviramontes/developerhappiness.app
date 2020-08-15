import React, { useState, useCallback, useEffect } from "react";
import { format } from "timeago.js";
import { useInterval } from "./hooks/useInterval";
import "./App.css";

interface IUser {
  name: string;
  active: boolean;
  isBot: boolean;
  email: string;
  imgUrl: string;
  timezone: string;
  isAdmin: boolean;
  isOwner: boolean;
}

const firstRender = Date.now();

const UserRow = ({
  name,
  active,
  isBot,
  email,
  imgUrl,
  timezone,
  isAdmin,
  isOwner,
}: IUser) => (
  <div className="row">
    <p>{name}</p>
    {active ? <CheckboxEmoji /> : <XboxEmoji />}
    {isBot ? <BotEmoji /> : <XboxEmoji />}
    <EmailButton email={email} />
    <p>{timezone}</p>
    <p>
      { imgUrl === "" ? "n/a" : <img src={imgUrl} alt="profile" />}
    </p>
    {isAdmin ? <CheckboxEmoji /> : <XboxEmoji />}
    {isOwner ? <CheckboxEmoji /> : <XboxEmoji />}
  </div>
);

export default function App() {
  let baseEndpoint: string;

  const refreshInterval: number = 1000 * 10;
  const [users, setUsers] = useState([]);
  const [refresh, setRefresh] = useState(firstRender);

  if (process.env.NODE_ENV === "production") {
    baseEndpoint = "";
  } else {
    baseEndpoint = "http://localhost:3000";
  }

  const fetchUsers = useCallback(async () => {
    const res = await fetch(`${baseEndpoint}/api/users`);
    const users = await res.json();
    setUsers(users);
  }, []);

  useEffect(() => {
    fetchUsers();
  }, []);

  useInterval(async () => {
    // TODO: last refreshed / time ago
    // const diff: number = Date.now() - firstRender;
    // setRefresh(Date.now() - diff);
    await fetchUsers();
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
        {users.map((user: any) => (
          <UserRow key={user.ID} {...user} />
        ))}
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
