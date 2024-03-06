import React from "react";

import { UseSelector, useDispatch, useSelector } from "react-redux";
import { login, logout } from "./jwtSlice";


export default function Jwt(){
    const jwt = useSelector((state) => state.jwt.token)

    return <div>
        {jwt == null ? "null" : jwt}
    </div>
}