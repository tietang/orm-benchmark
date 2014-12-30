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

```

#### Sample 2

```

```

#### Sample 3
```

```


### Contact

Maintain by [slene](https://github.com/slene)