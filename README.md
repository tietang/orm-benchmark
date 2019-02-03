# ORM Benchmark

A benchmark to compare the performance of golang orm package.

## Results (2014-12-30)

### Environment

* 2G RAM
* go version go1.3.3 linux/amd64
* [Go-MySQL-Driver Latest](https://github.com/go-sql-driver/mysql)
* MySQL 5.6.21-70.1-log Percona Server (GPL), Release 70.1, Revision 698

### ORMs

All package run in no-cache mode.

* [Beego ORM](http://beego.me/docs/mvc/model/overview.md) latest in branch [develop](https://github.com/astaxie/beego/tree/develop)
* [xorm](https://github.com/lunny/xorm) latest
* [gorm](https://github.com/jinzhu/gorm) latest
* [gorp](https://github.com/coopernurse/gorp) latest
* [modl](https://github.com/jmoiron/modl) latest
* [Hood](https://github.com/eaigner/hood) latest
* [Qbs](https://github.com/coocood/qbs) latest (Disabled stmt cache / [patch](https://gist.github.com/slene/8297019) / [full](https://gist.github.com/slene/8297565))
* [upper.io](https://upper.io/db) latest

### Run

```go
go get github.com/beego/orm-benchmark
orm-benchmark -multi=20 -orm=all
```

### Reports

#### Sample 1

```
hood
                   Insert:   2000    18.30s      9149968 ns/op   12090 B/op    199 allocs/op
      MultiInsert 100 row:    500     Not support multi insert
                   Update:   2000    19.83s      9914022 ns/op   12081 B/op    199 allocs/op
                     Read:   4000    35.64s      8910711 ns/op    4242 B/op     55 allocs/op
      MultiRead limit 100:   2000    32.98s     16491001 ns/op  232327 B/op   8765 allocs/op
raw
                   Insert:   2000     7.36s      3681234 ns/op     552 B/op     12 allocs/op
      MultiInsert 100 row:    500    17.59s     35184280 ns/op  110864 B/op    811 allocs/op
                   Update:   2000     7.09s      3543659 ns/op     616 B/op     14 allocs/op
                     Read:   4000    13.27s      3317756 ns/op    1432 B/op     37 allocs/op
      MultiRead limit 100:   2000    20.48s     10241508 ns/op   34704 B/op   1320 allocs/op
qbs
                   Insert:   2000     no primary key field
      MultiInsert 100 row:    500     Not support multi insert
                   Update:   2000     no primary key field
                     Read:   4000     no primary key field
      MultiRead limit 100:   2000     no primary key field
gorp
Error 1054: Unknown column 'id,pk' in 'field list'
                   Insert:   2000     0.00s      0.28 ns/op       0 B/op      0 allocs/op
      MultiInsert 100 row:    500     Not support multi insert
Error 1054: Unknown column 'id,pk' in 'where clause'
                   Update:   2000     0.00s      0.39 ns/op       0 B/op      0 allocs/op
Error 1054: Unknown column 'id,pk' in 'field list'
                     Read:   4000     0.00s      0.11 ns/op       0 B/op      0 allocs/op
Error 1054: Unknown column 'id,pk' in 'field list'
      MultiRead limit 100:   2000     0.00s      0.29 ns/op       0 B/op      0 allocs/op
upper.io
Error 1364: Field 'name' doesn't have a default value
                   Insert:   2000     0.00s      0.37 ns/op       0 B/op      0 allocs/op
      MultiInsert 100 row:    500     Not support multi insert
                   Update:   2000    13.92s      6960274 ns/op    5906 B/op    318 allocs/op
upper: no more rows in this result set
                     Read:   4000     0.00s      0.15 ns/op       0 B/op      0 allocs/op
Error 1364: Field 'name' doesn't have a default value
      MultiRead limit 100:   2000     0.00s      0.30 ns/op       0 B/op      0 allocs/op
dbx
                   Insert:   2000    13.93s      6966821 ns/op    1914 B/op     40 allocs/op
      MultiInsert 100 row:    500    15.27s     30534547 ns/op   69552 B/op    715 allocs/op
                   Update:   2000    12.32s      6158802 ns/op    2606 B/op     59 allocs/op
                     Read:   4000    27.26s      6816054 ns/op    2774 B/op     74 allocs/op
      MultiRead limit 100:   2000    26.33s     13166953 ns/op   78848 B/op   1737 allocs/op
orm
                   Insert:   2000    15.11s      7554026 ns/op    1937 B/op     40 allocs/op
      MultiInsert 100 row:    500    21.92s     43849013 ns/op  147170 B/op   1534 allocs/op
                   Update:   2000    14.95s      7475361 ns/op    1928 B/op     40 allocs/op
                     Read:   4000    30.66s      7664590 ns/op    2800 B/op     97 allocs/op
      MultiRead limit 100:   2000    27.51s     13753703 ns/op   85216 B/op   4287 allocs/op
xorm
                   Insert:   2000    13.89s      6943038 ns/op    2543 B/op     68 allocs/op
      MultiInsert 100 row:    500    17.43s     34853332 ns/op  233982 B/op   4751 allocs/op
                   Update:   2000    13.53s      6765132 ns/op    2800 B/op     96 allocs/op
                     Read:   4000    29.88s      7469699 ns/op    9307 B/op    243 allocs/op
      MultiRead limit 100:   2000    29.00s     14501894 ns/op  180009 B/op   8083 allocs/op
gorm
                   Insert:   2000    26.32s     13160785 ns/op    7336 B/op    149 allocs/op
      MultiInsert 100 row:    500     Not support multi insert
                   Update:   2000    42.39s     21196195 ns/op   19124 B/op    402 allocs/op
                     Read:   4000    28.19s      7047402 ns/op   11611 B/op    239 allocs/op
      MultiRead limit 100:   2000    41.23s     20615999 ns/op  250911 B/op   6225 allocs/op

Reports: 

  2000 times - Insert
       raw:     7.36s      3681234 ns/op     552 B/op     12 allocs/op
      xorm:    13.89s      6943038 ns/op    2543 B/op     68 allocs/op
       dbx:    13.93s      6966821 ns/op    1914 B/op     40 allocs/op
       orm:    15.11s      7554026 ns/op    1937 B/op     40 allocs/op
      hood:    18.30s      9149968 ns/op   12090 B/op    199 allocs/op
      gorm:    26.32s     13160785 ns/op    7336 B/op    149 allocs/op
      gorp:     0.00s      0.28 ns/op       0 B/op      0 allocs/op
  upper.io:     0.00s      0.37 ns/op       0 B/op      0 allocs/op
       qbs:     no primary key field

   500 times - MultiInsert 100 row
       dbx:    15.27s     30534547 ns/op   69552 B/op    715 allocs/op
      xorm:    17.43s     34853332 ns/op  233982 B/op   4751 allocs/op
       raw:    17.59s     35184280 ns/op  110864 B/op    811 allocs/op
       orm:    21.92s     43849013 ns/op  147170 B/op   1534 allocs/op
       qbs:     Not support multi insert
      gorp:     Not support multi insert
  upper.io:     Not support multi insert
      hood:     Not support multi insert
      gorm:     Not support multi insert

  2000 times - Update
       raw:     7.09s      3543659 ns/op     616 B/op     14 allocs/op
       dbx:    12.32s      6158802 ns/op    2606 B/op     59 allocs/op
      xorm:    13.53s      6765132 ns/op    2800 B/op     96 allocs/op
  upper.io:    13.92s      6960274 ns/op    5906 B/op    318 allocs/op
       orm:    14.95s      7475361 ns/op    1928 B/op     40 allocs/op
      hood:    19.83s      9914022 ns/op   12081 B/op    199 allocs/op
      gorm:    42.39s     21196195 ns/op   19124 B/op    402 allocs/op
      gorp:     0.00s      0.39 ns/op       0 B/op      0 allocs/op
       qbs:     no primary key field

  4000 times - Read
       raw:    13.27s      3317756 ns/op    1432 B/op     37 allocs/op
       dbx:    27.26s      6816054 ns/op    2774 B/op     74 allocs/op
      gorm:    28.19s      7047402 ns/op   11611 B/op    239 allocs/op
      xorm:    29.88s      7469699 ns/op    9307 B/op    243 allocs/op
       orm:    30.66s      7664590 ns/op    2800 B/op     97 allocs/op
      hood:    35.64s      8910711 ns/op    4242 B/op     55 allocs/op
      gorp:     0.00s      0.11 ns/op       0 B/op      0 allocs/op
  upper.io:     0.00s      0.15 ns/op       0 B/op      0 allocs/op
       qbs:     no primary key field

  2000 times - MultiRead limit 100
       raw:    20.48s     10241508 ns/op   34704 B/op   1320 allocs/op
       dbx:    26.33s     13166953 ns/op   78848 B/op   1737 allocs/op
       orm:    27.51s     13753703 ns/op   85216 B/op   4287 allocs/op
      xorm:    29.00s     14501894 ns/op  180009 B/op   8083 allocs/op
      hood:    32.98s     16491001 ns/op  232327 B/op   8765 allocs/op
      gorm:    41.23s     20615999 ns/op  250911 B/op   6225 allocs/op
      gorp:     0.00s      0.29 ns/op       0 B/op      0 allocs/op
  upper.io:     0.00s      0.30 ns/op       0 B/op      0 allocs/op
       qbs:     no primary key field

 

```


### Contact

Maintain by [slene](https://github.com/slene)