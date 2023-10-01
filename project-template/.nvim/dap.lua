local dap = require("dap")

dap.configurations.go = {
	{
		request = "launch",
		type = "chrome",
		url = "http://localhost:7353",
		webRoot = "${workspaceFolder}",
		runtimeExecutable = "/Applications/Brave Browser.app/Contents/MacOS/Brave Browser",
		userDataDir = true,
	},
}
