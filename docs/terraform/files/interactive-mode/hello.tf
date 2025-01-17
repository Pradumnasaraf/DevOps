resource "local_file" hello{
  filename = "./hello.txt"
  content  = var.say_hello
  file_permission = "0700"
}