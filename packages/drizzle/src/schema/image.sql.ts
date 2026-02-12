import { relations } from "drizzle-orm";
import * as pg from "drizzle-orm/pg-core";
import { common } from "./columns.helper";
import { product } from "./product.sql";

export const image = pg.pgTable("image", {
	...common,
	productId: pg
		.text("product_id")
		.references(() => product.id, { onDelete: "cascade" }),
	url: pg.text().notNull(),
});

export const imageRelations = relations(image, (r) => ({
	product: r.one(product, {
		fields: [image.productId],
		references: [product.id],
	}),
}));
