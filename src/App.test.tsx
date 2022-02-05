import { render, screen } from '@testing-library/react';

import App from './App';

test("renders 'clairBuoyant' in header", () => {
  render(<App />);
  const linkElement = screen.getByText(/clairBuoyant/i);
  expect(linkElement).toBeInTheDocument();
});
