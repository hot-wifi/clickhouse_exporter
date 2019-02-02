package collector

import (
	"strings"

	"github.com/iancoleman/strcase"
)

func fixMetricName(in string) string {
	s := strcase.ToSnake(in)
	return strings.Replace(s, ".", "_", -1)
}
