import React, { ReducerAction, memo, useEffect, useState } from "react";
import { Props } from "./type"

import { TextField, FormControl, Box, Button } from '@mui/material';
import { FormProps } from "react-bootstrap";

export const Form: React.FC<Props> = memo(
    ({
        size,
        label,
        text,
        type,
        id,
        onClick
    }) => {
        let len = 200
        switch (type) {
            case "回答":
                len = 100
            case "志望動機":
                len = 300
            case "良い点":
                len = 200
            case "懸念点":
                len = 200
            case "メモ":
                len = 200
            case "質問":
                len = 50
        }
        const [defaultText, setText] = useState(text)
        useEffect(()=>{
            collLocalStrage()
        },[])

        const onChange = (event: React.ChangeEvent<HTMLInputElement>) => {
            setText(event.target.value)
            localStorage.setItem(id, event.target.value)
        }
        const deleteLocalStrage = (event: any) => {
            setText(text)
            localStorage.setItem(id, text)
        }
        const collLocalStrage = async () => {
            let localtext = localStorage.getItem(id)
            if (localtext != null) {
                setText(localtext)
            } else {
                setText(text)
            }
        }
        

        return (
            <div className="form">
              
                    <TextField label={type} value={defaultText} onChange={onChange} />
                    <Button type="submit" variant="outlined" onClick={onClick}> SUBMIT </Button>
                    <Button variant="outlined" onClick={deleteLocalStrage}>変更を削除</Button>


            </div>
        )
    }
)