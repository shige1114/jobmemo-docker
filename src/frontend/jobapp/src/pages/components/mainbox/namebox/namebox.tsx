import React, { memo } from "react";
import { Props } from "./type"
import { Box, Typography, Grid, ThemeProvider } from "@mui/material";
import { lapTheme,smartTheme } from "./theme";
export const NameBar: React.FC<Props> = memo(
    ({ types, size, name, ...props }) => {
        let fontSize = "h1"
        let theme = lapTheme
        switch (size) {
            case ("l"):
                break
                
            case ("s"):
                theme = smartTheme
                break
        }


        return (
            <ThemeProvider theme={theme}>
                <Box>{name}</Box>
            </ThemeProvider>
        )
    })
