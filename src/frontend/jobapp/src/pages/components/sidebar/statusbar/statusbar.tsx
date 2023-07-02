import React, { memo, useEffect, useState } from "react";
import { Props } from "./type"
import { CardContent, Typography } from "@mui/material";

export const StatusBar: React.FC<Props> = memo(
    ({
        types,
        type,
        size,
        level,
        adjusting,
        reject,
        offer,
        pass,
        fail,
        date, ...props }) => {
        let color = ""
        let label = ""
        let content = ""
        switch (type) {
            case 0:
                label = "説明会"
            case 1:
                label = "書類選考"
            case 2:
                label = "試験"
            case 3:
                label = "面接" + level?.toString()
            case 4:
                label = "最終面接"
        }
        
        switch (types) {
            case "adjusting":
                color = "black"
                content = "日程調整中"
                break
            case "offer":
                content = "合格"
                color = "green"
                break
            case "reject":
                color = "red"
                content = "不合格"
                
            case "waiting":
                color="yellow"
                content = content = "予定日 "
                + date?.getMonth.toString() + "月"
                + date?.getDate().toString() + "日"
                + date?.getHours().toString() + ":"
                + date?.getMinutes().toString()
        }



        return (
            <CardContent>
                <Typography style={{color:color}}>{label} {content}</Typography>
            </CardContent>
        )
    })



// const StatusBar = (props: Sidebar) => {
//     const fether = async () => { }

//     return (
//         <div className="side-statusbar">
//             {TypeBar(props)}
//             {StatusFunction(props)}
//         </div>
//     )

// }
