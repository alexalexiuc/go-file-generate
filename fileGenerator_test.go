package fileGenerator

import (
	"os"
	"testing"
)

const (
	Byte            int64 = 1
	Kilobyte        int64 = Byte * 1024
	Megabyte        int64 = Kilobyte * 1024
	Gigabit         int64 = Megabyte * 1024
	removeTestFiles       = true
)

func Test_randStringRunes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test random string generator of size 1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "Test random string generator of size 1000",
			args: args{n: 1000},
			want: 1000,
		},
		{
			name: "Test random string generator of size 10000",
			args: args{n: 10000},
			want: 10000,
		},
		{
			name: "Test random string generator of size 100000",
			args: args{n: 100000},
			want: 100000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randStringRunes(tt.args.n); len(got) != tt.want {
				t.Errorf("randStringRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateFile(t *testing.T) {
	type args struct {
		fullPath string
		fileSize int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test file of 0 byte size",
			args:    args{fullPath: "./test1.file", fileSize: 0},
			wantErr: false,
		},
		{
			name:    "Test file of 1 byte size",
			args:    args{fullPath: "./test2.file", fileSize: Byte * 1},
			wantErr: false,
		},
		{
			name:    "Test file of 1 Megabyte size",
			args:    args{fullPath: "./test3.file", fileSize: Megabyte * 1},
			wantErr: false,
		},
		// {
		// 	name:    "Test file of 1 Gigabyte size",
		// 	args:    args{fullPath: "./test4.file", fileSize: Gigabit * 1},
		// 	wantErr: false,
		// },
		// {
		// 	name:    "Test file of 10 Gigabyte size",
		// 	args:    args{fullPath: "./test4.file", fileSize: Gigabit * 10},
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateFile(tt.args.fullPath, tt.args.fileSize); (err != nil) != tt.wantErr {
				t.Errorf("GenerateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			f, err := os.Open(tt.args.fullPath)
			if err != nil {
				t.Errorf("GenerateFile() unable to read generated file, error = %v", err)
			}
			defer func() {
				f.Close()
				if removeTestFiles {
					os.Remove(tt.args.fullPath)
				}
			}()
			fStats, err := f.Stat()
			if err != nil {
				t.Errorf("GenerateFile() unable to read file stats, error = %v", err)
			}
			fileSize := fStats.Size()
			if tt.args.fileSize != fileSize {
				t.Errorf("GenerateFile() file size = %d, want %d", fileSize, tt.args.fileSize)
			}
		})
	}
}
