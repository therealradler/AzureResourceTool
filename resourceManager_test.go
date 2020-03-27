package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
)

func TestGetGroups(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    resources.Group
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGroups(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
