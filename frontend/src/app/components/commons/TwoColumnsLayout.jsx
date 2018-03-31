import React from 'react'
import PropTypes from 'prop-types'

const TwoColumnsLayout = ({ children }) => (
  <div className='TwoColumnsLayout container'>
    {children}
  </div>
)

TwoColumnsLayout.propTypes = {
  children: PropTypes.node.isRequired,
}

export default TwoColumnsLayout;
