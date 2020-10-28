package flache

type memCache struct {

}

func NewMem()*memCache{
	return &memCache{}
}

func(m memCache)Get(key string)(interface{},error){
	return nil,nil
}

func(m memCache)Set(key string, value interface{})error{
	return nil
}