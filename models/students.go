package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type Student struct {
	StudentNO  string `json:"student_no"`
	FullNameFA string `json:"full_name_fa"`
	FullNameEN string `json:"full_name_en"`
	FatherName string `json:"father_name"`
	BirthDate  string `json:"birth_date"`
	Mobile     string `json:"mobile"`
	Major      string `json:"major"`
	Email      string `json:"email"`
}

var getCourseStudents = `
	SELECT student.student_no, full_name_fa, full_name_en, father_name, birth_date, mobile, major, email
		FROM student, course_takes, course
		WHERE student.student_no=course_takes.student_no 
		AND course_takes.course_id=course.course_id
		AND course.professor_no=?
		AND course.course_id=?
`

func GetCourseStudents(db *sql.DB, username, courseID string) ([]Student, error) {

	rows, err := db.Query(getCourseStudents, username, courseID)

	if err != nil {
		fmt.Println("error occured")
		fmt.Println(err)
		return nil, errors.New("some error occured!")
	}
	defer rows.Close()
	students := []Student{}

	for rows.Next() {
		var student Student
		if err := rows.Scan(
			&student.StudentNO,
			&student.FullNameFA,
			&student.FullNameEN,
			&student.FatherName,
			&student.BirthDate,
			&student.Mobile,
			&student.Major,
			&student.Email,
		); err != nil {
			return students, err
		}

		students = append(students, student)
	}
	if err = rows.Err(); err != nil {
		return students, err
	}
	return students, nil
}
