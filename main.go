package main

/* SupremeGarbanzo
Lightweight CRUD System Backed By GO

Features list :

	(Server Side)
		Auto Browser Launch(On Windows)
		Real-Time Diversed Server Log Reporting
		Blocks All Unauthorized Access
		Auto Redirect For Non-Existent Pages

	(Database)
		Two-Level Database System
		Main Database(TXT) For CRUD
		Temporary Database(ARRAY) For Manipulation (Searching, Sorting)

	(Sign In & Sign Up Page)
		Sign In & Sign Up In Single Page
		Input Error Handling: Blank, Whitespace, And Too Long String
		Database Error Handling:
			User Must Exist & Password Match(Sign In)
			User Must Not Exist(Sign Up)

	(Landing Page)
		Added "Hello $username!"

	(Dashboard Page)
		Two-Folding Page, Dashboard Is Below

		(Transaction)
			Able To Send & Receive Money
			Error Handling:
				Must Be A Multiple Of 50.000
				Minimum Of 50.000 And Maximum Of 10.000.000
				Target Account: Not(Blank, Whitespace, And Too Long String)

		(Account Management)
			Auto-Redirect To User New Updated Account Page

		(Record History)
			All-Field Search & Sort Capabilities
			Search Return All Matched Results
			Search Error Handling: If No Data Exist, Shows All Data By Default
			Two Sorting Modes: Ascending - Descending

	(About Page)
		Two-Folding Page, About Is Below
		Written With Dedication & Sweat

	(Sign Out)
		Destroy User Session
		Auto-Redirect To Sign In & Sign Up Page
*/

import (

	// Basic Interface - Database IO

	"fmt"
	"io/ioutil"

	// Server Logging

	"log"
	"time"

	// OS - Terminal Command Execution

	"os"
	"os/exec"

	// Standarized String Manipulation

	"strconv"
	"strings"

	// Networking - Web Interface Handling

	"html/template"
	"net/http"
)

const N = 20192020

// Data Nasabah - transaksi

type transaksi struct {
	Tanggal string
	Debit   int
	Kredit  int
	Saldo   int
	Ket     string
}

type ttransaksi struct {
	transaksi [N]transaksi
	n         int
}

type qidNasabah struct {
	Username string
	Password string
	Saldo    int
}

type nasabah struct {
	Id      qidNasabah
	Riwayat []transaksi
}

// End of Data Nasabah - Transaksi
// HTML Head

type Head struct {
	Title string
	CSS   []string
}

type teksL struct {
	data [N]string
	n    int
}

// End of HTML Head
// HTML Page

type Header struct {
	PageTitle string
}

type Menu struct {
	Menu string
	URL  string
}

type tMenu struct {
	data [N]Menu
	n    int
}

type Content struct {
	Title   string
	Content string
}

type tContent struct {
	data [N]Content
	n    int
}

type Footer struct {
	Footer string
}

type Body struct {
	Header  Header
	Menu    []Menu
	Content []Content
	Footer  Footer
}

// End of HTML Page
// Content

type Page struct {
	Head Head
	Body Body
	Data nasabah
}

// End of Content

// Global Variable

var (
	deHeder Header
	deCSS   []string
	deMenu  []Menu
	Foo     Footer

	data Page

	db []nasabah = []nasabah{{qidNasabah{"NGC", "1920", 1000000}, []transaksi{{"011119", 1000000, 0, 1000000, "StarterPack"}}}}

	dbTemp []byte = []byte("NGC 1920 1000000 011119 1000000 0 1000000 StarterPack")
)

// Dummy function
// Intended for testing purposes

func tryit(n string) (string) {
  return n + "!\n"
}

// Reinitialization
// Return Content To Ther Default Values

