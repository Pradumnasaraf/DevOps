output "pet_name" {
  value = random_pet.my-pet.id # The random pet name from main.tf
  description = "The random pet name generated by Terraform"
}