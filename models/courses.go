package models

import (
	"database/sql"
	"errors"
	"fmt"
)

var getUserCourses = `
	(
		SELECT course.course_id, course_name, professor_no 
		FROM course, course_takes 
		WHERE course.course_id=course_takes.course_id AND course_takes.student_no=?
	)
	UNION
	(
		SELECT course.course_id, course_name, professor_no 
		FROM course 
		WHERE course.professor_no=?
	)
`

type Course struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ProfessorNO string `json:"professor_no"`
}

func GetUserCourses(db *sql.DB, username string) ([]Course, error) {

	rows, err := db.Query(getUserCourses, username, username)

	if err != nil {
		fmt.Println("error occured")
		fmt.Println(err)
		return nil, errors.New("some error occured!")
	}
	defer rows.Close()
	courses := []Course{}

	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Name, &course.ProfessorNO); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}
	if err = rows.Err(); err != nil {
		return courses, err
	}
	return courses, nil
}
