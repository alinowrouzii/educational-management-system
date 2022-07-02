package main

import (
	"database/sql"
	"log"
)

// ****************Token table**********************
var createTokenTable = `
	CREATE TABLE token (
		id VARCHAR(512) NOT NULL,
		username VARCHAR(32) NOT NULL,
		role VARCHAR(32) NOT NULL,
		issue_at DATETIME NOT NULL,
		expired_at DATETIME NOT NULL,
		PRIMARY KEY(id)
	)
`

// ****************END of token table********************************

// ****************User login function*******************************
var loginUserFunc = `
	CREATE FUNCTION login_user ( username VARCHAR(7) , user_password VARCHAR(512) )
	RETURNS VARCHAR(16) DETERMINISTIC

	BEGIN

		DECLARE user_hashed_password VARCHAR(512);
		DECLARE LOGIN_STATUS int DEFAULT 0;
		DECLARE RETURN_VALUE VARCHAR(16) DEFAULT "FAIL";

		SET user_hashed_password := MD5(user_password);

		SELECT count(*) INTO LOGIN_STATUS
		FROM student 
		WHERE student.student_no=username AND student.password=user_hashed_password;

		IF LOGIN_STATUS=0 THEN
			SELECT count(*) INTO LOGIN_STATUS
			FROM professor 
			WHERE professor.professor_no=username AND professor.password=user_hashed_password;
			IF LOGIN_STATUS > 0 THEN
				SET RETURN_VALUE = "PROFESSOR";
			END IF;
		ELSEIF LOGIN_STATUS > 0 THEN
			SET RETURN_VALUE = "STUDENT";
		END IF;

		RETURN RETURN_VALUE;
	END;
`

// ****************End of User login function************************

// ***************Student TABLE***************
var createStudent = `
	CREATE TABLE student (
		national_code CHAR(10),
		student_no CHAR(7),
		full_name_fa VARCHAR(40) NOT NULL,
		full_name_en VARCHAR(40) NOT NULL,
		father_name VARCHAR(40) NOT NULL,
		birth_date VARCHAR(40) NOT NULL,
		mobile CHAR(11),
		major VARCHAR(64) NOT NULL,
		password VARCHAR(512),
		email VARCHAR(64),
		PRIMARY KEY (student_no),
		UNIQUE(national_code)
	)
`
var dropStudent = `DROP TABLE student`

// this trigger create email for user and also set default hashed password for user that starts with national_code + first_char of his first_name in capital form + first_char of last_name in lower form
var studentTriggerBeforeSave = `
	CREATE TRIGGER set_student_password_email BEFORE INSERT
	ON student
	FOR EACH ROW
	BEGIN
		DECLARE full_name_en VARCHAR(40);
		DECLARE _password VARCHAR(512);
		SET full_name_en := REPLACE(LOWER(NEW.full_name_en), " ", "");
		SET _password := CONCAT(NEW.national_code, UPPER(SUBSTRING(full_name_en, 1, 1)), LOWER(SUBSTRING(full_name_en, POSITION("-" IN full_name_en)+1, 1)));
		SET NEW.email = CONCAT(SUBSTRING(full_name_en, 1, 1), ".", SUBSTRING(full_name_en, POSITION("-" IN full_name_en)+1), "@aut.ac.ir");
		SET NEW.password = MD5(_password);
	END
`

var dropStudentTrigger = `DROP TRIGGER set_student_password_email`

var createStudentChangePasswordFunc = `
	CREATE FUNCTION change_student_password ( student_no VARCHAR(7) , student_password VARCHAR(512), student_new_password VARCHAR(512) )
	RETURNS INT DETERMINISTIC

	BEGIN

	DECLARE user_old_password VARCHAR(512);
	DECLARE user_new_password VARCHAR(512);
	DECLARE AFFECTED_ROWS int DEFAULT 0;
	declare ERROR_MESSAGE varchar(128);

	if student_new_password REGEXP '^[0-9]+$' or student_new_password REGEXP '^[A-Za-z]+$'  then
		set ERROR_MESSAGE = "Password should be alphanumeric";
		signal sqlstate '45000' set message_text = ERROR_MESSAGE;
	end if;

	if LENGTH(student_new_password) < 3 THEN
		set ERROR_MESSAGE = "Password is too short";
		signal sqlstate '45000' set message_text = ERROR_MESSAGE;
	END IF;

	if LENGTH(student_new_password) > 20 THEN
		set ERROR_MESSAGE = "Password is too long";
		signal sqlstate '45000' set message_text = ERROR_MESSAGE;
	END IF;

	SET user_old_password := MD5(student_password);
	SET user_new_password := MD5(student_new_password);

	UPDATE student
	SET student.password = user_new_password
	WHERE student.student_no = student_no AND student.password = user_old_password;

	SELECT ROW_COUNT() into AFFECTED_ROWS;

	RETURN AFFECTED_ROWS;
END;
`

