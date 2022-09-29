import type { BaseProps } from '@types';

import { globalStyles } from '../styles';

import Header from './header';

export default function BaseLayout({ children }: BaseProps) {
  globalStyles(); // TODO: move to centralized location like <App />
  return (
    <>
      <Header />
      {children}
    </>
  );
}
