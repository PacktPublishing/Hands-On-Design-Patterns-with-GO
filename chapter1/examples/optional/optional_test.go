package chapter1

import (
	"reflect"
	"testing"
)

func TestParseOptionsNone(t *testing.T) {
	defaultConfig, err := parseOptions()
	if err != nil {
		t.Errorf("Error occurred with no options: %s", err)
	}

	if reflect.DeepEqual(defaultConfig, &config{level: "low", environment: "home", command: "none"}) {
		t.Log("defaultConfig returned config as expected.")
	} else {
		t.Errorf("defaultConfig does not match! %#v", defaultConfig)
	}
}
func TestParseOptionsSingle(t *testing.T) {
	levelOption := optionLevel("> 9000")
	levelConfig, err := parseOptions(levelOption)
	if err != nil {
		t.Errorf("Error occurred when single option. %s", err)
	}

	if reflect.DeepEqual(levelConfig, &config{level: "> 9000", environment: "home", command: "none"}) {
		t.Log("levelConfig returned config as expected.")
	} else {
		t.Errorf("levelConfig does not match! %#v", levelConfig)
	}
}

func TestParseOptionsAll(t *testing.T) {

	levelOption := optionLevel("> 9000")
	envOption := optionEnvironment("space")
	commandOption := optionCommand("go run ./...")

	optionsConfig, err := parseOptions(levelOption, envOption, commandOption)
	if err != nil {
		t.Errorf("Error occurred when single option. %s", err)
	}

	if reflect.DeepEqual(optionsConfig, &config{level: "> 9000", environment: "space", command: "go run ./..."}) {
		t.Log("optionsConfig returned config as expected.")
	} else {
		t.Errorf("optionsConfig does not match! %#v", optionsConfig)
	}
}
