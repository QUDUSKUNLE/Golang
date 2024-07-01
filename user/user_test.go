package user

import (
	"testing"
)

func TestPerson_FullName(t *testing.T) {
	type fields struct {
		Name FullName
		Age  int
	}
	tests := []struct {
		name   string
		fields fields
		want   Name
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Person{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if got := n.FullName(); got != tt.want {
				t.Errorf("Person.FullName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_EditFirstName(t *testing.T) {
	type fields struct {
		Name FullName
		Age  int
	}
	type args struct {
		newName FullName
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
			n := &Person{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if err := n.EditFirstName(tt.args.newName); (err != nil) != tt.wantErr {
				t.Errorf("Person.EditFirstName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPerson_DeleteLastName(t *testing.T) {
	type fields struct {
		Name FullName
		Age  int
	}
	type args struct {
		lastName FullName
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
			n := &Person{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if err := n.DeleteLastName(tt.args.lastName); (err != nil) != tt.wantErr {
				t.Errorf("Person.DeleteLastName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPerson_GetFirstName(t *testing.T) {
	type fields struct {
		Name FullName
		Age  int
	}
	tests := []struct {
		name   string
		fields fields
		want   Name
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Person{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if got := n.GetFirstName(); got != tt.want {
				t.Errorf("Person.GetFirstName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_GetLastName(t *testing.T) {
	type fields struct {
		Name FullName
		Age  int
	}
	tests := []struct {
		name   string
		fields fields
		want   Name
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Person{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if got := n.GetLastName(); got != tt.want {
				t.Errorf("Person.GetLastName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatRandonNumbers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := StatRandonNumbers(tt.args.n)
			if got != tt.want {
				t.Errorf("StatRandonNumbers() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("StatRandonNumbers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSquaresOfSumAndDiff(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SquaresOfSumAndDiff(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("SquaresOfSumAndDiff() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SquaresOfSumAndDiff() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCompareLower4Bits(t *testing.T) {
	type args struct {
		m uint32
		n uint32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareLower4Bits(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("CompareLower4Bits() = %v, want %v", got, tt.want)
			}
		})
	}
}
