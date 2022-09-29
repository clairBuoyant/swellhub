import { styled } from '@stitches/react';

import type { IBaseProps } from '@types';

const StyledResponsiveBox = styled('div', {
  // backgroundColor: '$pink',
});

function ResponsiveBox({ children }: IBaseProps) {
  return <StyledResponsiveBox>{children}</StyledResponsiveBox>;
}

export default ResponsiveBox;
