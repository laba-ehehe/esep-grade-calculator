package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 50, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 50, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	calculator := NewGradeCalculator()
	calculator.AddGrade("Assignment", 70, Assignment)
	calculator.AddGrade("Exam", 70, Exam)
	calculator.AddGrade("Essay", 70, Essay)

	if actual := calculator.GetFinalGrade(); actual != "C" {
		t.Errorf("Expected C; got %s", actual)
	}
}

func TestGetGradeD(t *testing.T) {
	calculator := NewGradeCalculator()
	calculator.AddGrade("Assignment", 60, Assignment)
	calculator.AddGrade("Exam", 60, Exam)
	calculator.AddGrade("Essay", 60, Essay)

	if actual := calculator.GetFinalGrade(); actual != "D" {
		t.Errorf("Expected D; got %s", actual)
	}
}

func TestGradeTypeString(t *testing.T) {
	cases := []struct {
		gradeType GradeType
		expected  string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
	}

	for _, tc := range cases {
		if actual := tc.gradeType.String(); actual != tc.expected {
			t.Errorf("Expected %s; got %s", tc.expected, actual)
		}
	}
}

func TestMultipleGradesAndWeights(t *testing.T) {
	calculator := NewGradeCalculator()

	// Assignment average = 90.
	calculator.AddGrade("Assignment 1", 80, Assignment)
	calculator.AddGrade("Assignment 2", 100, Assignment)

	// Exam average = 60.
	calculator.AddGrade("Exam 1", 50, Exam)
	calculator.AddGrade("Exam 2", 70, Exam)

	// Essay average = 100.
	calculator.AddGrade("Essay 1", 100, Essay)
	calculator.AddGrade("Essay 2", 100, Essay)

	// Weighted score = 90(0.50) + 60(0.35) + 100(0.15) = 81.
	if actual := calculator.GetFinalGrade(); actual != "B" {
		t.Errorf("Expected B; got %s", actual)
	}
}

func TestPassFailEvaluation(t *testing.T) {
	cases := []struct {
		name     string
		score    int
		expected string
	}{
		{"A passes", 95, "Pass"},
		{"B passes", 85, "Pass"},
		{"C boundary passes", 70, "Pass"},
		{"Below C fails", 69, "Fail"},
		{"D fails", 60, "Fail"},
		{"F fails", 50, "Fail"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			calculator := NewGradeCalculator(PassFail)

			calculator.AddGrade("Assignment", tc.score, Assignment)
			calculator.AddGrade("Exam", tc.score, Exam)
			calculator.AddGrade("Essay", tc.score, Essay)

			actual := calculator.GetFinalGrade()

			if actual != tc.expected {
				t.Errorf("Expected %s; got %s", tc.expected, actual)
			}
		})
	}
}

func TestExplicitLetterGradeMode(t *testing.T) {
	calculator := NewGradeCalculator(LetterGrade)

	calculator.AddGrade("Assignment", 85, Assignment)
	calculator.AddGrade("Exam", 85, Exam)
	calculator.AddGrade("Essay", 85, Essay)

	actual := calculator.GetFinalGrade()

	if actual != "B" {
		t.Errorf("Expected B; got %s", actual)
	}
}
