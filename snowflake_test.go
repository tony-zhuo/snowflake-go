package snowflake

import (
	"fmt"
	"testing"
)

func TestNewSnowflake(t *testing.T) {
	type args struct {
		nodeID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				nodeID: 1,
			},
			wantErr: false,
		},
		{
			name: "over max node id",
			args: args{
				nodeID: maxNodeID + 1,
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewSnowflake(tt.args.nodeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSnowflake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSnowflake_NextID(t *testing.T) {
	type fields struct {
		nodeID int64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "",
			fields: fields{
				nodeID: 1,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sf, _ := NewSnowflake(tt.fields.nodeID)
			fmt.Printf("snowflake id: %d", sf.NextID())
		})
	}
}
