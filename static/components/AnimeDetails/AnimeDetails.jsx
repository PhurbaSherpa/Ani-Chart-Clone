import React, { useEffect, useState } from "react";
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
  console.log(characters);
  console.log(anime);

  return <div></div>;
}
