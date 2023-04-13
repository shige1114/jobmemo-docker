import React, { useState, useEffect } from "react";
import classNames from "classnames";
import PropTypes from "prop-types";
import { Container, Col } from "react-bootstrap"


export interface CompanyStatusProps extends React.HTMLAttributes<HTMLElement> {
    company_id?: string
    user_id?: string

    //company_status
    recruitment_status?: string
    //status
    is_accepting?: boolean
    is_scheduling?: boolean
}
const propTypes = {
    company_id: PropTypes.string,
    user_id: PropTypes.string,

    recruitment_status: PropTypes.string,
    is_accepting: PropTypes.bool,
    is_scheduling: PropTypes.bool,
}

const CompanyStatus = (props: CompanyStatusProps) => {
    const [recruteStatus, setStatus] = useState(props.recruitment_status)
    const [isAccepting, setAccepting] = useState(props.is_accepting)
    const [isScheduling,setScheduling] = useState(props.is_scheduling)
    useEffect(() => {
        if (isAccepting){

        }
        if (isScheduling){

        }

        setStatus("jal")
    }, [isAccepting,isScheduling])
    
    return (
        <div className='company-status'>
            <Container>
                <Col>
                    <p>{recruteStatus}</p>
                </Col>
                <Col>
                    
                </Col>
            </Container>
        </div>
    )
}
CompanyStatus.propTypes = propTypes

export default CompanyStatus