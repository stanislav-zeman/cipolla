package request

type getProgramRequest struct{}

func ParsegetProgramRequest(c *echo.Context) (r getProgramRequest, err error) {
    return getProgramRequest{}, nil
}

func (r *getProgramRequest) To() error {
    return nil
}
