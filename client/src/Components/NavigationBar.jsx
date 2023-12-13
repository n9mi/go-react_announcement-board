import React from "react";
import { NavLink } from "react-router-dom";
import { Nav } from "react-bootstrap";

const NavigationBar = () => {
    const styleActiveNav = (e) => {
        return { textDecoration: e.isActive ? "underline" : "none" }
    };

    return (
        <Nav className="justify-content-end mt-3 mb-3" activeKey="/">
            <Nav.Item>
                <NavLink className="navigation__item" 
                    to="/" 
                    style={styleActiveNav}>
                        Home</NavLink>
            </Nav.Item>
            <Nav.Item>
                <NavLink 
                    className="navigation__item" 
                    to="/create"
                    style={styleActiveNav} >
                        Create</NavLink>
            </Nav.Item>
        </Nav>
    );
}

export default NavigationBar;