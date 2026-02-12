import { relations } from "drizzle-orm";
import * as pg from "drizzle-orm/pg-core";
import { user } from "./auth.schema";
import { common } from "./columns.helper";
import { product } from "./product.sql";

export const order = pg.pgTable("order", {
	...common,
	userId: pg
		.text("user_id")
		.references(() => user.id, { onDelete: "no action" })
		.notNull(),
	isReceived: pg.boolean("is_received").default(false).notNull(),
	totalAmount: pg.doublePrecision("total_amount").notNull(),
});

export const orderItem = pg.pgTable("order_item", {
	...common,
	orderId: pg
		.text("order_id")
		.references(() => order.id, { onDelete: "cascade" }),
	productId: pg
		.text("product_id")
		.references(() => product.id, { onDelete: "cascade" }),
	quantity: pg.integer().notNull().default(1),
	unitPrice: pg.doublePrecision("unit_price").notNull(),
});

export const orderRelations = relations(order, (r) => ({
	orderItems: r.many(orderItem),
	user: r.one(user, {
		fields: [order.userId],
		references: [user.id],
	}),
}));

export const orderItemRelations = relations(orderItem, (r) => ({
	order: r.one(order, {
		fields: [orderItem.orderId],
		references: [order.id],
	}),
	product: r.one(product, {
		fields: [orderItem.productId],
		references: [product.id],
	}),
}));
