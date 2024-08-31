Kubedlv
---

Kubedlv is a small tool to forward command to golang `delve` command.

### The problem

In a Kubernetes container, some commands require numerous arguments and subcommands
to initiate. Starting a Delve instance can be challenging without modifying the
original YAML file. One approach is to have `delve` attach to the running process.
However, this option is often unavailable due to missing Docker capabilities.

### The solution

> You may use `go install github.com/imwithye/kubedlv@v1.0.0` to install `kubedlv`.

1. Install or copy `dlv` and `kubedlv` to the running container.
2. Open the YAML file of the running deployment, and in the command section,
   add `kubedlv` as the first command.

For example:

```yaml
# change this
command:
  - http-server

# to this
command:
  - kubedlv
  - http-server
```

Now you can attach to port 2345 to start debugging.

### How it works

The `kubedlv` simply pipes the original execution command to `dlv`. It translates

```bash
http-server --host 0.0.0.0 --port 8080
```

to

```bash
dlv --api-version=2 --listen=:2345 --headless=true --continue --accept-multiclient exec http-server -- --host 0.0.0.0 --port 8080
```