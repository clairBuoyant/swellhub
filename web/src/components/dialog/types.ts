import type { ForwardedRef, ReactNode } from 'react';

export type DialogRef = ForwardedRef<HTMLDivElement>;

export interface Props {
  children: ReactNode;
}
