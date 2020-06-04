package ball

import (
	"errors"
)

type Ball struct{}

func NewBall() *Ball {
	return &Ball{}
}

type Container struct {
	ID    int
	Max   int
	Balls []*Ball
}

func NewContainer(id int, max int, b []*Ball) *Container {

	balls := []*Ball{}
	for range b {
		balls = append(balls, NewBall())
	}
	return &Container{
		ID:    id,
		Max:   max,
		Balls: balls,
	}
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
