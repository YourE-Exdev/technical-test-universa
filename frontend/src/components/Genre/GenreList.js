/* Third party */
import React, { useEffect, useState } from 'react'
import { Link } from 'react-router-dom';

const MusicList = () => {
  const [genres, setGenres] = useState([])
  const id = 0;
  useEffect(() => {
    setGenres([
        { id: 1, genre_name: "POP"},
        { id: 2, genre_name: "Reggae"},
        { id: 3, genre_name: "Metal"},
    ]);
  }, []);
  
  return (
    <div className="row">
        {genres.map((genres, index)=>(
           <div className="col-sm-2 mb-3" key={index}>
                <div className="card">
                    <div className="card-body text-center">
                        <Link to={`/genres/${genres.id}`}>{genres.genre_name}</Link>
                    </div>
                </div>
            </div>
        ))}
    </div>
  )
}

export default MusicList