// ***************END of student TABLE*********************

// ***************professor TABLE*******************************
var createProfessor = `
	CREATE TABLE professor (
		national_code CHAR(10),
		professor_no CHAR(5),
		full_name_fa VARCHAR(40) NOT NULL,
		full_name_en VARCHAR(40) NOT NULL,
		father_name VARCHAR(40) NOT NULL,
		birth_date VARCHAR(40) NOT NULL,
		mobile CHAR(11),
		department VARCHAR(64) NOT NULL,
		title ENUM("استاد", "استادیار", "دانش‌یار") NOT NULL,
		password VARCHAR(512),
		email VARCHAR(64),
		PRIMARY KEY (professor_no),
		UNIQUE(national_code)
	)
`
var dropProfessor = `DROP TABLE professor`

var createProfessorChangePassword = `
	CREATE FUNCTION change_professor_password ( professor_no CHAR(5) , professor_password VARCHAR(512), professor_new_password VARCHAR(512) )
	RETURNS INT DETERMINISTIC
	BEGIN
		DECLARE user_old_password VARCHAR(512);
		DECLARE user_new_password VARCHAR(512);
		DECLARE AFFECTED_ROWS int DEFAULT 0;
		SET user_old_password := MD5(professor_password);
		SET user_new_password := MD5(professor_new_password);
		UPDATE professor
		SET professor.password = user_new_password
		WHERE professor.professor_no = professor_no AND professor.password = user_old_password;
		SELECT ROW_COUNT() into AFFECTED_ROWS;
		RETURN AFFECTED_ROWS;
	END;
`

// this trigger create email for user and also set default hashed password for user that starts with national_code + first_char of his first_name in capital form + first_char of last_name in lower form
var professorTriggerBeforeSave = `
	CREATE TRIGGER set_professor_password_email BEFORE INSERT
	ON professor
	FOR EACH ROW
	BEGIN
		DECLARE full_name_en VARCHAR(40);
		DECLARE _password VARCHAR(512);
		SET full_name_en := REPLACE(LOWER(NEW.full_name_en), " ", "");
		SET _password := CONCAT(NEW.national_code, UPPER(SUBSTRING(full_name_en, 1, 1)), LOWER(SUBSTRING(full_name_en, POSITION("-" IN full_name_en)+1, 1)));
		SET NEW.email = CONCAT(SUBSTRING(full_name_en, 1, 1), ".", SUBSTRING(full_name_en, POSITION("-" IN full_name_en)+1), "@aut.ac.ir");
		SET NEW.password = MD5(_password);
	END
`
var dropProfessorTrigger = `DROP TRIGGER set_professor_password_email`

// **************End of professor TABLE****************

// **************course TABLE**************************
var createCourse = `
	CREATE TABLE course (
		course_id CHAR(8),
		course_name VARCHAR(64),
		professor_no CHAR(5) NOT NULL,
		PRIMARY KEY(course_id),
		FOREIGN KEY(professor_no) REFERENCES professor(professor_no)
	)
`
var dropCourse = `DROP TABLE course`

// **************end of course TABLE**************************

var createCourseTakes = `
	CREATE TABLE course_takes(
		student_no CHAR(7),
		course_id CHAR(8),
		PRIMARY KEY(student_no, course_id),
		FOREIGN KEY(course_id) REFERENCES course(course_id),
		FOREIGN KEY(student_no) REFERENCES student(student_no)
	)
`
var dropCourseTakes = `DROP YABLE course_takes`

// ***************Exam question TABLE***********
var createExam = `
	CREATE TABLE exam(
		exam_id INT AUTO_INCREMENT,
		exam_name VARCHAR(32),
		start_date DATETIME,
		end_date DATETIME,
		duration INT,
		course_id CHAR(8),
		PRIMARY KEY (exam_id),
		FOREIGN KEY (course_id) REFERENCES course(course_id)
	)
`
var dropExam = `DROP TABLE exam`

