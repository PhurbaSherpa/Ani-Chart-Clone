import React, { useEffect, useState } from "react";
import { Navbar, SingleAnime } from "../../components";
import "./GenreList.css";
import axios from "axios";

export default function GenreList(props) {
  const { genre } = props.match.params;
  const [animes, setAnimes] = useState([]);

  useEffect(() => {
    const getAnimes = async genre => {
      const { data } = await axios.get(`/api/animes/genre/${genre}`);
      console.log(data);
      setAnimes(data);
    };
    getAnimes(genre);
  }, [genre]);

  console.log(animes);
  return (
    <div>
      <Navbar />
      <h1>{genre}</h1>
      <div className="list-container">
        {animes.length ? (
          animes.map((anime, index) => {
            return <SingleAnime key={index} anime={anime} />;
          })
        ) : (
          <div>No animes in this genre</div>
        )}
      </div>
    </div>
  );
}
