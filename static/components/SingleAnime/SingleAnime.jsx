import React from "react";
import "./SingleAnime.css";

export default function SingleAnime(props) {
  const { anime } = props;

  return (
    <div className="img-container">
      <img src={anime.Imageurl} />
    </div>
  );
}
