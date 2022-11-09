package decorater

type IDraw interface {
	Draw() string
}

type Square struct {
}

func (s Square) Draw() string {
	return "this is a square"
}

type ColorSquare struct {
	square IDraw
	color  string
}

func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{color: color, square: square}
}

func (c ColorSquare) Draw() string {
	return c.square.Draw() + ", color is" + c.color
}

//type HandlerFunc func(interface{})
//
//func normal(f HandlerFunc) HandlerFunc {
//	return func(i interface{}) {
//		f(i)
//	}
//}
//
//func min(f HandlerFunc) HandlerFunc {
//	return func(i interface{}) {
//		f(i)
//	}
//}
//func f(i interface{}) {
//}
