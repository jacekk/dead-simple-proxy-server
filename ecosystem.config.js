module.exports = {
	apps: [
		{
			// Options reference: https://pm2.io/doc/en/runtime/reference/ecosystem-file/
			args: "server",
			exec_interpreter: "none",
			exec_mode: "fork_mode",
			name: "dead-simple-proxy-server",
			script: "./main",
		}
	]
};
