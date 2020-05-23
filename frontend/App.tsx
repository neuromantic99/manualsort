import React from 'react';
import { Component } from 'react';
import { StyleSheet, Text, View} from 'react-native';
import { connect, sendMsg } from "./Api";
import HelloWorld from './HelloWorld';
import LargeButton from './LargeButton';


//  TO START APP: 'npm run web' //

export default class App extends Component<Props> {

  render() {
    return (

      <div className="App">
        <LargeButton />
      </div>
    );
  }
}


const styles = StyleSheet.create({
  container: {
    alignItems: 'center',
    justifyContent: 'center',
  },
});

