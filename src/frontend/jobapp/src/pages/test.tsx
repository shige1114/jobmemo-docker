import { Form } from "./components/mainbox/form/form"
import { Props } from "./components/mainbox/form/type"
export default function Test() {
    const test:Props = {
        id: "id",
        size: "l",
        label: "label",
        text: "text",
        type: "志望動機",
        onClick:()=>{},

    }
    return (<>
    
       <Form {...test}/> 
    
    </>)
}