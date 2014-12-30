package benchs

import (
	"database/sql"
	"fmt"

	"github.com/coopernurse/gorp"
)

var gr *gorp.DbMap

func init() {
	st := NewSuite("gorp")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, GorpInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, GorpInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, GorpUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, GorpRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, GorpReadSlice)

		dbDialect := gorp.MySQLDialect{"InnoDB", "UTF8"}
		db, _ := sql.Open("mysql", ORM_SOURCE)
		gr = &gorp.DbMap{Db: db, Dialect: dbDialect}
		gr.AddTableWithName(Model{}, "model").SetKeys(true, "Id")
	}
}

func GorpInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if err := gr.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GorpInsertMulti(b *B) {
	panic(fmt.Errorf("Not support multi insert"))
}

func GorpUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		gr.Insert(m)
	})

	for i := 0; i < b.N; i++ {
		if _, err := gr.Update(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GorpRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		gr.Insert(m)
	})

	for i := 0; i < b.N; i++ {
		if _, err := gr.Get(Model{}, m.Id); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GorpReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if err := gr.Insert(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if _, err := gr.Select(&models, "select * from model where id > 0 limit 100"); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
