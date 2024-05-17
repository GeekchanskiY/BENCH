import logo from '../img/logo_GPT.jpg'

import WhoAmI from './WhoAmI'
import { Link } from 'react-router-dom'
import '../styles/header.css'

export default function Header() {
    return <header>
        <div className='side-header part-header logo-header'>
            <Link to="/">
                <img className='headeritem' src={logo} alt="logo" />
                <h1 className='headeritem'>Bench</h1>
            </Link>

        </div>
        <div className='center-header part-header'>
            <Link to={'/cv'}>CV Builder</Link>
            <Link to={'/projects'}>Projects</Link>
            <Link to={'/ressearch'} >Ressearch</Link>
            <Link to={'/services'} >Services</Link>
            <Link to={'/tests'} >Tests</Link>
        </div>
        <div className='side-header whoami part-header'>
            <WhoAmI />
        </div>

    </header>
}