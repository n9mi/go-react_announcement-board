import { useEffect, useState } from "react";
import { Container, Form, Button, Alert, Row, Col } from "react-bootstrap";
import { useParams } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import { fetchData } from "../utils/api";
import Title from "./Title";

const Save = ({ apiHandler, pageTitle }) => {
    const { id } = useParams();
    const [ title, setTitle ] = useState("");
    const [ content, setContent ] = useState("");
    const [ errMessage, setErrMessage ] = useState(null);
    const [ show, setShowAlert ] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        if (id > 0) {
            fetchData(id)
            .then((res) => {
                if (res.status === "OK") {
                    setTitle(res.data.title);
                    setContent(res.data.content);
                }
            })
            .catch((err) => {
                navigate("/");
            });
        }
    }, []);
    
    const handleSubmit = (e) => {
        e.preventDefault();

        const method = id > 0 ? "PUT" : "POST"; 
        apiHandler(method, id, {
            title,
            content
        })
        .then((response) => {
            if (response.status === "OK") {
                navigate("/");
            }
        })
        .catch((err) => {
            setErrMessage(err.response.data.message);
        });
    }

    const handleDelete = (e) => {
        apiHandler("DELETE", id)
        .then((response) => {
            if (response.status === "OK") {
                navigate("/");
            }
        })
        .catch((err) => {
            setErrMessage(err.response.data.message);
        });
    }

    return (
        <Container>
            <Container className="w-75">
                <Row>
                    <Col md={ 10 }>
                        <Title title={ pageTitle }/>
                    </Col>
                    {
                        id && (
                            <Col md={ 2 }>
                                <Button type="button" 
                                    variant="outline-danger"
                                    className="w-100 mb-3"
                                    onClick={ handleDelete }> Hapus </Button>
                            </Col>
                        )
                    }
                </Row>
                { errMessage && (
                    <Alert variant="danger" onClose={() => setShowAlert(false)} dismissible>
                        { errMessage }
                    </Alert> 
                ) }
                <Form onSubmit={ handleSubmit }>
                    <Form.Group className="mb-3">
                        <Form.Label className="save__label">Title</Form.Label>
                        <Form.Control type="text" value={ title } className="save__form"
                            onChange={ (e) => setTitle(e.target.value) } />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label className="save__label">Content</Form.Label>
                        <Form.Control as="textarea" rows={3} value={ content } className="save__form"
                            onChange={ (e) => setContent(e.target.value) }/>
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Button type="submit" 
                            variant="outline-success" 
                            className="w-100" disabled={ title == "" | content == "" }>SUBMIT</Button>
                    </Form.Group>
                </Form>
            </Container>
        </Container>
    );
};

export default Save; 