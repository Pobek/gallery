import React from 'react';
import ReactDOM from 'react-dom';
import { Route, BrowserRouter as Router, Switch } from 'react-router-dom';
import { Navbar, Nav } from 'react-bootstrap';

import './Navbar.css';
import 'bootstrap/dist/css/bootstrap.min.css';

import BookStore from './BookStore';
import Login from './Login';
import NotFound from './NotFound';

import * as serviceWorker from './serviceWorker';

ReactDOM.render(
  <React.StrictMode>
    <Router>
      <Navbar bg="gray" expand="lg" sticky="top">
        <Navbar.Brand href="/">Book Store</Navbar.Brand>
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="mr-auto">
            <Nav.Link href="/">Home</Nav.Link>
            <Nav.Link href="/login">Login</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Navbar>
      <Switch>
        <Route exact path="/" component={BookStore}/>
        <Route path="/login" component={Login}/>

        <Route component={NotFound}/>
      </Switch>
    </Router>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
