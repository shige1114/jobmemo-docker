import React, { useState } from "react";
import classNames from "classnames";
import PropTypes from "prop-types";
import { Tab,Tabs } from "react-bootstrap";
import InfoOval, { InfoOvalProps } from "./InfoOval";
import { BoxProps } from "../../Grid/Box/Box";
import InfoBox, { InfoBoxProp } from "./InfoBox";


export interface BodyBarBoxProp extends BoxProps {
    datas: InfoBoxProp[]
    recrute_status: string
}
const propTypes = {
    datas: PropTypes.array,
    recrute_status: PropTypes.string
}

const BodyBarBox = (prop: BodyBarBoxProp) => {
    const [datas, setInfoDatas] = useState(prop.datas)
    const [key, setKey] = useState(prop.recrute_status)
    
    function Item(datas: InfoBoxProp[]) {
        if (datas != undefined) {
            if (datas.length != 0) {
                datas.forEach((item:any) => {
                    const { recrute_status,...info } = item
                    return (
                        <Tab
                        eventKey={recrute_status}
                        title={recrute_status}
                        >
                            <InfoBox {...info}/>
                        </Tab>
                    )
                })
            }
        }
        return (
            <></>
        )
        
    }

    return (
        <div className="box-dody">
            <Tabs
            id="company-info-bar"
            activeKey={key}
            onSelect={(k:any)=>setKey(k)}
            className="mb-3"
            >
                {Item(datas)}
            </Tabs>
        </div>
    )
}

BodyBarBox.propTypes = propTypes
export default BodyBarBox