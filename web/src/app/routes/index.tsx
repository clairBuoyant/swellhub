import type { QueryClient } from '@tanstack/react-query';
import { createBrowserRouter } from 'react-router-dom';

import { buoysLoader } from './app/buoys/buoys';
import AppRoot from './app/root';

export const createRouter = (queryClient: QueryClient) =>
  createBrowserRouter([
    {
      lazy: async () => {
        const { LandingRoute } = await import('./landing');
        return { Component: LandingRoute };
      },
      path: '/',
    },
    {
      lazy: async () => {
        const { RegisterRoute } = await import('./auth/register');
        return { Component: RegisterRoute };
      },
      path: '/auth/register',
    },
    {
      lazy: async () => {
        const { LoginRoute } = await import('./auth/login');
        return { Component: LoginRoute };
      },
      path: '/auth/login',
    },
    {
      children: [
        {
          lazy: async () => {
            const { BuoysRoute } = await import('./app/buoys/buoys');
            return { Component: BuoysRoute };
          },
          loader: buoysLoader(queryClient),
          path: 'buoys',
        },
        // {
        //   lazy: async () => {
        //     const { BuoyRoute } = await import('./app/buoys/buoy');
        //     return { Component: BuoyRoute };
        //   },
        //   loader: buoyLoader(queryClient),
        //   path: 'buoys/:buoyID',
        // },
        {
          lazy: async () => {
            const { ProfileRoute } = await import('./app/profile');
            return { Component: ProfileRoute };
          },
          path: 'profile',
        },
        {
          lazy: async () => {
            const { DashboardRoute } = await import('./app/dashboard');
            return { Component: DashboardRoute };
          },
          path: '',
        },
      ],
      element: (
        // TODO(@kylejb): wrap with <ProtectedRoute> and update all paths
        <AppRoot />
      ),
      path: '/app',
    },
    {
      lazy: async () => {
        const { NotFoundRoute } = await import('./not-found');
        return { Component: NotFoundRoute };
      },
      path: '*',
    },
  ]);
