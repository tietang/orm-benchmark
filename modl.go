package main

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/modl"
)

var mo *modl.DbMap

func init() {
	st := NewSuite("modl")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, ModlInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, ModlInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, ModlUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, ModlRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, ModlReadSlice)

		dbDialect := modl.MySQLDialect{"InnoDB", "UTF8"}
		db, _ := sql.Open(dbDialect.DriverName(), ORM_SOURCE)
		mo = modl.NewDbMap(db, dbDialect)

		mo.Dbx.SetMaxIdleConns(ORM_MAX_IDLE)
		mo.Dbx.SetMaxOpenConns(ORM_MAX_CONN)

		mo.AddTable(Model{}).SetKeys(true, "Id")
	}
}

func ModlInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if err := mo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func ModlInsertMulti(b *B) {
	panic(fmt.Errorf("Not support multi insert"))
}

func ModlUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		mo.Insert(m)
	})

	for i := 0; i < b.N; i++ {
		if _, err := mo.Update(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func ModlRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		mo.Insert(m)
	})

	for i := 0; i < b.N; i++ {
		if err := mo.Get(m, m.Id); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func ModlReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if err := mo.Insert(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if err := mo.Select(&models, "select * from model where id > 0 limit 100"); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
