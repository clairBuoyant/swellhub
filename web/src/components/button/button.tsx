import { style } from '@vanilla-extract/css';

import type { Props } from './types';

// TODO: support dynamic variants or rework with vanilla-extract in mind.
const styledButton = style({
  border: '2px solid turquoise',
  borderRadius: '$round',
  color: 'white',
  fontSize: '$4',
  padding: '$2 $3',
  selectors: {
    '&:hover': {
      backgroundColor: 'turquoise',
      color: 'black',
    },
  },
});

export default function Button({ content }: Props) {
  return (
    <button type="button" className={styledButton}>
      {content}
    </button>
  );
}
