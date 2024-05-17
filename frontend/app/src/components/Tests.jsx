import { useState, useEffect } from 'react';
import '../styles/tests.css'

export function SupremeButton() {

    const [supreme, setSupreme] = useState(false);
    let childs = [];
    useEffect(() => {
        let container = document.getElementById('supreme-container-1')
        for (let i = 0; i < 16; i++) {
            let child = document.createElement('div')
            child.className = 'supreme-filler'
            childs.push(child)
            container.appendChild(child)
        }
    }, [])

    function supreme_activate() {
        console.log("supreme")
        childs.forEach(element => {
            console.log(element)
            setTimeout(() => {
                element.style.backgroundColor = 'red';
                console.log("supreme")
            }, 100 * Math.floor(Math.random() * 3))
        });
    }




    return <div className='supreme-container' id='supreme-container-1'>
        <div className="supreme-prior">
            <button onClick={supreme_activate}>Supreme</button>
        </div>
    </div>
}

export default function Tests() {
    const [supreme, setSupreme] = useState(false);
    function supreme_activate() {
        console.log("supreme")
    }


  return (
    <div className="tests">
      <h1>Tests</h1>
      <h3>Supreme button</h3>
      <SupremeButton />
      
    </div>
  );
}