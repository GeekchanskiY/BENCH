import { UseSelector, useSelector } from "react-redux";

async function getRequest(url){
    
    return fetch(url)
    .then(res => res.json())
    .then(data => {
        console.log(data);
        return data;
    })
    .catch(rejected => {
        console.log(rejected);
        return {'error': rejected};
    });
}

async function getRequestAuth(url, jwt){
    return fetch(url, {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer ' + jwt
        }
    })
    .then(res => res.json())
    .then(data => {
        console.log(data);
        return data;
    })
    .catch(rejected => {
        console.log(rejected);
        return {'error': rejected};
    });
}

async function postRequest(url, body){
    return fetch(
        url,
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(body)
            
        }
    )
    .then(res => res.json())
    .then(data => {
        console.log(data);
        
        return data
    })
    .catch(rejected => {
        return {"error": rejected}
    });
    
}

export {postRequest, getRequest, getRequestAuth};