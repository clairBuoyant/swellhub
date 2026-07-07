import { render, screen } from '@testing-library/react';
import { expect, test, vi } from 'vitest';

import App from './app';

// AuthLoader blocks rendering until the authenticated-user query settles. That
// query goes through the axios api-client, so resolve it with no signed-in
// user; AuthLoader then renders its children (the router).
vi.mock('@lib/api-client', () => ({
  api: {
    get: vi.fn().mockResolvedValue(null),
    post: vi.fn().mockResolvedValue(undefined),
  },
}));

// Keep Connect RPCs off the network in jsdom: serve them from an in-memory
// router transport instead (no services registered — calls reject cleanly).
vi.mock('@lib/connect', async () => {
  const { createRouterTransport } = await import('@connectrpc/connect');
  return { transport: createRouterTransport(() => {}) };
});

test("renders 'clairBuoyant' in header", async () => {
  render(<App />);
  const heading = await screen.findByText(/clairBuoyant/i);
  expect(heading).toBeInTheDocument();
});
