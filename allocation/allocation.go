package allocation


type File struct {
	fd int
	name string
	dirinfo interface{}
	nepipe int
}


func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name, nil, 0}
}

var P *[]int = new([]int) // Pointer P
var V []int = make([]int, 100) // Allocate to V

