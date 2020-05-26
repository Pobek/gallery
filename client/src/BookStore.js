import React, { Component } from 'react';
import BookList from './BookList';

class BookStore extends Component{
  render(){
    return (
      <div className="App">
        <BookList/>
      </div>
    );
  }
}

export default BookStore;
