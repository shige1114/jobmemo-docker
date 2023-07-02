import React, { ReducerAction, memo, useEffect, useState } from "react";
import { Props } from "./type"

import { TextField, Typography, FormControl, Box, Button, createTheme, ThemeProvider } from '@mui/material';

const theme = createTheme({
    palette: {
        background: {
            paper: '#fff',
        },
        text: {
            primary: '#173A5E',
            secondary: '#46505A',
        },

    }

})
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
                break
            case "志望動機":
                len = 300
                break
            case "良い点":
            case "懸念点":
            case "メモ":
                len = 200
                break
            case "質問":
                len = 50
                break
        }
        const [defaultText, setText] = useState(text)
        useEffect(() => {
            collLocalStrage()
        }, [])

        const onChange = (event: React.ChangeEvent<HTMLInputElement>) => {
            if (event.target.value !== undefined) {

                setText(event.target.value)
            } else {
                setText('')
            }
            localStorage.setItem(id, defaultText)
        }
        const deleteLocalStrage = (event: any) => {
            setText(text)
            localStorage.setItem(id, text)
        }
        const collLocalStrage = async () => {
            let localtext = localStorage.getItem(id)
            if (localtext !== null) {
                setText(localtext)
            }
        }
        const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
            event.preventDefault()
            if (defaultText.length <= len) {
                onClick(defaultText)
            } else {
                alert(type.toString() + "は" + len.toString() + "以下で入力してください")
            }
        }


        return (
            <ThemeProvider theme={theme}>
                <Box
                    component="form"

                    onSubmit={handleSubmit}
                    sx={{
                        bgcolor: 'background.paper'
                    }}
                >
                    <Box sx={{ color: 'text.secondary' }}> {defaultText.length}/{len} </Box>
                    <TextField label={type} name="text" value={defaultText} onChange={onChange} multiline={true} />
                    <Button type="submit" variant="outlined"> SUBMIT </Button>
                    <Button variant="outlined" onClick={deleteLocalStrage}>変更を削除</Button>
                </Box>

            </ThemeProvider>

        )
    }
)