import React from 'react';
import PropTypes from 'prop-types';
import LatestBlocksResults from './LatestBlocksResults';

const LatestBlocks = ({ chain }) => (
  <div className='LatestBlocks container is-first-content'>
    <div className='LatestBlocks-top'>
      <h2 className='LatestBlocks-top-title'>Latest {chain} Blocks</h2>
      <a className='LatestBlocks-top-seemore'> See more...</a>
    </div>
    <div className='LatestBlocks-content'>
      <div className='LatestBlocks-content-head'>
        <span className='LatestBlocks-content-title'>#</span>
        <span className='LatestBlocks-content-title'>Age</span>
        <span className='LatestBlocks-content-title'>Transactions</span>
        <span className='LatestBlocks-content-title'>Mined by</span>
        <span className='LatestBlocks-content-title'>Reward</span>
      </div>
      <LatestBlocksResults
        number='55000'
        age='19 segs ago'
        transactions='100 Txs'
        minedBy='ETC Pool'
        reward='3.15 ETC'
      />
      <LatestBlocksResults
        number='55000'
        age='19 segs ago'
        transactions='100 Txs'
        minedBy='ETC Pool'
        reward='3.15 ETC'
      />
      <LatestBlocksResults
        number='55000'
        age='19 segs ago'
        transactions='100 Txs'
        minedBy='ETC Pool'
        reward='3.15 ETC'
      />
      <LatestBlocksResults
        number='55000'
        age='19 segs ago'
        transactions='100 Txs'
        minedBy='ETC Pool'
        reward='3.15 ETC'
      />
      <LatestBlocksResults
        number='55000'
        age='19 segs ago'
        transactions='100 Txs'
        minedBy='ETC Pool'
        reward='3.15 ETC'
      />
    </div>
  </div>
);

LatestBlocks.propTypes = {
  chain: PropTypes.string.isRequired,
};

export default LatestBlocks;
