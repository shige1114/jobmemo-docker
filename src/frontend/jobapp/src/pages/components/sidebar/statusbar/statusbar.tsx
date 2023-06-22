import React, { useEffect, useState } from "react";
import PropTypes from "prop-types";
import { Nav } from "react-bootstrap"
import { Sidebar, Statusbar } from "@/domain/domain";
import useSWR from "swr"

export interface Statusfunction {
    adjusting: boolean
    reject: boolean
    offer: boolean
    pass: boolean
    fail: boolean
    date: Date
}

const StatusFunction = (props: Statusfunction) => {
    const [adjusting, SetAdjusting] = useState(props.adjusting)
    const [offer, SetOffer] = useState(props.offer)
    const [reject, SetReject] = useState(props.reject)
    const [fail, SetFail] = useState(props.fail)
    const [pass, SetPass] = useState(props.pass)
    const [date, SetDate] = useState(props.date)

    const fether = async () => { }

    const swither = () => {
        if (adjusting == true) {
            return <></>
        } else {
            if (fail == true || pass == true) {
                return <></>
            } else if (offer == true || reject == true) {
                return <></>
            } else {
                return <></>
            }
        }
    }
    return (
        <div className="side-statufunction">
            {swither()}
        </div>
    )

}

interface Typebar {
    level: number
    type: number
}
const TypeBar = (props: Typebar) => {
    switch (props.type) {
        case 0:
            return <></>
        case 1:
            return <></>
        case 2:
            return <></>
        case 4:
            return <></>

    }
}


const StatusBar = (props: Sidebar) => {
    const fether = async () => { }

    return (
        <div className="side-statusbar">
            {TypeBar(props)}
            {StatusFunction(props)}
        </div>
    )

}
