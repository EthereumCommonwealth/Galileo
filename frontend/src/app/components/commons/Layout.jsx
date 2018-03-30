import React, { PureComponent } from 'react'
import PropTypes from 'prop-types'
import Loader from './Loader'

class Layout extends PureComponent {
  constructor(props) {
    super(props)
    this.state = {
      mounted: false
    }
  }

  componentWillMount() {
    setTimeout(() => {
      this.setState({ mounted: true })
    }, 1500);
  }

  render() {
    const { className, children } = this.props;
    if (this.state.mounted) {
      return (
        <div className={className}>
          {children}
        </div>
      )
    }
    return <Loader />
  }
}

Layout.propTypes = {
  children: PropTypes.element.isRequired,
  className: PropTypes.string,
}

Layout.defaultProps = {
  className: '',
}

export default Layout
