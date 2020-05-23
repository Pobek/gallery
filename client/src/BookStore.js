import React, { Component } from 'react';
import BookList from './BookList';
import Navbar from './Navbar';

class BookStore extends Component{
  render(){
    return (
      <div className="App">
        <Navbar/>
        <BookList/>
      </div>
    );
  }
}

export default BookStore;
