import { lazy } from 'react';

// eslint-disable-next-line import/no-unresolved
const AuthRoutes = lazy(() => import('@features/auth/routes')); // TODO: resolve eslint error

export const publicRoutes = [
  {
    element: <AuthRoutes />,
    path: '/auth/*',
  },
];
