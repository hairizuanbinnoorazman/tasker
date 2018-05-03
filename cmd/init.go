package cmd

type ToolConfig struct {
	Name             string
	Token            string
	DefaultProjectID string
}

type Config struct {
	PrimaryTool    ToolConfig
	SecondaryTools []ToolConfig
	Copy           bool
}
