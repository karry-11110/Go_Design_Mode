package chain

type SensitiveWordFilter interface {
	Filter(content string) bool
}
type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

func (c *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter) {
	c.filters = append(c.filters, filter)
}
func (c *SensitiveWordFilterChain) Filter(content string) bool {
	for _, filter := range c.filters {
		if filter.Filter(content) {
			return true
		}
	}
	return false
}

// AdSensitiveWordFilter 广告
type AdSensitiveWordFilter struct{}

// Filter 实现过滤算法
func (f *AdSensitiveWordFilter) Filter(content string) bool {
	// TODO: 实现算法
	return false
}

// PoliticalWordFilter 政治敏感
type PoliticalWordFilter struct{}

// Filter 实现过滤算法
func (f *PoliticalWordFilter) Filter(content string) bool {
	// TODO: 实现算法
	return true
}

//type Context struct {
//	// ...
//
//	// handlers 是一个包含执行函数的数组
//	// type HandlersChain []HandlerFunc
//	handlers HandlersChain
//	// index 表示当前执行到哪个位置了
//	index    int8
//
//	// ...
//}
//
//// Next 会按照顺序将一个个中间件执行完毕
//// 并且 Next 也可以在中间件中进行调用，达到请求前以及请求后的处理
//// Next should be used only inside middleware.
//// It executes the pending handlers in the chain inside the calling handler.
//// See example in GitHub.
//func (c *Context) Next() {
//	c.index++
//	for c.index < int8(len(c.handlers)) {
//		c.handlers[c.index](c)
//		c.index++
//	}
//}
