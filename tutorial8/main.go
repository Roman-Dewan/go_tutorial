package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["js"] = "javaScript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("languages: ", languages)
	fmt.Printf("Type of langugae: %T", languages)

	delete(languages, "rb")
	fmt.Println("languages: ", languages)

	for key, value := range languages {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

}
