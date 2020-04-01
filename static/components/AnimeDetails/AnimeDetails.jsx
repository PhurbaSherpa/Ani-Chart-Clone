import React, { useEffect, useState } from "react";
import { Navbar, Detailsbar, SingleCharacter } from "../../components";
import "./AnimeDetails.css";
import axios from "axios";

export default function AnimeDetails(props) {
  const { id } = props.match.params;
  const [characters, setCharacters] = useState([]);
  const [anime, setAnime] = useState([]);

  useEffect(() => {
    const getDetails = async id => {
      const res1 = await axios.get(`/api/characters/${id}`);
      setCharacters(res1.data);
      const res2 = await axios.get(`/api/animes/${id}`);
      setAnime(res2.data);
    };
    getDetails(id);
  }, [id]);

  return (
    <div>
      <Navbar />
      <div className="details-container">
        <div className="card mb-3 anime-info">
          <div className="row no-gutters">
            <div className="col-md-4">
              <img src={anime.Imageurl} className="card-img" alt="..." />
            </div>
            <div className="col-md-8">
              <div className="card-body">
                <h5 className="card-title detail-title">{anime.Title}</h5>
                <p className="card-text detail-description">
                  {anime.Description}
                </p>
                <p className="card-text">
                  <small className="text-muted detail-date">
                    Release Date: {anime.Date}
                  </small>
                </p>
              </div>
            </div>
          </div>
        </div>
        <Detailsbar />
        <div className="character-list">
          {characters.map((character, index) => {
            return <SingleCharacter key={index} character={character} />;
          })}
        </div>
      </div>
    </div>
  );
}
