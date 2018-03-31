import React, { PureComponent } from 'react';
import Layout from '../components/commons/Layout';
import Header from '../components/commons/Header';
import Footer from '../components/commons/Footer';

class Transaction extends PureComponent {
  constructor(props) {
    super(props);
    this.state = { redirect: false };
  }

  render() {
    const { match } = this.props;
    return (
      <Layout className='Transaction'>
        <Header chain={match.params.chain} />
        <Footer />
      </Layout>
    );
  }
}

export default Transaction
