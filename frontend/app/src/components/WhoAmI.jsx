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
            console.log(data)
            setResponse(data.name)
        })
    }

    useEffect(()=>{
               
    }, [])

    if (jwt != null){
        get_my_data()
        return <div className="whoAmI">
            <span>
                <Link to={'/me'}>{response}</Link>
            </span>
        </div> 
    } else {
        return <div className="whoAmI">
            <Link to={'/login'}>Login</Link>
        </div>
    }
}