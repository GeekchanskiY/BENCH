import React from "react";

import { UseSelector, useDispatch, useSelector } from "react-redux";
import { logout, login } from "./jwtSlice";


export default function Jwt(){
    const jwt = useSelector((state) => state.jwt.token)
    const dispatch = useDispatch()
    function flogin(){
        fetch(
            'http://0.0.0.0:80/users/auth',
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                "email": "SampleUserEmail2",
                "password": "SampleStrongPassword2"
                })
                
            }
        )
        .then(res => res.json())
        .then(data => {
            console.log(data)
            dispatch(
                login({
                    username: data.username,
                    expires_at: data.expires_at,
                    token: data.token
                })
            )
            
        })
    }

    return <div>
        {jwt == null ? "null" : jwt}
        <button onClick={flogin}>asdasd</button>
    </div>
}