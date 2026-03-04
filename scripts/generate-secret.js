const crypto = require("node:crypto");

function generateSecretKey() {
	console.log(crypto.randomBytes(16).toString("hex"));
}

generateSecretKey();
