import {useState, useEffect} from "react";

function ApiItem(props){
    console.log(props.active)
    if (props.active == "Success"){
        return <div className="headerItem">
            <span>{props.name} <div className="dot active"/></span>
        </div>
    } else {
        return <div className="headerItem">
           <span>{props.name} <div className="dot inactive"/> </span>
        </div>
    }
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
        
    function getServices(){
        if (response.services != undefined){
            return <div className="serviceStatuses">
            <span className="serviceStatusesHeader">Services: </span>
            {response.services.map(entry => 
                <span key={entry.name}>
                    <ApiItem name={entry.name} active={entry.response}></ApiItem>
                </span>
            )}
            </div>
        } else {
            return <span className="serviceStatusesHeader">Services data unavailable. Coordinator is not responding.</span>
        }
        
    }
        
    return <div>
        {getServices()}
    </div>
}

export default ApiCaller;