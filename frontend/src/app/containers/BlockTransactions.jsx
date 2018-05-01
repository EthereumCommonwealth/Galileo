import React, { PureComponent } from 'react';
import { hot } from 'react-hot-loader';
import Layout from '../components/commons/Layout';
import Header from '../components/commons/Header';
import Footer from '../components/commons/Footer';

class BlockTransactions extends PureComponent {
  constructor(props) {
    super(props);
    this.state = { redirect: false };
  }

  render() {
    const { match } = this.props;
    return (
      <Layout className='BlockTransactions'>
        <Header chain={match.params.chain} />
        <br/><br/>
        <br/><br/>
        <br/><br/>
        <br/><br/>
        Block Transactions
        <Footer />
      </Layout>
    );
  }
}

export default hot(module)(BlockTransactions);
