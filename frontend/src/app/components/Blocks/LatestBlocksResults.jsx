import React from 'react';
import PropTypes from 'prop-types';

const LatestBlocksResults = ({ number, age, transactions, minedBy, reward }) => (
  <div className='LatestBlocksResults'>
    <a className='LatestBlocksResults-result'>{number}</a>
    <span className='LatestBlocksResults-title'>{age}</span>
    <a className='LatestBlocksResults-result'>{transactions}</a>
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
};

export default LatestBlocksResults;
