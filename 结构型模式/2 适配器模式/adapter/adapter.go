package adapter

type Enginne struct {
}

func (e *Enginne) Get(url string, handler interface{}) {

}

type IGroup interface {
	gGet(url string, handler interface{})
}

// Group struct 实现了适配器
type Group struct {
	enginne *Enginne // 指向core结构
	prefix  string   // 这个group的通用前缀
}

// 实现Get方法
func (g *Group) gGet(uri string, handler interface{}) {
	//一些操作
	g.enginne.Get(uri, handler)
}
