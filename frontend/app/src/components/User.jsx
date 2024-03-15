import { useSelector } from "react-redux"
import { useNavigate } from "react-router-dom"
import { useEffect, useState } from "react"
import { getRequestAuth } from "../features/requests/requests"

import '../styles/user.css'
import logo from '../img/logo_GPT.jpg'

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
    

    return <div className="about-me">
        <img src={logo} alt="" />
        <div className="user-data">
            <span>Username: {userData.username}</span>
            <span>Ip Adress: {userData.ip}</span>
            <span>Email: {userData.email}</span>
        </div>
        
    </div>
}