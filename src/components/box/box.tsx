import { styled } from '@stitches/react';

import type { BaseProps } from '@types';

const StyledResponsiveBox = styled('div', {
  // backgroundColor: '$pink',
});

function ResponsiveBox({ children }: BaseProps) {
  return <StyledResponsiveBox>{children}</StyledResponsiveBox>;
}

export default ResponsiveBox;
