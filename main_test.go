package main

import (
	"reflect"
	"testing"
)

func TestLoadConfig_No_Config(t *testing.T) {
	_, err := loadConfig("no_config")
	if err != nil {
		t.Errorf("Error on loading config %v", err)
	}
}

func TestLoadConfig_Example_Config(t *testing.T) {
	_, err := loadConfig("config/config_example.yml")
	if err != nil {
		t.Errorf("Error on loading config %v", err)
	}
}

func TestLoadConfigContent_ParsingError(t *testing.T) {
	configFile := `
DUMMY
:FOO
`
	_, err := loadConfigContent([]byte(configFile))
	if err == nil {
		t.Errorf("Should have an error parsing unparseable content")
	}
}

func TestLoadConfigContent_Ok_Standard(t *testing.T) {
	configFile := `
resource_tags:
  - tag_selections:
      Client: "Alice"
      Env: "Prod"
    resource_type_selections:
      - "Microsoft.Web/serverfarms"
      - "Microsoft.Web/sites"
`
	want := Config{
		[]ResourceTag{
			ResourceTag{
				TagSelections: map[string]string{
					"Client": "Alice",
					"Env":    "Prod",
				},
				ResourceTypeSelections: []string{
					"Microsoft.Web/serverfarms",
					"Microsoft.Web/sites",
				},
			},
		},
	}
	got, err := loadConfigContent([]byte(configFile))
	if err != nil {
		t.Errorf("Error on loading config content %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Error in getting config Got:%v, Expected config:%v", got, want)
	}
}
