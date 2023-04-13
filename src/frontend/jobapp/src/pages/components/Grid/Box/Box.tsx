import React, { useState, useEffect } from "react";
import classNames from "classnames";
import PropTypes from "prop-types";

export interface BoxProps extends React.HTMLAttributes<HTMLElement> {
    children ?: React.ReactNode,
    

}

const propTypes = {
    className: PropTypes.string,
}
const Box = (props: BoxProps) => {
    return (
        <div className={classNames("box", props.className)}>
            {props.children}
        </div>
    )
}

Box.propTypes = propTypes

export default Box