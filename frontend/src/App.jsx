import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Landing from './pages/Landing';
import PlantPage from './pages/PlantPage';
import Glossary from './pages/Glossary';

const App = () => {
  console.log('test');
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Landing />} />
        <Route path="/glossary" element={<Glossary />} />
        <Route path="/plant_page" element={<PlantPage />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
