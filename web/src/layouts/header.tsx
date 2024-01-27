import { NavLink } from 'react-router-dom';

import Container from '@components/container';

import { styledNav, styledNavItem, styledNavItems, styledNavLink } from './nav.css';

export default function Header() {
  return (
    <nav className={styledNav}>
      <Container>
        <ul className={styledNavItems}>
          <li className={styledNavItem}>
            <NavLink className={styledNavLink} to="/">
              clairBuoyant
            </NavLink>
          </li>
          <li className={styledNavItem}>
            <NavLink className={styledNavLink} to="/home">
              Home
            </NavLink>
          </li>
          <li className={styledNavItem}>
            <NavLink className={styledNavLink} to="/about">
              About
            </NavLink>
          </li>
          <li className={styledNavItem}>
            <NavLink className={styledNavLink} to="/profile">
              Profile
            </NavLink>
          </li>
        </ul>
      </Container>
    </nav>
  );
}
