import React from 'react';
import { hot } from 'react-hot-loader';
import Layout from '../components/commons/Layout';
import Header from '../components/commons/Header';
import Footer from '../components/commons/Footer';

const NotFound = () => (
  <Layout className='NotFound'>
    <Header />
    <Footer />
  </Layout>
);

export default hot(module)(NotFound);
