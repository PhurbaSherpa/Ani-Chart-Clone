import React, { useEffect, useState } from "react";
import Routes from "./routes";
import { BrowserRouter as Router } from "react-router-dom";
import axios from "axios";

export default function App() {
  const [animes, setAnimes] = useState([]);
  useEffect(() => {
    const getAnimes = async () => {
      const { data: animes } = await axios.get("/api/animes");
      setAnimes(animes);
    };
    if (!animes.length) {
      getAnimes();
    }
  }, [animes]);
  return (
    <Router>
      <Routes animes={animes} />
    </Router>
  );
}
