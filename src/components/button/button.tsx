import { styled } from '@stitches/react';

import type { Props } from './types';

const StyledButton = styled('button', {
  '&:hover': {
    backgroundColor: '$turq',
    color: '$black',
  },
  border: '2px solid $turq',
  borderRadius: '$round',
  color: '$white',
  fontSize: '$4',
  padding: '$2 $3',
});

export function Button({ content }: Props) {
  return <StyledButton>{content}</StyledButton>;
}
