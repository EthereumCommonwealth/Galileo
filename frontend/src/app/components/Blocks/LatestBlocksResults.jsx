import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';

const LatestBlocksResults = ({ number, age, transactions, minedBy, reward, chain }) => (
  <div className='LatestBlocksResults'>
    <Link
      className='LatestBlocksResults-result'
      to={`/chain/${chain}/block/${number}`}
    >
      #{number}
    </Link>
    <span className='LatestBlocksResults-title'>{age}</span>
    <Link
      className='LatestBlocksResults-result'
      to={`/chain/${chain}/block/${number}/transactions`}
    >
      {transactions}
    </Link>
    <a className='LatestBlocksResults-result'>{minedBy}</a>
    <span className='LatestBlocksResults-title'>{reward}</span>
  </div>
);

LatestBlocksResults.propTypes = {
  number: PropTypes.string.isRequired,
  age: PropTypes.string.isRequired,
  transactions: PropTypes.string.isRequired,
  minedBy: PropTypes.string.isRequired,
  reward: PropTypes.string.isRequired,
  chain: PropTypes.string.isRequired,
};

export default LatestBlocksResults;
