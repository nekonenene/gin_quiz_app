import * as React from 'react';

export interface Props {
  user: JSON;
}

export default class Navbar extends React.Component<Props, {}> {
  constructor(props: any) {
    super(props)
    console.log(this.props.user);
  }

  public render() {
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
