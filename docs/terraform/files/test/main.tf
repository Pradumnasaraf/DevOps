# local.tf
resource "local_file" "pet" {
  for_each = var.files
  content  = "We love pets!"
  filename = each.value
}

# variables.tf
variable "files" {
  default = {
    pets = "./pets.txt"
    dogs = "./dogs.txt"
    cats = "./cats.txt"
  }
}