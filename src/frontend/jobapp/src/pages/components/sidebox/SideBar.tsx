import React, { useEffect, useState } from "react";
import classNames from "classnames";
import PropTypes from "prop-types";
import { Nav } from "react-bootstrap"
import ListOval, { ListOvalProps } from "./namebox/ListOval";

export interface SideBarProps {
    datas?: ListOvalProps[]
}
const propTypes = {
    datas: PropTypes.array
}
const SideBar = (props: SideBarProps) => {
    const [datas, setDatas] = useState<ListOvalProps[]>([])
    const checkDatas = (datas: any) => {
        if (datas != undefined) {
            return true
        }
        return false
    }
    const Item = () => {

        if (props.datas.length > 0) {

            return props.datas.map((data: ListOvalProps, index: number) => (
                <Nav.Item key={index}>
                    <ListOval
                        company_id={data.company_id}
                        company_name={data.company_name}
                        user_id={data.user_id}
                        recruitment_status={data.recruitment_status}
                    />
                </Nav.Item>
            ))
        }


        return (
            <p>失敗しました{props.datas?.length}</p>
        )
    }
    return (
        <div className='side-bar'>

            <Nav justify activeKey={"/index"} className="flex-column"
            >
                <Item />
            </Nav>
        </div>
    )
}

SideBar.propTypes = propTypes

export default SideBar


