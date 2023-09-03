import { style } from '@vanilla-extract/css';

import type { BaseProps } from '@types';

const styledResponsiveBox = style({
  // TODO: add responsive styles
});

export default function ResponsiveBox({ children }: BaseProps) {
  return <div className={styledResponsiveBox}>{children}</div>;
}
