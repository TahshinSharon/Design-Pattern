package main

import (
	"fmt"
	"os"
)

type Report struct{
	Data string
}
func(r Report)Generate()string{
	return fmt.Sprintf("Report %s",r.Data)
}
type FileSaver struct{

}
func( f FileSaver) Save(filename string,content string) error{
	return os.WriteFile(filename,[]byte(content),0644)
}
type EmailSender struct{}
func(e EmailSender)Send(email string,content string){
	fmt.Printf("Sending '%s' to %s\n",content,email)
}
func main() {

	report := Report{
		Data: "Sales Data In Go",
	} 
	content := report.Generate()
	file := FileSaver{}
	err := file.Save("report.txt", content)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
	mail := EmailSender{}
	mail.Send("abc@gmail.com", content)

}