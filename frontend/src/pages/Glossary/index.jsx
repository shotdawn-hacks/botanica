import React, { useState } from 'react';
import { NavLink } from 'react-router-dom';
import Icon from '@mdi/react';
import { mdiMagnify } from '@mdi/js';
import useTitle from '../../hooks/useTitle';
import { useGetPlantsQuery } from '../../state/plant';

export default function Glossary() {
  useTitle('Sth | Glossary');
  const [currentPlant, setCurrentPlant] = useState({
    name: 'Аир обыкновенный',
    scientific_name: 'Acorus calamus L.',
    family: 'куст',
    sowing: {
      start: 'апрель',
    },
    complexity: '35%',
  });


  const plants = useGetPlantsQuery();

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
        {plants.isFetching || plants.isLoading ? (
          <div />
        ) : (
          <div className="glossary-searchresults">
            {plants.map((plant) => (
              <div className="searchresult-card" key={plant._id} onClick={() => setCurrentPlant(plant)}>
                <div
                  className="img-icon"
                  style={{
                    backgroundImage: `http://localhost:9000/static/${plant.image}.jpg`,
                  }}
                />
                <div className="plant-label">
                  <p>plant.name</p>
                </div>
              </div>
            ))}
          </div>
        )}
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
              <div className="glossary-char-text">{currentPlant.family}</div>
            </div>
            <div />
            <div className="glossary-char">
              <div className="glossary-icon">
                <img
                  src="../../assets/time.png"
                  className="glossary-icon-img"
                />
              </div>
              <div className="glossary-char-text">{currentPlant.sowing.start}</div>
            </div>
            <div />
            <div className="glossary-char">
              <div className="glossary-icon"></div>
              <div className="glossary-char-text">{currentPlant.complexity}</div>
            </div>
          </div>
          <div className="glossary-description">
            <h2 className="glossary-description-text2">
              {currentPlant.name}
              <em style={{ color: 'rgba(255, 255, 255, 0.5)' }}>
                ({currentPlant.scientific_name})
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
