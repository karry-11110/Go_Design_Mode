package simplefactory

type IGroup interface {
	Get(interface{})
	Post(interface{})
	Put(interface{})
	Delete(interface{})
}

func NewIGroup(t int) IGroup {
	if t == 1 {
		return &Enginee{}
	} else if t == 2 {
		return &Core{}
	}
	return nil
}

type Enginee struct {
}

// 匹配GET 方法, 增加路由规则
// 匹配DELETE 方法, 增加路由规则
func (c *Enginee) Get(interface{}) {
}

// 匹配POST 方法, 增加路由规则
// 匹配DELETE 方法, 增加路由规则
func (c *Enginee) Delete(interface{}) {
}

// 匹配PUT 方法, 增加路由规则
// 匹配DELETE 方法, 增加路由规则
func (c *Enginee) Put(interface{}) {
}

// 匹配DELETE 方法, 增加路由规则
func (c *Enginee) Post(interface{}) {
}

type Core struct {
}

// 匹配GET 方法, 增加路由规则
// 匹配DELETE 方法, 增加路由规则
func (c *Core) Get(interface{}) {
}

// 匹配POST 方法, 增加路由规则
// 匹配DELETE 方法, 增加路由规则
func (c *Core) Delete(interface{}) {
}

// 匹配PUT 方法, 增加路由规则
// 匹配DELETE 方法, 增加路由规则
func (c *Core) Put(interface{}) {
}

// 匹配DELETE 方法, 增加路由规则
func (c *Core) Post(interface{}) {
}
