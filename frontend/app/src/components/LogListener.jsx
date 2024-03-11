import { useState, useEffect } from "react"
import io from 'socket.io-client';

const socket = new WebSocket("ws://0.0.0.0:80/ws");

export default function LogListener(){
    const [messages, setMessages] = useState([]);
    useEffect(() => {
        socket.onopen = function() {
            alert("Соединение установлено.");
          };
        socket.onmessage = (event) => {
          setMessages((prevMessages) => [...prevMessages, event.data]);
        };
      }, []);

      const handleSubmit = (event) => {
        event.preventDefault();
        const message = event.target.elements.message.value;
        socket.send('message', message);
        
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