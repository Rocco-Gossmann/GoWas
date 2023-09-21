require("lspconfig").gopls.setup {
    settings = {
        gopls = {
            buildFlags = { "-tags=js,wasm" }
        }
    }
}
