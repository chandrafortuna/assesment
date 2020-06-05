package homekey

import "log"

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s *Service) Solution(m MatrixPoint) (*MatrixPointResult, error) {
	matrixPoint := NewMatrixPoint(m.Matrix, m.StartPoint)
	result := NewMatrixPointResult([]*Point{}, 0, matrixPoint)
	pointCollection := NewPointCollection([]*Point{})

	pointCollection = getKey(pointCollection, matrixPoint, matrixPoint.StartPoint.X, matrixPoint.StartPoint.Y, "A")

	uniqPoints := []*Point{}
	totalPoint := 0
	for _, value := range pointCollection.Points {

		if !pointInSlice(value, uniqPoints) {
			log.Println("Print Point [", value.X, ",", value.Y, "]")
			uniqPoints = append(uniqPoints, NewPoint(value.X, value.Y))
			totalPoint++
		}
	}

	result = result.SetAvailableStep(uniqPoints)
	result = result.SetTotalStep(totalPoint)
	return result, nil
}

func pointInSlice(p *Point, list []*Point) bool {
	for _, v := range list {
		if (v.X == p.X) && (v.Y == p.Y) {
			return true
		}
	}
	return false
}

func getKey(res *PointCollection, matrixPoint *MatrixPoint, x int, y int, step string) *PointCollection {
	if !matrixPoint.ValidRange(x, y) {
		return res
	}

	if step == "A" {
		if matrixPoint.North(x, y) {
			res.AddPoint(NewPoint(x, y))
			getKey(res, matrixPoint, x-1, y, "A")
		}
		if matrixPoint.East(x, y) {
			res.AddPoint(NewPoint(x, y))
			getKey(res, matrixPoint, x, y+1, "B")
		}
	} else if step == "B" {
		if matrixPoint.East(x, y) {
			res.AddPoint(NewPoint(x, y))
			getKey(res, matrixPoint, x, y+1, "B")
		}
		if matrixPoint.South(x, y) {
			res.AddPoint(NewPoint(x, y))
			getKey(res, matrixPoint, x+1, y, "C")
		}
	} else if step == "C" {
		if matrixPoint.South(x, y) {
			res.AddPoint(NewPoint(x, y))
			getKey(res, matrixPoint, x+1, y, "C")
		}
	}
	res.AddPoint(NewPoint(x, y))
	return res
}
