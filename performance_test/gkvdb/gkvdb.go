package main

import (
    "fmt"
    "strconv"
    "gitee.com/johng/gkvdb/gkvdb"
    "gitee.com/johng/gf/g/util/gtime"
)

var db *gkvdb.DB

func init() {
    t := gtime.Microsecond()
    db, _ = gkvdb.New("/tmp/gkvdb", "test")
    fmt.Println("db init:", gtime.Microsecond() - t)
}

// 测试数据库写入
func TestSet(count int) {
    t := gtime.Microsecond()
    for i := 0; i < count; i++ {
        key   := []byte("k_" + strconv.Itoa(i))
        value := []byte("v_" + strconv.Itoa(i))
        if err := db.Set(key, value); err != nil {
            fmt.Println(err)
        }
    }
    fmt.Println("TestSet:", gtime.Microsecond() - t)
}

// 测试数据库查询
func TestGet(count int) {
    t := gtime.Microsecond()
    for i := 0; i < count; i++ {
        key := []byte("k_" + strconv.Itoa(i))
        if r := db.Get(key); r == nil {
            fmt.Println("TestGet value not found for index:", i)
        }
    }
    fmt.Println("TestGet:", gtime.Microsecond() - t)
}

// 测试数据库删除
func TestRemove(count int) {
    t := gtime.Microsecond()
    for i := 0; i < count; i++ {
        key := []byte("k_" + strconv.Itoa(i))
        if err := db.Remove(key); err != nil {
            fmt.Println(err)
        }
    }
    fmt.Println("TestRemove:", gtime.Microsecond() - t)
}


func main() {
    count := 1000000
    //TestSet(count)
    TestGet(count)
    TestRemove(count)
    select {

    }
}