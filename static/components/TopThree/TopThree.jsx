import React from "react";
import LatestAnime from "../LatestAnime/LatestAnime";
import "./TopThree.css";

export default function TopThree(props) {
  const { animes } = props;
  return (
    <div className="Topthree-container">
      <h1 className="top3-description">TOP ANIMES OUT RIGHT NOW</h1>
      <div className="top3-bg">
        <div className="top3-list">
          {animes.map((anime, index) => {
            return <LatestAnime key={index} anime={anime} />;
          })}
        </div>
      </div>
    </div>
  );
}
