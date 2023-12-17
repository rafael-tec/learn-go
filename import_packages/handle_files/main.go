package main

impot (
	"os"
	"fmt"
	"bufio"
)

func main() {
	writeFile("Hello Word!\nHello Brazil!\nHello São Paulo!")
	readFile("file.txt")
	bufferedFileRead("file.txt")
	removeFile("file.txt")
}

func writeFile(content string) {
	fmt.Println("\nStarting file creation")

	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	length, err := f.WriteString(content)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File created successfully, size in bytes: %d\n", length)

	f.Close()
}

func readFile(fileName string) {
	fmt.Println("\nStarting file reading")
	f, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File content: %s\n", string(f))
}

func bufferedFileRead(fileName string) {
	fmt.Println("\nStarting buffered file reading")

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
	
		fmt.Println(string(buffer[:n]))
	}
}

func removeFile(fileName string) {
	fmt.Println("\nStarting file remotion")

	err := os.Remove(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File %s removed\n", fileName)
}