import React from "react";
import { Navbar } from "./components";
import Routes from "./routes";
import { BrowserRouter as Router } from "react-router-dom";

export default function App() {
  return (
    <Router>
      <Navbar />
      <Routes />
    </Router>
  );
}
