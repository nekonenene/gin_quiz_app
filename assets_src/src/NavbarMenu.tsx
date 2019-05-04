import * as React from 'react';

export interface Props {
  isSignin: boolean;
}

export default class NavbarMenu extends React.Component<Props, {}> {
  public render() {
    const signinElement = this.props.isSignin
      ? <li><a href="/signout">ログアウト</a></li>
      : <li><a href="/oauth/google/signin">ログイン</a></li>;
    return (
      signinElement
    );
  }
}
