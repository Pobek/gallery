import React, { Component } from 'react';
import Book from './Book';
import PropTypes from 'prop-types';
import './BookList.css'

class BookList extends Component{
  static defaultProps = {
    books: [
      {
        title: "The Spook's",
        tags: ["thriller", "action"],
        img: "the_spooks_apprentice.jpg",
        description: "A wonderful and terrifying series by a new writer about a young boy training to be an exorcist. Thomas Ward is the seventh son of a seventh son and has been apprenticed to the local Spook. The job is hard, the Spook is distant and many apprentices have failed before Thomas."
      },
      {
        title: "The Spook's",
        tags: ["thriller", "action"],
        img: "the_spooks_apprentice.jpg",
        description: "A wonderful and terrifying series by a new writer about a young boy training to be an exorcist. Thomas Ward is the seventh son of a seventh son and has been apprenticed to the local Spook. The job is hard, the Spook is distant and many apprentices have failed before Thomas."
      },
      {
        title: "The Spook's",
        tags: ["thriller", "action"],
        img: "the_spooks_apprentice.jpg",
        description: "A wonderful and terrifying series by a new writer about a young boy training to be an exorcist. Thomas Ward is the seventh son of a seventh son and has been apprenticed to the local Spook. The job is hard, the Spook is distant and many apprentices have failed before Thomas."
      }
    ]
  }

  static propTypes = {
    books: PropTypes.arrayOf(PropTypes.object).isRequired
  }

  render(){
    const books = this.props.books.map((r, index) => (
      <Book 
        key={index}
        {...r}
      />
    ))

    return(
      <div className="book-list">
        {books}
      </div>
    );
  }
}

export default BookList;