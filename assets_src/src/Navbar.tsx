import * as React from 'react';
import User from './model/User';

export interface Props {
  user: User;
}

export default class Navbar extends React.Component<Props, {}> {
  constructor(props: any) {
    super(props)
  }

  componentDidUpdate(prevProps: Props) {
    console.log('prev', prevProps.user);
    console.log('current', this.props.user);
  }

  public render() {
    if (this.props.user.id > 0) {
      return (
        <nav className="red accent-3" role="navigation">
          <div className="nav-wrapper container">
            <a href="/" className="brand-logo">Quiz App</a>
            <ul id="nav-mobile" className="right hide-on-med-and-down">
              <li><a href="/signout">ログアウト</a></li>
            </ul>
          </div>
        </nav>
      );
    } else {
      return (
        <nav className="red accent-3" role="navigation">
          <div className="nav-wrapper container">
            <a href="/" className="brand-logo">Quiz App</a>
            <ul id="nav-mobile" className="right hide-on-med-and-down">
              <li><a href="/oauth/google/signin">ログイン</a></li>
            </ul>
          </div>
        </nav>
      );
    }
  }
}
