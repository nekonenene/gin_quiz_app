import * as React from 'react';
import User from './model/User';

export interface Props {
  user: User;
}

export default class Main extends React.Component<Props, {}> {
  private constructor(props: any) {
    super(props);
    console.log(this.props.user);
  }

  public render() {
    if (this.props.user.id > 0) {
      return (
        <div className="section no-pad-bot" id="index-banner">
          <div className="container">
            <h1 className="header center">Quiz App</h1>
            <div className="row center">
              <h5 className="header col s12 light">ログイン中だよ</h5>
            </div>
            <div className="row center">
              <a className="btn-large waves-effect waves-light blue" href="/signout">ログアウト</a>
            </div>
          </div>
        </div>
      );
    }
    return (
      <div className="section no-pad-bot" id="index-banner">
        <div className="container">
          <h1 className="header center">Quiz App</h1>
          <div className="row center">
            <h5 className="header col s12 light">まずはGoogleログインをしてみてください</h5>
          </div>
          <div className="row center">
            <a className="btn-large waves-effect waves-light red" href="/oauth/google/signin">Google ログイン</a>
          </div>
        </div>
      </div>
    );
  }
}
