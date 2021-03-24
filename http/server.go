package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Listen(address string) error {
	fmt.Println("Http Listen:", address)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"code": 200,
			"msg":  "success",
			"data": map[string]interface{}{
				"time": time.Now().Format("2006-01-02 15:04:05"),
				"ip":   r.RemoteAddr,
			},
		}

		str, _ := json.Marshal(&data)
		w.Header().Add("Content-Type", "application/json;charset=utf-8")
		_, _ = w.Write(str)
	})

	return http.ListenAndServe(address, nil)
}
