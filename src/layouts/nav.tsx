/* eslint-disable sort-keys-fix/sort-keys-fix */
import { styled } from '@stitches/react';
import { NavLink } from 'react-router-dom';

export const StyledNav = styled('nav', {
  backgroundColor: '#0e60dc',
  color: 'White',
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'space-between',
});

export const StyledNavTitle = styled('div', {
  fontSize: '1.5rem',
  marginLeft: '0.5rem',
  '& a': {
    color: 'White',
    display: 'block',
    textDecoration: 'none',
  },
});
export const StyledNavItems = styled('ul', {
  display: 'flex',
  listStyle: 'none',
  justifyContent: 'flex-end',
});

export const StyledNavItem = styled('li', {
  color: 'White',
  display: 'block',
  padding: '1rem',
  textDecoration: 'none',
  '&:hover': {
    backgroundColor: '#6a9de8',
  },
});

export const StyledNavLink = styled(NavLink, {
  color: 'White',
  display: 'block',
  fontSize: '1.2rem',
  padding: '1rem',
  textDecoration: 'none',
});
