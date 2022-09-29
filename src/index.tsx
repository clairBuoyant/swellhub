import React from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';

import renderRoutes, { allRoutes } from './routes';

const rootContainer = document.getElementById('root');
const entryPoint = createRoot(rootContainer!);

entryPoint.render(
  <React.StrictMode>
    <BrowserRouter>{renderRoutes(allRoutes)}</BrowserRouter>
  </React.StrictMode>,
);
