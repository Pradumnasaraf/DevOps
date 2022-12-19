## Kubescape

Kubescape is a tool that helps you scan your Kubernetes clusters for security misconfigurations. We can scan local cluster and config files.

- [Website](https://armosec.io/kubescape)
- [Docs](https://hub.armosec.io/docs)
- [GitHub](https://github.com/kubescape/kubescape)

 
### Usage

For more usages head over to [GitHub](https://github.com/kubescape/kubescape)

- Install Kubescape

Installations instructions can be found [here](https://github.com/kubescape/kubescape#install-on-macos)

- To scan local YAML/JSON files.

```bash
kubescape scan *.yaml
```

- To Scan Kubernetes manifest files from a git repository

```bash
kubescape scan https://github.com/kubescape/kubescape
```

- We can output the results in JSON, html, and markdown PDF.

```bash
kubescape scan *.yaml --output results.json
kubescape scan *.yaml --output results.pdf
kubescape scan *.yaml --output results.html
```