package ioc

type ApplicationContextInterface interface {
	SetBean(name string, v any)
	GetBeanByName(name string) (any, error)
	GetProxyBeanByName(name string) (any, error)
}

// BeforeFunc 前置处理
type BeforeFunc func()

// AfterFunc 后置处理
type AfterFunc func()

// Bean 每一个被注入的bean都应该具有 实际注入结构体的指针、前置处理函数集合、后置处理函数集合
type Bean struct {
	BeanPtr     any
	BeforeFuncs []BeforeFunc
	AfterFuncs  []AfterFunc
}

// Beans k结构体的名字
type Beans map[string]*Bean

type ApplicationContextDefaultImpl struct {
	Beans
}

func NewApplicationContextDefaultImpl(map[string]any) *ApplicationContextDefaultImpl {
	impl := ApplicationContextDefaultImpl{}
	impl.Beans = map[string]*Bean{} //避免空指针
	return &impl
}

func (a *ApplicationContextDefaultImpl) SetBean(name string, v any) {
	//SET BEAN
	//todo 存在一个问题，name是全局不能重复的，否则出现一个注入的bean错乱，这里暂时不要使用反射解决这个问题，
	//todo 所以对name提出一个开发要求，使用结构体的全路径+结构体名称作为name，可以避免这个问题，后期再思考怎么根治这个问题

	_, ok := a.Beans[name]
	if !ok {
		b := new(Bean)
		a.Beans[name] = b
	}

	//注入对象指针赋值
	a.Beans[name].BeanPtr = v
}

func (a *ApplicationContextDefaultImpl) GetBeanByName(name string) (any, error) {
	// get bean
	a2 := a.Beans[name]
	return a2, nil
}

func (a *ApplicationContextDefaultImpl) GetProxyBeanByName(name string) (any, error) {
	// get bean
	a.Beans
	//执行前置处理
	for _, beforeFunc := range a.Beans[name].BeforeFuncs {
		beforeFunc()
	}

	a2 := a.Beans[name]

	for _, afterFunc := range a.AfterFuncs {
		afterFunc()
	}
	return a2, nil
}
