// SPDX-License-Identifier: GPL-3.0-or-later

package iox_test

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/bassosimone/iox"
)

// ExampleCopyContext shows how to stream an HTTP response into a buffer using [iox].
func ExampleCopyContext() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "hello from httptest\n")
	}))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		log.Println("request error:", err)
		return
	}
	defer resp.Body.Close()

	buff := &bytes.Buffer{}
	writer := iox.NewLockedWriteCloser(iox.NopWriteCloser(buff))
	if _, err := iox.CopyContext(context.Background(), writer, resp.Body); err != nil {
		log.Println("copy error:", err)
		return
	}

	fmt.Println(buff.String())
	// Output:
	// hello from httptest
	//
}
