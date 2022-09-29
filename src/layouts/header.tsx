import Container from '@components/container';

import { StyledNav, StyledNavItem, StyledNavItems, StyledNavLink } from './nav';

export default function Header() {
  return (
    <StyledNav>
      <Container>
        <StyledNavItems>
          <StyledNavItem>
            {/* TODO: align this to left and fix dims */}
            <StyledNavLink to="/">clairBuoyant</StyledNavLink>
          </StyledNavItem>
          <StyledNavItem className="nav-link">
            <StyledNavLink to="/home">Home</StyledNavLink>
          </StyledNavItem>
          <StyledNavItem className="nav-link">
            <StyledNavLink to="/about">About</StyledNavLink>
          </StyledNavItem>
          <StyledNavItem className="nav-link">
            <StyledNavLink to="/profile">Profile</StyledNavLink>
          </StyledNavItem>
        </StyledNavItems>
      </Container>
    </StyledNav>
  );
}
