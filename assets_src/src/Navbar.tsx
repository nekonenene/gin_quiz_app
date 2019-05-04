import * as React from 'react';
import NavbarMenu from './NavbarMenu';
import User from './model/User';

export interface Props {
  user: User;
}

export default class Navbar extends React.Component<Props, {}> {
  public render() {
    const isSignin = this.props.user.id > 0;
    return (
      <nav className="red accent-3" role="navigation">
        <div className="nav-wrapper container">
          <a href="/" className="brand-logo">Quiz App</a>
          <ul className="right hide-on-small-and-down">
            <NavbarMenu
              isSignin={isSignin}
            />
          </ul>
          <ul className="right hide-on-med-and-up">
            <li><a className="dropdown-trigger show-on-small" href="#!" data-target="menuDropdown"><i className="material-icons">menu</i></a></li>
          </ul>
          <ul className="dropdown-content" id="menuDropdown">
            <NavbarMenu
              isSignin={isSignin}
            />
          </ul>
        </div>
      </nav>
    );
  }
}
