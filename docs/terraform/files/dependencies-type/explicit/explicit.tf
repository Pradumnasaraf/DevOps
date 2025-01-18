resource local_file "my-file" {
  content = "I love Pets"
  filename = "/tmp/my-file.txt"
  depends_on = [random_pet.my-pet]
}
resource random_pet "my-pet" {
  prefix = "Mrs"
  separator = "."
  length = "1"
}