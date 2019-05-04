import * as React from 'react';
import * as ReactDOM from 'react-dom';
import Navbar from './Navbar';
import Main from './Main';
import Footer from './Footer';
import { getCookieByName } from './util';
import User, { defaultUser } from './model/User';

export interface Props {
}

type AppState = {
  loading: boolean;
  user: User;
};

class App extends React.Component<Props, AppState> {
  constructor(props: any) {
    super(props);
    this.state = {
      loading: true,
      user: defaultUser
    };
  }

  private fetchCurrentUser() {
    fetch('/api/user/current', {
      credentials: 'same-origin'
    }).then(res => res.json())
      .then(resJson => {
        const userJson = resJson.user;
        const user: User = {
          id: userJson.id,
          name: userJson.name,
          email: userJson.email,
        }
        console.log('Success:', JSON.stringify(userJson))
        this.setState({
          loading: false,
          user: user
        })
      })
      .catch(error => {
        console.error('Error:', error);
        this.setState({
          loading: false,
        })
      });
  }

  componentDidMount() {
    this.fetchCurrentUser();
    console.log(getCookieByName('session_id'));
  }

  public render() {
    return (
      <div>
        <Navbar
          user={ this.state.user }
        />
        <Main
          user={ this.state.user }
        />
        <Footer
          author="ハトネコエ"
          link="https://twitter.com/nekonenene"
        />
      </div>
    );
  }
};

ReactDOM.render(
  <App />,
  document.getElementById('app'),
);
