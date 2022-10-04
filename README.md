## Prerequisites

In order to build and run this showcase you need to have a couple of things installed:

* The [Go Programming Language](https://go.dev)
* The [Iris Web Framework](https://www.iris-go.com)
* The [Docker Toolbox](https://www.docker.com/get-started) or native Docker, whatever you prefer
* The [Make tool](https://man7.org/linux/man-pages/man1/make.1.html) (optional)
* The [Wercker CLI](https://devcenter.wercker.com/development/cli/) (optional)

## Building the showcase

In case you have make installed you can simply issue the following command to build and
install the application:

```shell
$ go env -w CGO_ENABLED=0
$ go install -a ./... # or make install
```

In case you want to give Wercker a try (make sure you have the Wercker CLI installed) you
can issue the following command:

```shell
$ wercker dev --publish 8080 
```

This will build and run the microservice, and also watch for changes to the sources. If you only
want to run the Wercker build pipeline use the following command:

```shell
$ wercker build 
```

## Running the showcase

You have two options. Either build and run the showcase locally, or build and run the Docker image.

```shell
$ docker build -t cloud-native-go:1.0.2 . # or make docker
$ docker run --name cloud-native-go -it -p 18080:8080 cloud-native-go:1.0.2
```

Now open a browser or use somethin like `curl` or `HTTPie` and issue a GET request on the
URL http://localhost:18080/api/hello

## Deploying to Kubernetes

We will be using Minikube to run a small Kubernetes cluster locally. Make sure you have build
the Docker image locally.

```shell
$ kubectl create -f k8s-deployment.yml
$ kubectl get deployments pods
$ kubectl scale deployment cloud-native-go --replicas=3
$ kubectl get pods
$ kubectl get services
```

If you want to access the Go microservice you need to use the node port displayed by the last
command, e.g. `http GET localhost:32278/api/hello`

## License

This software is provided under the MIT open source license.

This is a ported version of M.-Leander Reimer (@lreimer), read [LICENSE file](https://github.com/lreimer/cloud-native-go/blob/master/LICENSE) for details. 
