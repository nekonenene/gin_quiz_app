import * as React from 'react';

export interface Props {
}

export default class Main extends React.Component<Props, {}> {
  public render() {
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
