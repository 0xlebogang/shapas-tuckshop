import fs from "node:fs";
import path from "node:path";

function main() {
	const rootDir = process.cwd();
	const files = fs
		.readdirSync(rootDir, { withFileTypes: true, recursive: true })
		.filter((f) => {
			if (f.name === ".env.example") {
				return true;
			}
			return false;
		})
		.filter((f) => !f.parentPath.includes("node_modules"))
		.map((f) => path.join(f.parentPath, f.name));

	for (const file of files) {
		fs.copyFileSync(file, path.join(path.dirname(file), ".env"));
	}

	console.log("Files copied successfully");
}

main();
