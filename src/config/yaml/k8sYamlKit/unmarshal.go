package k8sYamlKit

import "sigs.k8s.io/yaml"

var Unmarshal func(y []byte, o interface{}, opts ...yaml.JSONOpt) error = yaml.Unmarshal

func UnmarshalFromString(in string, out interface{}) error {
	return Unmarshal([]byte(in), out)
}
