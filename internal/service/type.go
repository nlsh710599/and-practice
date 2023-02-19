package service

type CreateRequest struct {
	Name  string
	Value string
}

type UpdateRequest struct {
	Name  string
	Value string
}

type DeleteRequest struct {
	Name string
}

type AddRequest struct {
	Addend  string
	Summand string
}

type SubtractRequest struct {
	Subtrahend string
	Minuend    string
}

type MultiplyRequest struct {
	Multiplicand string
	Multiplier   string
}

type DivideRequest struct {
	Dividend string
	Divisor  string
}
