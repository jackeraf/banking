# Banking project

This is an in progress project where a user can transfer, withdraw and check the amount of money he has on his fake bank account

Right now there is no user and the amount withdrawn is hardcoded when you curl the root directory

The main idea behind this project is the architecture not the app itself: dockerhub custom images, kubernetes namespaces, importing custom libraries, monitoring events and microservice resources...

TODOS:

- seed users in postgresql
- create more endpoints to transfer and withdraw money from users
- add a load balancer
- integration tests
- etc

### Stack used:

  kubernetes
  helm charts
  rabbitmq
  golang
  prometheus
  grafana

### Requirements:

You need to have minikube installed and also kubectl and helm:

https://kubernetes.io/docs/tasks/tools/install-minikube/

https://kubernetes.io/docs/tasks/tools/install-kubectl/

https://helm.sh/docs/intro/install/


### Setup:

1)

`minikube start`
2)

`kubectl apply -f k8s/namespaces`

3)

`
helm install rabbitmq \
  --set rabbitmqUsername=admin,rabbitmqPassword=secretpassword,managementPassword=anothersecretpassword,rabbitmqErlangCookie=secretcookie \
    stable/rabbitmq-ha --namespace messaging
`

4)

`kubectl apply -f k8s/backend`

5)

`helm install monitoring --namespace monitoring stable/prometheus-operator`

### Steps to run it:

1) Port forward rabbitmq

`kubectl port-forward --address=0.0.0.0 --namespace messaging service/rabbitmq-rabbitmq-ha 15672:15672`

2) Check the queues and messages. On the browser enter:

http://localhost:15672/

admin: admin
password: secretpassword

2) Get the minikube ip:

`minikube ip`

open the browser hit the url:

`<minikube_ip>::30080`

You'll get `Money withdrawn 5 €` as a response


### Other useful notes:

- enable node exporter
https://prometheus.io/docs/guides/node-exporter/

- Lint on development:

`golangci-lint run --config ./.golangci.yml`
