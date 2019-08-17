module.exports = {
	apps: [
		{
			// Options reference: http://pm2.keymetrics.io/docs/usage/application-declaration/
			args: "--worker --server",
			exec_interpreter: "none",
			exec_mode: "fork_mode",
			name: "dead-simple-proxy-server",
			script: "./main",
		}
	]
};
