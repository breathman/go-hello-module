package hello

import "fmt"

func Hello(name string, lang string) (string, error) {
	switch lang {
	case "en": return fmt.Sprintf("Hello %s! \n", name), nil
	case "ru": return fmt.Sprintf("Привет %s! \n", name), nil
	default:
		return "", fmt.Errorf("lang %q not supported", lang)
	}
}