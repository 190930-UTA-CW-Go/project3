package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	_ "github.com/lib/pq" // necessary
)

//Portfolio stores all of the information pulled from the portfolio form. This is stored as global variable.
type Portfolio struct {
	Information Information `json:"Information"`
	About       About       `json:"About"`
	Education   Education   `json:"Education"`
	Project     Project     `json:"Project"`
	PortStatus  PortStatus  `json:"PortStatus"`
}

//Information stores all of the relevant info on the person creating the portfolio.
type Information struct {
	Name  string `json:"Name"`
	Title string `json:"Title"`
	Email string `json:"Email"`
	Phone string `json:"Phone"`
}

//About stores the text the user enters in the "About Me" field.
type About struct {
	Aboutme string `json:"Aboutme"`
}

//Education stores the college the user attended and the degree they got, if any.
type Education struct {
	College string `json:"College"`
	Degree  string `json:"Degree"`
}

//Project stores information on the user's projects that they have completed, if any.
type Project struct {
	Name string `json:"Name"`
	Tech string `json:"Tech"`
	Desc string `json:"Desc"`
}

// PortStatus =
type PortStatus struct {
	Status  string `json:"Status"`
	Comment string `json:"Comment"`
}

// OpenDB = opens Postgres database
func OpenDB() *sql.DB {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		panic(err)
	}
	return db
}

// StartDB = starts the docker database
func StartDB() {
	fmt.Println("-> Starting Database")
	exec.Command("docker", "stop", "project3").Run()
	exec.Command("docker", "rmi", "project3").Run()
	exec.Command("docker", "build", "-t", "project3", "db/.").Run()
	exec.Command("docker", "run", "-p=5432:5432", "-d", "--rm", "--name=projec3", "project3").Run()
	fmt.Println("-> Running Container")
}

// Insert =
// Check if User "id" is in the database
// If "FALSE" then add employee to database
// If "TRUE" then update employee info for changes
func Insert(id, accountname, displayname, email string) {
	var identity string
	db := OpenDB()
	defer db.Close()

	// Check if employee "id" already exists
	row := db.QueryRow("select 1 from employees where id=$1", id)
	switch err := row.Scan(&identity); err {

	// "FALSE" = add employee to database
	case sql.ErrNoRows:
		fmt.Println("-> Adding New Employee")
		sql := `
		insert into employees (id, accountname, displayname, email)
		values ($1, $2, $3, $4)`

		_, err := db.Exec(sql, id, accountname, displayname, email)
		if err != nil {
			fmt.Println("ERROR:", err)
		}

	// "TRUE" = update employee info for changes
	case nil:
		fmt.Println("-> Employee Found Updating Information")
		sql := `
			update employees
			set accountname = $1, displayname = $2, email = $3
			where id = $4`

		_, err = db.Exec(sql, accountname, displayname, email, id)
		if err != nil {
			fmt.Println("ERROR:", err)
		}

	default:
		fmt.Println("ERROR:", err)
	}

}

// GetAccountName = get employee "accountname" using their "id"
func GetAccountName(id string) (result string) {
	db := OpenDB()
	defer db.Close()
	sql := `select accountname from employees where id = $1`
	row := db.QueryRow(sql, id)
	row.Scan(&result)
	return
}

// GetDisplayName = get employee "displayname" using their "id"
func GetDisplayName(id string) (result string) {
	db := OpenDB()
	defer db.Close()
	sql := `select displayname from employees where id = $1`
	row := db.QueryRow(sql, id)
	row.Scan(&result)
	return
}

/*
// GetEmail = get employee "email" using their "id"
func GetEmail(id string) (result string) {
	db := OpenDB()
	defer db.Close()
	sql := `select email from employees where id = $1`
	row := db.QueryRow(sql, id)
	row.Scan(&result)
	return
}
*/

// GetSize = get folder "size" using their "id"
func GetSize(id string) (result string) {
	db := OpenDB()
	defer db.Close()
	sql := `select size from folder where id = $1`
	row := db.QueryRow(sql, id)
	row.Scan(&result)
	return
}

// GetTime = get folder "time" using their "id"
func GetTime(id string) (result string) {
	db := OpenDB()
	defer db.Close()
	sql := `select time from folder where id = $1`
	row := db.QueryRow(sql, id)
	row.Scan(&result)
	return
}

// GetID = get employee "id" using their "email"
func GetID(email string) (result string) {
	db := OpenDB()
	defer db.Close()
	sql := `select id from employees where email = $1`
	row := db.QueryRow(sql, email)
	row.Scan(&result)
	return
}

