# ESEP Grade Calculator

A Go library that calculates a final letter grade from assignments,
exams, and essays.

## Requirements

Go 1.24.4 or a compatible newer version.

## Setup

Clone this repository and enter its Go module directory:

```bash
git clone https://github.com/YOUR_USERNAME/esep-grade-calculator.git
cd esep-grade-calculator/grade-calculator
```

## Usage

Within the esepunittests package:

```go
calculator := NewGradeCalculator()
calculator.AddGrade("Homework 1", 90, Assignment)
calculator.AddGrade("Exam 1", 80, Exam)
calculator.AddGrade("Essay 1", 100, Essay)
result := calculator.GetFinalGrade()
```

The required calculation averages each category, then applies
weights of 50% for assignments, 35% for exams, and 15% for essays.

## Run Tests

From the grade-calculator directory:

```bash
go test -v ./...
```

## Measure Coverage

```bash
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out
```

## Initial Project Status

The starter implementation and tests contain known defects.
Some provided tests fail before the assignment's corrections
are applied.