func reIns() {

	//HTML Parts

	deHeder = Header{"SupremeGarbanzo"}

	deCSS = []string{
		"/static/css/styles.css",
		"https://fonts.googleapis.com/css?family=Nunito:200,600",
	}

	deMenu = []Menu{Menu{"Landing", "/index/"}, Menu{"Dashboard", "/dashb/"}, Menu{"About", "/about/"}, Menu{"Sign Out", "/logout/"}}

	Foo = Footer{"2019 - 2020 | NGC"}

	//Database Reinitialization

	var nsbhTemp nasabah
	var qidTemp qidNasabah
	var riwayatTemp []transaksi

	// File listing

	files, err := ioutil.ReadDir("assets/users/")

	// When no file exist
	// Make new fake admin user

	if err == nil {

		if len(files) == 0 {

			err := ioutil.WriteFile("assets/users/0.txt", dbTemp, 0777)

			if err != nil {

				log.Fatal(err)
			} else {

				fmt.Println(time.Now().Format("02/01/06  03:04:05.000000000"), "\tDatabase Reinitialization Success")
			}

		} else if len(files) > 0 {

			// If There's At Least A File
			// Read From Each Files
			// Then Store It To Server-Cache

			db = nil

			// Temporary Variable

			var res, strTanggal, strKet string
			var strDebit, strKredit, strSaldo int
			var new []string

			// Gets All Files

			for _, file := range files {

				// Gets File By Name

				content, err := ioutil.ReadFile("assets/users/" + file.Name())

				if err == nil {

					res = fmt.Sprintf("%s", content)
					new = strings.Split(res, " ")

					// Store User Account Data To Temp

					qidTemp.Username = new[0]
					qidTemp.Password = new[1]
					qidTemp.Saldo, _ = strconv.Atoi(new[2])

					// Store All User Account Transaction History

					for i := 2; i < len(new)-3; {

						i++
						strTanggal = new[i]
						i++
						strDebit, _ = strconv.Atoi(new[i])
						i++
						strKredit, _ = strconv.Atoi(new[i])
						i++
						strSaldo, _ = strconv.Atoi(new[i])
						i++
						strKet = new[i]

						riwayatTemp = append(riwayatTemp, transaksi{strTanggal, strDebit, strKredit, strSaldo, strKet})
					}

					// Append To Temporary Variable

					nsbhTemp = nasabah{qidTemp, riwayatTemp}
					riwayatTemp = nil

					// Append To Temporary Database,
					// In This Case, Array-Like Database

					db = append(db, nsbhTemp)
				}
			}
		}
	}
}

// Checks If Username & Password Correct,
// Return True If Correct
// Used In SIgn In
// Manual Searching

func userChk(name, pass string) bool {

	for _, val := range db {
		if val.Id.Username == name && val.Id.Password == pass {

			return true
		}
	}

	return false
}

// Checks If User Exist,
// Return True If Said Eser Exist
// Used In Auth Page Request

func userExist(name string) bool {

	for _, val := range db {
		if val.Id.Username == name {

			return true
		}
	}

	return false
}

// Get User File
// Reading Files Manually
// If Match Said Username,
// Return File Name

func getFileName(n string) (string, bool) {

	var fn string

	// Gets All Files

	files, err := ioutil.ReadDir("assets/users/")

	if err == nil {

		for _, file := range files {

			// Gets File By Name

			content, err := ioutil.ReadFile("assets/users/" + file.Name())

			if err == nil {

				res := fmt.Sprintf("%s", content)
				new := strings.Split(res, " ")

				// If Username Matches,
				// Return True

				if new[0] == n {

					fn = file.Name()

					return fn, true
				}
			}
		}
	}

	return "", false
}

// When User Close Their Account,
// It Creates A Void In Database
// Return A Non-Existent File Name
// Or Next Available File Name,
// If There Aren't Any Void

func getFileGap() string {

	var i int
	var fName string

	// Gets All Files

	files, err := ioutil.ReadDir("assets/users/")

	if err == nil {

		// Make A Files List

		for _, file := range files {

			fName = strconv.Itoa(i) + ".txt"

			// If Name Not Matches,
			// Void Found
			// Return File Gap Name

			if fName != file.Name() {

				return fName
			} else {

				i++
			}
		}
	}

	// Else
	// Return Next Available File Name

	fName = strconv.Itoa(i) + ".txt"

	return fName
}

// Name Is Self-Explanatory
// Get User's File From getFileGap
// Return True If Operation Succeed

func createUser(name, pass string) bool {

	// User Musn't Exist

	if !userExist(name) {

		// Basic User Account Initialization

		day := fmt.Sprintf("%s", time.Now().Format("020106"))
		reg := []byte(name + " " + pass + " 1000000 " + day + " 1000000 0 1000000 StarterPack")

		// StarterPack
		// Free Money

		regData := nasabah{qidNasabah{name, pass, 1000000}, []transaksi{{day, 1000000, 0, 1000000, "StarterPack"}}}

		// Store It To Server Cache

		db = append(db, regData)

		// Then Make A Database Entry

		err := ioutil.WriteFile("assets/users/"+getFileGap(), reg, 0777)

		if err == nil {

			return true
		}
	}

	return false
}

// Removes A User From System
// Return True If Operation Succeed

