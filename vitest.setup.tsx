import "@testing-library/jest-dom/vitest";
import { type RenderOptions, render } from "@testing-library/react";
import type * as React from "react";
import Providers from "@/components/providers";

const AllProviders = ({ children }: { children: React.ReactNode }) => {
	return <Providers>{children}</Providers>;
};

const customRender = (
	ui: React.ReactElement,
	options?: Omit<RenderOptions, "wrapper">,
) => render(ui, { wrapper: AllProviders, ...options });

export { customRender as render };
