import { relations } from "drizzle-orm";
import * as pg from "drizzle-orm/pg-core";
import { common } from "./columns.helper";
import { image } from "./image.sql";
import { orderItem } from "./order.sql";

export const product = pg.pgTable("product", {
	...common,
	name: pg.varchar({ length: 100 }).notNull(),
	description: pg.text().notNull(),
	price: pg.doublePrecision().notNull(),
});

export const productRelations = relations(product, (r) => ({
	images: r.many(image),
	orderItems: r.many(orderItem),
}));
