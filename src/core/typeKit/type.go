package typeKit

import "fmt"

func GetTypeString(obj interface{}) string {
	return fmt.Sprintf("%T", obj)
}
