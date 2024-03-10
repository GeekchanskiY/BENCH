import { UseSelector, useSelector } from "react-redux";

async function getRequest(url){
    
    return fetch(url)
    .then(res => res.json())
    .then(data => {
        return {
            'response': data,
            'success': true
        }
    })
    .catch(rejected => {
        return {
            'response': rejected,
            'success': false
        }
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
        return {
            'response': data,
            'success': true
        }
    })
    .catch(rejected => {
        return {
            'response': rejected,
            'success': false
        }
    });
}

async function postRequest(url, body){
    let status_code = 404
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
    .then(res =>{ 
        status_code = res.status
        res.json()
    })
    .then(data => {
        if (status_code == 200){
            return {
                'response': data,
                'success': true
            }
        } else {
            return {
                'response': data,
                'success': false
            }
        }
    })
    .catch(rejected => {
        return {
            'response': rejected,
            'success': false
        }
    });
    
}

export {postRequest, getRequest, getRequestAuth};