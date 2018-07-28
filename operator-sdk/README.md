# Operator SDK

The app-operator was created with:

```shell
$ operator-sdk new app-operator --api-version=policy.k8s-go.openshift.org/v1beta1 --kind=HealthCheckPolicy
```

and can be launched locally via:

```shell
$ OPERATOR_NAME=policy-operator operator-sdk up local
```

or build as docker image via:

```shell
$ operator-sdk build <docker-image-name-including-registry>
```
