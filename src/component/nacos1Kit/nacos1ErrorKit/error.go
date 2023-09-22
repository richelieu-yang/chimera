package nacos1ErrorKit

// IsInstanceListEmpty
/**
 * 参考：nacos-sdk-go依赖中的naming_client.go.
 * 涉及的方法：selectInstances()、selectOneHealthyInstances().
 */
func IsInstanceListEmpty(err error) bool {
	return err != nil && err.Error() == "instance list is empty!"
}

// IsHealthyInstanceListEmpty
/**
 * 参考：nacos-sdk-go依赖中的naming_client.go.
 * 涉及的方法：selectOneHealthyInstances().
 */
func IsHealthyInstanceListEmpty(err error) bool {
	return err != nil && err.Error() == "healthy instance list is empty!"
}
