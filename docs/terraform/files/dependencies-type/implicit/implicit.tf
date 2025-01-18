resource local_file "my-file" {
  content = "My pet is ${random_pet.my-pet.id}"
  filename = "/tmp/my-file.txt"
}
resource random_pet "my-pet" {
  prefix = "Mrs"
  separator = "."
  length = "1"
}