package processor

import (
	"fmt"

	"github.com/stanislav-zeman/gonion/internal/dto"
)

func (p *Processor) processCommon(
	serviceName string,
	config dto.Config,
) error {
	data, err := p.templator.TemplateConfig(config)
	if err != nil {
		return fmt.Errorf("failed templating config: %w", err)
	}

	err = p.writer.WriteConfig(serviceName, data)
	if err != nil {
		return fmt.Errorf("failed writing config: %w", err)
	}

	return nil
}
