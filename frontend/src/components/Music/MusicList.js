/* Third party */
import React, { useEffect, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios'

const MusicList = () => {
  const navigate = useNavigate();
  const [musics, setMusics] = useState([])
  const id = 0;
  
const confirmDelete = async (id) => {
  const payload = {
      id: id.toString(),
  }
  await axios.post('http://localhost:4000/music/delete', JSON.stringify(payload))
  navigate('/musics')
}

  useEffect(() => {
    const fetchMusic = async () => {
      const result = await axios(`http://localhost:4000/music`)
      await setMusics(result.data.music)
    }
    fetchMusic()
  }, []);
  
  
  return (
    <>
    <div className="col-sm-2 mb-4 ">
      <Link to={`/music/create`} className='btn btn-primary'>
        Add
      </Link>
    </div>
    <div className="row">
        {musics.map((musics, index)=>( 
           <div className="col-sm-4 mb-4" key={index}>
                <div className="card">
                <div className="card-body">
                    <h5 className="card-title">{musics.name}</h5>
                    <p className="card-text">
                      Album : {musics.album} <br></br>
                      Singer : {musics.singer}
                    </p>
                    
                    <Link to={`/musics/detail/${musics.id}`} className='btn btn-secondary'>
                      Detail
                    </Link>

                    <Link to={`/music/edit/${musics.id}`} className='btn btn-warning mx-2'>
                      Edit
                    </Link>

                    <Link className='btn btn-danger mx-1' onClick={() => confirmDelete(musics.id)}>
                      Delete
                    </Link>
                </div>
                </div>
            </div>
        ))}
    </div>
    </>
  )
}

export default MusicList