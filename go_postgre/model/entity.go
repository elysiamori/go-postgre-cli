package model

type CStudent struct {
	ID          string
	FullName    string
	Email       string
	Age         int
	Programming string
}

// cstudents type cstudent list
type CStudents []CStudent

// constructor
func NewCStudent() *CStudent {
	return &CStudent{}
}
