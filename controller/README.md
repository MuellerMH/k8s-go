# Writing a controller

Run this controller via:

```shell
$ kubectl create -f config/healthcheckpolicy-crd.yaml
$ go run cmd/policy-controller/main.go
```

Create a test resource with:

```shell
$ kubectl create -f config/example-healthcheckpolicy.yaml
```

Rebuild generated code via:

```shell
hack/update-codegen.sh
```
