import logo from '../img/logo_GPT.jpg'
import ApiCaller from './ApiCaller'
import WhoAmI from './WhoAmI'

export default function Header(){
    return <header>
        <div className='side_header'>
            <img className='headeritem' src={logo} alt="logo" />
            <h1 className='headeritem'>Bench</h1>
        </div>
        <div className='side_header'>
            <ApiCaller/>
        </div>
        <div className='side_header whoami'>
            <WhoAmI/>
        </div>
        
    </header>
}