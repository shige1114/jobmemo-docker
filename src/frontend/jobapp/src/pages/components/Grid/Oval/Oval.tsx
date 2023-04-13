import React, { useState, useEffect } from "react";
import classNames from "classnames";
import PropTypes from "prop-types";

export interface OvalProps extends React.HTMLAttributes<HTMLElement> {
    className: string | "info" | "name" | "plus" | "list"
}
const propTypes = {
    className: PropTypes.string,
}

const Oval = (props: OvalProps) => {
   
    return (
        <div className={classNames("oval",props.className)}>
            {props.children}
        </div>
    )
}

Oval.propTypes = propTypes

export default Oval