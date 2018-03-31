import React from 'react';
import PropTypes from 'prop-types';

const MarketStatus = ({ chain }) => (
  <div className='MarketStatus'>
    <div className='MarketStatus-content'>
      MarketStatus content
    </div>
  </div>
);

MarketStatus.propTypes = {
  chain: PropTypes.string.isRequired,
};

export default MarketStatus;
