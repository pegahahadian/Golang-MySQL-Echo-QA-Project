package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type Datatype struct {
	SiteID          string `sql: "Site_ID" json:"SiteID" form:"SiteiD"`
	Scope           string `sql: "Scope" json:"Scope" form:"Scope"`
	Region          string `sql: "Region" json:"Region" form:Region`
	Province        string `sql: "Province" json:"Province" form:"Province"`
	Vendor          string `sql: "Vendor" json:"Vendor" form:"Vendor"`
	Supervisor      string `sql: "Supervisor" json:"Supervisor" form:"Supervisor"`
	SOAC_Date       string `sql: "SOAC_Date" json:"sOAC_Date" form:"SOAC_Date"`
	Last_Visit_Date string `sql: "Last_Visit_Date" json:"last_Visit_Date" form:"Last_Visit_Date"`
	Number_of_Visit string `sql: "Number_of_Visit" json:"Number_of_Visit" form:"Number_of_Visit"`
	COC_Date        string `sql: "COC_Date" json:"COC_Date" form:"COC_Date"`
	FAT_Date        string `sql: "FAT_Date" json:"FAT_Date" form:"FAT_Date"`
	FAC_Date        string `sql: "FAC_Date" json:"FAC_Date" form:"FAC_date"`
	Comment         string `sql: "Comment" json:"Comment" form:"Comment"`
}

func pegyDB(c echo.Context) error {

	datatype := Datatype{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&datatype)

	if err != nil {
		log.Printf("Failed!: %s", err)
		return c.String(http.StatusInternalServerError, "")

		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	// if err := c.Bind(datatype); err != nil {
	// 	return c.JSON(http.StatusBadRequest, "error!")
	// }
	log.Printf("QA DataBase: %s", datatype.Scope)
	return c.String(http.StatusOK, "Congrats on Response:)!")
}

func pegy(c echo.Context) error {
	db, err := sql.Open("mysql", "root:Abc@12345@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error while connecting to DB!")
	}
	defer db.Close()
	row, _ := db.Query("select Site_ID,Scope,region,province,vendor,supervisor,SOAC_Date,Last_Visit_Date,number_of_visit,COC_Date,FAT_Date,FAC_Date,Comment from test.db")
	var res []string
	for row.Next() {
		var db Datatype
		row.Scan(&db.SiteID, &db.Scope, &db.Region, &db.Province, &db.Vendor, &db.Supervisor, &db.SOAC_Date, &db.Last_Visit_Date, &db.Number_of_Visit, &db.COC_Date, &db.FAT_Date, &db.FAC_Date, &db.Comment)
		res = append(res, fmt.Sprintf("Site_Id: %s |Scope: %s|Region: %s |Province: %s|Vendor: %s |Supervisor: %s |SOAC_Date: %s|Last_Visit_date: %s|Number_of_Visit :%s|COC_Date: %s|FAT_Date: %s|FAC_Date: %s|Comment: %s", db.SiteID, db.Scope, db.Region, db.Province, db.Vendor, db.Supervisor, db.SOAC_Date, db.Last_Visit_Date, db.Number_of_Visit, db.COC_Date, db.FAT_Date, db.FAC_Date, db.Comment))
	}
	return c.JSON(http.StatusOK, res)
}

func main() {
	e := echo.New()
	e.GET("/", pegy)
	e.POST("/db", pegyDB)
	e.Logger.Fatal(e.Start(":7000"))
}
