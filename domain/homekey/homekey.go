package homekey

import (
	"errors"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewPoint(x int, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

type MatrixPoint struct {
	Matrix     [][]string `json:"matrix"`
	StartPoint *Point     `json:"startPoint"`
}

func NewMatrixPoint(matrix [][]string, point *Point) *MatrixPoint {
	return &MatrixPoint{
		Matrix:     matrix,
		StartPoint: point,
	}
}

func (m *MatrixPoint) Validate() error {
	if len(m.Matrix) <= 0 {
		return errors.New("Matrix size must be > 0")
	}

	return nil
}

func (m *MatrixPoint) ValidRange(x int, y int) bool {
	if x >= len(m.Matrix) {
		return false
	}
	if y >= len(m.Matrix[0]) {
		return false
	}
	if x < 0 {
		return false
	}
	if y < 0 {
		return false
	}
	return true
}

func (m *MatrixPoint) North(x int, y int) bool {
	x = x - 1
	if !m.ValidRange(x, y) {
		return false
	}

	if m.Matrix[x][y] == "." {
		return true
	}

	return false
}

func (m *MatrixPoint) East(x int, y int) bool {
	y = y + 1
	if !m.ValidRange(x, y) {
		return false
	}

	if m.Matrix[x][y] == "." {
		return true
	}

	return false
}

func (m *MatrixPoint) South(x int, y int) bool {
	x = x + 1
	if !m.ValidRange(x, y) {
		return false
	}

	if m.Matrix[x][y] == "." {
		return true
	}

	return false
}

type PointCollection struct {
	Points []*Point
}

func (pc *PointCollection) AddPoint(p *Point) *PointCollection {
	pc.Points = append(pc.Points, p)
	return pc
}

func (pc *PointCollection) PointInSlice(p *Point) bool {
	for _, v := range pc.Points {
		if v == p {
			return true
		}
	}
	return false
}

func NewPointCollection(p []*Point) *PointCollection {
	return &PointCollection{
		Points: p,
	}
}

type MatrixPointResult struct {
	TotalStep     int          `json:"totalStep"`
	AvailableStep []*Point     `json:"availableStep"`
	MatrixPoint   *MatrixPoint `json:"matrixPoint"`
}

func (mp *MatrixPointResult) SetAvailableStep(points []*Point) *MatrixPointResult {
	mp.AvailableStep = points
	return mp
}

func (mp *MatrixPointResult) SetTotalStep(total int) *MatrixPointResult {
	mp.TotalStep = total
	return mp
}

func NewMatrixPointResult(points []*Point, totalStep int, matrixPoint *MatrixPoint) *MatrixPointResult {
	return &MatrixPointResult{
		AvailableStep: points,
		TotalStep:     totalStep,
		MatrixPoint:   matrixPoint,
	}
}
