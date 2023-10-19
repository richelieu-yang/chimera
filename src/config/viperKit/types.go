package viperKit

type (
	ConfigurationData struct {
		// Data 配置文件内容
		Data []string `json:"data" yaml:"data"`

		// Type 配置文件类型，"yaml", "yml", "json", "toml", "hcl", "tfvars", "dotenv", "env", "properties", "props", "prop", "ini"
		Type string `json:"type" yaml:"type"`
	}
)
