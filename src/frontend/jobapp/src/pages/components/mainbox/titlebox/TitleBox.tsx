import React, { useState } from "react";
import classNames from "classnames";
import PropTypes from "prop-types";
import { Tab,Tabs } from "react-bootstrap";
import Box, { BoxProps } from "../../Grid/Box/Box";
import Oval,{  } from "../../Grid/Box/Box"

export interface TitleBoxProp{
    recrutment_status:string
    company_name:string
}

const propTypes= {
    company_name:PropTypes.string,
    recrutment_status:PropTypes.string,
}

const TitleBox = (prop:TitleBoxProp) =>{
    return (
        <Box className="title">
            <Oval className="name">
                {prop.company_name}
            </Oval>
            <Oval className="status">
                {prop.recrutment_status}
            </Oval>
        </Box>
    )

}
TitleBox.propTypes = propTypes

export default TitleBox