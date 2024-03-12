import { useState, useEffect } from "react"
import { useWebsocket } from "../hooks/useWebSocket";

const sockaddr = "ws://0.0.0.0:80/ws";

export default function LogListener(){
    const [messages, setMessages] = useState([]);
    const [currentMessage, setCurrentMessage] = useState('Message')
    const [ready, val, send] = useWebsocket(sockaddr)
    useEffect(() => {
      
    if (ready) {
      send("On my way!")
    }
  }, [ready])

    const handleSubmit = (event) => {
      event.preventDefault();
        if (ready && currentMessage) {
            send(currentMessage);
            setCurrentMessage('')
            
        }
      
      
      
    };
  
    return <div>
      <ul>
      {messages.map((message, index) => (
          <li key={index}>{message}</li>
      ))}
      </ul>
      <form onSubmit={handleSubmit}>
      <span>Ready: {JSON.stringify(ready)}, Value: {val}</span>
      <input type="text" value={currentMessage} onChange={(e) => setCurrentMessage(e.target.value)} style={{color: "#000"}}/>
      <button onClick={handleSubmit}>Send</button>
      </form>
      <button>Refresh socket</button>
  </div>
      
}