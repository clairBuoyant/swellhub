import { BrowserRouter, NavLink } from 'react-router-dom';

import { NAVBAR_PATH } from './navbar.constants';
import './navbar.css';

export const Navbar = () => {
  return (
    <div className="nav-container">
      <BrowserRouter>
        <nav className="nav">
          <p className="nav-title">
            <NavLink to="/">{NAVBAR_PATH.clairBuoyantTitle}</NavLink>
          </p>
          <ul className="nav-items">
            <li className="nav-link">
              <NavLink to="/">{NAVBAR_PATH.home}</NavLink>
            </li>
            <li className="nav-link">
              <NavLink to="/">{NAVBAR_PATH.about}</NavLink>
            </li>
            <li className="nav-link">
              <NavLink to="/">{NAVBAR_PATH.user}</NavLink>
            </li>
          </ul>
        </nav>
      </BrowserRouter>
    </div>
  );
};
