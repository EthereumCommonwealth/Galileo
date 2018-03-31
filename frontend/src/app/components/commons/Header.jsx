import React from 'react';
import PropTypes from 'prop-types'
import { Link } from 'react-router-dom';

const Header = ({ chain }) => (
  <header className='Header'>
    <div className='container'>
      <div className='Header-menu'>
        <Link
          className='Header-title'
          to='/'
        >
          <figure className='Header-logo'>
            <img src='/image/logo.png' alt='Galileo Logo' />
          </figure>
        </Link>
        <nav className='Header-nav'>
          <ul className='Header-nav-content no-margin no-padding'>
            <li className='Header-nav-content-elem'>
              <span
                className='Header-nav-content-elem-active'
              >
                Active chain: {chain}
              </span>
            </li>
            <li className='Header-nav-content-elem'>
              <Link
                className='Header-nav-content-elem-anchor'
                to='/blocks/'
              >
                Blocks
              </Link>
            </li>
            <li className='Header-nav-content-elem'>
              <Link
                className='Header-nav-content-elem-anchor'
                to='/page-2/'
              >
                Tokens
              </Link>
            </li>
            <li className='Header-nav-content-elem'>
              <div className='Header-nav-content-elem-search'>
                <input
                  type='text'
                  className='Header-nav-content-elem-search-input'
                  placeholder='Search...'
                />
                <i className='Header-nav-content-elem-search-icon fas fa-search' />
              </div>
            </li>
          </ul>
        </nav>
      </div>
    </div>
  </header>
)

Header.propTypes = {
  chain: PropTypes.string,
};

export default Header
