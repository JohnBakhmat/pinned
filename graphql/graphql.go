package graphql

type Project struct {
	name        string
	description string
	url         string
	stars       int
	forks       int
	language    []string
}

func GetProjects(username string) []Project{

    return []Project{}
}
