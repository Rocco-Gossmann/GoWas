importScripts("./wasm_exec.js");

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
