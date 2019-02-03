package main

import (
    "github.com/tietang/dbx"
    "github.com/tietang/orm-benchmark/benchs"
    "time"
)

func main() {
    settings := dbx.Settings{
        DriverName:      "mysql",
        User:            "root",
        Password:        "111111",
        Host:            "192.168.1.12:3306",
        Database:        "orm_bench",
        MaxOpenConns:    10,
        MaxIdleConns:    1,
        ConnMaxLifetime: time.Minute * 30,
        LoggingEnabled:  true,
        Options: map[string]string{
            "charset":   "utf8",
            "parseTime": "true",
        },
    }
    var err error
    db, err := dbx.Open(settings)
    if err != nil {
        panic(err)
    }
    m := benchs.NewModel()
    rs, _ := db.Insert(m)
    id, _ := rs.LastInsertId()
    db.GetOne(&benchs.Model{Id: int(id)})
    m.Id = int(id)
    db.Update(m)
}
