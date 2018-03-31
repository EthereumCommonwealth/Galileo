import React, { PureComponent } from 'react';
import Layout from '../components/commons/Layout';
import Header from '../components/commons/Header';

class Block extends PureComponent {
  constructor(props) {
    super(props);
    this.state = { redirect: false };
  }

  render() {
    return (
      <Layout className='Block'>
        <Header />
      </Layout>
    );
  }
}

export default Block
