package deferbug

import (
	"context"
	"errors"
	"fmt"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func crashCrash2(ctx context.Context, fnct func()) (rError error) {
	defer func() {
	}()
	defer func() {
		if r := recover(); r != nil {
			rError = errors.New(fmt.Sprint(r))
			return
		}
	}()
	fnct()
	return nil
}

func crashCrash() {
	crashCrash2(context.Background(), func() {
		defer func() {
			if r := recover(); r != nil {
				panic(r)
			}
		}()
		panic("hi")
	})
}

func TestDeferBug(t *testing.T) {
	var wg sync.WaitGroup
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		crashCrash()
		//c.JSON(200, gin.H{})
	})

	defer func() {}()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				func() {
					defer func() {
						recover()
					}()
					panic("hi")
				}()

				w := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "/", nil)
				server.ServeHTTP(w, req)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	wg.Wait()
}