// UpdateFolder = update "folder" info at "id"
func UpdateFolder(id string, size string, time string) {
	db := OpenDB()
	defer db.Close()
	sql := `
	update folder
	set size = $1, time = $2
	where id = $3`

	_, err := db.Exec(sql, size, time, id)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

// NewFolder =
// Create folder for employee using their "id" and return true
// If one already exists return false
func NewFolder(id string) bool {
	path := "../Portfolios/"
	path += id
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("-> Creating New Folder", id)
		os.Mkdir(path, 0700)
		AddFolder(id, path)
		return true
	}
	return false
}

// AddFolder = add "folder" to database
func AddFolder(id string, path string) {
	size, time := FolderStatus(path)
	db := OpenDB()
	defer db.Close()
	sql := `
	insert into folder (id, size, time)
	values ($1, $2, $3)`
	_, err := db.Exec(sql, id, size, time)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

// FolderStatus = retrieve "folder" size and time stats
func FolderStatus(path string) (size string, time string) {
	stat, _ := os.Stat(path)
	size = strconv.Itoa(int(stat.Size()))
	time = stat.ModTime().String()
	return size, time
}

// CompareStatus =
// Compare "folder" status in database with current status
// If there are changes, update "folder" status in database
func CompareStatus(id string) bool {
	initialSize := GetSize(id)
	initialTime := GetTime(id)

	path := "../Portfolios/"
	path += id
	currentSize, currentTime := FolderStatus(path)

	fmt.Println("InitialTime:", initialTime)
	fmt.Println("CurrentTime:", currentTime)

	if initialSize != currentSize || initialTime != currentTime {
		UpdateFolder(id, currentSize, currentTime)
		return true
	}
	return false
}

// FindFile =
func FindFile(email string) (string, string) {
	//hold := GetID(email)
	hold := ParseEmail(email)
	//fmt.Println(hold)

	path := "../Portfolios/"
	path += hold
	path += "/"
	path += "portfolio.json"
	fmt.Println("FindFile:", path)

	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Could not find the file")
		return "ERROR", "ERROR"
	}
	fmt.Println(string(file))

	data := Portfolio{}
	_ = json.Unmarshal([]byte(file), &data)

	return data.Information.Name, data.PortStatus.Status
}

// EditFile =
func EditFile(user string, status string) {
	path := "../Portfolios/"
	path += user
	path += "/"
	path += "portfolio.json"
	fmt.Println("PATH:", path)

	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Could not find the file")
	}
	fmt.Println(string(file))

	data := Portfolio{}
	_ = json.Unmarshal([]byte(file), &data)

	if !((status == "APPROVED" && data.PortStatus.Status == "APPROVED") || (status == "DENIED" && data.PortStatus.Status == "DENIED")) {
		//Delete old json file so new information isn't appended to existing file.
		cmd := exec.Command("rm", path)
		cmd.Run()

		//Create new json file of same name, all data is same except for modified status.
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return
		}
		defer f.Close()

		// Up to this point, "data" is a new json file that contains the exact information of the old json file.
		// If a choice 1 or 2 is made, Status will be changed before json file is remade.
		if status == "APPROVED" {
			data.PortStatus.Status = "APPROVED"
			fmt.Println("Status set to APPROVED")
		}

		if status == "DENIED" {
			data.PortStatus.Status = "DENIED"
			fmt.Println("Status set to DENIED")
		}

		b, _ := json.MarshalIndent(data, "", "    ") //new json file is remade with same name as old file.
		f.Write(b)
		f.Close()
	}
}

// ParsePayload =
func ParsePayload(payload string) (slice []string) {
	var result string
	result = strings.ReplaceAll(payload, "%7B", " ")
	result = strings.ReplaceAll(result, "%22", " ")
	result = strings.ReplaceAll(result, "%3A", " ")
	result = strings.ReplaceAll(result, "%2C", " ")
	result = strings.ReplaceAll(result, "%5C", " ")
	result = strings.ReplaceAll(result, "%2F", " ")
	result = strings.ReplaceAll(result, "%7D", " ")
	result = strings.ReplaceAll(result, "%5D", " ")
	result = strings.ReplaceAll(result, "%5B", " ")
	result = strings.ReplaceAll(result, "%27", "'")
	result = strings.ReplaceAll(result, "%40", "@")

	slice = strings.Fields(result)
	return slice

}

// GetEmail =
func GetEmail(slice []string) string {
	return slice[9]
}

// GetButton =
func GetButton(slice []string) string {
	return slice[5]
}

// ParseEmail =
func ParseEmail(email string) (user string) {
	index := strings.Index(email, "@")
	if index == -1 {
		return user
	}

	user = email[0:index]
	return user
}
