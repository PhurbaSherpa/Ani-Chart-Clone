import React, { useEffect } from "react";
import "./Navbar.css";
import { Link } from "react-router-dom";

export default function Navbar(props) {
  useEffect(() => {
    const componenet = document.getElementById("navbar");
    if (props.url !== "/") {
      componenet.classList.add("bg-black");
    } else {
      componenet.classList.remove("bg-black");
    }
  }, [props.url]);

  return (
    <div>
      <nav id="navbar">
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
            <Link to="/Winter">Winter</Link>
          </div>
          <div className="nav-item">
            <Link to="/Spring">Spring</Link>
          </div>
          <div className="nav-item">
            <Link to="/Summer">Summer</Link>
          </div>
          <div className="nav-item">
            <Link to="/Fall">Fall</Link>
          </div>
        </div>
      </nav>
    </div>
  );
}
