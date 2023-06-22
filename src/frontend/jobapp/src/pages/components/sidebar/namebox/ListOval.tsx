import React from "react";
import classNames from "classnames";
import PropTypes from "prop-types";
import { Row, Card, Container } from "react-bootstrap";
import { useRouter } from "next/router";

export interface ListOvalProps extends React.HTMLAttributes<HTMLElement> {
    company_name?: string
    company_id?: string
    user_id?: string
    recruitment_status?: string
}

const propTypes = {
    company_name: PropTypes.string,
    company_id: PropTypes.string,
    user_id: PropTypes.string,
    recruitment_status: PropTypes.string,
}
const ListOval = (props: ListOvalProps) => {
    console.log(">>>>>>>")
    console.log(props)
    
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
                query: { company_id: props.company_id, user_id: props.user_id }
            },
            undefined,
            options
        )
    }

    return (
        
        <div className='company-box' onClick={pageMove}>

            <p>{props.company_name}</p>

            <p>{props.recruitment_status}</p>

        </div>
        
    )
}

ListOval.propTypes = propTypes
export default ListOval
/*
const ListOval = React.forwardRef<HTMLDivElement, ListOvalProps>(
    (
        {
            className,
            fluid,
            ...props
        }
        , ref
    ) => {
        
        return (
            <div
                ref={ref}
                {...props}
                className={classNames(className, fluid)}
            >
            </div>
        );

    })
ListOval.displayName = 'ListOval';
ListOval.propTypes = propTypes;
*/