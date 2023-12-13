import React from "react";
import { Col } from "react-bootstrap";

const Title = ({ title }) => {
    return (
        <Col md="auto"><h1 className="page-title">{ title }</h1></Col>
    );
}

export default Title;