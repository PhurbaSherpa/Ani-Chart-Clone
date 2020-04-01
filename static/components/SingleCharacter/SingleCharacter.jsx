import React from "react";
import { Link } from "react-router-dom";
import "./SingleCharacter.css";

export default function SingleCharacter(props) {
  const { character } = props;
  console.log(character);
  return (
    <Link to={`/character/${character.Id}`}>
      <div className="card mb-3 single-character">
        <div className="row no-gutters">
          <div className="col-md-4">
            <img src={character.Imageurl} className="card-img" alt="..." />
          </div>
          <div className="col-md-8">
            <div className="card-body">
              <h5 className="card-title character-name">{character.Name}</h5>
              <p className="card-text">
                <small className="text-muted character-role">
                  {character.Role}
                </small>
              </p>
            </div>
          </div>
        </div>
      </div>
    </Link>
  );
}
