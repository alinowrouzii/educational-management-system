package main

import (
	"database/sql"
	"log"
)

var createStudent = `
	CREATE TABLE student (
		national_code CHAR(10),
		student_no CHAR(7),
		full_name_fa VARCHAR(40) NOT NULL,
		full_name_en VARCHAR(40) NOT NULL,
		father_name VARCHAR(40) NOT NULL,
		birth_date VARCHAR(40) NOT NULL,
		mobile CHAR(10),
		major VARCHAR(10) NOT NULL,
		password VARCHAR(512),
		email VARCHAR(64),
		PRIMARY KEY (student_no),
		UNIQUE(national_code),
		UNIQUE(student_no)
	)
`
var dropStudent = `DROP TABLE student`

var createMaster = `
	CREATE TABLE master (
		national_id CHAR(10),
		full_name VARCHAR(40) NOT NULL,
		personnel_code VARCHAR(10) NOT NULL,
		academic_level ENUM('assistant', 'associate'),
		PRIMARY KEY (personnel_code)
	)
`
var dropMaster = `DROP TABLE master`

var createCourse = `
	CREATE TABLE course (
		name VARCHAR(20),
		unit int UNSIGNED NOT NULL,
		code int NOT NULL,
		PRIMARY KEY(code)
	)
`
var dropCourse = `DROP TABLE course`

var createSection = `
	CREATE TABLE section(
		master_id VARCHAR(10) NOT NULL,
		course_code int NOT NULL,
		group_number VARCHAR(4) NOT NULL,
		year CHAR(4) NOT NULL,
		semester ENUM('FALL', 'SPRING'),
		PRIMARY KEY(course_code, group_number, year, semester),
		FOREIGN KEY(master_id) REFERENCES master(personnel_code),
		FOREIGN KEY(course_code) REFERENCES course(code)
	)
`
var dropSection = `DROP TABLE section`

var createCourseTakes = `
	CREATE TABLE course_takes(
		student_id CHAR(7),
		course_code int,
		group_number VARCHAR(4),
		year CHAR(4),
		semester ENUM('FALL', 'SPRING'),
		PRIMARY KEY(student_id, course_code, group_number, year, semester),
		FOREIGN KEY(course_code, group_number, year, semester) REFERENCES section(course_code, group_number, year, semester),
		FOREIGN KEY(student_id) REFERENCES student(student_id)
	)
`
var dropCourseTakes = `DROP YABLE course_takes`

var createExam = `
	CREATE TABLE exam (
		exam_id int AUTO_INCREMENT,
		-- section_id
		course_code int NOT NULL,
		group_number VARCHAR(4) NOT NULL,
		year CHAR(4) NOT NULL,
		semester ENUM('FALL', 'SPRING'),
		--        
		name VARCHAR(20),
		description VARCHAR(200),
		start_date DATE NOT NULL,
		exam_duration int UNSIGNED NOT NULL,
		PRIMARY KEY(exam_id),
		FOREIGN KEY(course_code, group_number, year, semester) REFERENCES section(course_code, group_number, year, semester)
	)
`
var dropExam = `DROP TABLE exam`

var createTestQuestion = `
	CREATE TABLE test_question(
		question_id int AUTO_INCREMENT,
		exam_id int NOT NULL,
		question VARCHAR(200) NOT NULL,
		correct_answer ENUM('A', 'B', 'C', 'D'),
		question_grade int UNSIGNED NOT NULL,
		PRIMARY KEY(question_id),
		FOREIGN KEY(exam_id) REFERENCES exam(exam_id)
	)
`
var dropTestQuestion = `DROP TABLE test_question`

var createShortQuestion = `
	CREATE TABLE short_question(
		question_id int,
		exam_id int NOT NULL,
		question VARCHAR(200) NOT NULL,
		correct_answer VARCHAR(200) NOT NULL,
		question_grade int UNSIGNED NOT NULL,
		PRIMARY KEY(question_id),
		FOREIGN KEY(exam_id) REFERENCES exam(exam_id)
	)
`
var dropShortQuestion = `DROP TABLE short_question`

var createTestAsnwer = `
	CREATE TABLE test_answer(
		student_id CHAR(7),
		test_question_id int,
		selected_option ENUM('A', 'B', 'C', 'D'),
		student_grade int DEFAULT 0,
		PRIMARY KEY(student_id, test_question_id),
		FOREIGN KEY(student_id) REFERENCES student(student_id),
		FOREIGN KEY(test_question_id) REFERENCES test_question(question_id)
	)
`
var dropTestAnswer = `DROP TABLE test_answer`

