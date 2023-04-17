import React from "react";
import classNames from "classnames";
import PropTypes from "prop-types";
import { Row, Card, Container } from "react-bootstrap";
import Nav from 'react-bootstrap/Nav'
import { BoxProps } from "../Grid/Box/Box";
import BodyBarBox, { BodyBarBoxProp } from "./bodybox/BodyBarBox";
import TitleBox, { TitleBoxProp } from "./titlebox/TitleBox";


interface MainBoxProps extends BoxProps {

}

const titleProps_t: TitleBoxProp = {
    "recrutment_status": "一次面接",
    "company_name": "株式会社カイシャ"
}
const barBoxProps_t: BodyBarBoxProp = {
    'datas': [
        {
            'datas': [
                {
                    "company_id": "kaisya",
                    "user_id": "go",
                    "title": "aiueo",
                    "body": "愛じゃflっjflアjdっかjあ"
                },
                {
                    "company_id": "kaisya",
                    "user_id": "go",
                    "title": "aiueo",
                    "body": "愛じゃflっjflアjdっかjあ"
                },
                {
                    "company_id": "kaisya",
                    "user_id": "go",
                    "title": "aiueo",
                    "body": "愛じゃflっjflアjdっかjあ"
                },
            ],
            "recrute_status":"1次面接"
        }

    ],
    "recrute_status": "djafljf;aljf"

}

const MainBox = () => {
    return (
        <div className="main-box">
            <BodyBarBox {...barBoxProps_t} />
        </div>
    )
}

export default MainBox