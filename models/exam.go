package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Exam struct {
	ExamID        string         `json:"exam_id"`
	ExamName      string         `json:"exam_name"  validate:"required"`
	StartDate     time.Time      `json:"start_date" validate:"required"`
	EndDate       time.Time      `json:"end_date"   validate:"required"`
	Duration      int64          `json:"duration"   validate:"required"`
	CourseID      string         `json:"course_id"  validate:"required"`
	ExamQuestions []ExamQuestion `json:"questions"`
}

type ExamQuestion struct {
	QuestionID          string `json:"question_id"`
	QuestionDescription string `json:"question_description"  validate:"required"`
	FirstChoice         string `json:"first_choice"  validate:"required"`
	SecondChoice        string `json:"second_choice"  validate:"required"`
	ThirdChoice         string `json:"third_choice"  validate:"required"`
	FourthChoice        string `json:"fourth_choice"  validate:"required"`
	Score               int64  `json:"score"  validate:"required"`
	CorrectAnswer       string `json:"correct_answer"  validate:"required"`
	ExamID              string `json:"exam_id"  validate:"required"`
}

type ExamAnswer struct {
	QuestionID string `json:"question_id"  validate:"required"`
	// ExamID     string `json:"exam_id"      validate:"required"`
	UserAnswer string `json:"user_answer"  validate:"required"`
	StudentNO  string `json:"student_no"`
}

var getCourseExam = `
	SELECT exam_id, exam_name, start_date, end_date, duration, exam.course_id
	FROM exam, course
	WHERE exam.course_id=?
	AND exam.course_id=course.course_id
	AND course.professor_no=?
`
var createCourseExam = `SELECT create_exam(?, ?, ?, ?, ?, ?)`

var addExamQuestion = "SELECT create_exam_question(?, ?, ?, ?, ?, ?, ?, ?, ?)"
var addExamAnswer = "SELECT submit_exam_answer(?, ?, ?)"

var getExamScore = "SELECT get_student_exam_score(?, ?)"

func GetCourseExams(db *sql.DB, professorNO, courseID string) ([]Exam, error) {

	rows, err := db.Query(getCourseExam, courseID, professorNO)

	if err != nil {
		fmt.Println("error occured")
		fmt.Println(err)
		return nil, errors.New("some error occured!")
	}
	defer rows.Close()
	exams := []Exam{}

	for rows.Next() {
		var exam Exam

		if err := rows.Scan(
			&exam.ExamID,
			&exam.ExamName,
			&exam.StartDate,
			&exam.EndDate,
			&exam.Duration,
			&exam.CourseID,
		); err != nil {
			return exams, err
		}
		exams = append(exams, exam)
	}
	if err = rows.Err(); err != nil {
		return exams, err
	}
	return exams, nil

}

func (exam *Exam) CreateCourseExam(db *sql.DB, professorNO string) error {

	createStatus := "FAIL"
	err := db.QueryRow(createCourseExam, professorNO, exam.ExamName, exam.StartDate, exam.EndDate, exam.Duration, exam.CourseID).Scan(&createStatus)

	if err != nil {
		fmt.Println("error occured")
		fmt.Println(err)
		return err
	}
	if createStatus != "SUCCESS" {
		return errors.New("Some error occured")
	}
	return nil
}

func (question *ExamQuestion) AddExamQuestion(db *sql.DB, professorNO string) error {

	createStatus := "FAIL"
	err := db.QueryRow(addExamQuestion,
		professorNO,
		question.QuestionDescription,
		question.FirstChoice,
		question.SecondChoice,
		question.ThirdChoice,
		question.FourthChoice,
		question.Score,
		question.CorrectAnswer,
		question.ExamID,
	).Scan(&createStatus)

	if err != nil {
		fmt.Println("error occured")
		fmt.Println(err)
		return err
	}

	if createStatus != "SUCCESS" {
		return errors.New("Some error occured")
	}

	return nil

}

func (answer *ExamAnswer) SubmitExamAnswer(db *sql.DB) error {

	createStatus := "FAIL"
	err := db.QueryRow(addExamAnswer, answer.StudentNO, answer.QuestionID, answer.UserAnswer).Scan(&createStatus)

	if err != nil {
		fmt.Println("error occured")
		fmt.Println(err)
		return err
	}
	if createStatus != "SUCCESS" {
		return errors.New("Some error occured")
	}

	return nil
}

func GetStudentExamScore(db *sql.DB, studentNO, examID string) (int, error) {

	examScore := 0
	err := db.QueryRow(getExamScore, studentNO, examID).Scan(&examScore)

	if err != nil {
		fmt.Println("error occured")
		fmt.Println(err)
		return 0, err
	}

	return examScore, nil
}
