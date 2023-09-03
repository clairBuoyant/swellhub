import * as DialogPrimitive from '@radix-ui/react-dialog';
import { style } from '@vanilla-extract/css';

import type { Props } from './types';

const styledOverlay = style({
  // overlay styles
});

export function Dialog({ children, ...props }: Props) {
  return (
    // eslint-disable-next-line react/jsx-props-no-spreading
    <DialogPrimitive.Root {...props}>
      <DialogPrimitive.Overlay className={styledOverlay} />
      {children}
    </DialogPrimitive.Root>
  );
}

export const DialogTrigger = DialogPrimitive.Trigger;
