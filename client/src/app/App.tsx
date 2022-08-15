import React from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import { RecoilRoot } from 'recoil';

import { Screen } from 'components';

import { AppRoutes } from './AppRoutes';
import 'antd/dist/antd.css';

export const App: React.FC = () => (
  <RecoilRoot>
    <Router>
      <Screen>
        <AppRoutes />
      </Screen>
    </Router>
  </RecoilRoot>
);
