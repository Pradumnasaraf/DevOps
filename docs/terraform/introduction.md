---
sidebar_position: 1
title: Terraform Introduction
description: Learn about Terraform and its use cases.
tags: ["Terraform", "Infrastructure as Code", "HashiCorp"]
keywords: ["Terraform", "Infrastructure as Code", "HashiCorp"]
slug: "/terraform"
---

Terraform is an open source infrastructure as code software tool created by HashiCorp. It allows users to define and provision infrastructure using a declarative configuration language. Terraform manages infrastructure across multiple cloud providers and on-premises environments. It uses a simple syntax called HashiCorp Configuration Language (HCL) to define the infrastructure components and their dependencies.

## Challenges with traditional infra

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
- **Server Templating Tools**: Tools like Packer and Docker are used to create machine images that can be deployed across different environments.
- **Provisioning Tools**: Tools like Terraform, CloudFormation, and ARM templates are used to provision and manage cloud resources.

![types of IaC tools](https://github.com/user-attachments/assets/346a952a-671b-478a-bbee-83d0ba1e301f)

## Why Terraform?

Terraform is a popular Infrastructure as Code tool that allows you to define and provision infrastructure using a declarative configuration language. With Terraform we can deploy infrastructure across multiple cloud providers like AWS, Azure, Google Cloud, and on-premises environments. And it supports a variety of providers like Cloudflare, Grafana, Auth0, and more.

![Terraform Providers](https://github.com/user-attachments/assets/e4fa9895-4cff-4b2d-b637-f1c97b38d599)

Terraform uses a simple syntax called **HashiCorp Configuration Language (HCL)** to define the infrastructure components and their dependencies. It's easy to learn and understand, and it allows you to create reusable modules and share them with others.

![Terraform Workflow](https://github.com/user-attachments/assets/c0697d6a-a73e-419b-8d52-6558cb9d1b63)

For example to create an AWS EC2 instance using Terraform, you can define the following configuration:

```hcl
resource "aws_instance" "example" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"
}
```

It start by creating a configuration file and then work in three phases, **init** to initialize the working directory, **plan** to create an execution plan, and **apply** to apply the changes (more info below)

Each object the Terraform manages is called a **resource**. Resources are defined in Terraform configuration files, and Terraform uses providers to interact with APIs of cloud providers and services. If any of resources go out of sync with the desired state, Terraform can update the resources to match the desired state.

### Terraform benefits

- **Declarative Configuration**: Terraform uses a declarative configuration language that allows you to define the desired state of your infrastructure. This means we define what we want, not how to get it.
- **Cloud Agnostic**: Terraform supports multiple cloud providers and services, allowing you to manage infrastructure across different environments.
- **Scalability**: Terraform can manage infrastructure at scale, making it suitable for small projects and large enterprises.
- **Integration**: Terraform can be integrated with other tools and services like CI/CD pipelines, monitoring systems, and configuration management tools.
- **Cost-Efficient**: Terraform helps optimize infrastructure costs by providing visibility into resource usage and enabling efficient capacity planning.
- **Security**: Terraform configurations can be audited, reviewed, and tested for security vulnerabilities, ensuring compliance with security best practices.
- **State Management**: Terraform stores the state of your infrastructure in a state file, allowing you to manage, version, and share the state across teams.

### HCL (HashiCorp Configuration Language)

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

### Terraform Workflow

To deploy a resource using Terraform, there is four step process:

1. **Configuration**: Define the infrastructure components and their dependencies in Terraform configuration files.
2. **Initialization**: Initialize the working directory with `terraform init` command to download the necessary providers and modules.
3. **Planning**: Create an execution plan with `terraform plan` command to preview the changes that Terraform will make.
4. **Apply**: Apply the changes with `terraform apply` command to create, update, or delete the resources as per the configuration.

So, let's see an example of creating a file locally with some content. First we need to create a configuration file:

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

### Terraform Providers

Terraform uses plugin based architecture to interact with cloud providers and services. Providers are responsible for understanding API interactions and exposing resources for managing infrastructure. Terraform supports a wide range of providers like AWS, Azure, Google Cloud, Kubernetes, Docker, and more. When we do `terraform init`, it downloads the necessary providers and modules. A list of available providers can be found on the [Terraform Registry](https://registry.terraform.io/browse/providers).

There are three types of providers:

1. **Official Providers**: These are maintained by HashiCorp and are considered stable and well-tested. For example, `aws`, `azure`, `google`, `kubernetes`.
2. **Partner Providers**: These are maintained by third-party companies and are officially supported by HashiCorp. For example, `datadog`, `newrelic`, `vault`.
3. **Community Providers**: These are maintained by the community and may not be as stable or well-tested as official providers.

Every time we do `terraform init`, it downloads the necessary providers and modules. Also, it shows the the version of the plugins that are being downloaded.

Every provider has its own set of resources and data sources that can be used to manage infrastructure. For example, the `aws` provider has resources like `aws_instance`, `aws_s3_bucket`, `aws_vpc`, and data sources like `aws_ami`, `aws_availability_zones`, `aws_regions`. And each resource has its own set of arguments and attributes that can be configured. 

For example we use `local_file` it has arguments like `filename`, `content`, `file_permission`.

### Configuration Files

Terraform configuration files have `.tf` extension and are written in HashiCorp Configuration Language (HCL). There are some naming conventions we use for configuration files:

- `main.tf`: Contain resource definition. Instead of creating multiple configuration files like `cat.tf`, `dog.tf`, `fish.tf`, we can define all the resources in a single file `main.tf`.
- `variables.tf`: Contains variable declarations and definitions.
- `outputs.tf`: Contains outputs from the resources.
- `provider.tf`: Contains provider definitions and configurations. 

#### Using Multiple providers

Terraform allows you to use multiple providers in a single configuration file. This is useful when you want to manage resources across different cloud providers or services. You can define multiple provider blocks in the configuration file and specify the provider alias for each block.

An example of using `local` and `random` providers:

```hcl
resource "local_file" "pet" {
  filename = "./pets.txt"
  content  = "We love pets!"
}

resource "random_pet" "my-pet" {
  prefix = "Mrs"
  separator = "."
  length = "1"
}
```

## Input Variables

Input variables allow us to reuse the same configuration with different values. We can define variables in a separate file `variables.tf` and use them in the configuration files. Variables can have default values, descriptions, and types.

### Creating and Using Variables

To create a variable, we define a variable block in the `variables.tf` file:

Syntax:
```hcl
variable "<variable_name>" {
  type        = <variable_type>
  default     = <default_value>
  description = "<description>"
}

An example of creating a variable for filename and content:

```hcl
# variables.tf
variable "filename" {
  default     = "./pets.txt"
}

variable "content" {
  default     = "We love pets!"
}

# main.tf
resource "local_file" "pet" {
  filename = var.filename
  content  = var.content
}
```

Now every we need to change the value of `filename` or `content`, we can simply update the `variables.tf` file and then apply the changes. We don't need to interfere with the main configuration file.

An example of ec2 instance creation using variables:

```hcl
# variables.tf
variable "ami" {
  default     = "ami-0c55b159cbfafe1f0"
}
variable "instance_type" {
  default     = "t2.micro"
}

# main.tf
resource "aws_instance" "webserver" {
  ami           = var.ami
  instance_type = var.instance_type
}
```

### Input Variable Types 

Terraform supports the following variable types:

- **String**: A sequence of characters enclosed in double quotes. For example, `"Hello, World!"`, "/path/to/file".
- **Number**: A numeric value without quotes. For example, `42`, `3.14`.
- **Bool**: A boolean value `true` or `false`. 
- **any**: Any data type. This is the default type if no type is specified.
- **list**: A list of values. For example, `["a", "b", "c"]`.
- **map**: A collection of key-value pairs. For example, `{ key1 = "value1", key2 = "value2" }`.
- **set**: A collection of unique values. For example, `["a", "b", "c"]`.
- **object**: A complex data type that represents a collection of attributes. For example, `{ name = "John", age = 30 }`.
- **truple**: A tuple is an ordered collection of elements. For example, `["a", 42, true]`.

### Using various types of variables:

**List:**

```hcl
# variables.tf
variable "prefix" {
  default     = ["Mr", "Mrs", "Miss", "Dr"]
  type       = list(string) # list of strings
}

# main.tf
resource "random_pet" "my-pet" {
  prefix = var.prefix[0] # Mr
}
```

We can change list type to `number`, `bool`, or `any`. It will fail if we try to use a different type of value.

Map:

```hcl
# variables.tf
variable "pets" {
  type = map(string) # map of strings
  default = {
    "statment1" = "We love pets!"
    "statment2" = "We love animals!"
  }
}

# main.tf
resource "local_file" "pet" {
  filename = "./pets.txt"
  content  = var.pets["statment1"] # We love pets!
}
```

We can change map type to `number`. It will fail if we try to use a different type of value.

**Set:**

```hcl
# variables.tf
variable "prefix" {
  default     = ["Mr", "Mrs", "Miss", "Dr"]
  type       = set(string) # set of strings
}

# main.tf
resource "random_pet" "my-pet" {
  prefix = var.prefix[0] # Mr
}
```

We can change set type to `number`. It will fail if we try to use a different type of value. The only caveat here is that it cannot have duplicate values.

**Object:**

```hcl
# variables.tf
variable "person" {
  type = object({
    name = string
    age  = number
    color = string
    food = list(string)
    favorite_pet = bool
  })
  default = {
    name = "John"
    age  = 30
    color = "blue"
    food = ["pizza", "burger"]
    favorite_pet = true
  }
}

# main.tf
resource "local_file" "person" {
  filename = "./person.txt"
  content  = var.person.name
}

resource "local_file" "person" {
  filename = "./person.txt"
  content  = var.person.age
}

resource "local_file" "person" {
  filename = "./person.txt"
  content  = var.person.color
}

resource "local_file" "bio" {
  filename = "./bio.txt"
  content  = "Name: ${var.person.name}\nAge: ${var.person.age}\nColor: ${var.person.color}\nFood: ${join(", ", var.person.food)}\nFavorite Pet: ${var.person.favorite_pet}"
}
```

Objects help us to define complex data types. We can define different types of values in a single object.

**Tuple:**

```hcl
# variables.tf
variable "person" {
  type = tuple([string, number, bool])
  default = ["John", 30, true]
}

# main.tf
resource "local_file" "person" {
  filename = "./person.txt"
  content  = var.person[0]
}
```

In Tuples, we can define different types of values unlike list, map, and set. We will get an error if we try to use a different type of value.

#### Interactive Mode

Terraform allows you to interactively input values for variables when running `terraform apply` command. We gives values from the CLI when we run `terraform apply` command. The way it works is we don't define value in the `variables.tf` file, instead we define the variable (name) and its type and then run `terraform apply` command. Terraform will prompt us to enter the value for the variable.

```hcl
# variables.tf
variable "say_hello" {
  description = "The message to write to the file"
  type        = string
}

# main.tf
resource "local_file" hello{
  filename = "./hello.txt"
  content  = var.say_hello
  file_permission = "0700"
}
```

Now, when we run `terraform apply` after the initialization, it will prompt us to enter the value for the variable `name`. Just like this: 

![terminal screenshot](https://github.com/user-attachments/assets/ca2f3fb5-97ab-493a-a2d4-e8b84de6836f)

In fact it will prompt us to enter the value for each variable defined in the configuration file first and then apply the changes.


#### Command Line Flags

We can also use the `-var` flag to pass the values from the CLI:

```bash
terraform apply -var 'say_hello=Hello'
```

Another example:

```bash
terraform apply -var 'filename=./pets.txt' -var 'content=We love pets!'
```

#### Environment Variables

We can also use environment variables to pass values to the variables. The convention is to use `TF_VAR_<variable_name>` to set the value of the variable. Terraform automatically reads the environment variables with the `TF_VAR_` prefix.

Using the previous example:

```bash
export TF_VAR_say_hello="Hello"
terraform apply
```

Another example:

```bash
export TF_VAR_filename="./pets.txt"
export TF_VAR_content="We love pets!"
terraform apply
```

#### Variable Definition Files

Another way to pass values to the variables is to use a variable definition file. We can create a file `terraform.tfvars` and define the values for the variables in the file. We can also use a different name for the file but the file should have `.tfvars` extension.

```hcl
# terraform.tfvars
say_hello = "Hello"
```

Then we can run `terraform apply` command and Terraform will automatically read the values from the `terraform.tfvars` file.

So the complete example will look like this:

```hcl
# variables.tf
variable "say_hello" {
  description = "The message to write to the file"
  type        = string
}

# main.tf
resource "local_file" hello{
  filename = "./hello.txt"
  content  = var.say_hello
  file_permission = "0700"
}

# terraform.tfvars
say_hello = "Hello"
```

And then run `terraform apply` command.

The file with naming convention `terraform.tfvars` or `terraform.tfvars.json` or `*.auto.tfvars` or `*.auto.tfvars.json` will be automatically loaded by Terraform. If we pass file with a different naming convention we can use `-var-file` flag to pass the file:

```bash
terraform apply -var-file="custom.tfvars"
```

### Variable Definition Precedence

Terraform follows a specific order of precedence when resolving variable values:

Command Line Flags > Variable Definition Files (*.auto.tfvars) > Variable Definition Files (terraform.tfvars) > Environment Variables > Default Values

An example:
 
```hcl
resource local_file animal {
  filename = var.filename
}
```

This will be the order of precedence:

![Variable Precedence](https://github.com/user-attachments/assets/5d696331-0831-4643-8297-78866cca3b41)

So, the value passed via command line flag will take precedence over the value defined in the variable definition file, which will take precedence over the value defined in the environment variable, and so on.

## Resource (Reference) Attributes 

Resource attributes are the properties of a resource that can be used to reference or manipulate the resource. Each resource has a set of attributes that can be used to access information about the resource. For example, the `random_pet` resource has attributes like `id`, which gives the random pet name generated by Terraform and we can pass it to other resources. We use interpolation syntax `${}` like `${random_pet.my-pet.id}` to reference the attribute.

The interpolation syntax for this will be `${resource_type.resource_name.attribute}`.

For example, to reference the `id` attribute of the `random_pet` resource:

![Resource Attributes](https://github.com/user-attachments/assets/2bbd4efc-3852-4c4f-8c11-c644f142201a)

### Resource Dependencies

Terraform automatically manages resource dependencies based on the order of resource definitions in the configuration file. If one resource depends on another resource, Terraform will create the resources in the correct order. There are two types of dependencies:

1. **Implicit Dependencies**: These are dependencies that are automatically created by Terraform based on the order of resource definitions in the configuration file. For example, if resource A depends on resource B, Terraform will create resource B first and then create resource A.

Like we saw an an example above, the `random_pet` resource depends on the `local_file` resource. So, Terraform will create the `local_file` resource first and then create the `random_pet` resource. And when we destroy the resources, Terraform will destroy the `random_pet` resource first and then destroy the `local_file` resource.

2. **Explicit Dependencies**: But there are times where we don't don't use the attributes of other resources and we still want to create the resources in a specific order. Because they might be indirectly dependent on each other. In such cases, we can use the `depends_on` argument to define explicit dependencies between resources. A real life example of this is when we want to create a VPC and then create an EC2 instance in that VPC. We can use the `depends_on` argument to define the dependency between the resources.

For our local file and random pet example, we can define the explicit dependency like this:

```hcl
resource "local_file" "pet" {
  filename = "./pets.txt"
  content  = "We love pets!"
  depends_on = [random_pet.my-pet]
}

resource "random_pet" "my-pet" {
  prefix = "Mrs"
  separator = "."
  length = "1"
}
```

## Output Values

Output values are the values that are returned by the resources after they are created. We can use output values to display information about the resources or pass them to other resources. We can define output values in a separate file `outputs.tf` and use them to display information about the resources.

For example, to output the `id` attribute of the `random_pet` resource:

```hcl
# outputs.tf
output "pet_name" {
  value = random_pet.my-pet.id
  description = "The random pet name generated by Terraform"
}

# main.tf
resource "random_pet" "my-pet" {
  prefix = "Mrs"
  separator = "."
  length = "1"
}
```

We can use the `terraform output` command to display the output values. We can also specify the name of the output value to display only that value like `terraform output pet_name`.

## Terraform State

Terraform uses a state file to store information about the resources it manages. The state file keeps track of the current state of the infrastructure and is used to plan and apply changes to the resources. The state file is stored locally by default, but it can also be stored remotely in a backend like Terraform Cloud, AWS S3, or HashiCorp Consul. More like a blueprint of the infrastructure or single source of truth.

When we run `terraform apply` command, it will create a state file `terraform.tfstate` in the working directory. The state file contains information about the resources managed by Terraform, their attributes, dependencies, and other metadata. The state file is used to track changes to the infrastructure and ensure that the desired state is maintained. 

The state file contains the following information:

- Resource IDs: The unique identifiers of the resources managed by Terraform.
- Resource Attributes: The properties of the resources like IP addresses, DNS names, and other metadata.
- Resource Dependencies: The relationships between the resources and their dependencies.
- Provider Configuration: The configuration settings for the providers used by Terraform.
- Output Values: The values returned by the resources after they are created.
- Lock Information: The lock information to prevent concurrent modifications to the state file.
- Version Information: The version of Terraform used to create the state file.

Every time we run `terraform apply` command, Terraform checks the state file to determine the current state of the infrastructure and apply the changes accordingly. If the state file is deleted or corrupted, Terraform will not be able to manage the resources and we will have to recreate the resources.

![Terraform State](https://github.com/user-attachments/assets/8826a8ad-5165-411c-ae9d-8aa8eeb804b0)

### Benefits of Terraform State

- **State Management**: Terraform stores the state of the infrastructure in a state file, allowing you to manage, version, and share the state across teams.
- **Collaboration**: Terraform state allows multiple team members to work on the same infrastructure and track changes made by each team member.
- **Rollback**: Terraform state allows you to rollback changes to the infrastructure by reverting to a previous state.
- **Performance**: Terraform state improves performance by caching the state of the infrastructure and only applying changes when necessary.

We can make use of this command `terraform plan -refresh=false`. With this, Terraform generates an execution plan but skips refreshing the state of resources in the real-world infrastructure.

One thing to note is state file contains sensitive information like passwords, access keys, and other secrets. So, it is important to secure the state file and avoid storing it in version control systems like Git. We can use remote backends like Terraform Cloud, AWS S3, or HashiCorp Consul to store the state file securely.

:::important
So, we store the resources files like the `main.tf`, `variables.tf`, `outputs.tf` in version control systems like Git and store the state file in a secure remote backend.

Another important things that we don't make direct changes to state file. We use state management commands to manage those changes.
:::

## Terraform Commands

You can find the list of Terraform commands and their usage [here](./commands.md).

## Mutable and Immutable Infrastructure

In the context of infrastructure management, there are two approaches to managing infrastructure:

### Mutable Infrastructure 

Mutable infrastructure, servers and resources are updated or modified in place. Changes are made to the existing infrastructure to update software, apply patches, or make configuration changes. This approach is common in traditional infrastructure management where servers are manually configured and updated.

This approach can lead to configuration drift, security vulnerabilities, and inconsistencies between environments. It can also make it difficult to scale and manage large, complex environments.

### Immutable Infrastructure

Immutable infrastructure, servers and resources are treated as disposable and are replaced with new instances when changes are required. Instead of updating existing servers, new servers are created with the desired configuration and the old servers are destroyed. This approach is common in cloud-native environments and is used to ensure consistency, reliability, and scalability.

Terraform follows the immutable infrastructure approach where infrastructure is defined as code and changes are applied by creating new resources and destroying old resources.

For example if we modify the file permissions of the file we created, Terraform will create a new file with the updated permissions and destroy the old file.
```diff
# local.tf
resource "local_file" "pet" {
  filename = "./pets.txt"
  content  = "We love pets!"
+ file_permission = "0700"
- file_permission = "0600"
}
```

## Lifecycle Rules

Without Lifecycle Rules, Terraform will destroy the old resource and create a new one. But with Lifecycle Rules, we can control the behavior of the resources. We can use `create_before_destroy` to create the new resource before destroying the old one. This can be useful when we want to avoid downtime or data loss.

```hcl
# local.tf
resource "local_file" "pet" {
  filename = "./pets.txt"
  content  = "We love pets!"
  file_permission = "0700"

  lifecycle {
    create_before_destroy = true
  }
}
```

There are other lifecycle rules like `prevent_destroy` to prevent a resource from being destroyed, this is useful to protect resources from accidental deletion. Note `terraform destroy` will still destroy the resource, it will prevent the resource from being destroyed when you make changes to the configuration and apply the changes.

```hcl
# local.tf
resource "local_file" "pet" {
  filename = "./pets.txt"
  content  = "We love pets!"
  file_permission = "0700"

  lifecycle {
    prevent_destroy = true
  }
}
```

Another lifecycle rule is `ignore_changes` to ignore changes to specific attributes of a resource. This can be useful when you want to update the configuration of a resource without affecting the resource itself.

```hcl
# local.tf
resource "local_file" "pet" {
  filename = "./pets.txt"
  content  = "We love pets!"
  file_permission = "0700"

  lifecycle {
    ignore_changes = [file_permission]
  }
}
```

Or an AWS instance example:

```hcl
# ec2.tf
resource "aws_instance" "webserver" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"
  tags = {
    Name = "webserver"
  }

  lifecycle {
   ignore_changes = [tags]
  }
}
```

## Data Sources

Data sources are used to fetch information from external sources like APIs, databases, and other resources. Data sources allow you to query external systems and use the information in your Terraform configuration. Data sources are read-only and do not create or manage resources.

The reason for data sources is that every resources will not be created and managed by Terraform. There are some resources that are created and managed by other methods like scripts, APIs, or other tools. In such cases, we can use data sources to fetch information about those resources and use them in our Terraform configuration.

![Data Sources](https://github.com/user-attachments/assets/29c914b3-8582-4d88-9698-b52859e55b05)

For example,

```hcl
resource "local_file" "pet" {
  filename = "./pets.txt"
  content  = data.local_file.dog.content
}

data "local_file" "dog" {
  filename = "./dogs.txt"
}
```

The data read from data source is available under data object. We can use the data object to access the attributes of the data source. We can also use interpolation syntax to reference the data source attributes. `{data.data_type.data_name.attribute}`.

### Resource vs Data Source

In simple words resources are used to create and manage infrastructure resources like EC2 instances, S3 buckets, and databases. Data sources are used to fetch information from external sources like APIs, databases, and other resources. 
![Resource vs Data Source](https://github.com/user-attachments/assets/44c291ea-6b98-4812-a40a-49a0ca9dd564)

## Meta-Arguments

Meta-arguments are special arguments that can be used with resources and data sources to control their behavior. Meta-arguments are used to define dependencies, manage lifecycle, and configure the behavior of resources and data sources.

We have already seen some meta-arguments like `depends_on` and `lifecycle` in the previous sections. Here are some other meta-arguments:

### Count

The `count` meta-argument allows you to create multiple instances of a resource or data source. It takes an integer value and creates that many instances of the resource or data source. This can be useful when you want to create multiple instances of the same resource with different configurations.

For example, to create multiple local files:

```hcl
resource "local_file" "pet" {
  count    = length(var.filenames)
  content  = "We love pets!"
  filename = var.filenames[count.index]
}

variable "filenames" {
  default = [
    "/root/pets.txt",
    "/root/dogs.txt",
    "/root/cats.txt"
  ]
}
```

This will create three local files with filenames `pets.txt`, `dogs.txt`, and `cats.txt`. We can use `count.index` to reference the index of the resource. The index starts from 0.

### For Each

The `for_each` meta-argument allows you to create multiple instances of a resource or data source based on a map or set of strings. It takes a map or set of strings and creates an instance of the resource or data source for each key or value in the map or set. 

There was a pain point with `count` meta-argument that we have to use the index to reference the values. For example if we remove an element from the top of list, all the indexes will change and files will be recreated. But with `for_each` meta-argument, we can use the key to reference the values.

For example, to create multiple local files:

```hcl
# local.tf
resource "local_file" "pet" {
  for_each = var.files
  content  = "We love pets!"
  filename = each.value
}

# variables.tf
variable "files" {
  default = {
    pets = "/root/pets.txt"
    dogs = "/root/dogs.txt"
    cats = "/root/cats.txt"
  }
}
```

Here `for_each` will take the keys of the map and create local files with filenames `pets.txt`, `dogs.txt`, and `cats.txt`. We can use `each.value` to reference the values of the map. This don't allow `list` type. We can use `map` or `set` type. If we are using `list` type, we can convert it to `set` type using `toset()` function. Or make the type to `set` in the variable definition.

For example, to create multiple local files:

```hcl
# local.tf
resource "local_file" "pet" {
  for_each = toset(var.files)
  content  = "We love pets!"
  filename = each.value
}

# variables.tf
variable "files" {
  type = list(string)
  default = [
    "/root/pets.txt",
    "/root/dogs.txt",
    "/root/cats.txt"
  ]
}
```

## Version Constraints

Terraform uses version constraints to specify the version of the provider to use. Version constraints allow you to define the range of versions that are compatible with your configuration. By default, Terraform uses the latest version of the provider, but you can specify a specific version or range of versions using version constraints.

Version constraints are defined using the following syntax:

```hcl
terraform {
  required_providers {
    local = {
      source  = "hashicorp/local"
      version = "1.4.0"
    }
  }
}

resource "local_file" "pet" {
  filename = "./pets.txt"
  content  = "We love pets!"
}
```

We can use version constraints to specify the version of the provider to use. We can use 

- `>=` to specify the minimum version. `version = ">= 1.4.0"`
- `<=` to specify the maximum version. 
- `~>` to specify the compatible version. It will allow the minor version to change but not the major version. `version = "~> 1.4.0"`
- `!=` to specify the version to exclude.
- `=` to specify the exact version.
- `>` to specify the greater than version.
- `<` to specify the less than version.
- `<=` to specify the less than or equal to version.
- `>=` to specify the greater than or equal to version.

We can also mix and match the version constraints to specify the range of versions that are compatible with your configuration.

For example here we are specifying the version of local provider to be greater than or equal to 1.4.0 and less than 2.0.0 and excluding version 1.5.0.

```hcl
terraform {
  required_providers {
    local = {
      source  = "hashicorp/local"
      version = ">= 1.4.0, < 2.0.0, != 1.5.0" 
    }
  }
}
```

## Working With AWS

AWS is on of most popular cloud provider in the world. It has 100s of services from compute to AI. It's most global coverage and has the most data centers around the world. We need to create an AWS account to work with AWS services. We can use the free tier account to get started.

### IAM

When we create an AWS account, it a root user account and it has all the privileges to create, update, and delete resources. But it is not recommended to use the root user account to manage resources, like an Linux Root user account. We should create an IAM user account and use that account to manage resources. When we create an user we have two kind of access, `Programmatic Access` and `Console Access`. We can use `Programmatic Access` to access AWS services using APIs and SDKs. We can use `Console Access` to access AWS services using the AWS Management Console.

The only ideal use case for root user account is to create an IAM user account and manage billing and other account level settings.

### IAM Policies

IAM policies are used to define permissions for IAM users, groups, and roles. IAM policies are JSON documents that specify the actions, resources, and conditions that are allowed or denied. We can attach policies to IAM users, groups, and roles to grant or restrict access to AWS services and resources.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "s3:*",
      "Resource": "arn:aws:s3:::my-bucket/*"
    }
  ]
}
```

Some other policies are:

![IAM Policies](https://github.com/user-attachments/assets/98237fdf-8ee8-4fee-bf5b-661dd50ea6ee)

For example, to create a policy that allows access to S3 buckets:

### IAM Groups

IAM groups are used to group IAM users and apply policies to multiple users at once. For example, you can create a group called `developers` and attach a policy that allows access to EC2 instances. Then you can add IAM users to the `developers` group to grant them access to EC2 instances. It great when we have multiple users with the same permissions.

![IAM Groups](https://github.com/user-attachments/assets/be4e6aec-2fd8-4dde-b886-40953d556f1e)

### IAM Roles

Let's understand with an example what if an EC2 instance to interact with S3 bucket? Creating policy with not make the job done. We need to create an IAM role and attach the policy to the role. Then we can attach the role to the EC2 instance. This is called IAM roles. IAM roles are used to grant permissions to AWS services like EC2 instances, Lambda functions, and ECS tasks. IAM roles are used to define the permissions that are allowed or denied to the service.

IAM roles are not just limited to AWS services, we can also use IAM roles to grant permissions to external services like third-party applications and services. We can use IAM roles to grant temporary access to external services without sharing access keys or credentials.

![IAM Roles](https://github.com/user-attachments/assets/1ad28ae1-c24d-427f-a882-79e0662a4095)

### AWS CLI

AWS CLI is a command-line tool that allows you to interact with AWS services using the command line. You can use the AWS CLI to manage resources, configure services, and automate tasks. You can use the AWS CLI to perform operations like creating EC2 instances, managing S3 buckets, and configuring IAM policies.

After installing we have to configure the AWS CLI with the access key and secret key. We can use `aws configure` command to configure the AWS CLI. We can also use `--profile` flag to create multiple profiles.

![AWS CLI](https://github.com/user-attachments/assets/75a4aa54-eb7e-4768-8f52-3b835b4a3496)

All the commands can be found [here](./commands.md#AWS-Commands).

### AWS S3 (Simple Storage Service)

Amazon S3 is a cloud storage service that allows you to store and retrieve data from anywhere on the web. S3 is highly scalable, durable, and secure. It great for storing files like images, videos, and documents. But not suitable for storing operating system files or databases.

Data is store in form of buckets. Everything under a bucket is an object. We can use the AWS CLI to create, update, and delete S3 buckets. We can also use the AWS Management Console to manage S3 buckets.

![AWS S3](https://github.com/user-attachments/assets/d452bd9e-1851-4840-8177-32e8be7eb934)

Once the bucket is created we can access it via unique URL. We can also use the bucket to host static websites. it's in format of `http://<bucket-name>.<region>.amazonaws.com`. For eg. `http://my-bucket.ap-south-1.amazonaws.com`.

We can access the files in the bucket using the URL `http://<bucket-name>.<region>.amazonaws.com/<object-key>`. For eg. `http://my-bucket.ap-south-1.amazonaws.com/image.jpg`.

![AWS S3 bucket](https://github.com/user-attachments/assets/3c7fa44c-3d3a-4b3f-b7bf-e2565be6bb79)

Any object stored in the bucket has the object data and the Metadata. The metadata contains information about the object like owner, size, last modified date, etc, in key-value pairs.

![AWS S3 object](https://github.com/user-attachments/assets/37c4c48a-011b-4fd5-8092-addcda8901e3)

### DynamoDB (NoSQL Database)

Amazon DynamoDB is a fully managed NoSQL database service that provides fast and predictable performance with seamless scalability. DynamoDB is highly scalable, single-digit millisecond latency, and fully managed.

Even tho it is NoSQL database, it has a table structure. Each table has a primary key that uniquely identifies each item in the table. We can use the AWS CLI to create, update, and delete DynamoDB tables. We can also use the AWS Management Console to manage DynamoDB tables.

![DynamoDB](https://github.com/user-attachments/assets/99b55a27-dab6-4a9f-b8ae-2790fe7590b7)


### EC2 (Elastic Compute Cloud)

Amazon EC2 is a web service that provides resizable compute capacity in the cloud. EC2 instances are virtual servers that can run applications and workloads. EC2 instances can be launched in minutes and scaled up or down based on demand.

EC2 provides Amazon Machine Images (AMIs) that contain the operating system, software, and configuration settings for the instance. The instance type determines the hardware specifications like CPU, memory, and storage. It can be different types like general purpose, compute optimized, memory optimized, storage optimized, and GPU instances.

![EC2](https://github.com/user-attachments/assets/bf28251f-7f29-4b2a-8400-14993f7e914)

The most common type is T2 General Purpose instances. It is suitable for web servers, small databases, and development environments. It has burstable performance and is cost-effective. It comes in a variety of sizes like:

![EC2 Instance Types](https://github.com/user-attachments/assets/2d0a3b31-18e2-4fc6-a0cc-5b67b6033d11)

The persistent storage for EC2 instances is provided by Elastic Block Store (EBS) volumes. EBS volumes are network-attached storage volumes that can be attached to EC2 instances. EBS volumes are durable, scalable, and high-performance. There are a variety of EBS volume tpo choose from on the basis of the data we want to persist. Some of them are:

![EBS Volume Types](https://github.com/user-attachments/assets/de915648-d65f-4352-aa80-92e33faa5d94)

We can pass user data to the EC2 instance to run scripts or commands when the instance is launched. We can use user data to install software, configure settings, and run custom scripts.

![EC2 User Data](https://github.com/user-attachments/assets/2794bb05-4e1a-4c06-8548-cf7b1e513042)

After everything is set up, we can access the EC2 instance using the public IP address or public DNS name. We can use SSH to connect to the instance and run commands. We can also use the AWS Management Console to manage EC2 instances.

## AWS and Terraform

Terraform uses the AWS provider to interact with AWS services. The AWS provider allows you to define resources like EC2 instances, S3 buckets, and IAM policies in your Terraform configuration. 

Creating an IAM user with Terraform:

```hcl
# main.tf
provider "aws" {
  region = "us-east-1"
}

resource "aws_iam_user" "admin-user" {
  name = "admin"
  tags = {
    Description = "Technical Team Leader"
  }
}
```

```bash
# .aws/credentials
[default]
aws_access_key_id = AKIAIOSFODNN7EXAMPLE
aws_secret_access_key = je7MtGbClwBF/2Zp9Utk/h3yCo8nvbEXAMPLEKEY
```

Terraform will automatically use the credentials from the `~/.aws/credentials` file to authenticate with AWS. We can also use environment variables to set the credentials. We can also use `profile` argument to specify the profile to use.

```hcl
provider "aws" {
  region = "us-east"
  profile = "default"
}
```

Another way to use `export` command to set the environment variables:

```bash
export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
export AWS_SECRET_ACCESS_KEY=je7MtGbClwBF/2Zp9Utk/h3yCo8nvbEXAMPLEKEY
export AWS_REGION=us-east-1
```

```hcl
resource "aws_iam_user" "admin-user" {
  name = "lucy"
  tags = {
    Description = "Technical Team Leader"
  }
}
```

Now we don't need to specify the providers in the configuration file. Terraform will automatically use the environment variables to authenticate with AWS.

Attaching a policy to the IAM user:

```hcl
# main.tf
resource "aws_iam_user" "admin-user
  name = "lucy"
  tags = {
    Description = "Technical Team Leader"
  }

resource "aws_iam_policy" "admin-user-policy" {
  name = "AdminUsers"
  policy = <<EOF
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Action": "*",
        "Resource": "*"
      }
    ]
  }
  EOF
}

resource "aws_iam_user_policy_attachment" "lucy-admin-access" {
  user = aws_iam_user.admin-user.name
  policy_arn = aws_iam_policy.admin-user-policy.arn
}
```

What we did is first we created an IAM user, then we created an IAM policy with full access to all the resources, and then we attached the policy to the IAM user. We use heredoc syntax to define the policy (`<<EOF`). It's not mandatory to use `EOF`, we can use any string. We can use `terraform plan` to see the changes and `terraform apply` to apply the changes.

In `aws_iam_user` resource, we are using the `name` argument to specify the name of the IAM user and the `tags` argument to add tags to the user. We use the `aws_iam_policy` resource to create an IAM policy. We are using the `name` argument to specify the name of the policy and the `policy` argument to specify the policy document. Then we used the `aws_iam_user_policy_attachment` resource to attach the policy to the user. We are using the `user` argument to reference the IAM user and the `policy_arn` argument to reference the ARN of the policy.

![terraform apply](https://github.com/user-attachments/assets/3f661e06-dc17-4bc6-a829-810cb4fd9ca5)

Another way to create an IAM policy is to create a JSON file and use the `file()` function to read the file.

```hcl
# main.tf
resource "aws_iam_user" "admin-user" {
  name = "lucy"
  tags = {
    Description = "Technical Team Leader"
  }
}

resource "aws_iam_policy" "admin-user-policy" {
  name = "AdminUsers"
  policy = file("policy.json")
}

resource "aws_iam_user_policy_attachment" "lucy-admin-access" {
  user = aws_iam_user.admin-user.name
  policy_arn = aws_iam_policy.adminUser.arn
}
```

```json
# policy.json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*"
    }
  ]
}
```

We can proceed with `terraform plan` and `terraform apply` to create the IAM user and attach the policy to the user.

Another example to create IAM user for a list of users:

```hcl
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
```

Here we are using the `count` meta-argument to create multiple instances of the resource. We are using the `count.index` to reference the index of the resource. The index starts from 0. We can use the `length` function to get the length of the list.

### S3

Here we are creating an S3 bucket, uploading a file to the bucket, and creating a bucket policy to allow access to the bucket.

```hcl
# main.tf
resource "aws_s3_bucket" "finance" {
  bucket = "finance-21092020 # optional, otherwise Terraform will create a unique name
  tags = {
    Name = "Finance and Payroll"
  }
}

resource "aws_s3_object" "finance-2020" {
  content = "/root/finance/finance-2020.doc"
  key = "finance-2020.doc"
  bucket = aws_s3_bucket.finance.id # reference to the 
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
```

Here, `aws_s3_bucket` resource is used to create an S3 bucket, `aws_s3_object` resource is used to upload a file to the bucket, and `aws_s3_bucket_policy` resource is used to create a bucket policy to allow access to a group of users to the bucket, which is defined using the `aws_iam_group` data source.

In the `aws_s3_bucket` resource, we are using the `bucket` argument to specify the name of the bucket and the `tags` argument to add tags to the bucket. We can use the `aws_s3_object` resource to upload a file to the bucket. We are using the `content` argument to specify the file to upload, the `key` argument to specify the name of the object, and the `bucket` argument to reference the bucket.

In the `aws_s3_bucket_policy` resource, we are using the `bucket` argument to reference the bucket, the `policy` argument to specify the bucket policy, 

The `data` block to fetch information about an IAM group. We are using the `aws_iam_group` data source to fetch information about an IAM group named `finance-analysts`. We are using the `group_name` argument to specify the name of the group. We can use the `arn` attribute to reference the ARN of the group.

:::Note
The bucket naming should not contain uppercase letters, underscores, or special characters. `ss_aa` is not allowed due to DNS compatibility.
:::

### DynamoDB

Here we are creating a DynamoDB table with a primary key and a sort key.

```hcl
# main.tf
resource "aws_dynamodb_table" "cars" {
  name = "cars"
  hash_key = "VIN"
  billing_mode = "PAY_PER_REQUEST"
  attribute {
    name = "VIN"
    type = "S"
  }
}
```

Here `hash_key` is the primary key and `attribute` is the sort key. We can use `billing_mode` to specify the billing mode for the table. We can use `PAY_PER_REQUEST` for on-demand capacity mode and `PROVISIONED` for provisioned capacity mode. In the attribute block, we can specify the name and type of the attribute. Here we are using `S` for string type. We can also use `N` for number type and `B` for binary type.

To insert data into the table, we can use the `aws_dynamodb_table_item` resource.

```hcl
# main.tf
resource "aws_dynamodb_table" "cars" {
  name = "cars"
  hash_key = "VIN"
  billing_mode = "PAY_PER_REQUEST"
  attribute {
    name = "VIN"
    type = "S"
  }
}

resource "aws_dynamodb_table_item" "car-items" {
  table_name = aws_dynamodb_table.cars.name
  hash_key = aws_dynamodb_table.cars.hash_key
  item = <<EOF
  {
    "Manufacturer": {"S": "Toyota"},
    "Model": {"S": "Corolla"},
    "Year": {"N": "2020"},
    "VIN": {"S": "JH4KA3240JC000000"}
  }
EOF
}
```

Here we are using the `aws_dynamodb_table_item` resource to insert data into the table. We are using the `table_name` to specify the name of the table, `hash_key` to specify the primary key, and `item` to specify the data to insert. We are using the `S` type for string, `N` type for number, and `B` type for binary.

### EC2 Instance

To create an EC2 instance, we can use the `aws_instance` resource. We can specify the AMI, instance type, key pair, and security group for the instance.

```hcl
# main.tf
resource "aws_instance" "webserver" {
  ami = "ami-0edab43bfb4c0e9b2"
  instance_type = "t2.micro"
  tags = {
    Name = "webserver"
    Description = "An Nginx web server on Ubuntu"
  }
  user_data = <<-EOF
  #!/bin/bash
  sudo apt update
  sudo apt install -y nginx
  systemctl enable nginx
  systemctl start nginx
  EOF
  key_name = aws_key_pair.web.id
  vpn_security_group_ids = [aws_security_group.ssh-access.id]
}

resource "aws_key_pair" "web" {
  public_key = file("/root/.ssh/id_rsa.pub")
}

resource "aws_security_group" "ssh-access" {
  name = "ssh-access"
  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

output "public_ip" {
  value = aws_instance.webserver.public_ip
}

# provider.tf
provider "aws" {
  region = "us-west-1"
}
```

In `user_data` value instead of heredoc syntax, we can also use the `file()` function to read the user data from a file.

```hcl
.
.
user_data = file("user-data.sh")
.
.
```

So, the way above config is working is that we are creating an EC2 instance with an Nginx web server on Ubuntu. We are using the `aws_instance` resource to create the instance, `aws_key_pair` resource to create the key pair, this will help us to SSH into the instance, `aws_security_group` resource to create the security group to allow SSH access to the instance, and `output` block to display the public IP address of the instance.

In `aws_instance` resource, we are using the `ami` that's the Amazon Machine Image, `instance_type` that's the type of the instance, `tags` to add tags to the instance, `user_data` to run the script when the instance is launched, `key_name` to specify the key pair to use to SSH into the instance, and `vpn_security_group_ids` to specify the security group to use for the instance (expected to be a list of security group IDs). 

In  `aws_key_pair` resource, we are using the `public_key` to specify the public key to use for the key pair. 

In `aws_security_group` resource, we are using the `name` to specify the name of the security group, `ingress` block to specify the inbound rules for the security group, `from_port` and `to_port` to specify the port range, `protocol` to specify the protocol, and `cidr_blocks` to specify the IP range to allow access. This is very similar to the security group in the AWS Management Console.

In `output` block, we are using the `value` to specify the value to display when the `terraform apply` command is run. We can use the `terraform output` command to display the value.

## Terraform Provisioners

Terraform provisioners are used to run scripts or commands on the local machine or on the remote machine. Provisioners are used to install software, configure settings, and run custom scripts. Provisioners are used to perform tasks that are not supported by Terraform resources.

There are two types of provisioners:

- Local-exec provisioner: It runs scripts or commands on the local machine where Terraform is running. It is used to install software, configure settings, and run custom scripts on the local machine.

- Remote-exec provisioner: It runs scripts or commands on the remote machine where the resource is created. It is used to install software, configure settings, and run custom scripts on the remote machine.

### Local-exec Provisioner

```hcl
resource "aws_instance "webserver" {
  ami = "ami-0edab43bfb4c0e9b2"
  instance_type = "t2.micro"
  provisioner "local-exec" {
    command = "echo ${aws_instance.webserver.public_ip} Created > instance.txt"
  }
}
```

In the `aws_instance` resource, we are using the `provisioner` block to specify the provisioner type, `local-exec` to specify the type of provisioner, and `command` to specify the command to run on the local machine. We can use the `${}` syntax to reference the attributes of the resource. We can use the `file()` function to read the public IP address from a file.

#### Destroy Time Provisioner

We can also run the provisioner when the resource is destroyed. We can use the `when` argument to specify the time to run the provisioner. We can use `create` to run the provisioner when the resource is created and `destroy` to run the provisioner when the resource is destroyed.

```hcl
resource "aws_instance "webserver" {
  ami = "ami-0edab43bfb4c0e9b2"
  instance_type = "t2.micro"
  provisioner "local-exec" {
    command = "echo ${aws_instance.webserver.public_ip} Created > instance.txt"
  }
  provisioner "local-exec" {
    when = destroy
    command = "echo ${aws_instance.webserver.public_dns} Destroyed > instance.txt"
  }
}
```

Above we are using the `when` argument to specify the time to run the provisioner. We are using `create` to run the provisioner when the resource is created and `destroy` to run the provisioner when the resource is destroyed. 

#### Provisioner Failure Behavior

The default behavior of Terraform is to stop (the applying state) the execution when the provisioner fails. The failure can occur due to an error in the script or command, file path, permissions, or network issues, etc. To change the default behavior of terraform when the provisioner fails, we can use `on_failure` argument. We can use `continue` to continue the execution, `fail` to stop the execution, and `ignore` to ignore the failure.

```hcl
resource "aws_instance "webserver" {
  ami = "ami-0edab43bfb4c0e9b2"
  instance_type = "t2.micro"
  provisioner "local-exec" {
    command = "echo ${aws_instance.webserver.public_ip} Created > instance.txt"
    on_failure = continue
  }
  provisioner "local-exec" {
    when = destroy
    command = "echo ${aws_instance.webserver.public_dns} Destroyed > instance.txt"
    on_failure = fail
  }
}
```

So here we are using the `on_failure` argument to specify the behavior of Terraform when the provisioner fails. We are using `continue` to continue the execution when the provisioner fails and `fail` to stop the execution when the provisioner fails.

### Remote-exec Provisioner

```hcl
resource "aws_instance" "webserver" {
  ami = "ami-0edab43bfb4c0e9b2"
  instance_type = "t2.micro"
  provisioner "remote-exec" {
    inline = [
      "echo 'Hello, World!'",
      "sudo apt update",
      "sudo apt install -y nginx",
      "systemctl enable nginx",
      "systemctl start nginx"
    ]
  }
  connection {
    type = "ssh"
    user = "ubuntu
    private_key = file("/root/.ssh/id_rsa")
    host = self.public_ip
  }
  key_name = aws_key_pair.web.id
  vpc_security_group_ids = [aws_security_group.ssh-access.id]
}

resource "aws_security_group" "ssh-access" {
  name = "ssh-access"
  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_key_pair" "web" {
  public_key = file("/root/.ssh/id_rsa.pub")
}
```

In the `aws_instance` resource, we are using the `provisioner` block to specify the provisioner type, `inline` to specify the commands to run on the remote machine, and `connection` block to specify the connection details like the type of connection, user, private key, and host. We are using the `self` object to reference the current resource. We can use the `file()` function to read the private key from a file.

### Provisioner Best Practices

Provisioner are great but they are not recommended to use. As we can run any command on the instance, it hard to have a structure and framework to manage the scripts. It's better to use arguments like `user_data` and `cloud_init` to run scripts on the instance. Provisioners are used when there is no other way to run the scripts.

Use custom AMIs to pre-install software and configure settings. Use user data to run scripts when the instance is launched. Like use an Ubuntu AMI with Nginx pre-installed.



## Remote State (Remote Backend)

Terraform stores the state file locally by default, but it can also store the state file remotely in a backend like Terraform Cloud, AWS S3, or HashiCorp Consul. Storing the state file remotely allows you to share the state file across teams, manage state locking, and store the state securely.

It becomes a necessity to store it remotely when working in a team. To avoid conflicts and to have a single source of truth. 

### State Locking

State locking is used to prevent concurrent modifications to the state file. When multiple users are working on the same infrastructure, state locking ensures that only one user can modify the state file at a time. This prevents conflicts and ensures that changes are applied in the correct order. This is not supported by version control systems like Git.

![State Locking](https://github.com/user-attachments/assets/52bdbf3f-aed8-444b-9407-4de323331c4e)

When a remote backend is used, Terraform automatically manages state locking, upload and download the state file, and store the state securely.

### AWS S3 and DynamoDB as Remote Backend

We can use AWS S3 and DynamoDB as a remote backend to store the state file. We use S3 to store the state file and DynamoDB to manage state locking. We can use the `backend` block to configure the remote backend in the configuration file. State locking is optional but it is recommended to use state locking to prevent concurrent modifications to the state file.

```hcl
# terraform.tf
terraform {
  backend "s3" {
    bucket = "pradumna-terraform-state-bucket"
    key = "finance/terraform.tfstate"
    region = "us-west-1"
    dynamodb_table = "state-locking"
  }
}

# main.tf
resource "local_file" "pet" {
  filename = "./pets.txt"
  content  = "We love pets!"
}
```

Once we initialize and the apply the changes, the state file will be stored in the S3 bucket and the state locking will be managed by DynamoDB table. After this we can delete the local state file. Now every time we run `terraform apply` command, Terraform will automatically download the state in memory, lock the state file, apply the changes, and upload the state file back to the S3 bucket.

![Remote State](https://github.com/user-attachments/assets/8579996c-640d-4649-a1f0-4f6d563d52e7)

### Terraform state commands

All the commands can be found [here](./commands.md).

## Terraform Taint 

When terraform apply get failed, it marks the resource as tainted. It means that the resource is not in the desired state. We can verify the resources is tainted by running `terraform plan` command. We can use `terraform taint` command to mark the resource as tainted. So when we mark a resource as tainted, Terraform will destroy the resource and recreate it. We can use `terraform untaint` command to undo the tainted mark from the resource.

The benefit of using `terraform taint` is that it allows us to configure the behavior of resources when they are not in the desired state. For example we are running an EC2 instance with nginx v1.16 and someone manually updated it to v1.18. We can mark the resource as tainted and Terraform will destroy the instance and recreate it with the desired state.

```bash
terraform taint aws_instance.webserver
```

## Terraform Debugging

Terraform already provides a lot of information in the output. But sometimes we need more information to debug the issues. 

We can use the `TF_LOG` environment variable to set the log level. We can use `TRACE`, `DEBUG`, `INFO`, `WARN`, `ERROR`, and `PANIC` log levels. 

![Terraform Log Levels](https://github.com/user-attachments/assets/441b4c46-0476-4b38-a95c-ababdd853ce9)

For example, to set the log level to `DEBUG`:

```bash
export TF_LOG=DEBUG
terraform plan
```

We can also use `TF_LOG_PATH` environment variable to specify the log file. We can use `TF_LOG_PATH` to specify the log file path. For example, to set the log file path to `/tmp/terraform.log`:

```bash
export TF_LOG_PATH=/tmp/terraform.log
terraform plan
```

we need to also set the `TF_LOG` environment variable to a level to enable logging to the file.

To disable logging we can unset the `TF_LOG` environment variable:

```bash
unset TF_LOG
```

## Terraform Import

So, there will not always be the case that every resource is being created and managed by Terraform. Sometimes we have to import it. Terraform import is used to import existing resources into Terraform state. It is used to import resources that were created outside of Terraform. For example some resources were created manually or using the AWS Management Console or using another Iac tool like Ansible or CloudFormation.

![Terraform Import](https://github.com/user-attachments/assets/2aee6966-4b1e-4c36-af55-a0b65f3c342c)

we have used `data` block, but that can only be used to fetch information about the resources. We can't manage the resources using the `data` block. With  `terraform import` command to import resources into Terraform state. We can use `terraform show` command to get the resource details.

```bash
terraform import <resource_type>.<resource_name> <resource_id>
terraform import aws_instance.webserver i-0c1c4b7b3b3e3b3b3
```

When we run this command, we will see an error. Terraform will not update the resource file.

To fix this, we have to manually add a resource block to the configuration file. We can define an empty resource block like this:

```hcl 
resource "aws_instance" "webserver-2"
# (resource block)
```

Now we can run the `terraform import` command again. This time Terraform will import the resource into the state file. We can now fill in the attributes in the resource block by looking into the `terraform show` command.

```hcl
resource "aws_instance" "webserver-2" {
  ami = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"
  tags = {
    Name = "webserver-2"
  }
}
```

Now, this resource is managed by Terraform. We can use `terraform plan` and `terraform apply` commands to manage the resource.


## Terraform Modules

Terraform modules are used to organize and reuse Terraform configurations. Modules are used to define reusable components like EC2 instances, S3 buckets, and IAM policies. Modules are used to create reusable components that can be shared across teams and projects. 

![Terraform Modules](https://github.com/user-attachments/assets/f7c060b4-6ad9-4c03-8c9a-80867cde33a0)

### Creating a Module

To create a module, we create a dir name `modules` and inside that dir we create a dir with the name of the module. Inside that dir we create resource files. 

Let's take an example to understand this better. We have an payroll app and we want to deploy it across different regions. We can create a module for the payroll app. Our structure will look like this:

```
terraform-project
├── modules
│   └── payroll-app
│       ├── app_server.tf
│       ├── s3_bucket.tf
│       ├── dynamodb_table.tf
```

The `app.server.tf` file contains the EC2 instance configuration, the `s3_bucket.tf` file contains the S3 bucket configuration, and the `dynamodb_table.tf` file contains the DynamoDB table configuration. 

```hcl
# app_server.tf
resource "aws_instance" "app-server" {
  ami = "var.ami"
  instance_type = "t2.micro"
  tags = {
    Name = "${var.app_region}-app-server"
  }
  depends_on = [aws_dynamodb_table.payroll_db,
                aws_s3_bucket.payroll_data
                ]
}

# s3_bucket.tf
resource "aws_s3_bucket" "payroll_data" {
  bucket = "${var.app_region}-${var.bucket}"
}

# dynamodb_table.tf
resource "aws_dynamodb_table" "payroll_db" {
  name = "user_data"
  hash_key = "EmployeeID"
  billing_mode = "PAY_PER_REQUEST"
  attribute {
    name = "EmployeeID"
    type = "N"
  }
}

# variables.tf
variable "ami" {
  type = string
}

variable "app_region" {
  type = string
}

variable "bucket" {
  type = "flexit-payroll-01-02"
}
```

In the `app_server.tf` file, we are using the `aws_instance` resource to create an EC2 instance. We are using the `ami` and `instance_type` arguments to specify the AMI and instance type, the `tags` argument to add tags to the instance, and the `depends_on` argument to specify the dependencies of the instance. We are using the `aws_dynamodb_table` and `aws_s3_bucket` resources to create the DynamoDB table and S3 bucket.

Also, we have intentionally hard-coded the values of the resources and use variables to pass the values when we use the module.

Now one the module is created, we can use it in the main configuration file. We can use the `module` block to use the module. 

### Using a Module

Now as we have created the module, we can use it in the main configuration file. We can use the `module` block to use the module. For example, we want to deploy the payroll app in the us region. We can use the module like this:

First, we will create a dir under the root dir with `us-payroll-app` under `terraform-project` dir. Inside that dir we will create a `main.tf` file. Now the structure will look like this:

```
terraform-project
├── us-payroll-app
│   └── main.tf
├── modules
│   └── payroll-app
│       ├── app_server.tf
│       ├── s3_bucket.tf
│       ├── dynamodb_table.tf
```

Our `main.tf` content will look like this:

```hcl
# /us-payroll-app/main.tf
provider "aws" {
  region = "us-west-1"
}

module "us_payroll" {
  source = "../modules/payroll-app"
  app_region = "us-east-1"
  ami = "ami-0c55b159cbfafe1f0"
}
```

In the `main.tf` file, we are using the `provider` block to specify the provider configuration, and the `module` block to use the module. We are using the `source` argument to specify the path to the module, and the `app_region` and `ami` arguments to pass the values to the module. The `us-payroll-app` directory is the root module and the `payroll-app` directory is the child module. 

Additionally we can also specify the `bucket` variable in the `main.tf` as we set it as a variable in the module, otherwise it will use the default value. Now we can use the `terraform init`, `terraform plan`, and `terraform apply` commands to deploy the payroll app in the us region.

Now we want to deploy to UK region as well. We can create a new dir `uk-payroll-app` under the root dir and create a `main.tf` file. The structure will look like this:

```
terraform-project
├── uk-payroll-app
│   └── main.tf
├── us-payroll-app
│   └── main.tf
├── modules
│   └── payroll-app
│       ├── app_server.tf
│       ├── s3_bucket.tf
│       ├── dynamodb_table.tf
```

The `main.tf` file will look like this:

```hcl
# /uk-payroll-app/main.tf
provider "aws" {
  region = "eu-west-1"
}

module "uk_payroll" {
  source = "../modules/payroll-app"
  app_region = "eu-west-2"
  ami = "ami-0c55b159cbfafe1f0avm"
}
```

Now we can use the `terraform init`, `terraform plan`, and `terraform apply` commands to deploy the payroll app in the UK region (london).

Since we are using modules the addressing syntax also gets changed. We use `module.<module_name>.<resource_type>.<resource_name>` to reference the resources in the module. For example, to reference the EC2 instance in the `payroll-app` module, we  use `module.us_payroll.aws_instance.app-server` or `module.uk_payroll.aws_instance.app-server` to reference the EC2 instance in the `us-payroll-app` and `uk-payroll-app` modules.

### Using Module from Registry

We can also use modules from the Terraform Registry. We can use the `source` argument to specify the path to the module in the Terraform Registry.

There are two types of modules in the Terraform Registry:

- Official modules (verified): These are modules created by HashiCorp and maintained by the Terraform team. They are verified and tested by HashiCorp.
- Community modules: These are modules created by the community and shared on the Terraform Registry. They are not verified by HashiCorp.

A modules can also have submodules. We can use the `source` argument to specify the path to the module in the Terraform Registry. For example, aws security group module:

```hcl
module "security_group_ssh" {
  source = "terraform-aws-modules/security-group/aws/modules/ssh"
  version = "3.16.0"
  vpc_id = "vpc-12345678"
  ingress_cidr_blocks = ["10.10.0.0/16"]
}
```

## Terraform Functions, Expressions, and Operators

Terraform provides a number of built-in functions that can be used to manipulate strings, numbers, lists, and maps. Functions are used to perform operations like string manipulation, arithmetic operations, and data transformation. Functions are used to generate dynamic values, format strings, and transform data.

We have already seen some of the functions like `file()`, `length()`, `count.index`, `var.<variable_name>`, `self.<attribute_name>`, etc.

We can use `terraform console` to test the functions. Like this:

![Terraform Console](https://github.com/user-attachments/assets/5e999dd2-ca82-4ff6-b880-258ae7e23da4)

There are many categories of functions like Number functions, String functions, List functions, Map functions, Filesystem functions, Date and Time functions, and Encoding functions.

### Numeric Functions

Numeric functions are used to perform arithmetic operations like addition, subtraction, min, max, and rounding.

- `max(1, 2, 3): Returns the maximum value from the list of numbers. In this case, it will return 3.
- `min(1, 2, 3)`: Returns the minimum value from the list of numbers. In this case, it will return 1.
- `ceil(10.1)`: Returns the smallest integer value greater than or equal to the number. In this case, it will return 11.
- `floor(10.9)`: Returns the largest integer value less than or equal to the number. In this case, it will return 10.

### String Functions

String functions are used to perform string manipulation operations like concatenation, substring, and formatting.

- `split(",", "ami-xyz, AMI-ABC, ami-pqr")`: Splits the string into a list of substrings based on the delimiter. In this case, it will return ["ami-xyz", "AMI-ABC", "ami-pqr"].
- `lower("AMI-XYZ")`: Converts the string to lowercase. In this case, it will return "ami-xyz".
- `upper("ami-xyz")`: Converts the string to uppercase. In this case, it will return "AMI-XYZ".
- `title("ami-xyz")`: Converts the string to title case. In this case, it will return "Ami-Xyz".
- `substr("ami-xyz", 0, 3)`: Returns a substring from the string. In this case, it will return "ami".
- `join(",", ["ami-xyz", "ami-abc", "ami-pqr"])`: Joins the list of strings into a single string. In this case, it will return "ami-xyz,ami-abc,ami-pqr".

### Collection Functions

Collection functions are used to perform operations on lists and maps like filtering, sorting, and merging.

- `length(["ami-xyz", "ami-abc", "ami-pqr"])`: Returns the length of the list. In this case, it will return 3.
- `index(["ami-xyz", "ami-abc", "ami-pqr"], "ami-abc")`: Returns the index of the element in the list. In this case, it will return 1.
- `element(["ami-xyz", "ami-abc", "ami-pqr"], 1)`: Returns the element at the index in the list. In this case, it will return "ami-abc".
- `contains(["ami-xyz", "ami-abc", "ami-pqr"], "ami-abc")`: Returns true if the element is present in the list. In this case, it will return true.

### Map Functions

Map functions are used to perform operations on maps like merging, filtering, and transforming.

- `keys({"ami-xyz": "ami-abc", "ami-abc": "ami-pqr"})`: Returns the keys of the map. In this case, it will return ["ami-xyz", "ami-abc"].
- `values({"ami-xyz": "ami-abc", "ami-abc": "ami-pqr"})`: Returns the values of the map. In this case, it will return ["ami-abc", "ami-pqr"].
- `lookup({"ami-xyz": "ami-abc", "ami-abc": "ami-pqr"}, "ami-abc")`: Returns the value of the key in the map. In this case, it will return "ami-pqr".
- `lookup({"ami-xyz": "ami-abc", "ami-abc": "ami-pqr"}, "ami-xyz", "default")`: Returns the value of the key in the map or the default value. In this case, it will return "ami-abc".

### Numeric Operators

Numeric operators are used to perform arithmetic operations like addition, subtraction, multiplication, and division.

- `1 + 2`: Addition
- `5 - 3`: Subtraction
- `2 * 3`: Multiplication
- `6 / 2`: Division
  
### Equality Operators

Equality operators are used to compare values like equal, not equal, greater than, and less than.

- `1 == 2`: Equal
- `1 != 2`: Not equal
- `1 > 2`: Greater than
- `1 < 2`: Less than
- `1 >= 2`: Greater than or equal
- `1 <= 2`: Less than or equal

### Logical Operators

Logical operators are used to perform logical operations like and, or, and not.

- 8 > 5 && 5 < 3: And (both conditions must be true)
- 8 > 5 || 5 < 3: Or (at least one condition must be true)
- !true: Not (negates the value)


### Conditional Operator

- condition ? true_value : false_value: Ternary operator (if the condition is true, return the true value, otherwise return the false value) `10 > 5 ? "true" : "false"`

```hcl
# main.tf
resource "random_password" "password" {
  length = var.length < 8 ? 8 : var.length
}
output "password" {
  value = random_password.password.result
}

# variables.tf
variable "length" {
  type = number
}
```

So, here if the `var.length` is less than 8, the password length will be 8, otherwise it will be the value of `var.length`.

## Terraform Workspaces

Terraform workspaces are used to manage multiple environments like development, staging, and production. Workspaces are used to create separate environments for different teams, projects, and regions. Workspaces are used to manage state files for different environments. Workspaces are used to create separate state files for different environments.

We can use the `terraform workspace` command to manage workspaces. We can use the `terraform workspace new` command to create a new workspace, the `terraform workspace list` command to list the workspaces, the `terraform workspace select` command to select a workspace, and the `terraform workspace delete` command to delete a workspace.

We have modify our configuration is a way that it can be used for multiple environments. 

```hcl
# main.tf
resource "aws_instance" "webserver" {
  ami = lookup(var.ami, terraform.workspace)
  instance_type = var.instance_type
  tags = {
    Name = terraform.workspace
  }
}

# variables.tf
variable "region" {
  default = "ca-central-1"
}

variable "instance_type" {
  default = "t2.micro"
}

variable "ami" {
  type = map
  default = {
    "ProjectA" = "ami-0c55b159cbfafe1f0"
    "ProjectB" = "ami-oe4b3bfb4c0e9b2"
  }
}
```

Now if we create a new workspace like `ProjectA` and run the `terraform apply` command, it will use the `ami-0c55b159cbfafe1f0` AMI. If we create a new workspace like `ProjectB` and run the `terraform apply` command, it will use the `ami-oe4b3bfb4c0e9b2` AMI. We can use the `terraform workspace list` command to list the workspaces and the `terraform workspace select` command to select a workspace.

The way terraform manages the the state files for different workspaces is that it creates a separate directory for each workspace and stores the state file in that directory. And all the directories are named as `terraform.tfstate.d/<workspace_name>`.

