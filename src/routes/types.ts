import type { ElementType } from 'react';
import type { NonIndexRouteObject } from 'react-router-dom';

export interface IRoute extends NonIndexRouteObject {
  layout?: ElementType;
}

export type TParentLayout = ElementType;
