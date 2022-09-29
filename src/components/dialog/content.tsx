import * as DialogPrimitive from '@radix-ui/react-dialog';
import { keyframes, styled } from '@stitches/react';
import { forwardRef } from 'react';

import type { DialogRef, Props } from './types';

const scaleIn = keyframes({
  from: { transform: 'scale(0.9)' },
  to: { transform: 'scale(1)' },
});

const fadeIn = keyframes({
  from: { opacity: 0 },
  to: { opacity: 1 },
});

const fadeout = keyframes({
  from: { opacity: 1 },
  to: { opacity: 0 },
});

const StyledContent = styled(DialogPrimitive.Content, {
  variants: {
    animation: {
      fade: {
        '&[data-state="closed"]': {
          animation: `${fadeout} 200ms ease-out`,
        },
        '&[data-state="open"]': {
          animation: `${fadeIn} 300ms ease-out`,
        },
      },
      scale: {
        animation: `${fadeIn} 300ms ease-out, ${scaleIn} 200ms ease-out`,
      },
    },
  },
});

const StyledCloseButton = styled(DialogPrimitive.Close, {
  // close button styles
});

export const DialogContent = forwardRef(
  ({ children, ...props }: Props, forwardedRef: DialogRef) => (
    // eslint-disable-next-line react/jsx-props-no-spreading
    <StyledContent {...props} ref={forwardedRef}>
      {children}
      <StyledCloseButton>O</StyledCloseButton>
    </StyledContent>
  ),
);
