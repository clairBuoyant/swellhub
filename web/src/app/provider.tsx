import { QueryClientProvider } from '@tanstack/react-query';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { Suspense } from 'react';
import { HelmetProvider } from 'react-helmet-async';

import { Notifications, NotificationProvider } from '@components/ui/notifications';
import { Spinner } from '@components/ui/spinner';
import { AuthLoader } from '@lib/auth';
import { queryClient } from '@lib/react-query';

type AppProviderProps = {
  children: React.ReactNode;
};

export default function AppProvider({ children }: AppProviderProps) {
  return (
    <Suspense
      fallback={
        <div className="flex h-screen w-screen items-center justify-center">
          <Spinner size="xl" />
        </div>
      }
    >
      <HelmetProvider>
        <QueryClientProvider client={queryClient}>
          {import.meta.env.DEV && <ReactQueryDevtools />}
          <NotificationProvider>
            <Notifications />
            <AuthLoader
              renderLoading={() => (
                <div className="flex h-screen w-screen items-center justify-center">
                  <Spinner size="xl" />
                </div>
              )}
            >
              {children}
            </AuthLoader>
          </NotificationProvider>
        </QueryClientProvider>
      </HelmetProvider>
    </Suspense>
  );
}
