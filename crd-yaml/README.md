# Custom Resource Definitions basics

The HealthCheckPolicy is pre-created. You can install it via:

```shell
$ kubectl create -f healthcheckpolicy-crd.yaml
```

Then the example can be created:

``` shell
$ kubectl create -f example-healthcheckpolicy.yaml
```

The other example need a CRD to be used. Also proper OpenAPI validation would be helpful.
