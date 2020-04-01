import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import "./SingleAnime.css";
import axios from "axios";

export default function SingleAnime(props) {
  const { anime } = props;
  const [genres, setGenres] = useState([]);

  useEffect(() => {
    const getGenres = async () => {
      const { data: genres } = await axios.get(`/api/genres/${anime.Id}`);
      setGenres(genres);
    };
    if (!genres.length) {
      getGenres();
    }
    console.log(genres);
  }, [genres.length]);

  return (
    <div className="pb-3 anime-box" style={{ width: "30%" }}>
      <Link to={`/animes/${anime.Id}`} className="card hover">
        <div className="row no-gutters">
          <div className="col-md-4">
            <img src={anime.Imageurl} className="card-img" alt="..." />
          </div>
          <div className="col-md-8">
            <div className="card-body">
              <h5 className="card-title">{anime.Title}</h5>
              <p className="card-text">{anime.Description}</p>
              <p className="card-text">
                <small className="text-muted">Release Date: {anime.Date}</small>
              </p>
            </div>
          </div>
        </div>
      </Link>
      <div className="card-footer">
        {genres.map((genre, index) => {
          return (
            <Link className="genre" to={`/genre/${genre}`} key={index}>
              <small>+{genre} </small>
            </Link>
          );
        })}
      </div>
    </div>
  );
}
