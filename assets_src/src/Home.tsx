import * as React from 'react';
import M from 'materialize-css';
import Navbar from './Navbar';
import Main from './HomeMain';
import Footer from './Footer';
import { getCookieByName } from './util';
import User, { defaultUser } from './model/User';

interface AppState {
  loading: boolean;
  user: User;
}

export default class App extends React.Component<{}, AppState> {
  public constructor(props: any) {
    super(props);
    this.state = {
      loading: true,
      user: defaultUser,
    };
  }

  public componentDidMount(): void {
    M.AutoInit();
    this.fetchCurrentUser();
    console.log(getCookieByName('session_id'));
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

  public render() {
    return (
      <div>
        <Navbar
          user={this.state.user}
        />
        <Main
          user={this.state.user}
        />
        <Footer
          author="ハトネコエ"
          link="https://twitter.com/nekonenene"
        />
      </div>
    );
  }
}
