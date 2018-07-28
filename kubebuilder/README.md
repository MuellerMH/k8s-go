# kubebuilder

The example was built using: 

```shell
$ kubebuilder init --domain k8s-go.openshift.org --license apache2 --owner "The workshop members"
$ kubebuilder create api 	--group policy --version v1beta1 --kind HealthCheckPolicy
```

Note, due to https://github.com/kubernetes-sigs/kubebuilder/issues/335 some imports had to be fixed manually.

The CRDs are installed via:

```shell
$ make install
```

The operator an be launched via:

```shell
$ go run cmd/manager/main.go
```
