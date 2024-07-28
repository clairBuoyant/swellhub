import { NavLink } from 'react-router-dom';

import { Container } from '@components/ui/container';

// TODO(@kylejb): support tailwindcss classNames
export default function Header() {
  return (
    <nav>
      <Container>
        <ul>
          <li>
            <NavLink to="/">clairBuoyant</NavLink>
          </li>
          <li>
            <NavLink to="/home">Home</NavLink>
          </li>
          <li>
            <NavLink to="/about">About</NavLink>
          </li>
          <li>
            <NavLink to="/profile">Profile</NavLink>
          </li>
        </ul>
      </Container>
    </nav>
  );
}
