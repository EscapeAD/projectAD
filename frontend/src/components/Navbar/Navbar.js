import React from 'react';

//font awesome
import { faTerminal, faAtom } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

const Navbar = () => {
    return (
    <div className="hero-head">
        <nav className="navbar">
          <div className="container">
            <div className="navbar-brand">
              <a href="#top" className="navbar-item">
                <span className="icon">
                    <FontAwesomeIcon icon={faAtom} fixedWidth /> 
                </span>
                <span>Adam Tse</span>
              </a>
              <span className="navbar-burger burger" data-target="navbarMenuHeroB">
                <span></span>
                <span></span>
                <span></span>
              </span>
            </div>
            <div id="navbarMenuHeroB" className="navbar-menu">
              <div className="navbar-end">
                <a href="#bio" className="navbar-item is-active">
                  Bio
                </a>
                <a href="#projects" className="navbar-item">
                  Projects
                </a>
                <a href="#blog" className="navbar-item">
                  Blog
                </a>
                <span className="navbar-item">
                  <a href="#login" className="button is-info is-inverted">
                    <span className="icon">
                        <FontAwesomeIcon icon={faTerminal} />
                    </span>
                    <span>Log in</span>
                  </a>
                </span>
              </div>
            </div>
          </div>
        </nav>
      </div>
    )
}

export default Navbar;
