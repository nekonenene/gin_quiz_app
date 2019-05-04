import * as React from 'react';
import User from './model/User';

export interface Props {
  user: User;
}

export default class SettingsMain extends React.Component<Props, {}> {
  private updateUsername(): void {
    fetch('/api/user/current', {
      credentials: 'same-origin',
    }).then(res => res.json())
      .then((resJson) => {
        const userJson = resJson.user;
        console.log('Success:', JSON.stringify(userJson));
      })
      .catch((error) => {
        console.error('Error:', error);
      });
  }

  public render() {
    if (this.props.user.id > 0) {
      return (
        <div className="section no-pad-bot" id="index-banner">
          <div className="container">
            <h4 className="header center">ユーザー設定</h4>
            <div className="row center">
              <div className="input-field col s12">
                <input type="text" id="username" className="validate" defaultValue={this.props.user.name} />
                <label className="active" htmlFor="username">ユーザー名</label>
              </div>
            </div>
            <div className="row center">
              <button className="btn-large waves-effect waves-light blue" onClick={this.updateUsername}>更新</button>
            </div>
          </div>
        </div>
      );
    } else {
      return <div></div>;
    }
  }
}
