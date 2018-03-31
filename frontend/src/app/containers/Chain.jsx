import React, { PureComponent } from 'react';
import PropTypes from 'prop-types'
import Layout from '../components/commons/Layout';
import TwoColumnsLayout from '../components/commons/TwoColumnsLayout';
import Header from '../components/commons/Header';
import Footer from '../components/commons/Footer';
import MarketStatus from '../components/MarketStatus';
import ChainList from '../components/Chain/ChainList';
import LatestBlocks from '../components/Blocks/LatestBlocks';
import TransactionsPerDay from '../components/Transactions/TransactionsPerDay';

class Chain extends PureComponent {
  constructor(props) {
    super(props);
    this.state = { redirect: false };
  }

  render() {
    const { match } = this.props;
    return (
      <Layout className='Chain'>
        <Header chain={match.params.chain} />
        <ChainList />
        <LatestBlocks chain={match.params.chain} />
        <TwoColumnsLayout>
          <TransactionsPerDay />
          <MarketStatus chain={match.params.chain} />
        </TwoColumnsLayout>
        <Footer />
      </Layout>
    );
  }
}

Chain.propTypes = {
  match: PropTypes.object,
};

export default Chain;
