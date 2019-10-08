import React from 'react';
import { BrowserRouter as Router } from "react-router-dom";
import './App.scss';
import SiteRouter from './router/SiteRouter'
import { Navbar, Footer } from './components'

function App() {
  return (
    <section className="hero is-info is-large">
      <Router>
        <Navbar />
        <SiteRouter />
        <Footer />
      </Router>
    </section>
  );
}

export default App;
