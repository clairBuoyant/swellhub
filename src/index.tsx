import React from 'react';
import { createRoot } from 'react-dom/client';

import App from './app';
import reportWebVitals from './reportWebVitals';

const rootContainer = document.getElementById('root');
const entryPoint = createRoot(rootContainer!);
entryPoint.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);

// TODO: https://bit.ly/CRA-vitals
reportWebVitals();
