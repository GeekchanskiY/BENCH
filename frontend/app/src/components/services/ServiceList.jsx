import { useState, useEffect } from "react"
import { getRequest, deleteRequestAuth } from "../../features/requests/requests"
import CreateService from "./CreateService"
import { useSelector } from "react-redux"



function Service(props){
    const jwt = useSelector((state) => state.jwt.token)
    async function deleteService(){
      let res = await deleteRequestAuth('http://0.0.0.0:80/services/service/'+props.service.id, jwt)
      props.setReload(prev => !prev)
    }
    return <div className="service_item">
        <h3>{props.service.name} {props.service.id}</h3>
        <p>{props.service.description}</p>
        <span>{props.service.url}</span> <br />
        <span>Status: {props.service.is_active ? "Active" : "Inactive"}</span> <br />
        <button onClick={deleteService}>Delete </button>
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
            return <Service key={service.id} setReload={setReload} service={service}/>
        })}
        </div>
    </div>
}