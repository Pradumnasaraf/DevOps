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

We can change list type to`number`, `bool`, or `any`. It will fail if we try to use a different type of value.

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

1. **Explicit Dependencies**: But there are times where we don't don't use the attributes of other resources and we still want to create the resources in a specific order. Because they might be indirectly dependent on each other. In such cases, we can use the `depends_on` argument to define explicit dependencies between resources. A real life example of this is when we want to create a VPC and then create an EC2 instance in that VPC. We can use the `depends_on` argument to define the dependency between the resources.

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

### Mutable Infrastructure

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

#### Lifecycle Rules

Without License Rules, Terraform will destroy the old resource and create a new one. But with Lifecycle Rules, we can control the behavior of the resources. We can use `create_before_destroy` to create the new resource before destroying the old one. This can be useful when we want to avoid downtime or data loss.

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

This will create three local files with filenames `pets-0.txt`, `pets-1.txt`, and `pets-2.txt`.

### For_each

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

Once the bucket is created we can access it via unique URL. We can also use the bucket to host static websites. it's in format of `http://<bucket-name>.<region>.amazonaws.com`. For eg. `http://my-bucket.s3.ap-south-1.amazonaws.com`.

We can access the files in the bucket using the URL `http://<bucket-name>.<region>.amazonaws.com/<object-key>`. For eg. `http://my-bucket.s3.ap-south-1.amazonaws.com/index.html`.

![AWS S3 bucket](https://github.com/user-attachments/assets/3c7fa44c-3d3a-4b3f-b7bf-e2565be6bb79)

Any object stored in the bucket has the object data and the Metadata. The metadata contains information about the object like owner, size, last modified date, etc, in key-value pairs.

![AWS S3 object](https://github.com/user-attachments/assets/37c4c48a-011b-4fd5-8092-addcda8901e3)


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

# .aws/credentials
[default]
aws_access_key_id = AKIAIOSFODNN7EXAMPLE
aws_secret_access_key = je7MtGbClwBF/2Zp9Utk/h3yCo8nvbEXAMPLEKEY
```

Terraform will automatically use the credentials from the `~/.aws/credentials` file to authenticate with AWS. We can also use environment variables to set the credentials. We can also use `profile` argument to specify the profile to use.

```hcl
provider "aws" {
  region = "us-east
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
resource "aws_iam_user" "admin-user" {
  name = "lucy"
  tags = {
    Description = "Technical Team Leader"
  }
}

resource "aws_iam_policy" "adminUser" {
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
  policy_arn = aws_iam_policy.adminUser.arn
}
```

We use heredoc syntax to define the policy (`<<EOF`). We can use `terraform plan` to see the changes and `terraform apply` to apply the changes.

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

resource "aws_iam_policy" "adminUser" {
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

resource "aws_s3_bucket_object" "finance-2020" {
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
```

Here, `aws_s3_bucket` resource is used to create an S3 bucket, `aws_s3_bucket_object` resource is used to upload a file to the bucket, and `aws_s3_bucket_policy` resource is used to create a bucket policy to allow access to the bucket. Additionally, we are using the `data` block to fetch information about an IAM group.