import { useState, useEffect } from "react"
import io from 'socket.io-client';

const socket = io('ws://coordinator:80/ws');

export default function LogListener(){
    const [messages, setMessages] = useState([]);
    useEffect(() => {
        socket.on('message', (data) => {
          setMessages((prevMessages) => [...prevMessages, data]);
        });
      }, []);

      const handleSubmit = (event) => {
        event.preventDefault();
        const message = event.target.elements.message.value;
        socket.emit('message', message);
        alert(message)
        event.target.elements.message.value = '';
      };
    
      return <div>
        <ul>
        {messages.map((message, index) => (
            <li key={index}>{message}</li>
        ))}
        </ul>
        <form onSubmit={handleSubmit}>
        <input type="text" name="message" />
        <button>Send</button>
        </form>
    </div>
      
}