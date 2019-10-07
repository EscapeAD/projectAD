import React from 'react';
import './App.scss';

import Navbar from './components/Navbar/Navbar'
import Portfolio from './components/Portfolio/Portfolio'
import Footer from './components/Footer/Footer'

function App() {
  return (
    <section className="hero is-info is-large">
      <Navbar />
      <Portfolio />
      <Footer />
    </section>
  );
}

export default App;
