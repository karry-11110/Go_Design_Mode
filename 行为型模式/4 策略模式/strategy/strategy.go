package strategy

type Strategy interface {
	Authenticate(name string) error
}

var strategys = map[string]Strategy{
	"qq":     qq{},
	"wechat": wechat{},
}

func NewStrategy(t string) Strategy {
	s, _ := strategys[t]
	return s
}

type wechat struct {
}

func (wc wechat) Authenticate(name string) error {
	return nil
}

type qq struct {
}

func (q qq) Authenticate(name string) error {
	return nil
}
