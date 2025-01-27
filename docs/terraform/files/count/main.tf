resource "local_file" "pet" {
  count   = length(var.pet_filenames)
  filename = var.pet_filenames[count.index]
  content  = "This is a file named ${var.pet_filenames[count.index]}"
}

output "pets" {
  value = local_file.pet
}

resource "local_file" "pets" {
  count    = length(var.animals)
  content  = "I love ${var.animals[count.index]}"
  filename = "./${var.animals[count.index]}.txt"
}


variable "animals" {
  default = [
    "rats",
    "dogs",
    "cats"
  ]
}