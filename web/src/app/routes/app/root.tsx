import { Suspense } from 'react';
import { Outlet } from 'react-router-dom';

import { DashboardLayout } from '@components/layouts/dashboard-layout';
import { Spinner } from '@components/ui/spinner';

export default function AppRoot() {
  return (
    <DashboardLayout>
      <Suspense
        fallback={
          <div className="flex size-full items-center justify-center">
            <Spinner size="xl" />
          </div>
        }
      >
        <Outlet />
      </Suspense>
    </DashboardLayout>
  );
}