// ***************END of exam question***********

var createExamQuestion = `
	CREATE TABLE exam_question (
		question_id INT AUTO_INCREMENT,
		question_description varchar(512) NOT NULL,
		first_choice  VARCHAR(512) NOT NULL,
		second_choice  VARCHAR(512) NOT NULL,
		third_choice VARCHAR(512) NOT NULL,
		fourth_choice VARCHAR(512) NOT NULL,
		score INT NOT NULL,
		correct_answer ENUM('A', 'B', 'C', 'D'),
		exam_id INT NOT NULL,
		PRIMARY KEY (question_id),
		FOREIGN KEY (exam_id) REFERENCES exam(exam_id)
	)
`
var dropExamQuestion = `DROP TABLE exam_question`

// ***************ExamAsnwer TABLE*********************
var createExamAnswer = `
CREATE TABLE exam_answer (
	question_id INT AUTO_INCREMENT,
	user_answer ENUM('A', 'B', 'C', 'D') NOT NULL,
	exam_id INT NOT NULL,
	student_no CHAR(7) NOT NULL,
	PRIMARY KEY(question_id),
	FOREIGN KEY(student_no) REFERENCES student(student_no),
	FOREIGN KEY(exam_id) REFERENCES exam(exam_id)
)

`

// ***************End of Exam TABLE**************
// var createqUESTIONAnswer = `
// 	CREATE TABLE exam_answer(
// 		student_id CHAR(7),
// 		short_question_id int,
// 		answer VARCHAR(200),
// 		student_grade int DEFAULT 0,
// 		PRIMARY KEY(student_id, short_question_id),
// 		FOREIGN KEY(student_id) REFERENCES student(student_id),
// 		FOREIGN KEY(short_question_id) REFERENCES short_question(question_id)
// 	)
// `
// var dropShortAnswer = `DROP TABLE short_answer`

// var createHomework = `
// 	CREATE TABLE hw(
// 		id int AUTO_INCREMENT,
// 		hw_number int UNSIGNED,
// 		-- section_id
// 		course_code int NOT NULL,
// 		group_number VARCHAR(4) NOT NULL,
// 		year CHAR(4) NOT NULL,
// 		semester ENUM('FALL', 'SPRING'),
// 		--
// 		description VARCHAR(200) NOT NULL,
// 		PRIMARY KEY(id),
// 		FOREIGN KEY(course_code, group_number, year, semester) REFERENCES section(course_code, group_number, year, semester)
// 	)
// `
// var dropHomework = `DROP TABLE hw`

// var createHomeworkParticipation = `
// 	CREATE TABLE hw_participation(
// 		student_id CHAR(7),
// 		hw_id int,
// 		date DATE not null,
// 		grade int UNSIGNED DEFAULT 0,
// 		file VARCHAR(400) NOT NULL,
// 		PRIMARY KEY(student_id, hw_id),
// 		FOREIGN KEY(student_id) REFERENCES student(student_id),
// 		FOREIGN KEY(hw_id) REFERENCES hw(id)
// 	)
// `
// var dropHomeworkParticipation = `DROP TABLE hw_participation`

var execs = []struct {
	stmt       string
	shouldFail bool
}{
	{
		stmt:       createStudent,
		shouldFail: false,
	},
	{
		stmt:       studentTriggerBeforeSave,
		shouldFail: false,
	},
	{
		stmt:       createStudentChangePasswordFunc,
		shouldFail: false,
	},
	{
		stmt:       loginUserFunc,
		shouldFail: false,
	},
	{
		stmt:       createTokenTable,
		shouldFail: false,
	},
	{
		stmt:       createProfessor,
		shouldFail: false,
	},
	{
		stmt:       professorTriggerBeforeSave,
		shouldFail: false,
	},
	{
		stmt:       createCourse,
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
		stmt:       createExamQuestion,
		shouldFail: false,
	},
	{
		stmt:       createExamAnswer,
		shouldFail: false,
	},
	// {
	// 	stmt:       dropStudent,
	// 	shouldFail: false,
	// },
	// {
	// 	stmt:       dropProfessor,
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
