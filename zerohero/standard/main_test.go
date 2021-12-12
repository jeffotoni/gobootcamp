package main

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestLogger(t *testing.T) {
	tests := []struct {
		name string
		want Middleware
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Logger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Logger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUse(t *testing.T) {
	type args struct {
		f           http.HandlerFunc
		middlewares []Middleware
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Use(tt.args.f, tt.args.middlewares...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Use() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Service(tt.args.w, tt.args.r)
		})
	}
}

func TestPost(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Post(tt.args.w, tt.args.r)
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Get(tt.args.w, tt.args.r)
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Delete(tt.args.w, tt.args.r)
		})
	}
}

func TestPut(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Put(tt.args.w, tt.args.r)
		})
	}
}

func TestZeroHero_InsertOne(t *testing.T) {
	type fields struct {
		Response    string
		ID          string
		UUID        string
		Name        string
		Powerstats  Powerstats
		Biography   Biography
		Appearance  Appearance
		Work        Work
		Connections Connections
		Image       Image
	}
	type args struct {
		collname string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zh := ZeroHero{
				Response:    tt.fields.Response,
				ID:          tt.fields.ID,
				UUID:        tt.fields.UUID,
				Name:        tt.fields.Name,
				Powerstats:  tt.fields.Powerstats,
				Biography:   tt.fields.Biography,
				Appearance:  tt.fields.Appearance,
				Work:        tt.fields.Work,
				Connections: tt.fields.Connections,
				Image:       tt.fields.Image,
			}
			if err := zh.InsertOne(tt.args.collname); (err != nil) != tt.wantErr {
				t.Errorf("ZeroHero.InsertOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFindOne(t *testing.T) {
	type args struct {
		name     string
		fatia    string
		collname string
	}
	tests := []struct {
		name    string
		args    args
		wantMzh interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMzh, err := FindOne(tt.args.name, tt.args.fatia, tt.args.collname)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMzh, tt.wantMzh) {
				t.Errorf("FindOne() = %v, want %v", gotMzh, tt.wantMzh)
			}
		})
	}
}

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteOne(tt.args.name, tt.args.collname); (err != nil) != tt.wantErr {
				t.Errorf("DeleteOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestZeroHero_UpdateOne(t *testing.T) {
	type fields struct {
		Response    string
		ID          string
		UUID        string
		Name        string
		Powerstats  Powerstats
		Biography   Biography
		Appearance  Appearance
		Work        Work
		Connections Connections
		Image       Image
	}
	type args struct {
		name     string
		collname string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zh := ZeroHero{
				Response:    tt.fields.Response,
				ID:          tt.fields.ID,
				UUID:        tt.fields.UUID,
				Name:        tt.fields.Name,
				Powerstats:  tt.fields.Powerstats,
				Biography:   tt.fields.Biography,
				Appearance:  tt.fields.Appearance,
				Work:        tt.fields.Work,
				Connections: tt.fields.Connections,
				Image:       tt.fields.Image,
			}
			if err := zh.UpdateOne(tt.args.name, tt.args.collname); (err != nil) != tt.wantErr {
				t.Errorf("ZeroHero.UpdateOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
