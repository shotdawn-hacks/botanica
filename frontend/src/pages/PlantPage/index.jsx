import React from 'react';
import { NavLink } from 'react-router-dom';
import Icon from '@mdi/react';
import { mdiMagnify, mdiSnowflake, mdiForest, mdiBookOpenBlankVariant } from '@mdi/js';
import useTitle from '../../hooks/useTitle';

export default function PlantPage() {
  useTitle('Search | Plant Box');

  return (
    <main className="plant-page">
      <div className="plant-container">
        <div className="plant-details">
        </div>
      </div>
    </main>
  );
}
