package similarity

import (
	bitconfig "OriginBoost/pkg/config"
	l "OriginBoost/app/loseweight"
	"flag"
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	// 配置初始化，依赖命令行 --env 参数
    var env string
    flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
    flag.Parse()
    bitconfig.InitConfig(env)
	teststring := "我曾经自认为是个自由主义者，但后来我发现，我之所以捍卫那些 我没有真正思考过的立场，是因为这些立场是自由主义信条的一部分。 只有立场却没有是非，这是不可取的。"
	reslice := l.Loseweight(teststring)
    Sortres(teststring, reslice)
	for i := 0; i < len(reslice); i++ {
		fmt.Println(reslice[i])
	}
}