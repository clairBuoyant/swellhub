import { useRoutes } from 'react-router-dom';

import WorldMap from '@views/world-map';

import { publicRoutes } from './public';

export default function AppRoutes() {
  // TODO: handle auth flow
  // const auth = useAuth();

  const commonRoutes = [
    {
      element: <WorldMap />,
      path: '/',
    },
  ];

  // const routes = auth.user ? protectedRoutes : publicRoutes;

  const element = useRoutes([...publicRoutes, ...commonRoutes]);

  return element;
}
