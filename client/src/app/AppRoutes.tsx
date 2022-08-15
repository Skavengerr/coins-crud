import React from 'react';
import {  Route, Routes } from 'react-router-dom';

import { Home } from 'containers';
import { CoinInfo } from 'components/CoinInfo/CoinInfo';

export const AppRoutes: React.FC = () => (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/:id" element={<CoinInfo />} />
    </Routes>
  );

