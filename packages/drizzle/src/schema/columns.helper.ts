import * as pg from "drizzle-orm/pg-core";
import { customAlphabet } from "nanoid";

const upperAlphabet = [...Array(26)].map((_, i) => String.fromCharCode(i + 65));
const lowerAlphabet = [...Array(26)].map((_, i) => String.fromCharCode(i + 97));
const digits = [...Array(10)].map((_, i) => i.toString());

export const common = {
	id: pg
		.text()
		.primaryKey()
		.$default(() => generateId()),
	createdAt: pg.timestamp("created_at").notNull().defaultNow(),
	updatedAt: pg
		.timestamp("updated_at")
		.$onUpdate(() => /* @__PURE__ */ new Date()),
};

export function generateId() {
	const characters = [...upperAlphabet, ...lowerAlphabet, ...digits];
	const generator = customAlphabet(characters.join(""), 21);
	return generator();
}
