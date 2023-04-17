import React, { useState } from "react";
import classNames from "classnames";
import PropTypes from "prop-types";
import { Tab, Tabs } from "react-bootstrap";
import InfoOval, { InfoOvalProps } from "./InfoOval";
import { BoxProps } from "../../Grid/Box/Box";


export interface InfoBoxProp extends BoxProps {
    datas?: InfoOvalProps[]
    recrute_status?: string
}

const propTypes = {
    datas: PropTypes.array,
    recrute_status: PropTypes.string
}

const InfoBox = (prop: InfoBoxProp) => {

    const [datas, setInfoDatas] = useState(prop.datas)
    console.log(datas)
    const Item = () => {
        if (datas != undefined) {
            if (datas.length != 0) {
                return (<>{
                    datas.map((item: InfoOvalProps) => {
                        return (
                            <InfoOval {...item} />
                        )
                        //{console.log(item.body)}
                    })
                }</>)
            }
        }
        return (
            <p>失敗しました</p>
        )
    }

    return (
        <div className="info-list">
            <Item />
        </div>
    )
}

InfoBox.propTypes = propTypes

export default InfoBox