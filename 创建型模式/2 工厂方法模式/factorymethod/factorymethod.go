package factorymethod

type IGroup interface {
	Get(interface{})
	Post(interface{})
	Put(interface{})
	Delete(interface{})
}

type Enginee struct {
}

// 匹配GET 方法, 增加路由规则
func (c Enginee) Get(interface{}) {
}


// 匹配DELETE 方法, 增加路由规则
func (c Enginee) Delete(interface{}) {
}

// 匹配PUT 方法, 增加路由规则
func (c Enginee) Put(interface{}) {
}

// 匹配DELETE 方法, 增加路由规则
func (c Enginee) Post(interface{}) {
}

type router interface {
	Group(interface{}) IGroup
}
type group struct {
}

func (g group) Group(interface{}) IGroup {
	return Enginee{}
}

func Newrouter(t int) router {
	return group{}
}

//func (g Enginee) Group(interface{}) IGroup {
//	return Enginee{}
//}
//
//func Newrouter(t int) router {
//	return Enginee{}
//}
