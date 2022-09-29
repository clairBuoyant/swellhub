export const NODE_ENV = {
  DEVELOPMENT: 'development',
  PRODUCTION: 'production',
};

export const API_VERSION = 'v1';
export const BASE_API_URI = `http://localhost:8888/api/${API_VERSION}`;
export const API_ROUTES = {
  BUOYS: `${BASE_API_URI}/buoys`,
  COASTLINES: `${BASE_API_URI}/coastlines`,
};
