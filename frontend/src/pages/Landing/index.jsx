import React from 'react';
import { NavLink } from 'react-router-dom';
import Icon from '@mdi/react';
import { mdiMagnify, mdiSnowflake, mdiForest, mdiBookOpenBlankVariant } from '@mdi/js';
import useTitle from '../../hooks/useTitle';

export default function Landing() {
  useTitle('Search | Plant Box');

  return (
    <div id="content-landing">
      <div className="top-part-search">
        <header className="landing-header-container">
          <div className="landing-header">
            <div className="logo" />
            <ul className="landing-header-links">
              <li>
                <NavLink to="/search">Поиск</NavLink>
              </li>
              <li>
                <NavLink to="/glossary">Глоссарий</NavLink>
              </li>
              <li>
                <NavLink to="/about_us">О нас</NavLink>
              </li>
            </ul>
          </div>
        </header>
        <main className="landing">
          <section className="landing-features">
            <div className="landing-features-container">
              <input type="search" className="searchbar" placeholder="Поиск растений по категориям"/>
              <button className="search-button">
                <Icon path={mdiMagnify} alt="magnifier" className="icon"/>
              </button>
            </div>
          </section>
        </main>
        <div className="search-results" >
          <div className="cards">
            <h1 className="wide-card">All search results that match "Name":</h1>
            <div className="card">
              <img src="../assets/stock_flower.jpg" className="card-img"/>
              <p className="card-title"> Name (Homo Neimo)</p>
              <div className="card-stats">
                <div className="card-stat" style={{backgroundColor: 'cyan'}}>
                  <Icon path={mdiSnowflake} alt="magnifier" className="card-icon"/>
                </div>
                <div className="card-stat" style={{backgroundColor: 'green'}}>
                  <Icon path={mdiForest} alt="magnifier" className="card-icon"/>
                </div>
                <div className="card-stat" style={{backgroundColor: 'red'}}>
                  <Icon path={mdiBookOpenBlankVariant} alt="magnifier" className="card-icon"/>
                </div>
              </div>
            </div>
            <div className="card">
              <img src="../assets/stock_flower.jpg" className="card-img"/>
              <p className="card-title"> Name (Homo Neimo)</p>
              <div className="card-stats">
                <div className="card-stat" style={{backgroundColor: 'cyan'}}>
                  <Icon path={mdiSnowflake} alt="magnifier" className="card-icon"/>
                </div>
                <div className="card-stat" style={{backgroundColor: 'green'}}>
                  <Icon path={mdiForest} alt="magnifier" className="card-icon"/>
                </div>
                <div className="card-stat" style={{backgroundColor: 'red'}}>
                  <Icon path={mdiBookOpenBlankVariant} alt="magnifier" className="card-icon"/>
                </div>
              </div>
            </div>
            <div className="card">
              <img src="../assets/stock_flower.jpg" className="card-img"/>
              <p className="card-title"> Name (Homo Neimo)</p>
              <div className="card-stats">
                <div className="card-stat" style={{backgroundColor: 'cyan'}}>
                  <Icon path={mdiSnowflake} alt="magnifier" className="card-icon"/>
                </div>
                <div className="card-stat" style={{backgroundColor: 'green'}}>
                  <Icon path={mdiForest} alt="magnifier" className="card-icon"/>
                </div>
                <div className="card-stat" style={{backgroundColor: 'red'}}>
                  <Icon path={mdiBookOpenBlankVariant} alt="magnifier" className="card-icon"/>
                </div>
              </div>
            </div>
            <div className="card">
              <img src="../assets/stock_flower.jpg" className="card-img"/>
              <p className="card-title"> Name (Homo Neimo)</p>
              <div className="card-stats">
                <div className="card-stat" style={{backgroundColor: 'cyan'}}>
                  <Icon path={mdiSnowflake} alt="magnifier" className="card-icon"/>
                </div>
                <div className="card-stat" style={{backgroundColor: 'green'}}>
                  <Icon path={mdiForest} alt="magnifier" className="card-icon"/>
                </div>
                <div className="card-stat" style={{backgroundColor: 'red'}}>
                  <Icon path={mdiBookOpenBlankVariant} alt="magnifier" className="card-icon"/>
                </div>
              </div>
            </div>

          </div>
        </div>
      </div>
    </div>
  );
}
