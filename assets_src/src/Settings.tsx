import * as React from 'react';
import { Link } from 'react-router-dom';
import Navbar from './Navbar';
import Main from './SettingsMain';
import Footer from './Footer';
import { getCookieByName } from './util';
import User, { defaultUser } from './model/User';

interface SettingsState {
  loading: boolean;
  user: User;
}

export default class Settings extends React.Component<{}, SettingsState> {
  private constructor(props: any) {
    super(props);
    this.state = {
      loading: true,
      user: defaultUser,
    };
  }

  private fetchCurrentUser(): void {
    fetch('/api/user/current', {
      credentials: 'same-origin',
    }).then(res => res.json())
      .then((resJson) => {
        const userJson = resJson.user;
        const user: User = {
          id: userJson.id,
          name: userJson.name,
          email: userJson.email,
        };
        console.log('Success:', JSON.stringify(userJson));
        this.setState({
          loading: false,
          user,
        });
      })
      .catch((error) => {
        console.error('Error:', error);
        this.setState({
          loading: false,
        });
      });
  }

  public componentDidMount(): void {
    this.fetchCurrentUser();
    console.log(getCookieByName('session_id'));
  }

  public render() {
    return (
      <div>
        <Navbar
          user={this.state.user}
        />
        <Main
          user={this.state.user}
        />
        <div className="center">
          <Link to="/" className="btn waves-effect waves-light red">トップへ</Link>
        </div>
        <Footer
          author="ハトネコエ"
          link="https://twitter.com/nekonenene"
        />
      </div>
    );
  }
}
