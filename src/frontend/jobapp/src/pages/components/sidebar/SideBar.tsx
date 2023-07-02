import React, { memo, useEffect, useState } from "react";
import { Props } from "./type"
import { Card ,createTheme,ThemeProvider} from "@mui/material";
import { NameBar } from "./namebar/namebar";
import { StatusBar } from "./statusbar/statusbar";
import { ColorType } from "./namebar/type";




export const SideBar: React.FC<Props> = memo(
    ({
        type,
        size,
        level,
        adjusting,
        reject,
        offer,
        pass,
        fail,
        date,
        name
    }) => {
        let types: ColorType = "offer"

        if (adjusting) {
            types = "adjusting"
        } else {
            if (!offer && !pass && !fail && !reject) {
                types = "waiting"
            }
            if (offer || pass) {
                types = "offer"
            }
            if (reject || fail) {
                types = "reject"
            }

        }
        return (
            <Card>
                <NameBar
                    size={size}
                    name={name}
                    types={types}
                />
                <StatusBar
                    type={type}
                    types={types}
                    size={size}
                    level={level}
                    adjusting={adjusting}
                    reject={reject}
                    offer={offer}
                    pass={pass}
                    fail={fail}
                    date={date}
                />
            </Card>
        )
    })


// (props: SideBarProps) => {

//     // useEffect(()=>{
//     //     checkDatas(props.datas)
//     // },[])
//     // function checkDatas(datas:any){
//     //     if (datas != undefined && datas.length > 0){
//     //         setDatas(datas)
//     //     }
//     // }
//     const Item = () => {

//         if (datas.length > 0) {

//             return (<>
//                 {datas.map((data: ListOvalProps, index: number) => (
//                         <Nav.Item key={index}>
//                             <ListOval
//                                 company_id={data.company_id}
//                                 company_name={data.company_name}
//                                 user_id={data.user_id}
//                                 recruitment_status={data.recruitment_status}
//                             />
//                         </Nav.Item>
//                     ))}
//             </>)
//         }
//         return (
//             <p>失敗しました{datas.length}</p>
//         )
//     }
//     return (
//         <div className='side-bar'>

//             <Nav justify activeKey={"/index"} className="flex-column"
//             >
//                 <Item />
//             </Nav>
//         </div>
//     )
// }


