// Package counters handles HTTP requests regarding counters.
package counters

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Counters
	}{
		{name: "new", want: &Counters{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCounter(t *testing.T) {
	type args struct {
		cp           *Counters
		i            int
		shouldChange bool
	}
	tests := []struct {
		name   string
		args   args
		want   Counter
		wantCp *Counters
	}{
		{
			name:   "empty counters slice",
			args:   args{cp: &Counters{}, i: 0, shouldChange: true},
			want:   Counter{Index: 0, Value: 0},
			wantCp: &Counters{Counter{0, 0}},
		},
		{
			name:   "counters slice with one matching element",
			args:   args{cp: &Counters{Counter{0, 0}}, i: 0, shouldChange: false},
			want:   Counter{Index: 0, Value: 0},
			wantCp: nil,
		},
		{
			name:   "counters slice with one non-matching element",
			args:   args{cp: &Counters{Counter{1, 0}}, i: 0, shouldChange: true},
			want:   Counter{Index: 0, Value: 0},
			wantCp: &Counters{Counter{0, 0}, Counter{1, 0}},
		},
		{
			name:   "counters slice with two elements",
			args:   args{cp: &Counters{Counter{0, 0}, Counter{1, 0}}, i: 0, shouldChange: false},
			want:   Counter{Index: 0, Value: 0},
			wantCp: nil,
		},
		{
			name:   "counters slice with two elements and sorting",
			args:   args{cp: &Counters{Counter{1, 0}, Counter{2, 0}}, i: 0, shouldChange: true},
			want:   Counter{Index: 0, Value: 0},
			wantCp: &Counters{Counter{0, 0}, Counter{1, 0}, Counter{2, 0}},
		},
	}
	for _, tt := range tests {
		orig := make(Counters, len(*tt.args.cp))
		copy(orig, *tt.args.cp)

		t.Run(tt.name, func(t *testing.T) {
			if got := GetCounter(tt.args.cp, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCounter() = %v, want %v", got, tt.want)
			}
			if !tt.args.shouldChange && !reflect.DeepEqual(*tt.args.cp, orig) {
				t.Errorf("GetCounters(): cp has changed")
			}
			if tt.args.shouldChange && reflect.DeepEqual(*tt.args.cp, orig) {
				t.Errorf("GetCounters(): cp has NOT changed")
			}
			if tt.wantCp != nil && !reflect.DeepEqual(tt.args.cp, tt.wantCp) {
				t.Errorf("GetCounters(): cp = %v, want %v", tt.args.cp, tt.wantCp)
			}
		})
	}
}

func TestSetCounter(t *testing.T) {
	type args struct {
		cp *Counters
		i  int
		v  int
	}
	tests := []struct {
		name    string
		args    args
		want    Counter
		wantErr bool
	}{
		{
			name:    "unset counter",
			args:    args{cp: &Counters{}, i: 0, v: 0},
			want:    Counter{},
			wantErr: true,
		},
		{
			name:    "one counter",
			args:    args{cp: &Counters{Counter{0, 0}}, i: 0, v: 42},
			want:    Counter{0, 42},
			wantErr: false,
		},
		{
			name:    "two counters",
			args:    args{cp: &Counters{Counter{0, 0}, Counter{1, 1}}, i: 1, v: 42},
			want:    Counter{1, 42},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetCounter(tt.args.cp, tt.args.i, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetCounter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIncrement(t *testing.T) {
	type args struct {
		cp *Counters
		i  int
		by int
	}
	tests := []struct {
		name    string
		args    args
		want    Counter
		wantCp  *Counters
		wantErr bool
	}{
		{
			name:    "unset counter",
			args:    args{cp: &Counters{}, i: 0, by: 1},
			want:    Counter{0, 1},
			wantCp:  &Counters{Counter{0, 1}},
			wantErr: false,
		},
		{
			name:    "one counter",
			args:    args{cp: &Counters{Counter{0, 1}}, i: 0, by: 42},
			want:    Counter{0, 43},
			wantCp:  &Counters{Counter{0, 43}},
			wantErr: false,
		},
		{
			name:    "two counters",
			args:    args{cp: &Counters{Counter{0, 1}, Counter{1, 1}}, i: 1, by: 42},
			want:    Counter{1, 43},
			wantCp:  &Counters{Counter{0, 1}, Counter{1, 43}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Increment(tt.args.cp, tt.args.i, tt.args.by)
			if (err != nil) != tt.wantErr {
				t.Errorf("Increment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Increment() = %v, want %v", got, tt.want)
			}
			if tt.wantCp != nil && !reflect.DeepEqual(tt.args.cp, tt.wantCp) {
				t.Errorf("Increment(): cp = %v, want %v", tt.args.cp, tt.wantCp)
			}
		})
	}
}

func TestDecrement(t *testing.T) {
	type args struct {
		cp *Counters
		i  int
		by int
	}
	tests := []struct {
		name    string
		args    args
		want    Counter
		wantCp  *Counters
		wantErr bool
	}{
		{
			name:    "unset counter",
			args:    args{cp: &Counters{}, i: 0, by: 1},
			want:    Counter{0, -1},
			wantCp:  &Counters{Counter{0, -1}},
			wantErr: false,
		},
		{
			name:    "one counter",
			args:    args{cp: &Counters{Counter{0, 1}}, i: 0, by: 42},
			want:    Counter{0, -41},
			wantCp:  &Counters{Counter{0, -41}},
			wantErr: false,
		},
		{
			name:    "two counters",
			args:    args{cp: &Counters{Counter{0, 1}, Counter{1, 1}}, i: 1, by: 42},
			want:    Counter{1, -41},
			wantCp:  &Counters{Counter{0, 1}, Counter{1, -41}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrement(tt.args.cp, tt.args.i, tt.args.by)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrement() = %v, want %v", got, tt.want)
			}
			if tt.wantCp != nil && !reflect.DeepEqual(tt.args.cp, tt.wantCp) {
				t.Errorf("Decrement(): cp = %v, want %v", tt.args.cp, tt.wantCp)
			}
		})
	}
}
