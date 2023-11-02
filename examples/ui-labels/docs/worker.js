importScripts("./wasm_exec.js");

const RESSOURCE_BASEDIR = "./ressources"

let ressourceHandleAutoIncrement = 0
//==============================================================================
// Fetching functions
//==============================================================================

/**
 * @param {string} file The Filename to fetch
 * @returns {Promise<Uint8Array>} - the data that was requested
 */
async function _fetchBinary(index, file) {
    const response = await fetch(RESSOURCE_BASEDIR + "/" + file)

    console.log("got binary response: ", response)

    if(response.status == 200) {
        let dataArray = new Uint8ClampedArray(await (await response.blob()).arrayBuffer());
        sendRessourceToGo(index, dataArray)
        delete dataArray

    } else {
        console.warn(`requested binary '${file}' not found`)
        markRessourceNotFoundInGo(index)

    }

    return response
}

//==============================================================================
// Go => Worker - API
//==============================================================================
/**
 * Called by bo to initialize the ressource load
 * @param {"binary"|"audio"} type
 * @param {string} filename filename within the projects `./ressource` folder
 * @returns {number} - the handle under which the Worker responds back to go (> 0 == success , -1 == invalid type)
 */
function requestRessource(type, filename) {

    const sanitizedFileName = filename.replace(/(^\.\.)*/)

    switch (type) {
        case "binary": {
            const index = ++ressourceHandleAutoIncrement;
            _fetchBinary(index, filename);
            return index;
        }

        case "audio": {
            const index = ++ressourceHandleAutoIncrement;




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
    .then(gowasm => {
        if (!gowasm) {
            console.error("failed to instantiate gowasm", go, gowasm);
            alert("technical error");
            return;
        }

        console.log("lets goooo !!!!!")
        return go.run(gowasm.instance)
    })
    .then(() => self.postMessage("wasm has ended"))

addEventListener("blur", () => {
    console.log("blur");
}) 
