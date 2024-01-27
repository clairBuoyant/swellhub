import type { ReactNode } from 'react';

export interface BaseBuoy {
  name: string;
  stationId: string;
}

export interface BaseProps {
  children: ReactNode;
  styles?: BaseStyles;
}

export interface BaseStyles {
  height?: string;
  width?: string;
}
