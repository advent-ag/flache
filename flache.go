package flache

type Flache interface {
 Get(key string)(interface{},error)
 Set(key string, value interface{})error
}

type Cachetype string

type Config struct {
	Cachetype Cachetype
}

func NewFlach(Cachetype Cachetype) Flache{
	switch Cachetype {
	case "mem":
       return NewMem()
	}
	return nil
}