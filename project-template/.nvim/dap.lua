local dap = require("dap")

local filetypes = { "javascript", "html", "go" }

for _, filetype in pairs(filetypes) do
	dap.configurations[filetype] = {
		{
			request = "launch",
			type = "chrome",
			url = "http://localhost:7353",
			webRoot = "${workspaceFolder}",
			runtimeExecutable = "/Applications/Brave Browser.app/Contents/MacOS/Brave Browser",
			userDataDir = true,
		},
	}
end
