import { useState, useEffect } from "react"
import { getRequest, postRequestAuth } from "../../features/requests/requests"
import CreateService from "./CreateService"




function Service(props){
    async function deleteService(){
      console.log("deleted")
    }
    return <div className="service_item">
        <h3>{props.service.name}</h3>
        <p>{props.service.description}</p>
        <span>{props.service.url}</span> <br />
        <span>Status: {props.service.is_active ? "Active" : "Inactive"}</span> <br />
        <button onClick={deleteService}>Delete</button>
    </div>
}

export default function ServiceList(){
    const [services, setServices] = useState([])
    const [reload, setReload] = useState(false)

    useEffect(() => {
        getRequest('http://0.0.0.0/services/')
        .then(data => {
            setServices(data.response)
        })
    }, [reload])
    return <div>
        <CreateService setReload={setReload}></CreateService>
        <div className="services">
        {services.map((service) => {
            return <Service service={service}/>
        })}
        </div>
    </div>
}