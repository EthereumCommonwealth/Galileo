import React, { PureComponent } from 'react';
import Layout from '../components/commons/Layout';
import Header from '../components/commons/Header';

class Home extends PureComponent {
  render() {
    return (
      <Layout>
        <Header />
      </Layout>
    );
  }
}

export default Home;
