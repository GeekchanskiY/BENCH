import { useState, useEffect } from "react"
import { getRequest } from "../features/requests/requests"

function Service(props){
    
    return <div>
        <span>{props.service.name}</span> <br />
        <span>{props.service.description}</span> <br />
        <span>{props.service.is_active ? "Active" : "Inactive"}</span> <br />
    </div>
}

export default function ServiceList(){
    const [services, setServices] = useState([])

    useEffect(() => {
        getRequest('http://0.0.0.0/services/')
        .then(data => {
            console.log(data)
            setServices(data)
        })
    }, [])
    return <div>
        {services.map((service) => {
        
            return <Service service={service}/>
        })}
    </div>
}