import React from 'react';
import { Link } from 'react-router-dom';

const SingleChain = ({ chain, logo, name }) => (
  <div className='SingleChain'>
    <figure className='SingleChain-figure'>
      <img src='/image/clo-logo.png' alt={name} />
    </figure>
    <span className='SingleChain-name'>{name} - ({chain})</span>
    <Link to={`/chain/${chain}`} className='SingleChain-anchor'/>
  </div>
);

export default SingleChain;
