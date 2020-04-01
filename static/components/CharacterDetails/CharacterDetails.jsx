import React, { useEffect, useState } from "react";
import { Navbar } from "../../components";
import axios from "axios";
import "./CharacterDetails.css";

export default function CharacterDetails(props) {
  const { id } = props.match.params;
  const [character, setCharacter] = useState({});

  useEffect(() => {
    const getDetails = async id => {
      const { data } = await axios.get(`/api/character/${id}`);
      setCharacter(data);
    };
    getDetails(id);
  }, [id]);

  console.log(character);

  return (
    <div>
      <Navbar />
      <div className="card mb-3 character-info">
        <div className="row no-gutters">
          <div className="col-md-4">
            <img src={character.Imageurl} className="card-img" alt="..." />
          </div>
          <div className="col-md-8">
            <div className="card-body">
              <h5 className="card-title">{character.Name}</h5>
              <p className="card-text detail-description">
                {character.Description}
              </p>
              <p className="card-text">
                <small className="text-muted">Role: {character.Role}</small>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
