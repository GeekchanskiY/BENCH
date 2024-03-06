import {useState, useEffect} from "react";

function ApiItem(props){
    return <div>
        <span>{props.name}: {props.active}</span>
    </div>
}

function ApiCaller(props){
    const [response, setResponse] = useState({})
    const [seconds, setSeconds] = useState(0)

    useEffect(()=>{
    var startDate = new Date();
     fetch('http://0.0.0.0:80/healthcheck')
    .then(res => res.json())
    .then(data => {
        console.log(data)
        setResponse(data)
    })
    .catch(rejected => {
        console.log(rejected);
    });
    
    var endDate = new Date();
    setSeconds((endDate.getTime() - startDate.getTime()) / 1000);
    }, [])
        
        
    return <div>
        {Object.keys(response).forEach(entry => {
            
            return <span>

                asdsadasd
                <ApiItem name={entry} active={response[entry]}></ApiItem>
            </span>
        })}
        <span>Request: {props.url}</span>
        <span>Response: {JSON.stringify(response)}</span>
        <span>Time: {seconds}</span>
        <span>Status: {response === null ? "Success" : "Fail"}</span>
    </div>
}

export default ApiCaller;