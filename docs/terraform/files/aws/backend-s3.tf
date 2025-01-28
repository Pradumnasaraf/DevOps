terraform {
  backend "s3" {
    bucket = "remote-state"
    key = "terraform.tfstate"
    region = "us-east-1"
  }
}