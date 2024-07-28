import { useQueryClient } from '@tanstack/react-query';
import { useMemo } from 'react';
import { RouterProvider } from 'react-router-dom';

import AppProvider from './provider';
import { createRouter } from './routes';

function AppRouter() {
  const queryClient = useQueryClient();

  const router = useMemo(() => createRouter(queryClient), [queryClient]);

  return <RouterProvider router={router} />;
}

export default function App() {
  return (
    <AppProvider>
      <AppRouter />
    </AppProvider>
  );
}
