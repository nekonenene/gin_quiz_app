import * as React from 'react';

export interface Props {
  content: string;
}

export default class MyComponent extends React.Component<Props, {}> {
  render() {
    return(
      <div className="container">
        <div className="section">
          <div className="row center">
            <p className="light">Powered by <a href="https://materializecss.com">Materialize.css</a></p>
          </div>
        </div>
      </div>
    )
  }
}
