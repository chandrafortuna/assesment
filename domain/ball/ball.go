package ball

import "errors"

type Ball struct{}

type Container struct {
	ID    int
	Max   int
	Balls []*Ball
}

func (c *Container) IsVerified() bool {
	return len(c.Balls) >= c.Max
}

type InitRequest struct {
	NumContainer int `json:"numContainer"`
	MaxLoad      int `json:"maxLoad"`
}

func (req *InitRequest) Validate() error {

	if req.NumContainer <= 0 {
		return errors.New("Number of container must be > 0")
	}

	if req.MaxLoad <= 0 {
		return errors.New("Max Load of container must be > 0")
	}

	return nil
}
