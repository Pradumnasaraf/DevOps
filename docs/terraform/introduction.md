---
sidebar_position: 1
title: Terraform Introduction
description: Learn about Terraform and its use cases.
tags: ["Terraform", "Infrastructure as Code", "HashiCorp"]
keywords: ["Terraform", "Infrastructure as Code", "HashiCorp"]
slug: "/terraform"
---

Terraform is an open-source infrastructure as code software tool created by HashiCorp. It allows users to define and provision data center infrastructure using a declarative configuration language. 

## Challenges with traditional infrastructure

In traditional infrastructure management, system administrators would manually configure servers, networks, and storage devices. This manual process was time-consuming, error-prone, and difficult to reproduce consistently across different environments. These come with the following challenges:

- **Slow Deployment**: Manual infrastructure provisioning is slow and error-prone, leading to delays in deploying new applications or services.
- **Expensive**: Manual infrastructure management requires significant human effort, leading to higher operational costs.
- **Limited Automation**: Traditional infrastructure management lacks automation, making it difficult to scale and manage large, complex environments.
- **Human Errors**: Manual configuration can lead to human errors, security vulnerabilities, and inconsistencies across environments.
- **Lack of Consistency**: Manual configuration can result in inconsistencies between development, testing, and production environments.
- **Wasted Resources**: Manual infrastructure management can lead to underutilization of resources and inefficient capacity planning.

To solve these challenges, organizations are adopting Infrastructure as Code (IaC) tools like Terraform.

## Infrastructure as Code (IaC)

Infrastructure as Code (IaC) is the practice of managing infrastructure using code and automation. With IaC, infrastructure configurations are defined in code files that can be version-controlled. In the Iac tooling space we have 3 types of tools:

- **Configuration Management Tools**: Tools like Ansible, Chef, and Puppet are used to automate the configuration of servers and applications.
- **Provisioning Tools**: Tools like Terraform, CloudFormation, and ARM templates are used to provision and manage cloud resources.
- **Server Templating Tools**: Tools like Packer and Docker are used to create machine images that can be deployed across different environments.

![types of IaC tools](https://github.com/user-attachments/assets/d81f948a-d79a-4eeb-8f19-31cc4db30d00)

## Why Terraform?

Terraform is a popular Infrastructure as Code tool that allows you to define and provision infrastructure using a declarative configuration language. With Terraform we can deploy infrastructure across multiple cloud providers like AWS, Azure, Google Cloud, and on-premises environments. And it supports a variety of providers like Cloudflare, Grafana, Auth0, and more.

![Terraform Providers](https://github.com/user-attachments/assets/63d359e3-3370-46ff-ba27-dd7de02aeda)

Terraform uses a simple syntax called HashiCorp Configuration Language (HCL) to define the infrastructure components and their dependencies. It's easy to learn and understand, and it allows you to create reusable modules and share them with others.

For example to create an AWS EC2 instance using Terraform, you can define the following configuration:

```hcl
resource "aws_instance" "example" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"
}
```

![Terraform Workflow](https://github.com/user-attachments/assets/c0697d6a-a73e-419b-8d52-6558cb9d1b63)

It start by creating a configuration file and then work in three phases, init to initialize the working directory, plan to create an execution plan, and apply to apply the changes.

Each object the Terraform manages is called a resource. Resources are defined in Terraform configuration files, and Terraform uses providers to interact with APIs of cloud providers and services. If any of resources go out of sync with the desired state, Terraform can update the resources to match the desired state.

### Terraform benefits

- **Declarative Configuration**: Terraform uses a declarative configuration language that allows you to define the desired state of your infrastructure. This means we define what we want, not how to get it.
- **Cloud Agnostic**: Terraform supports multiple cloud providers and services, allowing you to manage infrastructure across different environments.
- **Scalability**: Terraform can manage infrastructure at scale, making it suitable for small projects and large enterprises.
- **Integration**: Terraform can be integrated with other tools and services like CI/CD pipelines, monitoring systems, and configuration management tools.
- **Cost-Efficient**: Terraform helps optimize infrastructure costs by providing visibility into resource usage and enabling efficient capacity planning.
- **Security**: Terraform configurations can be audited, reviewed, and tested for security vulnerabilities, ensuring compliance with security best practices.
- **State Management**: Terraform stores the state of your infrastructure in a state file, allowing you to manage, version, and share the state across teams.

## HCL (HashiCorp Configuration Language)

HCL is a simple, human-readable language that is used to define infrastructure configurations in Terraform. It is designed to be easy to read and write, making it accessible to both developers and operators. 

Syntax:

```hcl
<block_type> "<block_label>" {
  key1 = value1
  key2 = value2
  ...
}
```

It consists of blocks, arguments, and expressions. Blocks contains information about the infrastructure platform and resources within that platform we want to create. Arguments are key-value pairs that define the configuration settings for a block. Expressions are used to reference variables, functions, and other resources within the configuration.

For example if want to create a file locally with some content, we can define the following configuration:

```hcl
# local.tf
resource "local_file" "pet" {
  filename = "/root/pets.txt"
  content  = "We love pets!"
}
```

Breakdown of the above configuration:

![HCL Example](https://github.com/user-attachments/assets/a958cd64-04c7-410a-91dc-4ca89e72ac67)

Another example of ec2 instance creation:

```hcl
# ec2.tf
resource "aws_instance" "example" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"
}
```

Another example of creating a s3 bucket:

```hcl
# s3.tf
resource "aws_s3_bucket" "example" {
  bucket = "my-unique-bucket"
  acl    = "private"
}
```

## Terraform Workflow

To deploy a resource using Terraform, there is four step process:

1. **Configuration**: Define the infrastructure components and their dependencies in Terraform configuration files.
2. **Initialization**: Initialize the working directory with `terraform init` command to download the necessary providers and modules.
3. **Planning**: Create an execution plan with `terraform plan` command to preview the changes that Terraform will make.
4. **Apply**: Apply the changes with `terraform apply` command to create, update, or delete the resources as per the configuration.

So, let's se an example of creating a file locally with some content. First we need to create a configuration file:

```hcl
# local.tf
resource "local_file" "pet" {
  filename = "./pets.txt"
  content  = "We love pets!"
}
```

Then next step is to initialize the working directory:

```bash
terraform init
```

This will initialize the working directory and download the necessary providers and modules. Next, we can create an execution plan:

```bash
terraform plan
```

This will show the changes that Terraform will make to create the file locally. Finally, we can apply the changes:

```bash
terraform apply
```

And then type `yes` to confirm the changes.

This will create the file locally with the specified content. You can verify the changes by checking the file `./pets.txt`.

We can use `terraform show` command to see the current state of the infrastructure managed by Terraform. 

Let's now update the file permissions of the file we created:

```hcl
# local.tf
resource "local_file" "pet" {
  filename = "./pets.txt"
  content  = "We love pets!"
  file_permission = "0700"
}
```

To check what changes Terraform will make, we can run `terraform plan`. And to apply the changes, we can run `terraform apply`.

To destroy the resources created by Terraform, we can run `terraform destroy` command. And type `yes` to confirm the deletion of resources.
