import React, { Fragment } from 'react';
import {BrowserRouter as Router, Routes, Route, Link, useLocation} from 'react-router-dom';
import Movies from './components/movies';
import Admin from './components/admin';
import Home from './components/home';
import Categories from './components/categories';
import OneMovie from './components/one_movie';

export default function App() {
  return (
    <Router>
    <div className="container">
      <div className="row">
        <h1 className="mt-3">
          Go Watch a Movie
        </h1>
        <hr className='mb-3'></hr>
      </div>
      <div className="row">
        <div className="col-md-2">
          <nav>
            <ul className="list-group">
              <li className="list-group-item">
                <Link to="/">Home</Link>
              </li>
              <li className="list-group-item">
                <Link to="/movies">Movies</Link>
              </li>
              <li className="list-group-item">
                <Link to="/by-category">Categories</Link>
              </li>
              <li className="list-group-item">
                <Link to="/admin">Manage Catalogue</Link>
              </li>
            </ul>
          </nav>
        </div>

        <div className="col-md-10">
          <Routes>
            {/* <Route path="/movies/:id" element={<Movie />} /> */}
            
            <Route exact path="/" element={<Home />} />
            <Route exact path="/movies" element={<Movies />} />
              <Route path="/movies/:id" element={<OneMovie/>} />
            <Route exact path="/by-category" element={<CategoryPage  />} />
              <Route path="/by-category/drama" element={<Categories title={`Drama`}/>} />
              <Route exact path="/by-category/comedy" element={<Categories title={`Comedy`}/>} />
      {/* Passing a map property called title to Route component Categories */}
            <Route exact path="/admin" element={<Admin />} />
          </Routes>

        </div>
      </div>
    </div>
    </Router>
  );
}

function CategoryPage(){
  let { pathname } = useLocation();
  return (
    <div>
      <h2>Categories</h2>

      <ul>
        <li><Link to={`${pathname}/comedy`}>Comedy</Link></li>
        <li><Link to={`${pathname}/drama`}>Drama</Link></li>
      </ul>
    </div>
  );


}