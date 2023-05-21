import React from 'react';
import { NavLink } from 'react-router-dom';
import Icon from '@mdi/react';
import { mdiMagnify } from '@mdi/js';
import useTitle from '../../hooks/useTitle';

export default function Glossary() {
  useTitle('Sth | Glossary');

  return (
    <main className="glossary-page">
      <div className="glossary-search">
        <div className="glossary-searchbar">
          <button className="searchbar-filter-button">
            <img src="../../assets/filter.png" className="filter-img" />
          </button>
          <input className="searchbar-input"></input>
          <button className="searchbar-button">
            <Icon path={mdiMagnify} className="icon searchbar-icon" />
          </button>
        </div>
        <div className="glossary-filters">
          <h3>Типы почв</h3>
          <h3>Ареалы</h3>
          <h3>Регионы</h3>
          <h3>Климат</h3>
          <h3>Активные в-ва</h3>
          <h3>Группы лекарств</h3>
        </div>
        <div className="glossary-searchresults">
          <div className="searchresult-card">
            <div className="img-icon"/>
            <div className="plant-label">
              <p>Name</p>
            </div>
          </div>

          <div className="searchresult-card">
            <div className="img-icon"/>
            <div className="plant-label">
              <p>Name</p>
            </div>
          </div>

          <div className="searchresult-card">
            <div className="img-icon"/>
            <div className="plant-label">
              <p>Name</p>
            </div>
          </div>

        </div>
      </div>
      <div className="glossary-container">
        <div className="glossary-details">
          <div className="glossary-chars">
            <div className="glossary-char">
              <div className="glossary-icon">
                <img
                  src="../../assets/plant.png"
                  className="glossary-icon-img"
                />
              </div>
              <div className="glossary-char-text">куст</div>
            </div>
            <div />
            <div className="glossary-char">
              <div className="glossary-icon">
                <img
                  src="../../assets/time.png"
                  className="glossary-icon-img"
                />
              </div>
              <div className="glossary-char-text">апрель</div>
            </div>
            <div />
            <div className="glossary-char">
              <div className="glossary-icon"></div>
              <div className="glossary-char-text">35%</div>
            </div>
          </div>
          <div className="glossary-description">
            <h2 className="glossary-description-text2">
              Аир обыкновенный
              <em style={{ color: 'rgba(255, 255, 255, 0.5)' }}>
                (Acorus calamus L.)
              </em>
            </h2>
            <h3 className="glossary-description-text">Обитание</h3>
            <hr />
            <h3 className="glossary-description-text">Посадка</h3>
            <hr />
            <h3 className="glossary-description-text">
              Химический состав и применение
            </h3>
            <hr />
          </div>
        </div>
      </div>
    </main>
  );
}
