# Zadanie

## Kubernetes

Nainstaluj si u seba `minikube` [1]. Pomocou `minikube` si v _Kubernetes_ nainstaluj nginx ingress controller [3].

## Aplikacie

Vybuilduj si kontainery ktore sa nachadzaju v priecinku docker

```
docker
├── backend
│   ├── Dockerfile
│   ├── Makefile
│   ├── README.md
│   ├── backend
│   └── main.go
└── nginx
    ├── Dockerfile
    ├── Makefile
    ├── default.conf
    ├── index.html
    └── start.sh
```

Zverejni svoje kontainery na [docker hub](https://hub.docker.com/) a nasledne ich pouzi pri
deploymente na Kubernetes. Na nasadenie aplikacie na kubernetes pouzi nasledovne veci:

1) Vytvor Deployment obsahujuci 2 containery v pode (nginx, backend)
    - nastav potrebne resources
    - pouzi liveness/readiness probu pre nginx a liveness probu pre php

2) Vytvor Service typu NodePort ukazujuci na Applikacny pod

  ```
    curl -v $(minikube service --url sample-app)
    curl -v $(minikube service --url sample-app)/app/test
  ```

  `sample-app` je meno naseho servicu.


3) Vytvor Ingress definiciu ukazujucu na aplikacny pod.

Over funkcnost celeho deploymentu pouzitim nasledovnych prikazov

```
echo "$(minikube ip) sample-app.info" | tee -a /usr/local/etc/hosts

curl -v sample-app.info/
curl -v sample-app.info/app/testasdasd
```

4) Nasad Kubernetes Dashboard, Heapster, Metric server [7]

5) Pouzi Helm chart na nasadenie aplikacie [2], [5]

6) Pre applikaciu vytvor (aj manifest aj chart) HPA [9] a nastav ho nasledovne:
    min: 1
    max: 5
    cpu: 15%

  Pomocou HTTP benchmark vytvor load na podoch tak aby HPA pridavalo nove pody.

7) Implementuj blue/green mode nasadenia updatu aplikacie. Popis aky je rozdiel medzi Rolling Update/Recreate. Ake     su ich vyhody/nevyhody, na co treba davat pozor a ake maju parametre. [8]

8) Pre build backend containera pouzi multistage docker build pattern [6]

9) Implementuj PodDistribution budget a vysvetli preco a kedy to ma zmysel.

## Vyhodnotenie

Cielom zadania nie je nevyhnutne aby si vsetko implementoval ale aby si preukazal ze si vies
nastudovat nove veci. Ulohy maju zvysujucu sa obtiaznost a kandidat je hodnoteny podla
vyhodnotenia a nie podla toho kolko spravil.

## Dokumentacia

[1] https://kubernetes.io/docs/setup/minikube/

[2] https://helm.sh/

[3] https://github.com/kubernetes/ingress-nginx

[5] https://github.com/kubernetes/charts

[6] https://docs.docker.com/develop/develop-images/multistage-build/

[7] https://github.com/kubernetes/minikube/blob/master/docs/addons.md

[8] https://www.ianlewis.org/en/bluegreen-deployments-kubernetes

[9] https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/