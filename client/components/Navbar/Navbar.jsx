import React from "react";
import "./Navbar.css";
import { Link } from "react-router-dom";

export default function Navbar() {
  return (
    <div>
      <nav className="navbar">
        <Link to="/" className="navbar-brand">
          <div className="logo">
            <div className="logo-firsthalf">ANI</div>
            <div className="logo-secondhalf">CHART</div>
          </div>
        </Link>
        <div className="nav-links">
          <div className="nav-item">
            <Link to="/animes">Browse</Link>
          </div>
          <div className="nav-item">
            <Link to="/">Winter</Link>
          </div>
          <div className="nav-item">
            <Link to="/">Spring</Link>
          </div>
          <div className="nav-item">
            <Link to="/">Summer</Link>
          </div>
          <div className="nav-item">
            <Link to="/">Fall</Link>
          </div>
        </div>
      </nav>
    </div>
  );
}
