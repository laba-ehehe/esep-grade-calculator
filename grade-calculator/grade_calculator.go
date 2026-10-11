package esepunittests

type EvaluationMode int

const (
	LetterGrade EvaluationMode = iota
	PassFail
)

type GradeCalculator struct {
	grades []Grade
	mode   EvaluationMode
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator(modes ...EvaluationMode) *GradeCalculator {
	mode := LetterGrade

	if len(modes) > 0 {
		mode = modes[0]
	}

	return &GradeCalculator{
		grades: make([]Grade, 0),
		mode:   mode,
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(
	name string,
	grade int,
	gradeType GradeType,
) {
	// Preserve the original behavior of ignoring invalid grade types.
	switch gradeType {
	case Assignment, Exam, Essay:
		gc.grades = append(gc.grades, Grade{
			Name:  name,
			Grade: grade,
			Type:  gradeType,
		})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignmentAverage := computeAverage(gc.grades, Assignment)
	examAverage := computeAverage(gc.grades, Exam)
	essayAverage := computeAverage(gc.grades, Essay)

	weightedGrade := float64(assignmentAverage)*0.50 +
		float64(examAverage)*0.35 +
		float64(essayAverage)*0.15

	return int(weightedGrade)
}

func computeAverage(grades []Grade, gradeType GradeType) int {
	sum := 0
	count := 0

	for _, grade := range grades {
		if grade.Type == gradeType {
			sum += grade.Grade
			count++
		}
	}

	return sum / count
}
