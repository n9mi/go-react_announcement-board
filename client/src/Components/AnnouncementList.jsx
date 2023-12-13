import React from "react";
import { Col, Row } from "react-bootstrap";
import Announcement from "./Announcement";

const AnnouncementList = ({ announcemets }) => {
    return (
        <Row>
            {
                announcemets.map((ann) => (
                    <Col md={ 6 } xs={ 12 } key={ ann.id }>
                        <Announcement {...ann}></Announcement>
                    </Col>
                ))
            }
        </Row>
    );
}

export default AnnouncementList;