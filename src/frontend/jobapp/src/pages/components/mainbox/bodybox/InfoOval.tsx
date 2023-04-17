import React from "react";
import classNames from "classnames";
import PropTypes from "prop-types";
import { Row, Col, Container } from "react-bootstrap";
import Oval from "../../Grid/Oval/Oval";


export interface InfoOvalProps {
    company_id?: string,
    user_id?: string,

    title?: string,
    body?: string,
}

export interface InfoOvalArrayProps {
    datas?: InfoOvalProps[]
}
const propTypes = {
    company_id: PropTypes.string,
    user_id: PropTypes.string,
    title: PropTypes.string,
    body: PropTypes.string,
}
/*
const propTypes = {
    company_id:PropTypes.string,
    user_id:PropTypes.string,

    title:PropTypes.string,
    body:PropTypes.string,
}
*/

const InfoOval = (props: InfoOvalProps) => {
    console.log(props.title)
    return (
        <Oval className="info">
            <div className="info-title">
                {props.title}
            </div>
            <div className="info-body">
                {props.body}
            </div>
        </Oval>
    )
}


InfoOval.propTypes = propTypes

export default InfoOval