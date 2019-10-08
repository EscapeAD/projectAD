import React from 'react';
import { BrowserRouter as Router } from "react-router-dom";
import './App.scss';
import SiteRouter from './router/SiteRouter'
import Navbar from './components/Navbar/Navbar'
import Footer from './components/Footer/Footer'

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
