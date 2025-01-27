resource "aws_s3_bucket" "finance" {
  bucket = "finance-21092020"
  tags = {
    Name = "Finance and Payroll"
  }
}

resource "aws_s3_object" "finance-2020" {
  content = "/root/finance/finance-2020.doc"
  key = "finance-2020.doc"
  bucket = aws_s3_bucket.finance.id # reference to the bucket
}

data "aws_iam_group" "finance-data" {
  group_name = "finance-analysts"
}

resource "aws_s3_bucket_policy" "finance-policy" {
  bucket = aws_s3_bucket.finance.id
  policy = <<EOF
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": "*",
        "Effect": "Allow",
        "Resource": "arn:aws:s3:::${aws_s3_bucket.finance.id}/*",
        "Principal": {
          "AWS": ["${data.aws_iam_group.finance-data.arn}"
          ]
        }
      }
    ]
  }
  EOF
}