func rmUser(n string) bool {

	// Gets All Files

	files, err := ioutil.ReadDir("assets/users/")

	if err == nil {

		for _, file := range files {

			// Checks For Each Files

			content, err := ioutil.ReadFile("assets/users/" + file.Name())

			if err == nil {

				res := fmt.Sprintf("%s", content)
				new := strings.Split(res, " ")

				// If Matches,
				// Delete Files

				if new[0] == n {

					err := os.Remove("assets/users/" + file.Name())

					if err == nil {

						// Read From Files Again
						// (Force Reinitialization)

						reIns()
						return true
					}
				}
			}
		}
	}

	return false
}

// Update User Account
// Return True If Operation Succeed

func editUser(old, new qidNasabah) bool {

	var dataTemp string

	// User Must Exist

	fName, exist := getFileName(old.Username)

	if exist && !userExist(new.Username) {

		content, err := ioutil.ReadFile("assets/users/" + fName)

		if err == nil {

			res := fmt.Sprintf("%s", content)
			text := strings.Split(res, " ")

			// Update The Account

			text[0] = new.Username
			text[1] = new.Password

			// Append User Transaction History
			// To The Updated Data

			for i, val := range text {

				if i != len(text)-1 {

					dataTemp = dataTemp + val + " "
				} else {

					dataTemp = dataTemp + val
				}
			}

			// Write Updated Account Data
			// To The User's Database Account

			write := ioutil.WriteFile("assets/users/"+fName, []byte(dataTemp), 0777)

			if write != nil {

				log.Fatal(err)

				data.Body.Content = []Content{Content{"Error: ", "Database Error"}}

				return false
			} else {

				fmt.Println(time.Now().Format("02/01/06  03:04:05.000000000"), "\t", old.Username, " Account Data Was Updated To ", new.Username)

				return true
			}
		} else {

			data.Body.Content = []Content{Content{"Error: ", "Couldn't Connect To Database"}}

			return false
		}
	}

	data.Body.Content = []Content{Content{"Error: ", "Data Is A Duplicate"}}

	return false
}

// Make A New Transaction
// Return True If Operation Succeed

func newTransaction(n, ref, sum string) bool {

	amount, _ := strconv.Atoi(sum)

	for _, val := range db {

		// Find The Right Account

		if val.Id.Username == n {

			// Extract Account Balance

			saldo := val.Riwayat[len(val.Riwayat)-1].Saldo

			// Depend On The Transaction Type

			if ref == "19NGC20" {

				saldo = saldo + amount
			} else {

				saldo = saldo - amount

				// Balance Must Be Sufficient

				if saldo < 50000 {

					data.Body.Content = []Content{Content{"Error: ", "Not Enough Balance Left"}}

					return false
				}
			}

			// Nake A Temporary Data
			// Append User Account Data

			regTemp := fmt.Sprintf("%s %s %s", val.Id.Username, val.Id.Password, strconv.Itoa(saldo))

			// Append All Account Transaction History

			for _, history := range val.Riwayat {

				regTemp = regTemp + fmt.Sprintf(" %s %s %s %s %s", history.Tanggal, strconv.Itoa(history.Debit), strconv.Itoa(history.Kredit), strconv.Itoa(history.Saldo), history.Ket)

			}

			// Find The Right File In Database

			fileName, err := getFileName(n)

			if !err {

				data.Body.Content = []Content{Content{"Error: ", "No Such User Exist"}}

				return false
			} else {

				// Append The Latest Transaction Details
				// And Write To The Database File

				if ref == "19NGC20" {

					regTemp = regTemp + fmt.Sprintf(" %s %s %s %s %s", time.Now().Format("020106"), sum, "0", strconv.Itoa(saldo), ref)

					err := ioutil.WriteFile("assets/users/"+fileName, []byte(regTemp), 0777)

					if err == nil {

						fmt.Println(time.Now().Format("02/01/06  03:04:05.000000000"), "\tA Balance Of", amount, "Has Been Received By", n)

						return true
					}
				} else {

					regTemp = regTemp + fmt.Sprintf(" %s %s %s %s %s", time.Now().Format("020106"), "0", sum, strconv.Itoa(saldo), ref)

					err := ioutil.WriteFile("assets/users/"+fileName, []byte(regTemp), 0777)

					if err == nil {

						fmt.Println(time.Now().Format("02/01/06  03:04:05.000000000"), "\tA Balance Of", amount, "Has Been Transfered From", n, "To", ref)

						if userExist(ref) {

							if newTransaction(ref, "19NGC20", sum) {

								return true
							} else {

								data.Body.Content = []Content{Content{"Error: ", "Transaction Failed"}}

								return false
							}
						} else {

							fmt.Println(time.Now().Format("02/01/06  03:04:05.000000000"), "\tA Balance Of", amount, "Has Been Moved From", n, "To", ref)

							return true
						}
					}
				}
			}
		}
	}

	data.Body.Content = []Content{Content{"Error: ", "Unknown Error"}}

	return false
}

