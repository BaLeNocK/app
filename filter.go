package main
 
import (
	"fmt"
	"os"
	"path/filepath"
)
 
func main() {
	for i := 1; i <= 6; i++ {
		// Создаем несколько тестовых файлов
		_, err := os.Create(fmt.Sprintf("./test.file%d", i))
		if err != nil {
			fmt.Println(err)
		}
	}
 
	// Все файлы которые заканчиваются с расширением .file1, .file2, .file3
	m, err := filepath.Glob("test.file[1-3]")
	if err != nil {
		panic(err)
	}
 
	for _, val := range m {
		fmt.Println(val)
	}
 
	// Удаляем временные файлы
	for i := 1; i <= 6; i++ {
		err := os.Remove(fmt.Sprintf("./test.file%d", i))
		if err != nil {
			fmt.Println(err)
		}
	}
}
