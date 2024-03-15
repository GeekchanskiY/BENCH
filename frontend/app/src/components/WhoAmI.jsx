import {useState, useEffect} from "react";
import { useSelector } from "react-redux";
import { getRequestAuth } from "../features/requests/requests";
import { Link } from "react-router-dom";

export default function WhoAmI(){
    const jwt = useSelector((state) => state.jwt.token)
    const [response, setResponse] = useState('')
    
    function get_my_data(){

        getRequestAuth('http://0.0.0.0/users/whoami', jwt)
        .then(data => {
            setResponse(data.response.name)
        })
    }

    useEffect(()=>{
               
    }, [])

    if (jwt != null){
        get_my_data()
        return <div className="about-me">
            <Link to={'/me'}>{response}</Link>
        </div> 
    } else {
        return <div className="login-register">
            <Link to={'/login'}>Login</Link>
            <Link to={'/register'}>Register</Link>
        </div>
    }
}