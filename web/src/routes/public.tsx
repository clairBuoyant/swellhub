import { lazy } from 'react';

const AuthRoutes = lazy(() => import('@features/auth/routes'));

export const publicRoutes = [
  {
    element: <AuthRoutes />,
    path: '/auth/*',
  },
];
