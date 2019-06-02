package otus_shortener

import (
	"fmt"
	"hash/fnv"
)

type Storage map[string]string

type Shortener interface {
	Shorten(string) string
	Resolve(string) string
}

func main() {
	rawString := "http://localhost:7070/pkg/hash/fnv/"
	storage := Storage{}

	hashString := storage.Shorten(rawString)
	storage[hashString] = rawString
	resolvedURL := storage.Resolve(hashString)

	fmt.Println("Hash:", hashString)
	fmt.Println("URL:", resolvedURL)
}

func (s Storage) Shorten(str string) string {
	h := fnv.New32()
	h.Write([]byte(str))
	return fmt.Sprintf("%d", h.Sum32())
}

func (s Storage) Resolve(str string) string {
	return s[str]
}
