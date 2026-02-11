import "dotenv/config";
import { drizzle } from "drizzle-orm/node-postgres";
import { Pool } from "pg";
import * as schema from "./schema";

const pool = new Pool({
	connectionString: process.env.DATABASE_URL as string,
	ssl: process.env.NODE_ENV === "production",
});

export const db = drizzle({ client: pool, schema });
