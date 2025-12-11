package processor

import (
	"fmt"
	"log"

	"github.com/stanislav-zeman/cipolla/internal/dto"
)

func (p *Processor) processAPIRESTControllers(
	serviceName string,
	logger dto.Logger,
	framework dto.Framework,
	controllers []dto.Controller,
) error {
	for _, controller := range controllers {
		controller.Logger = logger
		controller.Import = dto.Import{
			Module:  p.config.Module,
			Service: serviceName,
		}
		controller.Framework = framework

		log.Printf("Generating API controller: %v\n", controller)

		data, err := p.templator.TemplateRESTAPIController(controller)
		if err != nil {
			return fmt.Errorf("failed templating service: %w", err)
		}

		err = p.writer.WriteAPIController(serviceName, controller.Name, data)
		if err != nil {
			return fmt.Errorf("failed writing service: %w", err)
		}

		log.Printf("Generating application service interface: %v\n", controller)

		requests := make([]dto.Request, 0, len(controller.Methods))
		for _, method := range controller.Methods {
			requests = append(requests, dto.NewRequest(controller.Name, method, framework, controller.Import))
		}

		err = p.processAPIRESTRequests(serviceName, requests)
		if err != nil {
			return fmt.Errorf("failed processing requests: %w", err)
		}

		responses := make([]dto.Response, 0, len(controller.Methods))
		for _, method := range controller.Methods {
			responses = append(responses, dto.NewResponse(
				controller.Name,
				method,
				framework,
				controller.Import,
			))
		}

		err = p.processAPIRESTResponses(serviceName, responses)
		if err != nil {
			return fmt.Errorf("failed processing responses: %w", err)
		}
	}

	return nil
}

func (p *Processor) processAPIRESTRequests(serviceName string, requests []dto.Request) error {
	for _, request := range requests {
		log.Printf("Generating request: %v\n", request)

		data, err := p.templator.TemplateRESTAPIRequest(request)
		if err != nil {
			return fmt.Errorf("failed templating request: %w", err)
		}

		err = p.writer.WriteAPIRequest(serviceName, request.Method, request.Name, data)
		if err != nil {
			return fmt.Errorf("failed writing request: %w", err)
		}
	}

	return nil
}

func (p *Processor) processAPIRESTResponses(serviceName string, responses []dto.Response) error {
	for _, response := range responses {
		log.Printf("Generating response: %v\n", response)

		data, err := p.templator.TemplateRESTAPIResponse(response)
		if err != nil {
			return fmt.Errorf("failed templating response: %w", err)
		}

		err = p.writer.WriteAPIResponse(serviceName, response.Method, response.Name, data)
		if err != nil {
			return fmt.Errorf("failed writing response: %w", err)
		}
	}

	return nil
}
