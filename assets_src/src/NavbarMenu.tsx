import * as React from 'react';
import { Link } from 'react-router-dom';

export interface Props {
  isSignin: boolean;
}

export default class NavbarMenu extends React.Component<Props, {}> {
  public render() {
    var elements;
    if(this.props.isSignin) {
      elements =
        <div>
          <li><a href="/signout">ログアウト</a></li>
          <li><Link to="/settings">設定</Link></li>
        </div>
    } else {
      elements =
        <div>
          <li><a href="/oauth/google/signin">ログイン</a></li>
        </div>
    }

    return (
      <div>
        {elements}
      </div>
    );
  }
}
