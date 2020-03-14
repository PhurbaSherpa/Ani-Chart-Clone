import React from "react";
import SingleAnime from "../SingleAnime/SingleAnime";

export default function AllAnimes(props) {
  const { animes } = props;

  return (
    <div>
      {animes.map(anime => {
        return <SingleAnime anime={anime} />;
      })}
    </div>
  );
}
