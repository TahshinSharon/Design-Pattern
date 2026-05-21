class Report:
    def __init__(self,data):
        self.data=data
    def generate(self):
        return f"report: {self.data}"

class FileSaver:
    def save(self,filename,content):
        with open(filename,"w") as f:
            f.write(content)
class EmailSender:
    def send(self,email,content):
        print(f"Sending '{content}' to {email}")
report=Report("Sales Data")
content=report.generate()
file=FileSaver()
file.save("report.txt",content)
mail=EmailSender()
mail.send("abc@gmail.com",content)