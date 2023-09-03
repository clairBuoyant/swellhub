import type { BaseProps } from '@types';

import Header from './header';

export default function BaseLayout({ children }: BaseProps) {
  return (
    <>
      <Header />
      {children}
    </>
  );
}
