import BaseLayout from '@layouts/base-layout';

import AppProvider from './providers/app';
import AppRoutes from './routes';
import './styles/reset.css';

// TODO: rework handling of layout components for standardization
export default function App() {
  return (
    <AppProvider>
      <BaseLayout>
        <AppRoutes />
      </BaseLayout>
    </AppProvider>
  );
}
