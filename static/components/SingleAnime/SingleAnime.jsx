import React, { useEffect, useState } from "react";
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
    <div className="pb-3" style={{ width: "30%" }}>
      <div className="card">
        <div className="row no-gutters">
          <div className="col-md-4">
            <img src={anime.Imageurl} className="card-img" alt="..." />
          </div>
          <div className="col-md-8">
            <div className="card-body">
              <h5 className="card-title">{anime.Title}</h5>
              <p className="card-text">{anime.Description}</p>
              <p className="card-text">
                <small className="text-muted">{anime.Date}</small>
              </p>
            </div>
          </div>
        </div>
      </div>
      <div className="card-footer">
        {genres.map((genre, index) => {
          return (
            <small key={index} className="text-muted">
              +{genre}{" "}
            </small>
          );
        })}
      </div>
    </div>
  );
}
