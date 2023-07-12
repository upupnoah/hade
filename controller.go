package main

import (
	"context"
	"fmt"
	"github.com/upupnoah/hade/framework"
	"log"
	"time"
)

//func FooControllerHandler(ctx *framework.Context) error {
//	return ctx.Json(200, map[string]any{
//		"code": 0,
//	})
//}

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan any, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), 1*time.Second)
	defer cancel()

	// mu := sync.Mutex{}
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// Do real action
		time.Sleep(10 * time.Second)
		err := c.Json(200, "ok")
		if err != nil {
			return
		}
		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		err := c.Json(500, "panic")
		if err != nil {
			return err
		}
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		err := c.Json(500, "time out")
		if err != nil {
			return err
		}
		c.SetHasTimeout()
	}
	return nil
}

// 未封装自定义 Context 的控制器代码
//func Foo(request *http.Request, response http.ResponseWriter) {
//	obj := map[string]any{
//		"errno":  50001,
//		"errmsg": "inner error",
//		"data":   nil,
//	}
//
//	response.Header().Set("Content-Type", "application/json")
//
//	foo := request.PostFormValue("foo")
//	if foo == "" {
//		foo = "10"
//	}
//	fooInt, err := strconv.Atoi(foo)
//	if err != nil {
//		response.WriteHeader(500)
//		return
//	}
//	obj["data"] = fooInt
//	byt, err := json.Marshal(obj)
//	if err != nil {
//		response.WriteHeader(500)
//		return
//	}
//
//	_, err = response.Write(byt)
//	if err != nil {
//		response.WriteHeader(500)
//		return
//	}
//	response.WriteHeader(200)
//}

// 使用自定义 Context 的控制器代码

//func Foo2(ctx *framework.Context) error {
//	obj := map[string]any{
//		"errno": 50001,
//		"errmsg": "inner error",
//		"data": nil,
//	}
//	fooInt := ctx.FormInt("foo", 10)
//	obj["data"] = fooInt
//	return ctx.Json(http.StatusOK, obj)
//}

//func Foo3(ctx *framework.Context) error {
//	rdb := redis.NewClient(&redis.Options{
//		Addr: "localhost:6379",
//		Password: "", // no password set
//		DB: 0, // use default DB
//	})
//	return rdb.Set(ctx, "key", "value", 0).Err()
//}
