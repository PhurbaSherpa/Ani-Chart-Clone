import React from "react";
import SingleAnime from "../SingleAnime/SingleAnime";
import Navbar from "../Navbar/Navbar";
import "./AllAnimes.css";

export default function AllAnimes(props) {
  const { animes } = props;
  return (
    <div className="all-animes">
      <Navbar url={props.match.path} />
      <div className="list-container">
        {animes.map((anime, index) => {
          return <SingleAnime key={index} anime={anime} />;
        })}
      </div>
    </div>
  );
}
