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

  return <div></div>;
}
