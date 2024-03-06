import {useState, useEffect} from "react";
import { useSelector } from "react-redux";
import { getRequestAuth } from "../features/requests/requests";
export default function WhoAmI(){
    const jwt = useSelector((state) => state.jwt.token)
    const [response, setResponse] = useState('')
    // const 
    useEffect(()=>{
        
        getRequestAuth('http://0.0.0.0/users/whoami', jwt)
        .then(data => setResponse(data.detail))
    }, [])
    return <div>
        {response}
    </div>
}