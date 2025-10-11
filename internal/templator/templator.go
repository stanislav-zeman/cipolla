package templator

import (
	"bytes"
	"errors"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/stanislav-zeman/gonion/internal/dto"
	"github.com/stanislav-zeman/gonion/internal/layers"
)

var ErrFailedParsingTemplate = errors.New("failed parsing template")

// Templator holds all of the object templates and exposes
// an API to template them.
type Templator struct {
	// Domain templates.
	entity           *template.Template
	value            *template.Template
	domainInterface  *template.Template
	domainService    *template.Template
	domainRepository *template.Template

	// Application templates.
	command      *template.Template
	query        *template.Template
	appInterface *template.Template
	appService   *template.Template

	// Application templates.
	infraRepository *template.Template

	// API templates.
	apiController *template.Template
	apiRequest    *template.Template
	apiResponse   *template.Template
}

func New(assetsDirector string) (t Templator, err error) { //nolint: cyclop
	fp := filepath.Join(assetsDirector, layers.DomainLayer, "entity.tmpl")
	entity, err := template.ParseFiles(fp)
	if err != nil {
		return Templator{}, fmt.Errorf("%w: %w", ErrFailedParsingTemplate, err)
	}

	fp = filepath.Join(assetsDirector, layers.DomainLayer, "value.tmpl")
	value, err := template.ParseFiles(fp)
	if err != nil {
		return Templator{}, fmt.Errorf("%w: %w", ErrFailedParsingTemplate, err)
	}

	fp = filepath.Join(assetsDirector, layers.DomainLayer, "interface.tmpl")
	domainInterface, err := template.ParseFiles(fp)
	if err != nil {
		return Templator{}, fmt.Errorf("%w: %w", ErrFailedParsingTemplate, err)
	}

	fp = filepath.Join(assetsDirector, layers.DomainLayer, "service.tmpl")
	domainService, err := template.ParseFiles(fp)
	if err != nil {
		return Templator{}, fmt.Errorf("%w: %w", ErrFailedParsingTemplate, err)
	}

	fp = filepath.Join(assetsDirector, layers.DomainLayer, "repository.tmpl")
	domainRepository, err := template.ParseFiles(fp)
	if err != nil {
		return Templator{}, fmt.Errorf("%w: %w", ErrFailedParsingTemplate, err)
	}

	fp = filepath.Join(assetsDirector, layers.ApplicationLayer, "command.tmpl")
	command, err := template.ParseFiles(fp)
	if err != nil {
		return Templator{}, fmt.Errorf("%w: %w", ErrFailedParsingTemplate, err)
	}

	fp = filepath.Join(assetsDirector, layers.ApplicationLayer, "query.tmpl")
	query, err := template.ParseFiles(fp)
	if err != nil {
		return Templator{}, fmt.Errorf("%w: %w", ErrFailedParsingTemplate, err)
	}

	fp = filepath.Join(assetsDirector, layers.ApplicationLayer, "interface.tmpl")
	appInterface, err := template.ParseFiles(fp)
	if err != nil {
		return Templator{}, fmt.Errorf("%w: %w", ErrFailedParsingTemplate, err)
	}

	fp = filepath.Join(assetsDirector, layers.ApplicationLayer, "service.tmpl")
	appService, err := template.ParseFiles(fp)
	if err != nil {
		return Templator{}, fmt.Errorf("%w: %w", ErrFailedParsingTemplate, err)
	}

	fp = filepath.Join(assetsDirector, layers.InfrastructureLayer, "repository.tmpl")
	infraRepository, err := template.ParseFiles(fp)
	if err != nil {
		return Templator{}, fmt.Errorf("%w: %w", ErrFailedParsingTemplate, err)
	}

	fp = filepath.Join(assetsDirector, layers.APILayer, "rest/controller.tmpl")
	apiController, err := template.ParseFiles(fp)
	if err != nil {
		return Templator{}, fmt.Errorf("%w: %w", ErrFailedParsingTemplate, err)
	}

	fp = filepath.Join(assetsDirector, layers.APILayer, "rest/request.tmpl")
	apiRequest, err := template.ParseFiles(fp)
	if err != nil {
		return Templator{}, fmt.Errorf("%w: %w", ErrFailedParsingTemplate, err)
	}

	fp = filepath.Join(assetsDirector, layers.APILayer, "rest/response.tmpl")
	apiResponse, err := template.ParseFiles(fp)
	if err != nil {
		return Templator{}, fmt.Errorf("%w: %w", ErrFailedParsingTemplate, err)
	}

	t = Templator{
		entity:           entity,
		value:            value,
		domainInterface:  domainInterface,
		domainService:    domainService,
		domainRepository: domainRepository,

		command:      command,
		query:        query,
		appInterface: appInterface,
		appService:   appService,

		infraRepository: infraRepository,

		apiController: apiController,
		apiRequest:    apiRequest,
		apiResponse:   apiResponse,
	}

	return t, nil
}

// ----------------------------------------------------------------------------

func (t *Templator) TemplateEntity(e dto.Entity) (data []byte, err error) {
	return templateObject(t.entity, e)
}

func (t *Templator) TemplateValue(e dto.Value) (data []byte, err error) {
	return templateObject(t.value, e)
}

func (t *Templator) TemplateDomainInterface(s dto.Service) (data []byte, err error) {
	return templateObject(t.domainInterface, s)
}

func (t *Templator) TemplateDomainService(s dto.Service) (data []byte, err error) {
	return templateObject(t.domainService, s)
}

func (t *Templator) TemplateDomainRepository(r dto.Repository) (data []byte, err error) {
	return templateObject(t.domainRepository, r)
}

// ----------------------------------------------------------------------------

func (t *Templator) TemplateCommand(c dto.Command) (data []byte, err error) {
	return templateObject(t.command, c)
}

func (t *Templator) TemplateQuery(q dto.Query) (data []byte, err error) {
	return templateObject(t.query, q)
}

func (t *Templator) TemplateApplicationInterface(s dto.Service) (data []byte, err error) {
	return templateObject(t.appInterface, s)
}

func (t *Templator) TemplateApplicationService(s dto.Service) (data []byte, err error) {
	return templateObject(t.appService, s)
}

// ----------------------------------------------------------------------------

func (t *Templator) TemplateInfrastructureRepository(r dto.Repository) (data []byte, err error) {
	return templateObject(t.infraRepository, r)
}

// ----------------------------------------------------------------------------

func (t *Templator) TemplateRESTAPIController(r dto.Controller) (data []byte, err error) {
	return templateObject(t.apiController, r)
}

func (t *Templator) TemplateRESTAPIRequest(r dto.Request) (data []byte, err error) {
	return templateObject(t.apiRequest, r)
}

func (t *Templator) TemplateRESTAPIResponse(r dto.Response) (data []byte, err error) {
	return templateObject(t.apiResponse, r)
}

func templateObject(t *template.Template, object any) (data []byte, err error) {
	b := bytes.NewBuffer(make([]byte, 0))

	err = t.Execute(b, object)
	if err != nil {
		return nil, fmt.Errorf("failed executing template: %w", err)
	}

	data = b.Bytes()
	return data, nil
}
