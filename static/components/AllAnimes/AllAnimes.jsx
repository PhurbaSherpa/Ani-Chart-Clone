import React from "react";
import SingleAnime from "../SingleAnime/SingleAnime";
import { Navbar } from "..";

export default function AllAnimes(props) {
  const { animes } = props;
  return (
    <div>
      <Navbar url={props.match.path} />
      <div className="list-container">
        <div className="anime-list">
          {animes.map((anime, index) => {
            return <SingleAnime key={index} anime={anime} />;
          })}
        </div>
      </div>
    </div>
  );
}
