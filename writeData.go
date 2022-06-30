package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// CREATE TABLE student (
// 	national_code CHAR(10),
// 	student_no CHAR(7),
// 	full_name_fa VARCHAR(40) NOT NULL,
// 	full_name_en VARCHAR(40) NOT NULL,
// 	father_name VARCHAR(40) NOT NULL,
// 	birth_date VARCHAR(40) NOT NULL,
// 	mobile CHAR(10),
// 	major VARCHAR(10) NOT NULL,
// 	password VARCHAR(512),
// 	email VARCHAR(64),
// 	PRIMARY KEY (student_no),
// 	UNIQUE(national_code),
// 	UNIQUE(student_no)
// )

var insertIntoStudents = `INSERT INTO student(national_code, student_no, full_name_fa, full_name_en, father_name, birth_date, mobile, major) `

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

	res, _ := stmt.Exec(vals...)

	fmt.Println("result after inserting student", res)
}

func WriteDataToDatabsae(db *sql.DB) {
	data := readJsonData()
	// fmt.Println(data["students"])
	writeStudentsData(data["students"], db)

}
