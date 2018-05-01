import React, { Component } from 'react';
import { Platform, StyleSheet, View, Image } from 'react-native';

type Props = {};
class App extends Component<Props> {
  render() {
    return (
      <View style={styles.container}>
        <Image
          style={styles.logo}
          source={require('../images/logo.png')}
        />
      </View>
    );
  }
}

export default App;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#F5FCFF',
  },
  logo: {
    width: 220,
    height: 75,
  },
});
