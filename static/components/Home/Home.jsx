import React from "react";
import "./home.css";
import TopThree from "../TopThree/TopThree";
import Navbar from "../Navbar/Navbar";

export default function Home(props) {
  let animes =
    props.animes.length >= 3 ? props.animes.slice(0, 3) : props.animes;

  return (
    <div>
      <div className="hero">
        <Navbar url={props.match.path} />
        <div className="home-container">
          <div className="home-description">
            Discover, Explore and Experience Anime
          </div>
        </div>
        <div className="icon">
          <a href="#topThree">
            <i className="fas fa-angle-double-down"></i>
          </a>
        </div>
      </div>
      <section id="topThree">
        <TopThree animes={animes} />
      </section>
    </div>
  );
}
