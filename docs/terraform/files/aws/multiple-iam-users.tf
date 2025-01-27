# main.tf
variable "dev-team" {
  type = list(string)
  default = ["lucy", "john", "jane"]
}

resource "aws_iam_user" "admin-user" {
  name = var.dev-team[count.index]
  count = length(var.dev-team)
  tags = {
    Description = "Technical Team Leader"
  }
}