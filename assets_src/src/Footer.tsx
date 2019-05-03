import * as React from 'react';

export interface Props {
  author: string;
  link: string;
}

export default class Footer extends React.Component<Props, {}> {
  public render() {
    return (
      <div className="container">
        <div className="section">
          <div className="row center">
            <p className="light">Created by <a href={this.props.link}>{this.props.author}</a></p>
          </div>
        </div>
      </div>
    );
  }
}
