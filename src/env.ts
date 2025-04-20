import { createEnv } from "@t3-oss/env-core";
import z from "zod";

export const env = createEnv({
	server: {
		PORT: z.string().min(1).max(5).default("80"),
		GITHUB_TOKEN: z.string().min(1),
	},
	runtimeEnv: process.env,
});
