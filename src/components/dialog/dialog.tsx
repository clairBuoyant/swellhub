import * as DialogPrimitive from '@radix-ui/react-dialog';
import { styled } from '@stitches/react';

import type { Props } from './types';

const StyledOverlay = styled(DialogPrimitive.Overlay, {
  // overlay styles
});

export function Dialog({ children, ...props }: Props) {
  return (
    // eslint-disable-next-line react/jsx-props-no-spreading
    <DialogPrimitive.Root {...props}>
      <StyledOverlay />
      {children}
    </DialogPrimitive.Root>
  );
}

export const DialogTrigger = DialogPrimitive.Trigger;
