import React, { Component } from 'react';
import { connect } from 'react-redux';
import '../client/styles/App.css';

class App extends Component {

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

export default connect(mapStateToProps)(App);
