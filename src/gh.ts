import { gql, request } from "graphql-request";
import z from "zod";
import { env } from "./env";

export const getProjects = async (user: string) => {
	const query = gql`
query {
  user(login: "${user}") {
    pinnedItems(types: [REPOSITORY], first: 6) {
      nodes {
        ... on Repository {
          id
          name
          description
          url
          forkCount
          stargazerCount
          languages(first: 5) {
            nodes {
              ... on Language {
                name
              }
            }
          }
        }
      }
    }
  }
}
	`;

	console.dir(env);
	const response = await request(
		"https://api.github.com/graphql",
		query,
		undefined,
		{
			Authorization: `Bearer ${env.GITHUB_TOKEN}`,
		},
	);
	// return response;

	const schema = z.object({
		user: z.object({
			pinnedItems: z.object({
				nodes: z.array(
					z.object({
						id: z.string(),
						name: z.string(),
						description: z.string().nullable(),
						url: z.string(),
						forkCount: z.number(),
						stargazerCount: z.number(),
						languages: z.object({
							nodes: z.array(
								z.object({
									name: z.string(),
								}),
							),
						}),
					}),
				),
			}),
		}),
	});

	const data = await schema.parseAsync(response);

	const projects = data.user.pinnedItems.nodes.map((node) => ({
		id: node.id,
		name: node.name,
		description: node.description,
		url: node.url,
		forks: node.forkCount,
		stars: node.stargazerCount,
		languages: node.languages.nodes.map((node) => node.name),
	}));

	return projects;
};

/*
user(login: "${user}") {
			pinnedItems(types: [REPOSITORY], first: 6) {
				nodes {
					... on Repository {
						id
						name
						description
						url
						forkCount
						stargazerCount
						languages(first: 5) {
							nodes {
								... on Language {
									name
								}
							}
						}
					}
				}
			}
		}

*/
