package config

import "github.com/zeromicro/go-zero/rest"

// 用于 yaml 配置文件
type Config struct {
	rest.RestConf
	Gender string
	Point  Point1 
}

// 注意：属性名不要设置为单个字母，设置为单个字母好像无法正确解析，可能是bug
type Point1 struct {
	XValue int 
	YValue int 
}

// 用于 json 配置文件
// 由于在 json 中首字母是小写的，和这里首字母大写不一样，因此需要设置tag
// type Config struct {
// 	rest.RestConf
// 	Address string `json:"address,optional"`
// 	Age int `json:"age,default=18"`
// 	Point Point `json:"point"`
// }

type Point2 struct {
	X int `json:"x"`
	Y int `json:"y"`
}
