package operations

type Runbook struct {
	Name   string `json:"name" example:"Deploy to staging" doc:"The name of the runbook"`
	Inputs []Input
	Jobs   []Job
}

type Input struct {
	ID          string
	DisplayName string
	Required    bool
	Type        string
}

type Job struct {
	JobType         string // "github-action"
	GitHubJobConfig GitHubJobConfig
}

type GitHubJobConfig struct {
	Owner         string
	Repo          string
	Ref           string
	WorkflowID    string
	FieldInputMap map[string]string // map from input ID to field name
}
