/* eslint-disable sort-keys-fix/sort-keys-fix */
export const reset = {
  '*, *::before, *::after': {
    boxSizing: 'border-box',
  },
  '*': {
    margin: '0',
  },
  'html, body': {
    height: '100%',
  },
  body: {
    WebkitFontSmoothing: 'antialiased',
    lineHeight: '1.5',
  },
  'img, picture, video, canvas, svg': {
    display: 'block',
    maxWidth: '100%',
  },
  'input, button, textarea, select': {
    font: 'inherit',
  },
  'p, h1, h2, h3, h4, h5, h6': {
    overflowWrap: 'break-word',
  },
  '#root': {
    isolation: 'isolate',
  },
};