// Checks User Session
// Reinitialize User Page For Each Request
// Redirect To Login,
// For Non-Existent User
// Or Unauthorized Access

func sessChk(w http.ResponseWriter, r *http.Request) {

	reIns()
	auth := r.URL.Query().Get("auth")

	if userExist(auth) {

		deMenu = []Menu{Menu{"Landing", "/index/?auth=" + auth}, Menu{"Dashboard", "/dashb/?auth=" + auth}, Menu{"About", "/about/?auth=" + auth}, Menu{"Sign Out", "/logout/"}}
	} else {

		fmt.Println(time.Now().Format("02/01/06  03:04:05.000000000"), "\tJust Casually Parsing Request")

		http.Redirect(w, r, "/login/", http.StatusMovedPermanently)
	}
}

// The Redirect Handler
// All Unauthorized Access,
// Shall Be Redirected To The Index Page

func rHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(time.Now().Format("02/01/06  03:04:05.000000000"), "\tApplication Running Fine...So Far")

	auth := r.URL.Query().Get("auth")

	if auth != "" {

		http.Redirect(w, r, "/index/?auth="+auth, http.StatusMovedPermanently)
	} else {

		http.Redirect(w, r, "/index/", http.StatusMovedPermanently)
	}
}

// The Login Handler

func loginHandler(w http.ResponseWriter, r *http.Request) {

	// Reinitialization + File Parsing

	reIns()

	tmpl := template.Must(template.ParseFiles("login.html"))

	data = Page{Head{"Sign In", deCSS}, Body{deHeder, deMenu, []Content{Content{"", ""}}, Foo}, fetchData("")}

	pass := r.FormValue("username")
	key := r.FormValue("password")

	// If User Is Signing In

	if r.FormValue("SignIn") == "Sign In" && r.FormValue("SignUp") == "" {

		if !strings.Contains(pass, " ") && !strings.Contains(key, " ") {

			if userChk(pass, key) {

				fmt.Println(time.Now().Format("02/01/06  03:04:05.000000000"), "\t", pass, "Just Signed In")
				http.Redirect(w, r, "/index/?auth="+pass, http.StatusMovedPermanently)

			} else {

				data.Body.Content = []Content{Content{"Error: ", "Username and/or Password Do Not Match"}}
			}
		} else {

			data.Body.Content = []Content{Content{"Error: ", "Username and/or Password Cannot Contain Whitespace"}}
		}

		// If User Is Signing Up

	} else if r.FormValue("SignUp") == "Sign Up" && r.FormValue("SignIn") == "" {

		if !strings.Contains(pass, " ") && !strings.Contains(key, " ") {

			if createUser(pass, key) {

				fmt.Println(time.Now().Format("02/01/06  03:04:05.000000000"), "\t", pass, "Just Signed Up")
				http.Redirect(w, r, "/index/?auth="+pass, http.StatusMovedPermanently)

			} else {

				data.Body.Content = []Content{Content{"Error: ", "User already exist"}}
			}
		} else {

			data.Body.Content = []Content{Content{"Error: ", "Username and/or Password Cannot Contain Whitespace"}}
		}
	}

	tmpl.Execute(w, data)
}

// Get User Data
// From Server-Cache

func fetchData(n string) nasabah {

	var res nasabah

	for _, val := range db {

		if val.Id.Username == n {

			return val
		}
	}

	return res
}

// All Field Search
// Return Matched Data

func searchData(data []transaksi, key string) []transaksi {

	var res []transaksi
	hasil, err := strconv.Atoi(key)

	for _, val := range data {

		if err == nil {

			if val.Debit == hasil || val.Kredit == hasil || val.Saldo == hasil {

				res = append(res, val)
			}
		} else {

			if val.Ket == key {

				res = append(res, val)
			}
		}

		// If If Else Fails

		if val.Tanggal == key {

			res = append(res, val)
		}
	}

	return res
}

// All Field Sort
// Return Sorted Data
// Insertion + Merge Sort

