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
    }

    async function export_command() {
        let res = await fetch("http://0.0.0.0:3001/v1/utils/backup")
        console.log(res)
        let data = await res.json()
        console.log(data)
        await copyToClipboard(JSON.stringify(data))
        document.getElementById("export_res_status").innerHTML = "Copied!"
        document.getElementById("export_res_status").style.color = "green"
        setTimeout(() => {
            document.getElementById("export_res_status").innerHTML = "Result will be copied to clipboard"
            document.getElementById("export_res_status").style.color = "#fff"
        }, 2000)
    }

    return <div className="cv_builder_commands">
        <h3>Commands</h3>
        
        <input type="text" value={imported} onChange={(e) => setImported(e.target.value)} />
        <button onClick={import_command}>Import</button>
        <br />
        <button onClick={export_command}>Export</button>
        <span id="export_res_status">Result will be copied to clipboard</span>
        
        
    </div>
}