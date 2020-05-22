import React, { Component } from 'react';
import Book from './Book';

class BookStore extends Component{
  render(){
    return (
      <div className="App">
        <Book 
          title="The Spook's"
          tags={["thriller", "action"]}
          img="the_spooks_apprentice.jpg"
          description="A wonderful and terrifying series by a new writer about a young boy training to be an exorcist. Thomas Ward is the seventh son of a seventh son and has been apprenticed to the local Spook. The job is hard, the Spook is distant and many apprentices have failed before Thomas."
        />
      </div>
    );
  }
}

export default BookStore;
