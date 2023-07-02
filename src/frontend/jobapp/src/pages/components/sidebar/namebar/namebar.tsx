import React, { memo } from "react";
import { Props } from "./type"
import { CardHeader } from "@mui/material";

export const NameBar: React.FC<Props> = memo(
    ({ types, size, name, ...props }) => {
        let bg = "primary"
        let color = "wthite"

        switch (types) {
            case "adjusting":
                bg = ""
                color = "yellow"
                break
            case "offer":
                bg = ""
                color = "green"
                break
            case "reject":
                bg = ""
                color = "red"
            case "waiting":
                bg= ""
                color=""
        }


        return (
            <CardHeader title={name} style={{ color: color }} />
        )
    })
