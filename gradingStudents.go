package hackerrank

func gradingStudents(grades []int32) []int32 {
	var c int32
	for i := range grades {
		if grades[i] > 35 {
			c = grades[i] % 5
			if c > 2 {
				grades[i] = grades[i] + (5 - c)
			}
		}
	}
	return grades
}
