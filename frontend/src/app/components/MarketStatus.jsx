import React from 'react';
import PropTypes from 'prop-types';

const MarketStatus = ({ chain }) => (
  <div className='MarketStatus'>
    <div>
      <h2 className='MarketStatus-title'>Market</h2>
      <hr className='MarketStatus-divider' />
    </div>
    <div className='MarketStatus-content'>
      <div className='MarketStatus-content-element'>
        <span className='MarketStatus-content-element-chain'>1 {chain} </span>
        <span className='MarketStatus-content-element-value'> $80 USD</span>
      </div>
      <div className='MarketStatus-content-element'>
        <span className='MarketStatus-content-element-chain'>1 {chain} </span>
        <span className='MarketStatus-content-element-value'> 0.1 ETH</span>
      </div>
      <div className='MarketStatus-content-element'>
        <span className='MarketStatus-content-element-chain'>1 {chain} </span>
        <span className='MarketStatus-content-element-value'> 0.01 BTC</span>
      </div>
      <div className='MarketStatus-content-element'>
        <span className='MarketStatus-content-element-chain'>Market Cap</span>
        <span className='MarketStatus-content-element-value'>$200.000.000</span>
      </div>
    </div>
  </div>
);

MarketStatus.propTypes = {
  chain: PropTypes.string.isRequired,
};

export default MarketStatus;
