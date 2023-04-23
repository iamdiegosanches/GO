// primeiro vai precisar de fazer go mod init "nome do pacote" no terminal para criar o arquivo de modulos

package stuff

import(
	"strconv"
)

var Name string  = "Derek"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, i :=  range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}
