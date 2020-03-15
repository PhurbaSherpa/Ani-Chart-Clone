import React, { useEffect, useState } from "react";
import axios from "axios";

export default function AnimeDetails(props) {
  const [characters, setCharacters] = useState([]);

  useEffect(() => {
    const getCharacters = async id => {
      const { data: characters } = await axios.get(`/api/characters/${id}`);
      setCharacters(characters);
    };

    getCharacters(props.anime.id);
  }, [props]);

  console.log(characters);
  return <div></div>;
}
