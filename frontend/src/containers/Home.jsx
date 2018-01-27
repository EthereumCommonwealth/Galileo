import React, { Component } from 'react';
import { connect } from 'react-redux';
import '../client/styles/Home.css';

class Home extends Component {

  render() {
    return (
      <div className="App">
        {this.props.name} App
      </div>
    );
  }
}

const mapStateToProps = (state) => {
  return {
    name: state.name,
  };
};

export default connect(mapStateToProps)(Home);
