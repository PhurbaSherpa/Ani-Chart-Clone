import React, { useEffect, useState } from "react";
import "./SeasonList.css";
import { Navbar, SingleAnime } from "../../components";

export default function SeasonList(props) {
  const season = props.match.path.slice(1);
  const allAnimes = props.animes;
  const [animes, setAnimes] = useState([]);

  useEffect(() => {
    const getAnimesBySeason = async season => {
      let filtered = allAnimes.filter(anime => {
        return anime.Season === season;
      });
      setAnimes(filtered);
    };

    getAnimesBySeason(season);

    console.log(animes);
  }, [season]);

  return (
    <div>
      <Navbar />
      <h1>{season}</h1>
      <div className="list-container">
        {animes.length ? (
          animes.map((anime, index) => {
            return <SingleAnime key={index} anime={anime} />;
          })
        ) : (
          <div>No animes this season</div>
        )}
      </div>
    </div>
  );
}
