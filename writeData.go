package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var insertIntoStudents = `INSERT INTO student(national_code, student_no, full_name_fa, full_name_en, father_name, birth_date, mobile, major) VALUES `
var insertIntoProfessors = `INSERT INTO professor(national_code, professor_no, full_name_fa, full_name_en, father_name, birth_date, mobile, department, title) VALUES `
var insertIntoCourses = `INSERT INTO course(course_id, course_name, professor_no) VALUES `
var insertIntoCourseTakes = `INSERT INTO course_takes(student_no, course_id) VALUES `

func readJsonData() map[string]interface{} {
	// Open our jsonFile
	pwd, _ := os.Getwd()
	fmt.Println(pwd + "/data/data.json")
	jsonFile, err := os.Open(pwd + "/data/data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json", jsonFile)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// https://stackoverflow.com/questions/31398044/got-error-invalid-character-%C3%AF-looking-for-beginning-of-value-from-json-unmar
	byteValue = bytes.TrimPrefix(byteValue, []byte("\xef\xbb\xbf"))

	var result map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		fmt.Println(err)
	}

	return result
}

func writeStudentsData(data interface{}, db *sql.DB) {

	students := data.([]interface{})
	vals := []interface{}{}

	for _, s := range students {
		student := s.(map[string]interface{})
		national_code := student["national_code"].(string)
		student_no := student["student_no"].(string)
		name_fa := student["name_fa"].(string)
		name_en := student["name_en"].(string)
		father_name := student["father_name"].(string)
		birth_date := student["birth_date"].(string)
		mobile := student["mobile"].(string)
		major := student["major"].(string)

		//INSERT INTO student(national_code, student_no, full_name_fa, full_name_en, father_name, birth_date, mobile, major) `
		insertIntoStudents += " (?, ?, ?, ?, ?, ?, ?, ?),"
		vals = append(vals, national_code, student_no, name_fa, name_en, father_name, birth_date, mobile, major)
	}
	// trim last colon
	insertIntoStudents = insertIntoStudents[0 : len(insertIntoStudents)-1]

	stmt, _ := db.Prepare(insertIntoStudents)

	res, err := stmt.Exec(vals...)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("result after inserting student", res)
}
func writeProfessorsData(data interface{}, db *sql.DB) {

	professors := data.([]interface{})
	vals := []interface{}{}

	for _, p := range professors {
		professor := p.(map[string]interface{})
		national_code := professor["national_code"].(string)
		professor_no := professor["professor_no"].(string)
		name_fa := professor["name_fa"].(string)
		name_en := professor["name_en"].(string)
		father_name := professor["father_name"].(string)
		birth_date := professor["birth_date"].(string)
		mobile := professor["mobile"].(string)
		department := professor["department"].(string)
		title := professor["title"].(string)

		insertIntoProfessors += " (?, ?, ?, ?, ?, ?, ?, ?, ?),"
		vals = append(vals, national_code, professor_no, name_fa, name_en, father_name, birth_date, mobile, department, title)
	}
	// trim last colon
	insertIntoProfessors = insertIntoProfessors[0 : len(insertIntoProfessors)-1]

	stmt, _ := db.Prepare(insertIntoProfessors)

	res, err := stmt.Exec(vals...)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("result after inserting professors", res)
}

func writeCoursesData(data interface{}, db *sql.DB) {

	courses := data.([]interface{})
	vals := []interface{}{}

	for _, c := range courses {
		course := c.(map[string]interface{})
		course_id := course["id"].(string)
		course_name := course["name"].(string)
		professor_no := course["professor_no"].(string)

		insertIntoCourses += " (?, ?, ?),"
		vals = append(vals, course_id, course_name, professor_no)
	}
	// trim last colon
	insertIntoCourses = insertIntoCourses[0 : len(insertIntoCourses)-1]

	stmt, _ := db.Prepare(insertIntoCourses)

	res, err := stmt.Exec(vals...)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("result after inserting courses", res)
}

func writeCourseTakesData(data interface{}, db *sql.DB) {

	takes := data.([]interface{})
	vals := []interface{}{}

	for _, t := range takes {
		studentTakes := t.(map[string]interface{})
		student_no := studentTakes["student_no"].(string)
		course_id := studentTakes["course_id"].(string)

		insertIntoCourseTakes += " (?, ?),"
		vals = append(vals, student_no, course_id)
	}
	// trim last colon
	insertIntoCourseTakes = insertIntoCourseTakes[0 : len(insertIntoCourseTakes)-1]

	stmt, _ := db.Prepare(insertIntoCourseTakes)

	res, err := stmt.Exec(vals...)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("result after inserting course_takes", res)
}

func WriteDataToDatabsae(db *sql.DB) {
	data := readJsonData()
	// fmt.Println(data["students"])
	writeStudentsData(data["students"], db)
	writeProfessorsData(data["faculty"], db)
	writeCoursesData(data["courses"], db)
	writeCourseTakesData(data["classrooms"], db)

}
