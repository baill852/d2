package parser

import (
	"regexp"
	"strings"
)

type parser struct {
}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) Parse(data map[string]interface{}) {
	for key, value := range data {
		if p.isInvalidKey(key) {
			delete(data, key)
			continue
		}

		switch v := value.(type) {
		case []interface{}:
			for _, vv := range v {
				if nextData, ok := vv.(map[string]interface{}); ok {
					p.Parse(nextData)
				}
			}
		case map[string]interface{}:
			p.Parse(v)
		}
	}
}

func (p *parser) isInvalidKey(key string) bool {
	var snakeCase = regexp.MustCompile("^([a-z]*)_([a-z]*_){0,}[a-z]*$")

	return snakeCase.MatchString(strings.ToLower(key))
}
