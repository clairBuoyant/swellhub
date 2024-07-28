export const NODE_ENV = {
  DEVELOPMENT: 'development',
  PRODUCTION: 'production',
};

const BASE_API_URL = import.meta.env.DEV ? 'http://127.0.0.1:4000' : 'https://clairbuoyant.live';

export const API_ROUTES = {
  BUOYS: `${BASE_API_URL}/buoys`,
  COASTLINES: `${BASE_API_URL}/coastlines`,
};
