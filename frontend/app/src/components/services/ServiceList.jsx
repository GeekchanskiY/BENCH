import { useState, useEffect } from "react"
import { getRequest, deleteRequestAuth } from "../../features/requests/requests"
import CreateService from "./CreateService"
import { useSelector } from "react-redux"
import { Link } from "react-router-dom"



function Service(props) {
    const jwt = useSelector((state) => state.jwt.token)
    async function deleteService() {
        let res = await deleteRequestAuth('http://0.0.0.0:80/services/service/' + props.service.id, jwt)
        props.setReload(prev => !prev)
    }
    return <div className="service-item">
        <h3>{props.service.id}: {props.service.name}</h3>
        <p>{props.service.description}</p>
        <span>{props.service.url}</span>
        <span>Status: {props.service.is_active ? "Active" : "Inactive"}</span>
        <div className="service-controls">
            <button className="delete-button" onClick={deleteService}>Delete</button>
            <button>Refresh</button>
        </div>
        
    </div>
}

export default function ServiceList() {
    const [services, setServices] = useState([])
    const [reload, setReload] = useState(false)
    const jwt = useSelector((state) => state.jwt.token)


    useEffect(() => {
        getRequest('http://0.0.0.0/services/')
            .then(data => {
                setServices(data.response)
            })
    }, [reload])
    return <div className="service-list">
        {jwt != null ? <CreateService setReload={setReload}></CreateService> : <Link to={'/login'}>Login to create new services</Link>}

        <div className="services">
            {services.map((service) => {
                return <Service key={service.id} setReload={setReload} service={service} />
            })}
        </div>
    </div>
}