import React, { PureComponent } from 'react';
import Layout from '../components/commons/Layout';
import { Redirect } from 'react-router'
import Header from '../components/commons/Header';

class Home extends PureComponent {
  constructor(props) {
    super(props)
    this.state = { redirect: false }
  }

  componentWillMount() {
    if (this.props.match.path === '/') {
      this.setState({ redirect: true })
    }
  }

  render() {
    if (this.state.redirect) {
      return <Redirect to='chain/CLO' />
    }
    return (
      <Layout className='Home'>
        <Header />
      </Layout>
    );
  }
}

export default Home;