func sortData(temp nasabah, n, t string) nasabah {

	var ganti bool = false
	data := temp

	var cekI int

	for i := 0; i < len(data.Riwayat); i++ {

		// Insertion Sort - Optimized

		ganti = false

		if i > 0 {

			if n == "Date" {

				if t == "Asc" {

					if data.Riwayat[i].Tanggal < data.Riwayat[i-1].Tanggal {

						ganti = true
					}
				} else if t == "Desc" {

					if data.Riwayat[i].Tanggal > data.Riwayat[i-1].Tanggal {

						ganti = true
					}
				}
			} else if n == "Notes" {

				if t == "Asc" {

					if data.Riwayat[i].Ket < data.Riwayat[i-1].Ket {

						ganti = true
					}
				} else if t == "Desc" {

					if data.Riwayat[i].Ket > data.Riwayat[i-1].Ket {

						ganti = true
					}
				}
			}
		}

		if ganti {

			data.Riwayat[i], data.Riwayat[i-1] = data.Riwayat[i-1], data.Riwayat[i]
			i -= 2
		}

		// Selection Sort

		cekI = i

		if n == "Debit" {

			if t == "Asc" {

				cek := data.Riwayat[i].Debit

				for j := i + 1; j < len(data.Riwayat); j++ {

					if cek > data.Riwayat[j].Debit {

						cekI = j
						cek = data.Riwayat[j].Debit
					}
				}

			} else if t == "Desc" {

				cek := data.Riwayat[i].Debit

				for j := i + 1; j < len(data.Riwayat); j++ {

					if cek < data.Riwayat[j].Debit {

						cekI = j
						cek = data.Riwayat[j].Debit
					}
				}
			}
		} else if n == "Credit" {

			if t == "Asc" {

				cek := data.Riwayat[i].Kredit

				for j := i + 1; j < len(data.Riwayat); j++ {

					if cek > data.Riwayat[j].Kredit {

						cekI = j
						cek = data.Riwayat[j].Kredit
					}
				}

			} else if t == "Desc" {

				cek := data.Riwayat[i].Kredit

				for j := i + 1; j < len(data.Riwayat); j++ {

					if cek < data.Riwayat[j].Kredit {

						cekI = j
						cek = data.Riwayat[j].Kredit
					}
				}
			}
		} else if n == "Balance" {

			if t == "Asc" {

				cek := data.Riwayat[i].Saldo

				for j := i + 1; j < len(data.Riwayat); j++ {

					if cek > data.Riwayat[j].Saldo {

						cekI = j
						cek = data.Riwayat[j].Saldo
					}
				}

			} else if t == "Desc" {

				cek := data.Riwayat[i].Saldo

				for j := i + 1; j < len(data.Riwayat); j++ {

					if cek < data.Riwayat[j].Saldo {

						cekI = j
						cek = data.Riwayat[j].Saldo
					}
				}
			}
		}

		if i != cekI {

			data.Riwayat[i], data.Riwayat[cekI] = data.Riwayat[cekI], data.Riwayat[i]
		}

	}

	return data
}

// The Home / Index Page Landing

func idHandler(w http.ResponseWriter, r *http.Request) {

	sessChk(w, r)

	auth := r.URL.Query().Get("auth")
	tmpl := template.Must(template.ParseFiles("index.html"))

	data = Page{Head{"Landing", deCSS}, Body{deHeder, deMenu, []Content{Content{"", ""}}, Foo}, fetchData(auth)}

	tmpl.Execute(w, data)
}

// The Dashboard Handler

