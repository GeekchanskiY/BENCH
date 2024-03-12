import { useState } from "react"
import { useRef, useEffect } from "react"

export const useWebsocket = (url) => {
    const [isReady, setIsReady] = useState(false)
    const [val, setVal] = useState(null)
    const ws = useRef(null)
  
    useEffect(() => {
      const socket = new WebSocket(url)
  
      socket.onopen = () => setIsReady(true)
      socket.onclose = () => setIsReady(false)
      socket.onmessage = (event) => setVal(event.data)
      
  
      ws.current = socket
  
      return () => {
        socket.close()
      }
    }, [url])
  
    return [isReady, val, ws.current?.send.bind(ws.current)]
  }