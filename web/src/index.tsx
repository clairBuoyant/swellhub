import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';

import App from './app';
import './index.css';

const rootContainer = document.getElementById('root');
const entryPoint = createRoot(rootContainer!);

entryPoint.render(
  <StrictMode>
    <App />
  </StrictMode>,
);
