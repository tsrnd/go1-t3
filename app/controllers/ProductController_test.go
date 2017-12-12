package controller

import (
	"net/http"
	"testing"
)

func TestProductController_Show(t *testing.T) {
	type fields struct {
		Render render
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &ProductController{
				Render: tt.fields.Render,
			}
			this.Show(tt.args.w, tt.args.r)
		})
	}
}
