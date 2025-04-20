import { env } from "./env";
import { getProjects } from "./gh";

Bun.serve({
	port: env.PORT,
	routes: {
		"/user/:username": async (req) => {
			const { username } = req.params;

			const projects = await getProjects(username);
			if (!projects) {
				return new Response("User not found", { status: 404 });
			}
			return Response.json(projects);
		},
	},

	fetch(req) {
		return new Response("OK");
	},
});
