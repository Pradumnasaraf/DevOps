---
title: Kustomize Introduction
sidebar_position: 1
---

Kustomize is a tool that lets you customize raw, template-free YAML files for multiple purposes, leaving the original YAML untouched and usable as is.

Kustomize uses the `apiVersion`, `kind`, `name` and `namespace` to identify resources. It appends a suffix to the name and namespace fields to differentiate between different versions of the same resource.

## Resources

