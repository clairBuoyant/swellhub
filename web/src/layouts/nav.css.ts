import { style } from '@vanilla-extract/css';

export const styledNav = style({
  backgroundColor: '#0e60dc',
  color: 'White',
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'space-between',
});

export const styledNavTitle = style({
  fontSize: '1.5rem',
  marginLeft: '0.5rem',
});
export const styledNavItems = style({
  display: 'flex',
  justifyContent: 'flex-end',
  listStyle: 'none',
});

export const styledNavItem = style({
  color: 'White',
  display: 'block',
  padding: '1rem',
  selectors: {
    '&:hover': {
      backgroundColor: '#6a9de8',
    },
  },
  textDecoration: 'none',
});

export const styledNavLink = style({
  color: 'White',
  display: 'block',
  fontSize: '1.2rem',
  padding: '1rem',
  textDecoration: 'none',
});
