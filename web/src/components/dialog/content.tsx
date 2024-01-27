import * as DialogPrimitive from '@radix-ui/react-dialog';
import { keyframes, style } from '@vanilla-extract/css';
import { forwardRef } from 'react';

import type { DialogRef, Props } from './types';

const fadeIn = keyframes({
  from: { opacity: 0 },
  to: { opacity: 1 },
});

const fadeout = keyframes({
  from: { opacity: 1 },
  to: { opacity: 0 },
});

const styledContent = style({
  selectors: {
    '&[data-state="closed"]': {
      animation: `${fadeout} 200ms ease-out`,
    },
    '&[data-state="open"]': {
      animation: `${fadeIn} 300ms ease-out`,
    },
  },
});

const styledCloseButton = style({
  // close button styles
});

export const DialogContent = forwardRef(
  ({ children, ...props }: Props, forwardedRef: DialogRef) => (
    // eslint-disable-next-line react/jsx-props-no-spreading
    <DialogPrimitive.Content {...props} className={styledContent} ref={forwardedRef}>
      {children}
      <DialogPrimitive.Close className={styledCloseButton}>O</DialogPrimitive.Close>
    </DialogPrimitive.Content>
  ),
);
