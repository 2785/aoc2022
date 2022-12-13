package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func mustLoadInput(d int) string {
	fileName := fmt.Sprintf("source/d%d", d)

	content, err := os.ReadFile(fileName)
	c(err)

	return string(content)
}

func splitAndSanitize(content string, splitter string) []string {
	lines := strings.Split(content, splitter)
	lines = lo.FilterMap(lines, func(line string, _ int) (string, bool) {
		o := strings.TrimSpace(line)
		return o, o != ""
	})

	return lines
}

func NamedCaptureGroup(pattern *regexp.Regexp, text string) (map[string]string, error) {
	match := pattern.FindStringSubmatch(text)
	if match == nil {
		return nil, errors.Errorf("%s did not match %s", text, pattern.String())
	}

	captures := make(map[string]string)
	for i, name := range pattern.SubexpNames() {
		if i != 0 && name != "" {
			captures[name] = match[i]
		}
	}

	return captures, nil
}

// nolint // go away
func p[T any](v T) *T {
	return &v
}

func lcm(nums ...int) int {
	res := 1
	for _, n := range nums {
		res = res * n / gcd(res, n)
	}

	return res
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}
