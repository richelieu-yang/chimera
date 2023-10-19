package viperKit

type (
	Data struct {
		// Content 配置文件的内容
		Content []byte `json:"data" yaml:"data"`

		// Type 配置文件的类型，"yaml", "yml", "json", "toml", "hcl", "tfvars", "dotenv", "env", "properties", "props", "prop", "ini"
		Type string `json:"type" yaml:"type"`
	}
)
