import React from "react";
import "./LatestAnime.css";

export default function SingleAnime(props) {
  const { anime } = props;

  return (
    <div className="top3-container">
      <div className="single-row">
        <img src={anime.Imageurl} />
        <div className="top3info-container">
          <div className="top3-title">{anime.Title}</div>
          <div className="top3-date">
            Coming this {anime.Season}, {anime.Date}
          </div>
        </div>
      </div>
    </div>
  );
}
