import React, { memo } from "react";
import { Props } from "./type"
import { Box, Typography, Grid } from "@mui/material";
export const NameBar: React.FC<Props> = memo(
    ({ types, size, name, ...props }) => {
        let fontSize = "h1"

        switch (size) {
            case ("l"):

            case ("s"):
        }


        return (
            <Typography variant="h1">{name}</Typography>
        )
    })
