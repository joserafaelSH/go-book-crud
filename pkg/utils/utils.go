package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	fmt.Println("AAAAAAAAAAAAAAAAA")
	if err != nil {
		fmt.Println("BBBBBBBBBBBBBBBBB")
		return err
	}

	if err := json.Unmarshal(body, &x); err != nil {
		fmt.Println("CCCCCCCCCCCCCCCCC")
		return err
	}

	fmt.Println("x", x)

	return nil
}
