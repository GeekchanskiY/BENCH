import { useState } from "react"


export default function Commands() {

    const [imported, setImported] = useState("")

    // https://stackoverflow.com/questions/51805395/navigator-clipboard-is-undefined
    async function copyToClipboard(textToCopy) {
        if (navigator.clipboard && window.isSecureContext) {
            await navigator.clipboard.writeText(textToCopy);
        } else {
            const textArea = document.createElement("textarea");
            textArea.value = textToCopy;
                
            textArea.style.position = "absolute";
            textArea.style.left = "-999999px";
                
            document.body.prepend(textArea);
            textArea.select();
    
            try {
                document.execCommand('copy');
            } catch (error) {
                console.error(error);
            } finally {
                textArea.remove();
            }
        }
    }

    async function import_command() {
        await copyToClipboard(imported)
        alert("Commands imported!")
    }

    return <div className="cv_builder_commands">
        <h3>Commands</h3>
        
        <input type="text" value={imported} onChange={(e) => setImported(e.target.value)} />
        <button onClick={import_command}>Import</button>
        <br />
        <button>Export</button>
        
        
    </div>
}