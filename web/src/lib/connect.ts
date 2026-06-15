import { createConnectTransport } from '@connectrpc/connect-web';

/**
 * Same-origin Connect transport for the clairBuoyant API.
 *
 * In dev, Vite proxies the Connect path (`/clairbuoyant.*`) to the Go API
 * (see vite.config.ts), so a relative baseUrl keeps the browser same-origin and
 * avoids CORS. In production the SPA is embedded and served by the API, so the
 * origin is already correct.
 */
export const transport = createConnectTransport({
  baseUrl: '/',
});
