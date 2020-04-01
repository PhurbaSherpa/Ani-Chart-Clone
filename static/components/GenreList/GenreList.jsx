import React, { useEffect, useState } from "react";
import { Navbar } from "../../components";
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
    </div>
  );
}
