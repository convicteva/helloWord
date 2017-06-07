package file

import "os"

const (
	javaCodeRetractSpace_1 = "    "
	javaCodeRetractSpace_2 = "        "
	javaCodeRetractSpace_3 = "            "
)

var pathSeparator string

func init() {
	pathSeparator = string(os.PathSeparator)
}
