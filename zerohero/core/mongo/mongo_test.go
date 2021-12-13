package main

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
)

var pathFile = "../../json"

// go test -v -run ^TestZeroHeroHandlers$
// go test -v -run ^TestZeroHeroHandlers$ --count=10

// go test -coverprofile coverage.out
// go tool cover -html=coverage.out

// go test -bench . -benchmem

// openFileBuffer
// openFileBuffer(file string) *bytes.Buffer , error
func openFileBuffer(file string) (*bytes.Buffer, error) {
	dat, err := os.ReadFile(pathFile + "/" + file + ".json")
	if err != nil {
		return bytes.NewBuffer([]byte("")), err
	}
	return bytes.NewBuffer(dat), nil
}

var c Config

func init() {
	c = Config{
		Srv:     "mongodb",
		DB:      "zerohero",
		Host:    "localhost:27017",
		User:    "root",
		Pass:    "senha123",
		Options: "authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false",
	}
}

// TestZeroHeroHandlers
func TestZeroHeroHandlers(t *testing.T) {
	type args struct {
		method string
		ctype  string
		Header map[string]string
		url    string
		body   func() *bytes.Buffer
	}
	tests := []struct {
		name string
		args args
		want int //status code
		show bool
	}{
		// TODO: Add test cases.
		{
			name: "test_service_post",
			args: args{
				method: "POST",
				ctype:  "application/json",
				Header: nil,
				url:    "/api",
				body: func() (dat *bytes.Buffer) {
					dat, err := openFileBuffer("batgirl")
					if err != nil {
						t.Errorf("Error ReadFile batgirl:%s", err)
						return
					}
					return
				},
			},
			want: 201,
			show: true,
		},
		{
			name: "test_service_post",
			args: args{
				method: "POST",
				ctype:  "application/json",
				Header: nil,
				url:    "/api",
				body: func() (dat *bytes.Buffer) {
					dat, err := openFileBuffer("hulk")
					if err != nil {
						t.Errorf("Error ReadFile hulk:%s", err)
						return
					}
					return
				},
			},
			want: 201,
			show: true,
		},
		{
			name: "test_service_post",
			args: args{
				method: "POST",
				ctype:  "application/json",
				Header: nil,
				url:    "/apix",
				body: func() (dat *bytes.Buffer) {
					dat, err := openFileBuffer("batgirl")
					if err != nil {
						t.Errorf("Error ReadFile batgirl:%s", err)
						return
					}
					return
				},
			},
			want: 404,
			show: true,
		},
		{
			name: "test_service_post",
			args: args{
				method: "POST",
				ctype:  "application/json",
				Header: nil,
				url:    "/api/noname",
				body: func() *bytes.Buffer {
					return bytes.NewBuffer([]byte(""))
				},
			},
			want: 404,
			show: true,
		},
		{
			name: "test_service_get",
			args: args{
				method: "GET",
				ctype:  "application/json",
				Header: nil,
				url:    "/api/batgirl",
				body: func() *bytes.Buffer {
					return bytes.NewBuffer([]byte(""))
				},
			},
			want: 200,
			show: false,
		},
		{
			name: "test_service_get",
			args: args{
				method: "GET",
				ctype:  "application/json",
				Header: nil,
				url:    "/api/hulk",
				body: func() *bytes.Buffer {
					return bytes.NewBuffer([]byte(""))
				},
			},
			want: 200,
			show: false,
		},
		{
			name: "test_service_get",
			args: args{
				method: "GET",
				ctype:  "application/json",
				Header: nil,
				url:    "/api/hulk/work",
				body: func() *bytes.Buffer {
					return bytes.NewBuffer([]byte(""))
				},
			},
			want: 200,
			show: true,
		},
		{
			name: "test_service_get",
			args: args{
				method: "GET",
				ctype:  "application/json",
				Header: nil,
				url:    "/api/batgirl/biography",
				body: func() *bytes.Buffer {
					return bytes.NewBuffer([]byte(""))
				},
			},
			want: 200,
			show: true,
		},
		{
			name: "test_service_get",
			args: args{
				method: "GET",
				ctype:  "application/json",
				Header: nil,
				url:    "/api/batgirl/image",
				body: func() *bytes.Buffer {
					return bytes.NewBuffer([]byte(""))
				},
			},
			want: 200,
			show: true,
		},
		{
			name: "test_service_get",
			args: args{
				method: "GET",
				ctype:  "application/json",
				Header: nil,
				url:    "/api/batgirl/appearance",
				body: func() *bytes.Buffer {
					return bytes.NewBuffer([]byte(""))
				},
			},
			want: 200,
			show: true,
		},
		{
			name: "test_service_get",
			args: args{
				method: "GET",
				ctype:  "application/json",
				Header: nil,
				url:    "/api/batgirl/connections",
				body: func() *bytes.Buffer {
					return bytes.NewBuffer([]byte(""))
				},
			},
			want: 200,
			show: true,
		},
		{
			name: "test_service_get",
			args: args{
				method: "GET",
				ctype:  "application/json",
				Header: nil,
				url:    "/api/batgirl/speed",
				body: func() *bytes.Buffer {
					return bytes.NewBuffer([]byte(""))
				},
			},
			want: 400,
			show: false,
		},
		{
			name: "test_service_get",
			args: args{
				method: "GET",
				ctype:  "application/json",
				Header: nil,
				url:    "/api/batgirl/power",
				body: func() *bytes.Buffer {
					return bytes.NewBuffer([]byte(""))
				},
			},
			want: 400,
			show: false,
		},
		{
			name: "test_service_get",
			args: args{
				method: "GET",
				ctype:  "application/json",
				Header: nil,
				url:    "/api/batgirl/powerstats",
				body: func() *bytes.Buffer {
					return bytes.NewBuffer([]byte(""))
				},
			},
			want: 200,
			show: false,
		},
		{

			name: "test_service_delete",
			args: args{
				method: "DELETE",
				ctype:  "application/json",
				Header: nil,
				url:    "/api/batgirl",
				body: func() *bytes.Buffer {
					return bytes.NewBuffer([]byte(""))
				},
			},
			want: 204,
			show: false,
		},
		{

			name: "test_service_delete",
			args: args{
				method: "DELETE",
				ctype:  "application/json",
				Header: nil,
				url:    "/api/hulk",
				body: func() *bytes.Buffer {
					return bytes.NewBuffer([]byte(""))
				},
			},
			want: 204,
			show: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()
			w := httptest.NewRecorder()
			req := httptest.NewRequest(
				tt.args.method,
				tt.args.url,
				tt.args.body())

			req.Header.Add("Content-type", tt.args.ctype)
			for key, val := range tt.args.Header {
				req.Header.Add(key, val)
			}

			fmt.Fprint(w, "")
			// Service(w, req)
			// resp := w.Result()
			// defer resp.Body.Close()
			// assert.Equal(t, tt.want, resp.StatusCode)
			// t.Logf("\033[1;42mstatus=>[%d]\033[0m", resp.StatusCode)
			// if tt.show {
			// 	body, _ := io.ReadAll(resp.Body)
			// 	t.Logf("\nbody=>%s", string(body))
			// }
		})
	}
}
