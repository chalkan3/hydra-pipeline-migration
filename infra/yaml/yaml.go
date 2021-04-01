package yamlcustom

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// YamlCustom yaml
type YamlCustom struct {
}

// Unmarshal custom unmarshal
func (y *YamlCustom) Unmarshal(in []byte, out interface{}) {
	yaml.Unmarshal(in, out)
}

// Marshal custom marshal
func (y *YamlCustom) Marshal(in interface{}) []byte {
	ii, err := yaml.Marshal(in)
	if err != nil {
		panic(err)
	}

	return ii
}

// GetYAML get a string of payload
func (y *YamlCustom) GetYAML(payload []byte) string {
	return string(payload)
}

// PrintYAML print yaml
func (y *YamlCustom) PrintYAML(payload []byte) {
	fmt.Printf("--- t dump:\n%s\n\n", string(payload))
}

// NewYamlCustom Ioc
func NewYamlCustom() *YamlCustom {
	return &YamlCustom{}
}
