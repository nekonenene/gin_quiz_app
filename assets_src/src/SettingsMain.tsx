import * as React from 'react';
import { Link } from 'react-router-dom';
import M from 'materialize-css';
import User from './model/User';

export interface Props {
  user: User;
}

interface SettingsState {
  username: string;
}

export default class SettingsMain extends React.Component<Props, SettingsState> {
  constructor(props: any) {
    super(props);
    this.state = {
      username: this.props.user.name,
    };

    this.onChangeUsername = this.onChangeUsername.bind(this);
    this.updateUsername = this.updateUsername.bind(this);
  }

  private updateUsername(): void {
    fetch('/api/user/update', {
      method: "POST",
      credentials: 'same-origin',
      headers: {
        "Content-Type": "application/json; charset=utf-8",
      },
      body: JSON.stringify({
        id: this.props.user.id,
        name: this.state.username || this.props.user.name,
        email: this.props.user.email,
      }),
    }).then(res => res.json())
      .then((resJson) => {
        if (resJson.error != null) {
          M.toast({ html: `ユーザーの更新に失敗しました (${resJson.error})`, classes: 'red darken-1' });
          return;
        }

        console.log('Success:', JSON.stringify(resJson));
        M.toast({ html: 'ユーザーの更新に成功しました', classes: 'green darken-1' });
      })
      .catch((error) => {
        console.error('Error:', error);
        M.toast({ html: 'ユーザーの更新に失敗しました', classes: 'red darken-1' });
      });
  }

  private onChangeUsername(event: any): void {
    this.setState({
      username: event.target.value
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
                <input type="text" id="username" className="validate" defaultValue={this.props.user.name} onChange={this.onChangeUsername} />
                <label htmlFor="username">ユーザー名</label>
              </div>
            </div>
            <div className="row center">
              <button className="btn-large waves-effect waves-light blue" onClick={this.updateUsername}>更新</button>
            </div>
            <div className="row center">
              <Link to="/" className="btn waves-effect waves-light red">トップへ</Link>
            </div>
          </div>
        </div>
      );
    } else {
      return <div></div>;
    }
  }
}
