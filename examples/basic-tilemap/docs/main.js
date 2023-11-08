const ioState = new Uint32Array(3 + 8);

const keyMap = new Map([
	["Backspace", 8],
	["Tab", 9],
	["Enter", 13],
	["ShiftLeft", 16],
	["ControlLeft", 17],
	["AltLeft", 18],
	["Pause", 19],
	["CapsLock", 20],
	["Escape", 27],
	["Space", 32],
	["PageUp", 33],
	["PageDown", 34],
	["End", 35],
	["Home", 36],
	["ArrowLeft", 37],
	["ArrowUp", 38],
	["ArrowRight", 39],
	["ArrowDown", 40],
	["PrintScreen", 44],
	["Insert", 45],
	["Delete", 46],
	["Digit0", 48],
	["Digit1", 49],
	["Digit2", 50],
	["Digit3", 51],
	["Digit4", 52],
	["Digit5", 53],
	["Digit6", 54],
	["Digit7", 55],
	["Digit8", 56],
	["Digit9", 57],
	["KeyA", 65],
	["KeyB", 66],
	["KeyC", 67],
	["KeyD", 68],
	["KeyE", 69],
	["KeyF", 70],
	["KeyG", 71],
	["KeyH", 72],
	["KeyI", 73],
	["KeyJ", 74],
	["KeyK", 75],
	["KeyL", 76],
	["KeyM", 77],
	["KeyN", 78],
	["KeyO", 79],
	["KeyP", 80],
	["KeyQ", 81],
	["KeyR", 82],
	["KeyS", 83],
	["KeyT", 84],
	["KeyU", 85],
	["KeyV", 86],
	["KeyW", 87],
	["KeyX", 88],
	["KeyY", 89],
	["KeyZ", 90],
	["MetaLeft", 91],
	["OSLeft", 91],
	["ContextMenu", 93],
	["Numpad0", 96],
	["Numpad1", 97],
	["Numpad2", 98],
	["Numpad3", 99],
	["Numpad4", 100],
	["Numpad5", 101],
	["Numpad6", 102],
	["Numpad7", 103],
	["Numpad8", 104],
	["Numpad9", 105],
	["NumpadMultiply", 106],
	["NumpadAdd", 107],
	["NumpadSubtract", 109],
	["NumpadDecimal", 110],
	["NumpadDivide", 111],
	["NumpadEnter", 13],
	["F1", 112],
	["F2", 113],
	["F3", 114],
	["F4", 115],
	["F5", 116],
	["F6", 117],
	["F7", 118],
	["F8", 119],
	["F9", 120],
	["F10", 121],
	["F11", 122],
	["F12", 123],
	["F13", 124],
	["F14", 125],
	["F15", 126],
	["F16", 127],
	["F17", 128],
	["F18", 129],
	["F19", 130],
	["F20", 131],
	["F21", 132],
	["F22", 133],
	["F23", 134],
	["F24", 135],
	["NumLock", 144],
	["ScrollLock", 145],
	["Semicolon", 186],
	["Equal", 187],
	["Comma", 188],
	["Minus", 189],
	["Dash", 189],
	["Period", 190],
	["Slash", 191],
	["Backquote", 192],
	["BracketLeft", 219],
	["Backslash", 220],
	["BracketRight", 221],
	["Quote", 222],
]);

if (WebAssembly) {
  const worker = new Worker("./worker.js");

  const canvases = new Map();
  const contexts = new Map();

  worker.addEventListener("message", (ev) => {
    if (ev.data instanceof Array) {
      const canvasid = ev.data[1];
      const canv = canvases.get(canvasid);
      const ctx = contexts.get(canvasid);
      switch (ev.data[0]) {
        case "createCanvas":
          {
            const canv = document.createElement("canvas");
            canv.width = ev.data[2];
            canv.height = ev.data[3];
            canv.className = "go-wasm-canvas";

            canvases.set(canvasid, canv);
            contexts.set(canvasid, canv.getContext("2d"));

            document.body.appendChild(canv);
            reactToScreenSize(canv);

            handleMouseInput(canv);
            handleKeyboardInput(canv);
          }
          break;

        case "vblank":
          {
            const imgDat = new ImageData(ev.data[2], canv.width, canv.height);
            window.requestAnimationFrame(() => {
              ctx.putImageData(imgDat, 0, 0);
              const io = new Uint32Array(ioState);
              worker.postMessage(["vblankdone", canvasid, io], [io.buffer]);
            });
          }
          break;

        case "destroyCanvas":
          {
            canv.parentNode.removeChild(canv);
            canvases.delete(canvasid);
            contexts.delete(canvasid);
          }
          break;

        default: {
          console.log("[WORKER MESSAGE]", ev.data[0], " in ", ev);
        }
      }
    } else console.log("unknwon worker message", ev.data);
  }, { capture: true });
} else alert("Your Browser does not support WebAssembly");

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
    } else {
      canv.style.width = "initial";
      canv.style.height = "100vh";
    }
  };
  window.addEventListener("resize", handleScreenSize);
  handleScreenSize();
}

/**
 * Function handles all the stuff nessessary to keep track of the mouseInput on the canvas
 * @param {HTMLCanvasElement} canv
 */
function handleMouseInput(canv) {
  canv.addEventListener("contextmenu", (ev) => {
    ev.preventDefault();
    return ev;
  }, false);

  canv.addEventListener("click", (ev) => {
    ev.preventDefault();
    return ev;
  }, false);

  canv.addEventListener("mousemove", (ev) => {
    ev.preventDefault();
    ioState[0] = Math.max(
      0,
      Math.floor(
        ((ev.clientX - ev.target.offsetLeft) / ev.target.offsetWidth) *
          canv.width,
      ),
    );
    ioState[1] = Math.max(
      0,
      Math.floor(
        ((ev.clientY - ev.target.offsetTop) / ev.target.offsetHeight) *
          canv.height,
      ),
    );
  }, true);

  canv.addEventListener("mousedown", (ev) => {
    ev.preventDefault();
    const which = 1 << ev.button;
    ioState[2] |= which;
  }, true);

  canv.addEventListener("mouseup", (ev) => {
    ev.preventDefault();
    const which = 1 << ev.button;
    ioState[2] &= ~which;
  }, true);

  canv.addEventListener("mouseout", (ev) => {
    ev.preventDefault();
    ioState[2] = 0;
  }, true);
}

function handleKeyboardInput() {
  window.addEventListener("keyup", (ev) => {
    ev.preventDefault();
    const which = keyMap.get(ev.code) || 0;
    const offset = which % 32;
    const stack = (which - offset) / 32;

    console.log(ev.code, which);

    ioState[3 + stack] &= ~(1 << offset);

    return ev;
  });
  window.addEventListener("keydown", (ev) => {
    ev.preventDefault();

    const which = keyMap.get(ev.code) || 0;
    const offset = which % 32;
    const stack = (which - offset) / 32;

    ioState[3 + stack] |= 1 << offset;

    return ev;
  });
}
