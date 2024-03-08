import {useState, useEffect} from "react";
import { useSelector } from "react-redux";
import { getRequestAuth } from "../features/requests/requests";
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
        get_my_data()
        
    }, [])
    return <div className="whoAmI">
        <span>
        {response}
        </span>
        
        <button onClick={get_my_data}>refresh</button>
    </div>
}