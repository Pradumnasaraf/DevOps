---
sidebar_position: 3
title: Terraform Commands
description: Terraform commands and their usage
tags: ["Terraform", "Infrastructure as Code", "HashiCorp"]
keywords: ["Terraform", "Infrastructure as Code", "HashiCorp"]
slug: "/terraform/commands"
---

1. Terraform Init

It is used to initialize a working directory containing Terraform configuration files. 

```bash
terraform init
```

2. Terraform Plan

It is used to create an execution plan. It shows what Terraform will do when you call apply. 

```bash
terraform plan
```

3. Terraform Apply

It is used to apply the changes required to reach the desired state of the configuration. 

```bash
terraform apply
```

4. Terraform Validate

We can run this command before applying the changes to check whether the configuration is syntactically valid and internally consistent. 

```bash
terraform validate
```

It will print the exact in the console if there is any error in the configuration file.

5. Terraform Format

It is used to rewrite Terraform configuration files to canonical format and style. 

```bash
terraform fmt
```

6. Terraform Show

It is used to provide human-readable output of current state of the resources. 

```bash
terraform show
```

We can also output the state in JSON format by using the following command:

```bash
terraform show -json
```

7. Terraform Providers

To know the list of providers used in the configuration file, we can use the following command:

```bash
terraform providers
```

To mirror the provider configurations from the configuration file, we can use the following command:

```bash
terraform providers mirror ./path/to/new_local_file
```

8. Terraform Output

It is used to extract the output variables from the state file. 

```bash
terraform output
```

9. Terraform Refresh

It is used to update the state file according to the real-world infrastructure. 

```bash
terraform plan
```

or 

```bash
terraform apply -refresh-only
```

10. Terraform Graph

It is used to generate a visual representation of the configuration and state file. 

```bash
terraform graph
```

11. Terraform Destroy

It is used to destroy the Terraform-managed infrastructure. 

```bash
terraform destroy
```