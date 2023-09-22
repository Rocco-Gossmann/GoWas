const mouseIOState = new Uint32Array(3);

if (WebAssembly) {

    const worker = new Worker("./worker.js");

    const canvases = new Map();
    const contexts = new Map();

    worker.addEventListener("message", (ev) => {
        if (ev.data instanceof Array) {
            const canvasid = ev.data[1];
            const canv = canvases.get(canvasid)
            const ctx = contexts.get(canvasid);
            switch (ev.data[0]) {

                case "createCanvas": {
                    const canv = document.createElement("canvas");
                    canv.width = ev.data[2];
                    canv.height = ev.data[3];
                    canv.className="go-wasm-canvas";

                    canvases.set(canvasid, canv);
                    contexts.set(canvasid, canv.getContext("2d"));

                    document.body.appendChild(canv);
                    reactToScreenSize(canv)

                    handleMouseInput(canv);
                } break;


                case "vblank": {
                    const imgDat = new ImageData(ev.data[2], canv.width, canv.height);
                    window.requestAnimationFrame( () => { 
                        ctx.putImageData(imgDat, 0, 0);
                        const io = new Uint32Array(mouseIOState);
                        worker.postMessage(["vblankdone", canvasid, io], [io.buffer]) 
                    }); 
                } break;


                case "destroyCanvas": {
                    canv.parentNode.removeChild(canv);
                    canvases.delete(canvasid);
                    contexts.delete(canvasid);
                } break;

                default:{
                    console.log( "[WORKER MESSAGE]", ev.data[0], " in ", ev);
                }
            }
        }
        else console.log("unknwon worker message", ev.data);
    }, {capture: true})

}
else alert("Your Browser does not support WebAssembly");


/**
 * Reacts to screesize changes and makes sure the canvas always fills
 * the Screen, as much as possible without distortion
 * @param {HTMLCanvasElement} canv 
 */
function reactToScreenSize(canv) {
    const handleScreenSize = () => {
        const winWidth = window.innerWidth;
        const winHeight = window.innerHeight;

        const aspectwin = winWidth / winHeight;
        const aspectcanv = canv.width / canv.height;

        if (aspectwin < aspectcanv) {
            canv.style.width = "100vw";
            canv.style.height = "initial";
        }
        else {
            canv.style.width = "initial";
            canv.style.height = "100vh";
        }
    }
    window.addEventListener("resize", handleScreenSize)
    handleScreenSize();
}

/**
 * Function handles all the stuff nessessary to keep track of the mouseInput on the canvas
 * @param {HTMLCanvasElement} canv 
 */
function handleMouseInput(canv) {

    canv.addEventListener("contextmenu", (ev) => {
        ev.preventDefault()
        return ev;
    }, false)

    canv.addEventListener("click", (ev) => {
        ev.preventDefault()
        return ev;
    }, false)

    canv.addEventListener("mousemove", (ev) => {
        ev.preventDefault();
        mouseIOState[0] = Math.max(0, Math.floor(((ev.clientX - ev.target.offsetLeft) / ev.target.offsetWidth) * canv.width));
        mouseIOState[1] = Math.max(0, Math.floor(((ev.clientY - ev.target.offsetTop) / ev.target.offsetHeight) * canv.height));
    }, true)

    canv.addEventListener("mousedown", (ev) => {
        ev.preventDefault();
        const which = 1 << ev.button;
        mouseIOState[2] |= which;
    }, true)

    canv.addEventListener("mouseup", (ev) => {
        ev.preventDefault();
        const which = 1 << ev.button;
        mouseIOState[2] &= ~which;
    }, true)
    
    canv.addEventListener("mouseout", (ev) => {
        ev.preventDefault();
        mouseIOState[2] = 0;
    }, true)
}
