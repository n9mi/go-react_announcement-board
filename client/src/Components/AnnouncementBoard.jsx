import React, { useEffect, useState } from "react";
import { Container, Row, Alert } from "react-bootstrap"
import { useFetch } from "../utils/api";
import Title from "./Title";
import AnnouncementList from "./AnnouncementList";

const AnnouncementBoard = () => {
    const { data: announcements, isPending, err } = useFetch();
    const [ displayErr, setDiplayErr ] = useState(err != null ? true : false);
    
    useEffect(() => {
        setDiplayErr(err != null ? true : false);
    }, [ err ]);

    return (
        <Container>
            <Row className="main-title justify-content-md-center mt-3">
               <Title title={ "Announcement Board" } />
               { displayErr && ( <Alert variant="danger"> { "Terdapat kesalahan!" } </Alert> ) }
               { isPending && ( <i className="fa-solid fa-circle-notch fa-spin"></i> ) }
               { (announcements && !displayErr ) && ( <AnnouncementList announcemets={ announcements } /> ) }
            </Row>
        </Container>
    );
}

export default AnnouncementBoard;