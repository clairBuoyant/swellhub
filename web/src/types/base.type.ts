import type { ReactNode } from 'react';

export interface BaseBuoy {
  id: string;
  latitude: number;
  longitude: number;
  name: string;
}

export interface BaseProps {
  children?: ReactNode;
  styles?: BaseStyles;
}

export interface BaseStyles {
  height?: string;
  width?: string;
}
