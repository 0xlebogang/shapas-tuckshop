import { serve } from "@hono/node-server";
import app from "./app";

serve(
	{
		fetch: app.fetch,
		port: Number.parseInt(process.env.PORT || "5000", 10),
	},
	(info) => {
		console.log(`Server listening at ${info.address}${info.port}`);
	},
);
