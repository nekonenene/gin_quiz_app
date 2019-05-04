import * as React from 'react';
import * as ReactDOM from 'react-dom';
import Navbar from './Navbar';
import Main from './Main';
import Footer from './Footer';
import { getCookieByName } from './util';

export interface Props {
}

type AppState = {
  user: JSON
};

class App extends React.Component<Props, AppState> {
  constructor(props: any) {
    super(props);
    this.state = {
      user: JSON.parse('{}')
    };
  }

  private fetchCurrentUser() {
    fetch('/api/user/current', {
      credentials: 'same-origin'
    }).then(res => res.json())
      .then(resJson => {
        const userJson = resJson.user;
        this.setState({
          user: userJson
        })
        console.log('Success:', JSON.stringify(userJson))
      })
      .catch(error => console.error('Error:', error));
  }

  componentDidMount() {
    this.fetchCurrentUser()
      console.log(getCookieByName('session_id'));
  }

  public render() {
    return (
      <div>
        <Navbar
          user = {this.state.user}
        />
        <Main />
        <Footer
          author = "ハトネコエ"
          link = "https://twitter.com/nekonenene"
        />
      </div>
    );
  }
};

ReactDOM.render(
  <App />,
  document.getElementById('app'),
);
