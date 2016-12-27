package main

import (
	"regexp"
	"strings"

	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Check function for handling errors
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// HandleYAMLMetaData function for parse yaml data
func HandleYAMLMetaData(datum []byte) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	err := yaml.Unmarshal(datum, &m)
	return m, err
}

// SplitYaml function for spliting yaml frontmatter and body
func SplitYaml(datum []byte) ([]byte, []byte) {
	re := regexp.MustCompile("(?s)^---\n.*\n---\n")
	yaml := re.FindString(string(datum))
	body := strings.TrimPrefix(string(datum), yaml)
	return []byte(yaml), []byte(body)
}

// ParseFile function for parsing fileName and provide yalResult and html
func ParseFile(filename string) (map[string]interface{}, []byte) {
	input, err := ioutil.ReadFile(filename)
	Check(err)

	yaml, markdown := SplitYaml(input)

	yamlResult, err := HandleYAMLMetaData(yaml)
	Check(err)

	return yamlResult, markdown
}
