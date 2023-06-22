import React from "react";
import { useRouter } from "next/router";
import { Namebar } from "@/domain/domain";


const NameBar = (props: Namebar) => {
    
    const router = useRouter()
    const options = {
        shalow: false,
        locale: 'jp',
        scroll: false,
    }
    const pageMove = (e: any) => {
        e.preventDefault()
        router.push(
            {
                pathname: '/index',
                query: { company_id: props.id }
            },
            undefined,
            options
        )
    }

    return (
        <div className='side-namebar' onClick={pageMove}>
            <h1>{props.name}</h1>
        </div>
    )
}

export default NameBar