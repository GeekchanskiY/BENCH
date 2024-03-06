import logo from '../img/logo_GPT.jpg'
export default function Header(){
    return <header>
        <div className='side_header'>
            <img className='headeritem' src={logo} alt="logo" />
            <h1 className='headeritem'>Bench</h1>
        </div>
        <div className='side_header'>
            <div>
                Microservice: t
            </div>
        </div>
        
    </header>
}