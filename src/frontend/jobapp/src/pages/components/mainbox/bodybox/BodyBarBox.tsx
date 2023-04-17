import React, { useEffect, useState } from "react";
import classNames from "classnames";
import PropTypes, { func } from "prop-types";
import { Tab, Tabs } from "react-bootstrap";
import InfoOval, { InfoOvalProps } from "./InfoOval";
import { BoxProps } from "../../Grid/Box/Box";
import InfoBox, { InfoBoxProp } from "./InfoBox";


export interface BodyBarBoxProp extends BoxProps {
    datas?: InfoBoxProp[]
    recrute_status: string
}
const propTypes = {
    datas: PropTypes.array,
    recrute_status: PropTypes.string
}

const BodyBarBox = (prop: BodyBarBoxProp) => {
    const [datas, setInfoDatas] = useState<InfoBoxProp[]>([])
    const [key, setKey] = useState(prop.recrute_status)
    useEffect(()=>{
        checkDatas(prop.datas)
    },[])
    function checkDatas(datas:any){
        if (datas != undefined && datas.length > 0){
            setInfoDatas(datas)
        }
    }

    function Item() {
        if (datas.length > 0) {
            const html = (datas.map((item: InfoBoxProp) => {
                return (
                    <Tab
                        eventKey={item.recrute_status}
                        title={item.recrute_status}
                        key={item.recrute_status}
                    >
                        <InfoBox {...item} />
                    </Tab>
                )
            }))
            return html
        } else {
            return <p>失敗しました</p>;
        }
    }

    return (
        <div className="box-dody">
            <Tabs
                id="company-info-bar"
                defaultActiveKey={key}
                onSelect={(k: any) => setKey(k)}
                className="mb-3"
            >
                {Item()}

            </Tabs>
        </div>
    )
}
//<Item datas={datas}/>
BodyBarBox.propTypes = propTypes
export default BodyBarBox