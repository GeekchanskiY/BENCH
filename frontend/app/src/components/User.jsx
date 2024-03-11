import { useSelector } from "react-redux"
import { useNavigate } from "react-router-dom"
import { useEffect, useState } from "react"
import { getRequestAuth } from "../features/requests/requests"

export default function User(){
    const navigate = useNavigate()
    const jwt = useSelector((state) => state.jwt.token)
    const [userData, setUserData] = useState({
        'username': '',
        'email': '',
        'ip': ''
    }) 
    useEffect(() => {
        if (jwt == null){
            console.info('Not authentificated! Redirecting')
            navigate('/login')
        }
        getRequestAuth('http://0.0.0.0/users/whoami', jwt)
        .then(data => {
            console.log(data)
            setUserData({
                'email': data.response.email,
                'username': data.response.name,
                'ip': data.response.ip_adress
            })
        })
    }, [])
    

    return <div>
        <span>Username: {userData.username}</span> <br />
        <span>Ip Adress: {userData.ip}</span> <br />
        <span>Email: {userData.email}</span> <br />
    </div>
}