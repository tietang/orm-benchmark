package benchs

import (
    "fmt"
    "github.com/tietang/dbx"
    "strings"
    "time"
)

var database *dbx.Database

func init() {
    st := NewSuite("dbx")
    st.InitF = func() {
        st.AddBenchmark("Insert", 2000*ORM_MULTI, DbxIOInsert)
        st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, DbxIOInsertMulti)
        st.AddBenchmark("Update", 2000*ORM_MULTI, DbxIOUpdate)
        st.AddBenchmark("Read", 4000*ORM_MULTI, DbxIORead)
        st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, DbxIOReadSlice)

        settings := dbx.Settings{
            DriverName:      "mysql",
            User:            "root",
            Password:        "111111",
            Host:            "192.168.1.12:3306",
            Database:        "orm_bench",
            MaxOpenConns:    ORM_MAX_CONN,
            MaxIdleConns:    ORM_MAX_IDLE,
            ConnMaxLifetime: time.Minute * 30,
            LoggingEnabled:  true,
            Options: map[string]string{
                "charset":   "utf8",
                "parseTime": "true",
            },
        }
        var err error
        database, err = dbx.Open(settings)
        if err != nil {
            panic(err)
        }
        database.SetLogging(false)
        //database.RegisterTable(&Model{}, "model")
    }
}

func DbxIOInsert(b *B) {
    var m *Model
    wrapExecute(b, func() {
        initDB()
        m = NewModel()
    })

    for i := 0; i < b.N; i++ {
        m.Id = 0
        if _, err := database.Insert(m); err != nil {
            fmt.Println(err)
            b.FailNow()
        }
    }
}

func DbxIOInsertMulti(b *B) {
    var ms []*Model
    wrapExecute(b, func() {
        initDB()

        ms = make([]*Model, 0, 100)
        for i := 0; i < 100; i++ {
            ms = append(ms, NewModel())
        }
    })

    for i := 0; i < b.N; i++ {
        nFields := 7
        query := rawInsertBaseSQL + strings.Repeat(rawInsertValuesSQL+",", len(ms)-1) + rawInsertValuesSQL
        args := make([]interface{}, len(ms)*nFields)
        for j := range ms {
            offset := j * nFields
            args[offset+0] = ms[j].Name
            args[offset+1] = ms[j].Title
            args[offset+2] = ms[j].Fax
            args[offset+3] = ms[j].Web
            args[offset+4] = ms[j].Age
            args[offset+5] = ms[j].Right
            args[offset+6] = ms[j].Counter
        }
        _, _, err := database.Execute(query, args...)
        if err != nil {
            fmt.Println(err)
            b.FailNow()
        }
    }
}

func DbxIOUpdate(b *B) {
    var m *Model

    wrapExecute(b, func() {
        initDB()
        m = NewModel()
        rs, _ := database.Insert(m)
        id, _ := rs.LastInsertId()
        m.Id = int(id)

    })

    for i := 0; i < b.N; i++ {
        if _, err := database.Update(m); err != nil {
            fmt.Println(err)
            b.FailNow()
        }
    }
}

func DbxIORead(b *B) {
    var m *Model
    var id int64

    wrapExecute(b, func() {
        initDB()
        m = NewModel()
        rs, _ := database.Insert(m)
        id, _ = rs.LastInsertId()
    })

    q := &Model{Id: int(id)}
    for i := 0; i < b.N; i++ {
        if err := database.GetOne(q); err != nil {
            fmt.Println(err)
            b.FailNow()
        }
    }
}

func DbxIOReadSlice(b *B) {
    var m *Model
    wrapExecute(b, func() {
        initDB()
        m = NewModel()

        for i := 0; i < 100; i++ {
            m.Id = 0
            if _, err := database.Insert(m); err != nil {
                fmt.Println(err)
                b.FailNow()
            }
        }
    })

    for i := 0; i < b.N; i++ {
        var models []Model
        if err := database.Find(&models, "select * from model where id>? limit 100", 0); err != nil {
            fmt.Println(err)
            b.FailNow()
        }
    }
}
