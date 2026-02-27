const fs = require("fs");
const path = require("path");

/**
 * Recursively finds all .env.example files and creates .env copies
 */
function setupEnvFiles(rootDir = process.cwd()) {
	let count = 0;

	function walk(dir) {
		try {
			const files = fs.readdirSync(dir);

			for (const file of files) {
				const filePath = path.join(dir, file);
				const stat = fs.statSync(filePath);

				// Skip node_modules and other common directories
				if (stat.isDirectory()) {
					if (
						!file.startsWith(".") &&
						file !== "node_modules" &&
						file !== "dist" &&
						file !== "build"
					) {
						walk(filePath);
					}
				} else if (file === ".env.example") {
					const envPath = path.join(dir, ".env");

					// Only create .env if it doesn't already exist
					if (!fs.existsSync(envPath)) {
						const content = fs.readFileSync(filePath, "utf-8");
						fs.writeFileSync(envPath, content, "utf-8");
						console.log(`✓ Created ${envPath}`);
						count++;
					} else {
						console.log(`ℹ Skipped ${envPath} (already exists)`);
					}
				}
			}
		} catch (err) {
			// Skip directories we can't read
			if (err.code !== "EACCES") {
				console.error(`Error reading directory ${dir}:`, err.message);
			}
		}
	}

	walk(rootDir);
	console.log(`\nSetup complete. Created ${count} .env file(s).`);
}

setupEnvFiles();
