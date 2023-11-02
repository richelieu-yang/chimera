package pushKit

func PushToAll(data []byte) error {
	if err := isAvailable(); err != nil {
		return err
	}

	return nil
}

func PushToGroup(data []byte, group string) error {
	if err := isAvailable(); err != nil {
		return err
	}

	return nil
}

func PushToUser(data []byte, user string) error {
	if err := isAvailable(); err != nil {
		return err
	}

	return nil
}

func PushToBsid(data []byte, bsid string) error {
	if err := isAvailable(); err != nil {
		return err
	}

	return nil
}
