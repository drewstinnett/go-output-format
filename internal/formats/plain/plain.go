package plain

import (
	"fmt"

	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
)

// PlainFormatter Just output in raw go format
type plug struct{}

// Format Do the actual formatting here
func (p *plug) Format(data interface{}, config *config.Config) ([]byte, error) {
	// b := []byte(fmt.Sprintf("%+v", data.(interface{})))
	b := []byte(fmt.Sprintf("%+v", data))
	return b, nil
}

// Output Capture the output
func (p *plug) Output(data interface{}, config *config.Config) ([]byte, error) {
	b, _ := p.Format(data, config)
	return b, nil
}

func init() {
	formatter.Add("plain", func() formatter.Formatter {
		return &plug{}
	})
}
