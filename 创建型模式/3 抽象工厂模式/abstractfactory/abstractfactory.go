package abstractfactory

type IGroup interface {
	Get(interface{})
	Post(interface{})
	Put(interface{})
	Delete(interface{})
}

type Enginee struct {
}

func (c Enginee) Get(interface{}) {
}
func (c Enginee) Delete(interface{}) {
}
func (c Enginee) Put(interface{}) {
}
func (c Enginee) Post(interface{}) {
}

type IGroups interface {
	Gets(interface{})
	Posts(interface{})
	Puts(interface{})
	Deletes(interface{})
}
type Core struct {
}

func (c Core) Gets(interface{}) {
}
func (c Core) Deletes(interface{}) {
}
func (c Core) Puts(interface{}) {
}
func (c Core) Posts(interface{}) {
}

type factory interface {
	http() IGroup
	https() IGroups
}
type object struct{}

func (ob object) http() IGroup {
	return Enginee{}
}
func (ob object) https() IGroups {
	return Core{}
}
func Newfactory() factory {
	return object{}
}
