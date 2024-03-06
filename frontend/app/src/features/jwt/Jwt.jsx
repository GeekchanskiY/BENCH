import React from "react";

import { UseSelector, useDispatch, useSelector } from "react-redux";
import { login, logout } from "./jwtSlice";


export default function Jwt(){
    const jwt = useSelector((state) => state.jwt.token)

    function login(){
        fetch(
            'http://0.0.0.0:80/healthcheck',
            {
                method: 'POST',
                body: JSON.stringify({
                    'email': 'SampleUserEmail',
                    'passoword': 'SampleStrongPassword'
                })
                
            }
        )
        .then(res => res.json())
        .then(data => {
            
        })
    }

    return <div>
        {jwt == null ? "null" : jwt}
    </div>
}