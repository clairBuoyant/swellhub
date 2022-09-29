import type { ElementType } from 'react';
import type { RouteObject } from 'react-router-dom';

export interface IRoute extends RouteObject {
  layout?: ElementType;
}

export type TParentLayout = ElementType;
