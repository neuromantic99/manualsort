import React from 'react';
import Button from 'react-bootstrap/Button'
import { StyleSheet, TextInput, View, Text } from 'react-native';
import { connect, sendMsg } from "./Api";
import 'bootstrap/dist/css/bootstrap.min.css';

export default class LargeButton extends React.Component {
  
  constructor(props) {
    super(props);
    connect();
  }

  render() {

  function clickedYes() {
      sendMsg("Yes")
  }
  function clickedNo() {
      sendMsg("No")
  }

  return (
      <View style={styles.container}>
        <View style={styles.button}>
          <Button variant="primary" onClick={clickedYes} size="lg">Yes</Button>
       </View>
        <View style={styles.button}>
          <Button variant="primary" onClick={clickedNo} size="lg">No</Button>
     </View>
     </View>
    );
  }
 }; 

const styles = StyleSheet.create({

  container: {
    justifyContent: 'center',
    alignItems: 'center',
    flexDirection: 'row',
    //flex: 1
  },
  button: {
    //flex: 1,
    padding: 20,
    width: '20%',
  }

}
)


