package windows

import "golang.org/x/sys/windows/registry"

// getUserPath retrieves the user PATH from the Windows registry.
func getUserPath() (string, error) {
	// Open the Environment key in the registry for the current user
	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.READ)
	if err != nil {
		return "", err
	}
	defer k.Close()

	// Retrieve the "Path" value from the registry
	path, _, err := k.GetStringValue("Path")
	if err != nil {
		return "", err
	}
	return path, nil
}

// setUserPath sets the user PATH in the Windows registry.
func setUserPath(newPath string) error {
	// Open the Environment key in the registry for the current user with SET_VALUE permission
	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	// Set the new user PATH in the registry
	return k.SetStringValue("Path", newPath)
}
