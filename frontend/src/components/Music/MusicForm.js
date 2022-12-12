import React, { useEffect } from 'react'
import { useForm } from 'react-hook-form'
import axios from 'axios'
import { useNavigate, useParams } from 'react-router-dom'

const MusicForm = () => {
  const { id } = useParams()
  const isAddMode = !id
  const { register, handleSubmit, setValue } = useForm()

  const field = [
    'id',
    'name',
    'album',
    'albumart',
    'singer',
    'publishdate',
    'createat',
    'updateat',
  ]

  const fetchMusic = async (id) => {
    try{
        const result = await axios(`http://localhost:4000/music/${id}`)
        result.data.music.id = result.data.music.id.toString()
        field.forEach((field) => setValue(field, result.data.music[field]))
    } catch (err) {
        console.log(err.response.data)
    }
  }

  useEffect(() => {
    if(!isAddMode){
        fetchMusic(id)
    }
  }, [isAddMode]);
  

  const navigate = useNavigate()
  const onSubmit = async data => {
    if(isAddMode){
        const result = await axios.post(
            'http://localhost:4000/music/create',
            JSON.stringify(data)
        )
    } else {
        const result = await axios.post(
            'http://localhost:4000/music/edit',
            JSON.stringify(data)
        )
    }
    navigate('/musics')
  }
  return (
    <>
        <h2>Music Form</h2>
        <hr/>
        <form onSubmit={handleSubmit(onSubmit)}>
            <div className='mb-3'>
                <label htmlFor='' className='form-label'>
                    Music Name
                </label>
                <input type='text' className='form-control' id='name' name='name' {...register('name', { required: true })}/>
            </div>
            
            <div className='mb-3'>
                <label htmlFor='' className='form-label'>
                    Album
                </label>
                <input type='text' className='form-control' id='album' name='album' {...register('album', { required: true })}/>
            </div>
            
            <div className='mb-3'>
                <label htmlFor='' className='form-label'>
                    Album Art
                </label>
                <input type='text' className='form-control' id='albumart' name='albumart' {...register('albumart', { required: true })}/>
            </div>

            <div className='mb-3'>
                <label htmlFor='' className='form-label'>
                    Singer
                </label>
                <input type='text' className='form-control' id='singer' name='singer' {...register('singer', { required: true })}/>
            </div>

            
            <div className='mb-3'>
                <label htmlFor='' className='form-label'>
                    Publish Date
                </label>
                <input type='text' className='form-control' id='publishdate' name='publishdate' {...register('publishdate', { required: true })}/>
            </div>

            <div className='mb-3'>
                <button type='submit' className='btn btn-success'>Submit</button>
            </div>
        </form>
    </>
  )
}

export default MusicForm