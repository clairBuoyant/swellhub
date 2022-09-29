import { Fragment, Suspense } from 'react';
import { Route, Routes } from 'react-router-dom';

import BaseLayout from '@layouts/base-layout';
import WorldMap from '@views/world-map';

import type { IRoute, TParentLayout } from './types';

export const allRoutes: IRoute[] = [
  {
    children: [
      {
        path: '/about',
      },
      {
        path: '/home',
      },
      {
        path: '/profile',
      },
    ],
    element: <WorldMap />,
    index: true,
    layout: BaseLayout,
    path: '/',
  },
];

export function renderRoutes(routes: IRoute[] = [], parentLayout: TParentLayout = Fragment) {
  return (
    <Suspense fallback={<>...</>}>
      <Routes>
        {routes.map((route, idx) => {
          const routeKey = `r-${idx}-${route.path}`;
          const Layout = route.layout || parentLayout;
          if (route.children) renderRoutes(route.children, Layout);
          return (
            <Route element={<Layout>{route.element}</Layout>} key={routeKey} path={route.path!} />
          );
        })}
      </Routes>
    </Suspense>
  );
}
