/* third party */
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

/* internal source */
import Home from './pages/Home'
import Musics from './pages/Musics'
import DetailMusic from './pages/Musics/DetailMusic'
import Genres from './pages/Genres'
import MusicForm from './components/Music/MusicForm'

/* style */
import './App.css';

function App() {
  return (
    <Router>
      <div className="container">
        <div className="row">    
          <h1 className="mt-3 text-center mb-5">Youre Music</h1>
          <hr className="mb-3" />
        </div>
        <div className="row">
          <div className="col-12">
            <Routes>
              <Route path='/' element={<Home/>} />
              <Route path='/musics' element={<Musics/>} />
              <Route exact path='/musics/detail/:id' element={<DetailMusic/>} />
              <Route exact path='/music/create' element={<MusicForm/>} />
              <Route exact path='/music/edit/:id' element={<MusicForm/>} />
              <Route path='/genres' element={<Genres/>} />
            </Routes>
          </div>
        </div>
      </div>
    </Router>
  );
}

export default App;
