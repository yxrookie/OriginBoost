package loseweight

import "OriginBoost/app/loseweight/utils"


// return the result after reducing similarity
func Loseweight(query string) []string {
	resString := make([]string, 7)
	resString[0] = utils.Method1(query)
	resString[1] = utils.Method2(query)
	resString[2] = utils.Method3(query)
	resString[3] = utils.Method4(query)
	resString[4] = utils.Method5(query)
	resString[5] = utils.Method6(query)
	resString[6] = utils.Method7(query)
	return resString
}

