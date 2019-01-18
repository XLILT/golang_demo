/*************************************************************************
* COPYRIGHT NOTICE
*  Copyright (c) 2018, Wuhan Youxuan Stream Education Co., Ltd.
*  All rights reserved.
*
*  @version : 1.0
*  @author : mxl
*  @E-mail : xiaolongicx@gmail.com
*  @date : 2019-01-18 17:43
*
*  Revision Notes :
*/

package main;

import (
    "fmt"
    "net/http"
)

func get1() {
    resp, err := http.Get("https://www.baidu.com")

    if err != nil {
        fmt.Println(err.Error())

        return
    }

    defer resp.Body.Close()

    headers := resp.Header

    for k, v := range headers {
        fmt.Println(k, "=", v)
    }

    fmt.Println("resp status %s,statusCode %d\n", resp.Status, resp.StatusCode)

    fmt.Printf("resp Proto %s\n", resp.Proto)

    fmt.Printf("resp content length %d\n", resp.ContentLength)

    fmt.Printf("resp transfer encoding %v\n", resp.TransferEncoding)

    fmt.Printf("resp Uncompressed %t\n", resp.Uncompressed)
}

func main() {
    get1()

    return
}