func dashHandler(w http.ResponseWriter, r *http.Request) {

	sessChk(w, r)

	auth := r.URL.Query().Get("auth")
	tmpl := template.Must(template.ParseFiles("dashboard.html"))

	data = Page{Head{"Dashboard", deCSS}, Body{deHeder, deMenu, []Content{Content{"", ""}}, Foo}, fetchData(auth)}

	// If User Decides To Delete Their Account

	if r.FormValue("Delete") == "Close Account" {

		if rmUser(auth) {

			fmt.Println(time.Now().Format("02/01/06  03:04:05.000000000"), "\t", auth, "Is No Longer Within Our Reach")

			http.Redirect(w, r, "/index/", http.StatusMovedPermanently)
		} else {

			fmt.Println(time.Now().Format("02/01/06  03:04:05.000000000"), "\t", auth, "Tried To Stop Being Our Subject")
		}
	}

	// If User Decides Fo Update Their Account

	if r.FormValue("Update") == "Update Account" {

		saldo, _ := strconv.Atoi(r.FormValue("saldo"))
		oldData := qidNasabah{r.FormValue("oldUsername"), r.FormValue("oldPassword"), saldo}
		newData := qidNasabah{r.FormValue("newUsername"), r.FormValue("newPassword"), saldo}

		if !strings.Contains(r.FormValue("newUsername"), " ") || !strings.Contains(r.FormValue("newPassword"), " ") {

			if editUser(oldData, newData) {

				http.Redirect(w, r, "/dashb/?auth="+r.FormValue("newUsername")+"#transaction", http.StatusMovedPermanently)

			} else {

				data.Body.Content = []Content{Content{"Error: ", "Data Is A Duplicate"}}
			}
		} else {

			data.Body.Content = []Content{Content{"Error: ", "Referal Code Cannot Contain Whitespace"}}
		}
	}

	// If User Makes A Transaction

	if r.FormValue("Transaction") == "GO!" {

		if !strings.Contains(r.FormValue("Code"), " ") {

			if newTransaction(auth, r.FormValue("Code"), r.FormValue("Jumlah")) {

				reIns()

				data = Page{Head{"Dashboard", deCSS}, Body{deHeder, deMenu, []Content{Content{"", ""}}, Foo}, fetchData(auth)}

				data.Body.Content = []Content{Content{"Great: ", "Transaction Success!"}}
			} else {

				data.Body.Content = []Content{Content{"Error: ", "Transaction Failed"}}
			}
		} else {

			data.Body.Content = []Content{Content{"Error: ", "Referal Code Cannot Contain Whitespace"}}
		}
	}

	// If User Want To Search Data

	if r.FormValue("Search") == "Search" {

		key := r.FormValue("Keywords")

		data.Data.Riwayat = searchData(data.Data.Riwayat, key)

		if len(data.Data.Riwayat) > 0 {

			data.Body.Content = []Content{Content{"Found: ", "Showing All Results Of " + fmt.Sprintf("%q", key)}}
		} else {

			reIns()

			data = Page{Head{"Dashboard", deCSS}, Body{deHeder, deMenu, []Content{Content{"", ""}}, Foo}, fetchData(auth)}

			data.Body.Content = []Content{Content{"Not Found: ", "No Such Data"}}
		}
	}

	// If User Want To Sort Data

	if r.FormValue("Date") != "" || r.FormValue("Debit") != "" || r.FormValue("Credit") != "" || r.FormValue("Balance") != "" || r.FormValue("Notes") != "" {

		var n, t string

		if r.FormValue("Date") != "" {

			n = "Date"
		} else if r.FormValue("Debit") != "" {

			n = "Debit"
		} else if r.FormValue("Credit") != "" {

			n = "Credit"
		} else if r.FormValue("Balance") != "" {

			n = "Balance"
		} else if r.FormValue("Notes") != "" {

			n = "Notes"
		}

		t = r.FormValue(n)

		data.Body.Content = []Content{Content{"Yeay: ", "Data Is Now Sorted"}}
		data.Data = sortData(data.Data, n, t)

	}

	tmpl.Execute(w, data)
}

// The About Page Handler

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	sessChk(w, r)

	auth := r.URL.Query().Get("auth")
	tmpl := template.Must(template.ParseFiles("about.html"))

	data = Page{Head{"About", deCSS}, Body{deHeder, deMenu, []Content{Content{"", ""}}, Foo}, fetchData(auth)}

	tmpl.Execute(w, data)
}

// The Sign Out Handler

func logoutHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/login/", http.StatusMovedPermanently)
}

// The Server Interface

func main() {

	// Launch Default Browser
	// On Windows

	launchBrowser := exec.Command("cmd", "cmd /c start http://localhost:8080/login/").Start()

	fmt.Printf("\nProject's official page:\nhttps://github.com/Dithmarschen/supreme-garbanzo\n\a")
	fmt.Printf("\nCurrent Status: Online!\n\a")

	if (launchBrowser) != nil {

		fmt.Printf("\nOpen one of these links in your browser:\nhttp://127.0.0.1:8080/index/ or http://localhost:8080/index/\n\n\a")
	}

	// Make Local Directory Accessible
	// In Server

	fs := http.FileServer(http.Dir("assets/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// The Handlers

	http.HandleFunc("/", rHandler)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/index/", idHandler)
	http.HandleFunc("/dashb/", dashHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/logout/", logoutHandler)

	// Execute On http://localhost

	log.Fatal(http.ListenAndServe(":8080", nil))
}
