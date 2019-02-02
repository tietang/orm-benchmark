package benchs

import (
    "database/sql"
    "fmt"
    "upper.io/db.v3"
    "upper.io/db.v3/lib/sqlbuilder"
    "upper.io/db.v3/mysql"
)

var ui sqlbuilder.Database

func init() {
    st := NewSuite("upper.io")
    st.InitF = func() {
        st.AddBenchmark("Insert", 2000*ORM_MULTI, UpperIOInsert)
        st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, UpperIOInsertMulti)
        st.AddBenchmark("Update", 2000*ORM_MULTI, UpperIOUpdate)
        st.AddBenchmark("Read", 4000*ORM_MULTI, UpperIORead)
        st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, UpperIOReadSlice)

        settings, _ := mysql.ParseURL(ORM_SOURCE)
        ui, _ = sqlbuilder.Open(mysql.Adapter, settings)

        driver := ui.Driver().(*sql.DB)
        driver.SetMaxIdleConns(ORM_MAX_IDLE)
        driver.SetMaxOpenConns(ORM_MAX_CONN)
    }
}

func UpperIOInsert(b *B) {
    var m *Model
    wrapExecute(b, func() {
        initDB()
        m = NewModel()
    })

    for i := 0; i < b.N; i++ {
        m.Id = 0
        collection := ui.Collection("model")
        if _, err := collection.Insert(m); err != nil {
            fmt.Println(err)
            b.FailNow()
        }
    }
}

func UpperIOInsertMulti(b *B) {
    panic(fmt.Errorf("Not support multi insert"))
}

func UpperIOUpdate(b *B) {
    var m *Model
    var r db.Result

    wrapExecute(b, func() {
        initDB()
        m = NewModel()
        c := ui.Collection("model")
        id, _ := c.Insert(m)
        r = c.Find("id =?", id)
    })

    for i := 0; i < b.N; i++ {

        if err := r.Update(m); err != nil {
            fmt.Println(err)
            b.FailNow()
        }
    }
}

func UpperIORead(b *B) {
    var m *Model
    var id interface{}

    wrapExecute(b, func() {
        initDB()
        m = NewModel()
        c := ui.Collection("model")
        id, _ = c.Insert(m)
    })

    for i := 0; i < b.N; i++ {
        collection := ui.Collection("model")
        if err := collection.Find("id=?", id).One(m); err != nil {
            fmt.Println(err)
            b.FailNow()
        }
    }
}

func UpperIOReadSlice(b *B) {
    var m *Model
    wrapExecute(b, func() {
        initDB()
        m = NewModel()

        c := ui.Collection("model")
        for i := 0; i < 100; i++ {
            m.Id = 0
            if _, err := c.Insert(m); err != nil {
                fmt.Println(err)
                b.FailNow()
            }
        }
    })

    for i := 0; i < b.N; i++ {
        var models []Model
        collection := ui.Collection("model")
        if err := collection.Find("id > ?", 0).Limit(100).All(&models); err != nil {
            fmt.Println(err)
            b.FailNow()
        }
    }
}
