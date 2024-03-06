import React from "react";

import { UseSelector, useDispatch, useSelector } from "react-redux";
import { logout, login } from "./jwtSlice";
import { postRequest } from "../requests/requests";

export default function Jwt(){
    const jwt = useSelector((state) => state.jwt.token)
    const dispatch = useDispatch()
    async function do_login(){
        let data = await postRequest(
            'http://0.0.0.0:80/users/auth',
            {
                "email": "SampleUserEmail2",
                "password": "SampleStrongPassword2"
            }
        )
        
        dispatch(
        login({
            username: data.username,
            expires_at: data.expires_at,
            token: data.token
        })
        )
    }

    return <div>
        {jwt == null ? "null" : jwt}
        <button onClick={do_login}>asdasd</button>
    </div>
}