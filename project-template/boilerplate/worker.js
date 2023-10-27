importScripts("./wasm_exec.js");

/**
 * Called by bo to initialize the ressource load
 * @param {"binary"|"audio"} type
 * @param {string} filename filename within the projects `./ressource` folder
 * @returns {number} - the handle under which the Worker responds back to go (> 0 == success , -1 == invalid type)
 */

let ressourceHandleAutoIncrement = 0

function requestRessource(type, filename) {
    switch(type) {
        case "binary": { const index = ++ressourceHandleAutoIncrement;




            return index;
        }

        case "audio": { const index = ++ressourceHandleAutoIncrement;




            return index;
        }

        default: return -1;
    }
}




//==============================================================================
// Let's Gooooo !!!
//==============================================================================
const go = new Go();
WebAssembly.instantiateStreaming(fetch("./main.wasm"), go.importObject)
    .then( gowasm => {
        if(!gowasm) {
            console.error("failed to instantiate gowasm", go, gowasm);
            alert("technical error");
            return;
        }

        console.log("lets goooo !!!!!")
        return go.run(gowasm.instance)
    })      
    .then( () => self.postMessage("wasm has ended"))

addEventListener("blur", () => {
    console.log("blur");
}) 
