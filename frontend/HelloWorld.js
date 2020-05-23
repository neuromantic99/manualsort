import React from 'react';

const HelloWorld = () => {
  
  function sayHello() {
    alert('Fuck you');
  }
  
  return (
    <button onClick={sayHello}>Click me!</button>
  );
};

export default HelloWorld;
