import React from 'react';
import SingleChain from './SingleChain';

const ChainList = () => (
  <div className='ChainList container is-first-content'>
    <h2 className='ChainList-title'>Available Chains</h2>
    <div className='ChainList-content'>
      <SingleChain
        chain='CLO'
        logo='/image/clo-logo.png'
        name='Callisto'
      />
    </div>
  </div>
);

export default ChainList;
