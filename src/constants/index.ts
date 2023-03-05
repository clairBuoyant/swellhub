export const NODE_ENV = {
  DEVELOPMENT: 'development',
  PRODUCTION: 'production',
};

const BASE_API_URL = import.meta.env.DEV ? 'http://127.0.0.1:8888' : 'https://clairbuoyant.live';

export const API_VERSION = 'v1';
export const BASE_API_URI = `${BASE_API_URL}/api/${API_VERSION}`;
export const API_ROUTES = {
  BUOYS: `${BASE_API_URI}/buoys`,
  COASTLINES: `${BASE_API_URI}/coastlines`,
};
