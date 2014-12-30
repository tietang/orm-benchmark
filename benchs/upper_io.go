package benchs

import (
	"database/sql"
	"fmt"

	"upper.io/db"
	"upper.io/db/mysql"
)

var ui db.Database

func init() {
	st := NewSuite("upper.io")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, UpperIOInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, UpperIOInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, UpperIOUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, UpperIORead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, UpperIOReadSlice)

		settings, _ := mysql.ParseURL(ORM_SOURCE)
		ui, _ = db.Open(mysql.Adapter, settings)

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
		collection, _ := ui.Collection("model")
		if _, err := collection.Append(m); err != nil {
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
		c, _ := ui.Collection("model")
		id, _ := c.Append(m)
		r = c.Find(db.Cond{"id =": id})
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
		c, _ := ui.Collection("model")
		id, _ = c.Append(m)
	})

	for i := 0; i < b.N; i++ {
		collection, _ := ui.Collection("model")
		if err := collection.Find(db.Cond{"id": id}).One(m); err != nil {
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

		c, _ := ui.Collection("model")
		for i := 0; i < 100; i++ {
			m.Id = 0
			if _, err := c.Append(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []Model
		collection, _ := ui.Collection("model")
		if err := collection.Find(db.Cond{"id > ": 0}).Limit(100).All(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
