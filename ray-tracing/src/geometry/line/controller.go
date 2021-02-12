package line

// ExtractLine is a function to extract a line from 2 points.
//
// Parameters:
// 	p0 - the starting point.
//  p1 - the target point.
//
// Returns:
// 	the Line.
//
//func ExtractLine(p0, p1 Point) Line {
//	v := ExtractVector(&p0, &p1)
//	line := Line{Start: p0, Director: v}
//	return line
//}

// FindPos is a function to get the position of a line at a given t.
//
// Parameters:
// 	t - the t parameter.
//
// Returns:
// 	the Point.
//
//func (line Line) FindPos(t float64) Point {
//	v := utils.CMultVector(&line.Director, t)
//	pos := InitPoint(3)
//	for i := 0; i < 3; i++ {
//		pos.Coordinates[i] = line.Start.Coordinates[i] + v.Coordinates[i]
//	}
//	return pos
//}
