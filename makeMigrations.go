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

// ***************Exam TABLE***********
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

var createExamFunction = `
	CREATE FUNCTION create_exam (
		professor_no CHAR(5), 
		exam_name VARCHAR(32), 
		start_date DATETIME, 
		end_date DATETIME, 
		duration INT, 
		course_id CHAR(8)
	)
	RETURNS VARCHAR(32) DETERMINISTIC
	BEGIN
		DECLARE COURSE_FOUND int DEFAULT 0;
		DECLARE RETURN_VALUE VARCHAR(32) DEFAULT "FAIL";
		DECLARE ERROR_MESSAGE varchar(128);

		SELECT COUNT(*) INTO COURSE_FOUND
		FROM course
		WHERE course.course_id=course_id AND course.professor_no=professor_no;

		IF COURSE_FOUND=0 THEN
			set ERROR_MESSAGE = "course not found for professor";
			signal sqlstate '45000' set message_text = ERROR_MESSAGE;
		ELSE
			INSERT INTO exam (exam_name, start_date, end_date, duration, course_id) VALUES (
				exam_name,
				start_date,
				end_date,
				duration,
				course_id
			);
			SET RETURN_VALUE="SUCCESS";
		END IF;
		
		RETURN RETURN_VALUE;
	END;
`

// ***************END of exam TABLE***********

// ***************exam question***********
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

var createExamQuestionFunction = `
	CREATE FUNCTION create_exam_question (
		professor_no VARCHAR(5),
		question_description varchar(512),
		first_choice  VARCHAR(512),
		second_choice  VARCHAR(512),
		third_choice VARCHAR(512),
		fourth_choice VARCHAR(512),
		score INT,
		correct_answer ENUM('A', 'B', 'C', 'D'),
		exam_id INT
	)
	RETURNS VARCHAR(32) DETERMINISTIC
	BEGIN
		DECLARE EXAM_FOUND int DEFAULT 0;
		DECLARE RETURN_VALUE VARCHAR(32) DEFAULT "FAIL";
		DECLARE ERROR_MESSAGE varchar(128);

		SELECT COUNT(*) INTO EXAM_FOUND
		FROM exam, course
		WHERE exam.exam_id = exam_id 
			AND exam.course_id=course.course_id 
			AND course.professor_no=professor_no;

		IF EXAM_FOUND=0 THEN
			set ERROR_MESSAGE = "exam not found for professor";
			signal sqlstate '45000' set message_text = ERROR_MESSAGE;
		ELSE
			INSERT INTO exam_question (question_description, first_choice, second_choice, third_choice, fourth_choice, score, correct_answer, exam_id) VALUES (
				question_description,
				first_choice,
				second_choice,
				third_choice,
				fourth_choice,
				score,
				correct_answer,
				exam_id
			);
			SET RETURN_VALUE="SUCCESS";
		END IF;
		RETURN RETURN_VALUE;
	END;
`

// ***************END of exam question***********

// ***************ExamAsnwer TABLE*********************
var createExamAnswer = `
CREATE TABLE exam_answer (
	question_id INT,
	student_no CHAR(7),
	user_answer ENUM('A', 'B', 'C', 'D') NOT NULL,
	score INT DEFAULT 0,
	PRIMARY KEY(question_id, student_no),
	FOREIGN KEY(student_no) REFERENCES student(student_no),
	FOREIGN KEY(question_id) REFERENCES exam_question(question_id)
)
`

// exam_id INT NOT NULL,

