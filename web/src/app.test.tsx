import { render, screen } from '@testing-library/react';
import { expect, test } from 'vitest';

import App from './app';

// TODO(@kylejb): stale since the AppProvider/router restructure — App renders a
// Spinner until the auth query settles, which never happens in jsdom without a
// mocked Connect transport. Re-enable once a transport mock exists.
test.skip("renders 'clairBuoyant' in header", () => {
  render(<App />);
  const linkElement = screen.getByText(/clairBuoyant/i);
  expect(linkElement).toBeInTheDocument();
});
