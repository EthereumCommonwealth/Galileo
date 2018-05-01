import React from 'react';
import { hot } from 'react-hot-loader';
import Layout from '../components/commons/Layout';
import Header from '../components/commons/Header';
import Footer from '../components/commons/Footer';

const Block = ({ match }) => (
  <Layout className='Block'>
    <Header chain={match.params.chain} />
    <br/><br/>
    <br/><br/>
    <br/><br/>
    <br/><br/>
    Block
    <Footer />
  </Layout>
)

export default hot(module)(Block);
