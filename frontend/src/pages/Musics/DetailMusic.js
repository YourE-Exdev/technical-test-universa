import axios from 'axios'
import React, { useEffect, useState } from 'react'
import { Link, useParams } from 'react-router-dom'

const DetailMusic = () => {
  let { id } = useParams();
  const [music, setMusic] = useState([]);

  useEffect(() => {
    const fetchMusic = async () => {
        const result = await axios(`http://localhost:4000/music/${id}`);
        await setMusic(result.data.music);
    };
        fetchMusic();
    }, [id]);
  return (
    <>
        <ul>
            { music ? (
                <div>
                    <h2> Music {music.name} ({music.publishdate}) </h2> 
                    <h5>Album : {music.album}</h5>
                </div>
            ) : ( 
                <p>Oops... There's no movies data.</p>
            )}
        </ul>
    </>
  )
}

export default DetailMusic