var submitExamAnswer = `
CREATE FUNCTION submit_exam_answer (
	student_no VARCHAR(7),
	question_id INT,
	user_answer CHAR(1)
)
RETURNS VARCHAR(32) DETERMINISTIC
BEGIN
	DECLARE QUESTION_FOUND int DEFAULT 0;
	DECLARE RETURN_VALUE VARCHAR(32) DEFAULT "FAIL";
	DECLARE ERROR_MESSAGE varchar(128);
	DECLARE EXAM_TIME_IS_NOT_OVER INT DEFAULT 1;
	DECLARE EXAM_TIME_HAS_BEGUN INT DEFAULT 0;

	SELECT COUNT(*) INTO QUESTION_FOUND
	FROM exam_question, exam, course_takes
	WHERE 
		exam_question.question_id = question_id 
		AND exam_question.exam_id=exam.exam_id
		AND exam.course_id=course_takes.course_id
		AND course_takes.student_no=student_no;

	IF QUESTION_FOUND=0 THEN
		set ERROR_MESSAGE = "question not found for student";
		signal sqlstate '45000' set message_text = ERROR_MESSAGE;
	ELSE
		SELECT COUNT(*) INTO EXAM_TIME_IS_NOT_OVER
		FROM exam, exam_question
		WHERE 
			exam_question.question_id=question_id
			AND exam_question.exam_id=exam.exam_id
			AND exam.end_date > NOW();

		IF EXAM_TIME_IS_NOT_OVER=0 THEN
			set ERROR_MESSAGE = "exam time is over dude!";
			signal sqlstate '45000' set message_text = ERROR_MESSAGE;
		END IF;

		SELECT COUNT(*) INTO EXAM_TIME_HAS_BEGUN
		FROM exam, exam_question
		WHERE 
			exam_question.question_id=question_id
			AND exam_question.exam_id=exam.exam_id
			AND exam.start_date < NOW();

		IF EXAM_TIME_HAS_BEGUN=0 THEN
			set ERROR_MESSAGE = "exam has not begin yet dude!";
			signal sqlstate '45000' set message_text = ERROR_MESSAGE;
		END IF;

		INSERT INTO exam_answer (question_id, student_no, user_answer) VALUES (
			question_id,
			student_no,
			user_answer
		);
		SET RETURN_VALUE="SUCCESS";
		
	END IF;
	RETURN RETURN_VALUE;
END;
`
var dropSubmitExamAnswer = `DROP FUNCTION submit_exam_answer`

//*****************************8

var getStudentExamScore = `
CREATE FUNCTION get_student_exam_score (student_no VARCHAR(7), exam_id INT)
RETURNS INT DETERMINISTIC
BEGIN
	DECLARE score int DEFAULT 0;
	DECLARE EXAM_IS_OVER INT DEFAULT 0;
	DECLARE ERROR_MESSAGE VARCHAR(64);

	SELECT COUNT(*) INTO EXAM_IS_OVER
	FROM exam
	WHERE 
		exam.exam_id=exam_id 
		AND exam.end_date < NOW();

	IF EXAM_IS_OVER=0 THEN
		set ERROR_MESSAGE = "Exam is not over yet!";
		signal sqlstate '45000' set message_text = ERROR_MESSAGE;
	END IF;

	SELECT COALESCE(SUM(
		CASE 
			WHEN exam_answer.user_answer = exam_question.correct_answer
			THEN exam_question.score 
			ELSE 0 
		END
	), -1) INTO score
	FROM exam_answer, exam_question, exam, course_takes
	WHERE 
		exam_answer.question_id=exam_question.question_id
		AND exam_question.exam_id=exam_id
		AND exam_question.exam_id=exam.exam_id
		AND exam.course_id=course_takes.course_id
		AND course_takes.student_no=student_no;

	RETURN score;
END;
`
var dropGetStudentExamScoreFunc = `DROP FUNCTION get_student_exam_score`

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
		stmt:       createExamQuestionFunction,
		shouldFail: false,
	},
	{
		stmt:       createExamAnswer,
		shouldFail: false,
	},
	{
		stmt:       dropSubmitExamAnswer,
		shouldFail: false,
	},
	{
		stmt:       submitExamAnswer,
		shouldFail: false,
	},
	{
		stmt:       createExamFunction,
		shouldFail: false,
	},
	{
		stmt:       dropGetStudentExamScoreFunc,
		shouldFail: false,
	},
	{
		stmt:       getStudentExamScore,
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
