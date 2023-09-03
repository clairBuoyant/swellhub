import { Suspense } from 'react';
import { BrowserRouter } from 'react-router-dom';

type AppProviderProps = {
  children: React.ReactNode;
};

export default function AppProvider({ children }: AppProviderProps) {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <BrowserRouter>{children}</BrowserRouter>
    </Suspense>
  );
}
