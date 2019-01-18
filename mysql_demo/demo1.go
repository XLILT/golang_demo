package main

import (
    "database/sql"
    "time"
    "fmt"
    "math/rand"
    _ "github.com/go-sql-driver/mysql"
)

type DbWorker struct {
    // mysql data source name
    Dsn string
}

func query1(db * sql.DB) {
    stmt, err := db.Prepare(`SELECT Id,Ip From gate_online_video`)

    if err != nil {
        fmt.Printf(err.Error())

        return
    }

    defer stmt.Close()

    rows, err := stmt.Query()

    if err != nil {
        fmt.Printf(err.Error())

        return
    }

    defer rows.Close()

    for rows.Next() {
        var Id int
        var Ip sql.NullString

        rows.Scan(&Id, &Ip)

        if err != nil {
            fmt.Printf(err.Error())

            continue
        }

        fmt.Println("get data, Id: ", Id, " IP: ", Ip.String)
    }
}

func insert1(db * sql.DB) {
    stmt, err := db.Prepare("INSERT INTO gate_online_video(IP) VALUES(?)")

    if err != nil {
        fmt.Printf(err.Error())

        return
    }

    defer stmt.Close()

    rand.Seed(time.Now().Unix())
    res, err := stmt.Exec(fmt.Sprintf("%d.%d.%d.%d", (rand.Intn(254) + 1), (rand.Intn(254) + 1), (rand.Intn(254) + 1), (rand.Intn(254) + 1)))

    if err != nil {
        fmt.Printf(err.Error())

        return
    }

    lastInsertId, err := res.LastInsertId()

    if err != nil {
        fmt.Printf(err.Error())

        return
    }

    affectedRow, err := res.RowsAffected()

    if err != nil {
        fmt.Printf(err.Error())

        return
    }

    fmt.Println("lastInsertId: ", lastInsertId, "affectedRow: ", affectedRow);
}

func main() {
    dbw := DbWorker{
        // Dsn: "user:password@tcp(127.0.0.1:3306)/test",
        Dsn: "tars:tars2015@tcp(172.16.116.50:3307)/slonline",
    }

    db, err := sql.Open("mysql", dbw.Dsn)

    if err != nil {
        panic(err)

        return
    }

    defer db.Close()

    insert1(db);
    query1(db);
}

