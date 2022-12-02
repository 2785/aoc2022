package main

import (
	"fmt"
	"os"
	"strings"

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
