package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

// go test -v -run ^TestZeroHeroHandlers$
// go test -v -run ^TestZeroHeroHandlers$ --count=10

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
				body: func() (buff *bytes.Buffer) {
					dat, err := os.ReadFile("../json/batgirl.json")
					if err != nil {
						t.Errorf("Error ReadFile batgirl:%s", err)
						return nil
					}
					return bytes.NewBuffer(dat)
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
				body: func() (buff *bytes.Buffer) {
					dat, err := os.ReadFile("../json/hulk.json")
					if err != nil {
						t.Errorf("Error ReadFile hulk:%s", err)
						return nil
					}
					return bytes.NewBuffer(dat)
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
				body: func() *bytes.Buffer {
					dat, err := os.ReadFile("../json/batgirl.json")
					if err != nil {
						t.Errorf("Error ReadFile:%s", err)
						return nil
					}
					return bytes.NewBuffer(dat)
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

			Service(w, req)

			resp := w.Result()
			defer resp.Body.Close()
			assert.Equal(t, tt.want, resp.StatusCode)
			t.Logf("\033[1;42mstatus=>[%d]\033[0m", resp.StatusCode)
			if tt.show {
				body, _ := io.ReadAll(resp.Body)
				t.Logf("\nbody=>%s", string(body))
			}
		})
	}
}

// TestInsertOne add mongodb
func TestInsertOne(t *testing.T) {
	type args struct {
		collname string
		zerohero func() ZeroHero
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_insertOne",
			args: args{
				collname: CollHeros,
				zerohero: func() (zh ZeroHero) {
					dat, err := os.ReadFile("../json/hulk.json")
					if err != nil {
						t.Errorf("Error ReadFile hulk:%s", err)
						return
					}
					err = json.Unmarshal(dat, &zh)
					if err != nil {
						t.Errorf("Error Unmarshal hulk:%s", err)
						return
					}
					return
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zh := tt.args.zerohero()
			if err := zh.InsertOne(tt.args.collname); (err != nil) != tt.wantErr {
				t.Errorf("ZeroHero.InsertOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestFindOne findOne mongodb
func TestFindOne(t *testing.T) {
	var work Work
	type args struct {
		name     string
		fatia    string
		collname string
	}
	tests := []struct {
		name    string
		args    args
		wantMzh Work
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_findOne",
			args: args{
				name:     "hulk",
				fatia:    "work",
				collname: CollHeros,
			},
			wantMzh: work,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMzh, err := FindOne(tt.args.name, tt.args.fatia, tt.args.collname)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotMzh, tt.wantMzh) {
			t.Logf("FindOne() = %v, want %v", gotMzh, tt.wantMzh)
			//}
		})
	}
}

// TestDeleteOne remove mongodb
func TestDeleteOne(t *testing.T) {
	type args struct {
		name     string
		collname string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_deleteOne",
			args: args{
				name:     "hulk",
				collname: CollHeros,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteOne(tt.args.name, tt.args.collname); (err != nil) != tt.wantErr {
				t.Errorf("DeleteOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateOne(t *testing.T) {
	type args struct {
		name     string
		collname string
		zerohero func() ZeroHero
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_updateOne",
			args: args{
				name:     "hulk",
				collname: CollHeros,
				zerohero: func() (zh ZeroHero) {
					dat, err := os.ReadFile("../json/hulk.json")
					if err != nil {
						t.Errorf("Error ReadFile hulk:%s", err)
						return
					}
					err = json.Unmarshal(dat, &zh)
					if err != nil {
						t.Errorf("Error Unmarshal hulk:%s", err)
						return
					}
					return
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zh := tt.args.zerohero()
			if err := zh.UpdateOne(tt.args.name, tt.args.collname); (err != nil) != tt.wantErr {
				t.Errorf("ZeroHero.UpdateOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
