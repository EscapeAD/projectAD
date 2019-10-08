import React from 'react';
import { NavLink } from "react-router-dom";

//font awesome
import { faTerminal, faAtom } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

const Navbar = () => {
    return (
    <div className="hero-head">
        <nav className="navbar">
          <div className="container">
            <div className="navbar-brand">
              <NavLink to="/" exact className="navbar-item">
                <span className="icon">
                    <FontAwesomeIcon icon={faAtom} fixedWidth /> 
                </span>
                <span>Adam Tse</span>
              </NavLink>              
                <span className="navbar-burger burger" data-target="navbarMenuHeroB">
                <span></span>
                <span></span>
                <span></span>
              </span>
            </div>
            <div id="navbarMenuHeroB" className="navbar-menu">
              <div className="navbar-end">
                <NavLink to="/" exact activeClassName="is-active" className="navbar-item">Bio</NavLink>
                <span className="navbar-item">
                  <NavLink to="/login" exact activeClassName="is-active" className="button is-info is-inverted">
                    <span className="icon">
                      <FontAwesomeIcon icon={faTerminal} />
                    </span>
                    <span>Log in</span>
                  </NavLink>
                </span>
              </div>
            </div>
          </div>
        </nav>
      </div>
    )
}

export default Navbar;
