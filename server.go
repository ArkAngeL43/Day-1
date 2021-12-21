package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/bndr/gotabulate"
	_ "github.com/lib/pq"
)

const (
	v           = "1.0 BETA"
	DB_USER     = "user1"
	DB_PASSWORD = "fancy-bear-2021-fuckg"
	DB_NAME     = "testdb"
	host        = "localhost"
	port        = 5432
	user        = "user1"
	password    = "fancy-bear-2021-fuckg"
	dbname      = "testdb"
)

var (
	clear_hex = "\x1b[H\x1b[2J\x1b[3J"
	BLK       = "\033[0;30m"
	RED       = "\033[0;31m"
	GRN       = "\033[0;32m"
	YEL       = "\033[0;33m"
	BLU       = "\033[0;34m"
	MAG       = "\033[0;35m"
	CYN       = "\033[0;36m"
	WHT       = "\033[0;37m"
	BBLK      = "\033[1;30m"
	BRED      = "\033[1;31m"
	BGRN      = "\033[1;32m"
	BYEL      = "\033[1;33m"
	BBLU      = "\033[1;34m"
	BMAG      = "\033[1;35m"
	BCYN      = "\033[1;36m"
	BWHT      = "\033[1;37m"
	UBLK      = "\033[4;30m"
	URED      = "\033[4;31m"
	UGRN      = "\033[4;32m"
	UYEL      = "\033[4;33m"
	UBLU      = "\033[4;34m"
	UMAG      = "\033[4;35m"
	UCYN      = "\033[4;36m"
	UWHT      = "\033[4;37m"
	BLKB      = "\033[40m"
	REDB      = "\033[41m"
	GRNB      = "\033[42m"
	YELB      = "\033[43m"
	BLUB      = "\033[44m"
	MAGB      = "\033[45m"
	CYNB      = "\033[46m"
	WHTB      = "\033[47m"
	BLKHB     = "\033[0;100m"
	REDHB     = "\033[0;101m"
	GRNHB     = "\033[0;102m"
	YELHB     = "\033[0;103m"
	BLUHB     = "\033[0;104m"
	MAGHB     = "\033[0;105m"
	CYNHB     = "\033[0;106m"
	WHTHB     = "\033[0;107m"
	HBLK      = "\033[0;90m"
	HRED      = "\033[0;91m"
	HGRN      = "\033[0;92m"
	HYEL      = "\033[0;93m"
	HBLU      = "\033[0;94m"
	HMAG      = "\033[0;95m"
	HCYN      = "\033[0;96m"
	HWHT      = "\033[0;97m"
	BHBLK     = "\033[1;90m"
	BHRED     = "\033[1;91m"
	BHGRN     = "\033[1;92m"
	BHYEL     = "\033[1;93m"
	BHBLU     = "\033[1;94m"
	BHMAG     = "\033[1;95m"
	BHCYN     = "\033[1;96m"
	BHWHT     = "\033[1;97m"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func test_connection_psqql() bool {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(WHT, "\t\t[", REDHB, "INFO", WHT, "] \033[31mGET STAT ERROR  |=> ", err)
		os.Exit(1)
		return false

	}
	fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[31mAUTH\033[34m: PSQL PASS |=> CONNECTION PSQL TRUE")
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return true
}

func select_testdb_table_test() {
	flag.Parse()
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()
	fmt.Println("------ gathering information from db -----")
	rows, err := db.Query("SELECT * FROM Test2")
	checkErr(err)
	for rows.Next() {
		var uid int
		var NamePh string
		var NumberPh string
		err = rows.Scan(&uid, &NamePh, &NumberPh)
		checkErr(err)
		row_1 := []interface{}{NamePh, uid, NumberPh}
		t := gotabulate.Create([][]interface{}{row_1})
		t.SetHeaders([]string{"ID", "Name", "Number"})
		t.SetEmptyString("None")
		t.SetAlign("right")
		fmt.Println(t.Render("grid"))
		fmt.Println(uid, NamePh, NumberPh)
	}
}

func clear(clear_hex string) {
	fmt.Println(clear_hex)
}

func print() {
	fmt.Println("f")
}

func is_online(website string) bool {
	get, err := http.Get(website)
	if err != nil {
		fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET STAT ERROR  |=> ", get.StatusCode)
		return false
	} else {
		fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET STAT PASSED |=> ", get.StatusCode)
		return true
	}
}

func r_table_1_web(script string, w http.ResponseWriter, r *http.Request) {
	c := "Rscript"
	c1 := "main.r"
	c2 := script
	cf := exec.Command(c, c1, c2)
	stdout, err := cf.Output()
	checkErr(err)
	fmt.Print(w, string(stdout), BLU)
}

func logtxt(file string) {
	filepath := file
	pathmain, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	checkErr(err)
	defer pathmain.Close()
	c, err := fmt.Fprintln(pathmain, "[ INFO ] -> USER LOGGED IN AT => ", time.Now())
	if err != nil {
		fmt.Print(c)
		os.Exit(1)
	}
	fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mACCESS WRITTEN TO -> ", filepath)
}

func startp() {
	c := "sudo"
	c1 := "service"
	c2 := "postgresql"
	c3 := "start"
	cf := exec.Command(c, c1, c2, c3)
	stdout, err := cf.Output()
	checkErr(err)
	fmt.Print(string(stdout))
}

func process(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
		fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET STAT  |=> ", http.StateActive)
		fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET STAT  |=> ", http.StatusOK)
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		name := r.FormValue("name")
		password := r.FormValue("occupation")
		table := r.FormValue("table")
		command := r.FormValue("command")
		apassword := r.FormValue("Acting Password")
		fmt.Fprintln(w, name, password, table, command, apassword)
		if command == "lo-get" {
			resp, err := http.Get("http://localhost:8080")
			checkErr(err)
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET DATA: |=> ", resp.Header.Get("Accept-Ranges"))
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET DATA: |=> ", resp.Header.Get("Content-Length"))
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET DATA: |=> ", resp.Header.Get("Content-Type"))
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET DATA: |=> ", resp.Header.Get("Last-Modified"))
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mGET DATA: |=> ", resp.Header.Get("Date"))
			fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Accept-Ranges"))
			fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Content-Length"))
			fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Content-Type"))
			fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Last-Modified"))
			fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Date"))
			logtxt("log/logger.txt")
		}
		if apassword == "fancy-bear-2021-fuckg" {
			logtxt("log/logger.txt")
			fmt.Print("\n")
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mAUTH: PASSWORD    |-> ", apassword, " = TRUE ")
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mDATA: DB-NAME     |-> ", name)
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mDATA: DB-PASSWORD |-> ", password)
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mDATA: TABLE       |-> ", table)
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mDATA: COMMAND-RUN |-> ", command)

		} else {
			fmt.Println(WHT, "\t\t[", BRED, "INFO", WHT, "] \033[34mAUTH: PASSWORD -> ", apassword, " ENTERED CAME BACK NEGATIVE, FAILED PASSWORD AUTH ")
			os.Exit(1)
		}
		if command == "ver" {
			psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
			db, err := sql.Open("postgres", psqlInfo)
			checkErr(err)
			defer db.Close()
			var version string
			err = db.QueryRow("select version()").Scan(&version)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintln(w, "\033[32m[ + ] Current VERSION => ", version)
			logtxt("log/logger.txt")
			fmt.Print("\n")
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mAUTH: PASSWORD    |-> ", apassword, " = TRUE ")
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mDATA: DB-NAME     |-> ", name)
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mDATA: DB-PASSWORD |-> ", password)
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mDATA: TABLE       |-> ", table)
			fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mDATA: COMMAND-RUN |-> ", command)
		}
		if command == "chconn" {
			resp, err := http.Get("https://google.com")
			checkErr(err)
			if resp.StatusCode >= 100 {
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Accept-Ranges"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Content-Length"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Content-Type"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Last-Modified"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Date"))
				logtxt("log/logger.txt")
			} else {
				fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mDATA=>GET  code not parsed")
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Accept-Ranges"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Content-Length"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Content-Type"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Last-Modified"))
				fmt.Fprintln(w, "\t\t[  INFO  ] GET DATA: |=> ", resp.Header.Get("Date"))
				logtxt("log/logger.txt")
			}
		}
		if command == "help" {
			content, err := ioutil.ReadFile("log/help.txt")
			checkErr(err)
			fmt.Fprintln(w, string(content))
			logtxt("log/logger.txt")
			c := "Rscript"
			c1 := "main.r"
			c2 := "log/help.txt"
			cf := exec.Command(c, c1, c2)
			stdout, err := cf.Output()
			checkErr(err)
			fmt.Fprintln(w, string(stdout))
		}
		if command == "log" {
			content, err := ioutil.ReadFile("log/logger.txt")
			checkErr(err)
			fmt.Fprintln(w, string(content))
			logtxt("log/logger.txt")
			c := "Rscript"
			c1 := "main.r"
			c2 := "log/logger.txt"
			cf := exec.Command(c, c1, c2)
			stdout, err := cf.Output()
			checkErr(err)
			fmt.Fprintln(w, string(stdout))
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	port := ":8080"
	clear(clear_hex)
	http.HandleFunc("/", process)
	content, err := ioutil.ReadFile("log/banner.txt")
	checkErr(err)
	fmt.Println(WHT, string(content))
	is_online("https://www.google.com")
	startp()
	test_connection_psqql()
	fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mINITIATING SERVER")
	fmt.Println(WHT, "\t\t[", BLU, "INFO", WHT, "] \033[34mSERVER URL https://localhost", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