var createShortAnswer = `
	CREATE TABLE short_answer(
		student_id CHAR(7),
		short_question_id int,
		answer VARCHAR(200),
		student_grade int DEFAULT 0,
		PRIMARY KEY(student_id, short_question_id),
		FOREIGN KEY(student_id) REFERENCES student(student_id),
		FOREIGN KEY(short_question_id) REFERENCES short_question(question_id)
	)
`
var dropShortAnswer = `DROP TABLE short_answer`

var createHomework = `
	CREATE TABLE hw(
		id int AUTO_INCREMENT,
		hw_number int UNSIGNED,
		-- section_id
		course_code int NOT NULL,
		group_number VARCHAR(4) NOT NULL,
		year CHAR(4) NOT NULL,
		semester ENUM('FALL', 'SPRING'),
		--    
		description VARCHAR(200) NOT NULL,
		PRIMARY KEY(id),
		FOREIGN KEY(course_code, group_number, year, semester) REFERENCES section(course_code, group_number, year, semester)
	)
`
var dropHomework = `DROP TABLE hw`

var createHomeworkParticipation = `
	CREATE TABLE hw_participation(
		student_id CHAR(7),
		hw_id int,
		date DATE not null,
		grade int UNSIGNED DEFAULT 0,
		file VARCHAR(400) NOT NULL,
		PRIMARY KEY(student_id, hw_id),
		FOREIGN KEY(student_id) REFERENCES student(student_id),
		FOREIGN KEY(hw_id) REFERENCES hw(id)
	)
`
var dropHomeworkParticipation = `DROP TABLE hw_participation`

var execs = []struct {
	stmt       string
	shouldFail bool
}{
	{
		stmt:       createStudent,
		shouldFail: false,
	},
	{
		stmt:       createMaster,
		shouldFail: false,
	},
	{
		stmt:       createCourse,
		shouldFail: false,
	},
	{
		stmt:       createSection,
		shouldFail: false,
	},
	{
		stmt:       createCourseTakes,
		shouldFail: false,
	},
	{
		stmt:       createExam,
		shouldFail: false,
	},
	{
		stmt:       createTestQuestion,
		shouldFail: false,
	},
	{
		stmt:       createShortQuestion,
		shouldFail: false,
	},
	{
		stmt:       createTestAsnwer,
		shouldFail: false,
	},
	{
		stmt:       createShortAnswer,
		shouldFail: false,
	},
	{
		stmt:       createHomework,
		shouldFail: false,
	},
	{
		stmt:       createHomeworkParticipation,
		shouldFail: false,
	},
	// {
	// 	stmt:       dropStudent,
	// 	shouldFail: false,
	// },
	// {
	// 	stmt:       dropMaster,
	// 	shouldFail: false,
	// },
	// {
	// 	stmt:       dropCourse,
	// 	shouldFail: false,
	// },
	// {
	// 	stmt:       dropSection,
	// 	shouldFail: false,
	// },
	// {
	// 	stmt:       dropCourseTakes,
	// 	shouldFail: false,
	// },
	// {
	// 	stmt:       dropExam,
	// 	shouldFail: false,
	// },
	// {
	// 	stmt:       dropTestQuestion,
	// 	shouldFail: false,
	// },
	// {
	// 	stmt:       dropShortQuestion,
	// 	shouldFail: false,
	// },
	// {
	// 	stmt:       dropTestAnswer,
	// 	shouldFail: false,
	// },
	// {
	// 	stmt:       dropShortAnswer,
	// 	shouldFail: false,
	// },
	// {
	// 	stmt:       dropHomework,
	// 	shouldFail: false,
	// },
	// {
	// 	stmt:       dropHomeworkParticipation,
	// 	shouldFail: false,
	// },
}

func MakeMigrations(db *sql.DB) {

	for _, exec := range execs {
		_, err := db.Exec(exec.stmt)
		hasFailed := err != nil
		if exec.shouldFail != hasFailed {
			expected := "succeed"
			if exec.shouldFail {
				expected = "fail"
			}
			log.Printf("'%s' should have %sed but did not: %s", exec.stmt, expected, err)
		} else if exec.shouldFail {
			log.Printf("'%s' failed as expected: %s", exec.stmt, err)
		}
	}
	log.Println("finish!")
}
