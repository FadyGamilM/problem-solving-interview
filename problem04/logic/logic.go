package logic

// arg1(arg2(input))
func Compose(arg1, arg2 func())

type OneArgFunc func(arg1 interface{}) interface{}

// x : composition f after g
// x : call f after you call g : call g(param) then call f(result_of_calling g(param))

// compose(square, inc) => h
// h(6) = sqaure(inc(6))

// -> compose(fn1, fn2) returns fn1(fn2())
