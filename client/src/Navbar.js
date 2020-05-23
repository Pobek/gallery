import React, { Component } from 'react';
import './Navbar.css';

class Navbar extends Component{
  render(){
    return(
      <header>
        <h2><a>Book Store</a></h2>
        <nav>
          <li><a>New Book</a></li>
          <li><a>Home</a></li>
          <li><a>About</a></li>
        </nav>
      </header>
    );
  }
}

export default Navbar;