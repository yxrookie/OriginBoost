package main

import (
	"OriginBoost/config"
	bitconfig "OriginBoost/pkg/config"
	"flag"
	
	l "OriginBoost/app/loseweight"
	s "OriginBoost/app/similarity"
	"OriginBoost/app/table"
)



func init() {
	config.Initialize()
	
}

func main() {
	// 配置初始化，依赖命令行 --env 参数
    var env string
    flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
    flag.Parse()
    bitconfig.InitConfig(env)
	teststring := "多巴胺的增加会让人热情参与自己原本认为不重要的事务。例如， 有报道称一些吸大麻的人会站在水池前，看着水龙头不断地滴水，这一 平平无奇的景象让他们看得十分着迷。当吸食大麻的人迷失在自己的思 想中，漫无目的地在自己创造的想象世界中遨游时，多巴胺的增强效应 也就变得更明显。但在某些情况下，大麻也会抑制多巴胺，模仿当下分 子的作用。这种情况下，通常与渴望和动机有关的活动，如工作、学习 或洗澡，看起来就不那么重要了。"
	tempslice := l.Loseweight(teststring)
    res := s.Sortres(teststring, tempslice)
	table.Table(res)
	
}


