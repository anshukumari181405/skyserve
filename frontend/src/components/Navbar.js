// Navbar.js
import React from 'react';
import { NavLink } from 'react-router-dom';
import '../CSS/Navbar.css'; // Make sure to adjust the path as needed

const Navbar = () => {
  return (
    <nav className="navbar">
      <div>
        <NavLink className="navbar-link" to="/home" activeClassName="active">Home</NavLink>
        <NavLink className="navbar-link" to="/login" activeClassName="active">Login</NavLink>
        <NavLink className="navbar-link" to="/Signup" activeClassName="active">SignUp</NavLink>
        <NavLink className="navbar-link" to="/forgot-password" activeClassName="active">Forget Password</NavLink>

        <NavLink className="navbar-link" to="/upload" activeClassName="active">Upload file</NavLink>
        <NavLink className="navbar-link" to="/map" activeClassName="active">Map</NavLink>



        

      </div>
    </nav>
  );
};

export default Navbar;
