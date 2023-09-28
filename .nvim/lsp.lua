return function(lspconfig, on_attach, capabilities)

	lspconfig.gopls.setup{
		on_attach = on_attach,
		capabilities = capabilities,
		settings = {
			gopls = {
				buildFlags = { "-tags=js,wasm" }
			}
		}
	}

end

--require("lspconfig").gopls.setup {
--    settings = {
--        gopls = {
--            buildFlags = { "-tags=js,wasm" }
--        }
--    }
--}
