import React, { Component } from 'react';
import './Book.css'

class Book extends Component{
  render(){
    const {title, img, description} = this.props;
    const tags = this.props.tags.map((tag, index) => (
      <li key={index}>
        {tag}
      </li>
    ));

    return (
      <div className="book">
        <div className="book-image">
          <img src={img} alt={title}/>
        </div>
        <div className="book-content">
          <h2 className="book-title">{title}</h2>
          <p>{description}</p>
          <table>
          <tr>
            <td>Tags</td>
            <td>
              <ul>{tags}</ul>
            </td>
          </tr>
        </table>
        </div>
      </div>
    );
  }
}

export default Book;
