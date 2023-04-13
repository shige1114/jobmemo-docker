import React from "react";
import classNames from "classnames";
import PropTypes from "prop-types";
import { Row, Col, Container } from "react-bootstrap";
import SideBar from "./sidebox/SideBar";



const Layout = () => {
    return (
        <>
            <Container>
                <Row className='side-bar'>
                    <SideBar/>
                </Row>
                <Row className='main-box'>

                </Row>
            </Container>
        </>
    )
}

