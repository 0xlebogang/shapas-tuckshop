// Be sure to run this script from the root of the project (same directory as "pnpm-lock.yaml")

import fs from "node:fs";
import path from "node:path";

function main() {
	const rootDir = process.cwd();

	const files = fs
		// Read all files and directories in the project while recursing through
		// all sub-directories.
		.readdirSync(rootDir, { withFileTypes: true, recursive: true })

		// Filter all files to only the ".env.example" files excluding
		// those in a node_modules directory.
		.filter((f) => {
			if (f.name === ".env.example" && !f.parentPath.includes("node_modules")) {
				return true;
			}
			return false;
		})

		// write all file objects as an string representing the absolute path
		.map((f) => path.join(f.parentPath, f.name));

	// copy each file to the same directory and rename to ".env"
	for (const file of files) {
		fs.copyFileSync(file, path.join(path.dirname(file), ".env"));
	}

	console.log("Files copied successfully");
}

main();
