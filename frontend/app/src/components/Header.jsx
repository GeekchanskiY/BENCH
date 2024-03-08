import logo from '../img/logo_GPT.jpg'
import ApiCaller from './ApiCaller'
import WhoAmI from './WhoAmI'
import { Link } from 'react-router-dom'

export default function Header(){
    return <header>
        <div className='side_header'>
            <Link to="/">
                <img className='headeritem' src={logo} alt="logo" />
                <h1 className='headeritem'>Bench</h1>
            </Link>
            
        </div>
        <div className='side_header'>
            <ApiCaller/>
        </div>
        <div className='side_header whoami'>
            <WhoAmI/>
        </div>
        
    </header>
}