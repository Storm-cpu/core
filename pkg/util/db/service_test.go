package dbutil

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNew(t *testing.T) {
	type args struct {
		dbPsn string
		cfg   *gorm.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.DB
		wantErr bool
	}{
		{
			name: "Valid connection string",
			args: args{
				dbPsn: "postgresql://localhost:5432/postgres?user=postgres&password=moneyflow123",
				cfg:   &gorm.Config{},
			},
			wantErr: false,
		},
		{
			name: "Invalid connection string",
			args: args{
				dbPsn: "invalid-connection-string",
				cfg:   &gorm.Config{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.dbPsn, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(&gorm.DB{})) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
