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
*  jozefhalgas/pxlfrontend:1.0
* jozefhalgas/pxlbackend:latest
1) Vytvor Deployment obsahujuci 2 containery v pode (nginx, backend)
* sample-deployent.yaml
-
    - nastav potrebne resources
    - pouzi liveness/readiness probu pre nginx a liveness probu pre php
2) Vytvor Service typu NodePort ukazujuci na Applikacny pod
* sample-deployent.yaml
  ```
    curl -v $(minikube service --url sample-app)
    curl -v $(minikube service --url sample-app)/app/test
  ```

  `sample-app` je meno naseho servicu.

3) Vytvor Ingress definiciu ukazujucu na aplikacny pod.
* minikube addons enable ingress
* sample-deployent.yaml
Over funkcnost celeho deploymentu pouzitim nasledovnych prikazov

```
echo "$(minikube ip) sample-app.info" | tee -a /usr/local/etc/hosts

curl -v sample-app.info/
curl -v sample-app.info/app/testasdasd
```

4) Nasad Kubernetes Dashboard, Heapster, Metric server [7]
* minikube dashboard
* minikube addons enable heapster
* minikube addons open heapster
* minikube addons enable metrics-server

5) Pouzi Helm chart na nasadenie aplikacie [2], [5]
* helm lint ./sample-app
* helm install --dry-run --debug ./sample-app
* helm install --name sample-app ./sample-app
* helm package ./sample-app
* helm install --name sample-app sample-app-0.1.0.tgz
* helm ls
* heml delete sample-app
6) Pre applikaciu vytvor (aj manifest aj chart) HPA [9] a nastav ho nasledovne:
    min: 1
    max: 5
    cpu: 15%
* kubectl autoscale deployment sample-app --cpu-percent=15 --min=1 --max=5
* sucast menifestu a chartu
  Pomocou HTTP benchmark vytvor load na podoch tak aby HPA pridavalo nove pody.
* ab -n 200000 -c 100 http://sample-app.info/_nginx_status
* ab -n 200000 -c 100 http://sample-app.info/app/test

7) Implementuj blue/green mode nasadenia updatu aplikacie. Popis aky je rozdiel medzi Rolling Update/Recreate. Ake     su ich vyhody/nevyhody, na co treba davat pozor a ake maju parametre. [8]

Pridal som verziu do mena deploymentu - sample-app -> sample-app-1.0
Aby som vedel odlisit veriu podu do tmplatu som pridal label version: "1.0"
Na images v kontaineroch som nastavil tag 1.0, podla verzie (robil som to len pre nginx frontend kvoli jednosuchosti)
Pri service selectore pre service sample-app som zadefinoval verziu 1.0
Spustil som otestoval som. To bola moja blue.
Vytvoril som novy verziu frontendu 1.1, upravil som index
Vytvoril som novy deployment zo suboru sample-deployment-green.yml s menom sample-app-1.1
kubectl create -f sample-deployment-green.yml
Tento deployment ma verziu templatu 1.1 s frontend imagom pxlfrontend:1.1
Pre otestovanie som vytvoril novy service sample-app-green ktory ukazuje na deployment sample-app-1.1
Po otestovani som upravil service sample-app selector version  na "1.1" a uplatnil zmeni
kubectl apply -f sample-deployment.yaml
Pokial je vsetko ok mozme vymazat stary deployment a test service

Vyhodu Recreate vydim v tom ze vsetko sa vytvori na novo, to znamena ze pri kazdom update si tim prejde vsetko odznova a takze ma potom viac sebavedomie so spravovanim aplikacie a prostredia. Navyse ak sa vsetko vytvara odznova tak sa vymazu aj vsetky "docasne" nastavenia a ak nieco bolo dvolezite tak sa to znova ukaze a tim na to moze reagovat a pripravit sa do buducna aby sa to neopakovalo. Nevyhoda je downtime, ktory pri niecom takomto asi trva dlhsie. Podla man je nieco take rozumne urobit na dev a test prostrediach alebo aj v prode ak si vieme dovilit nejaky downtime, pretoze v zavislosti od aplikacie to moze byt celkom zlozite robit blue/green alebo rolling update ak sa meni napr db struktura, alebo sa tam robia zmeny v kode ktore nie su kompatibilne.
Blue/green vyhoda je v tom, ze si to pred nasadenim vieme dobre otestovat a jednoducho prepnut medzi verziami bez dlheho downtimu. Moze to byt ale narocne na zdroje ak chceme na testovanie green verzie pouzivat rovnake prostredie.
Rolling update vyhoda je ze to robi postupne a ak tam je nejaky problem tak to nemusi ovplivnit vsetkych uzivatelov. Zavisi od aplikacie ci zvladne ak bezi na roznych verziach sucasne a od rollbacku ake je komplikovane. Treba mysliet na to, aby uzivatel, ktory raz uz ide na novej verzii nepropol zase do starej. Rolling udpate my pride rozumne vyuzit napriklad ak aktualizujem napr nginx, alebo nejake os aktualizacie.


8) Pre build backend containera pouzi multistage docker build pattern [6]
* implementovane v Dockerfile pre backend
9) Implementuj PodDistribution budget a vysvetli preco a kedy to ma zmysel.
* sucast menifestu
O PodDistribution budget som nenasiel nic ale predpokladam, ze ste mali na mysli Pod Disruption Budget. Implementoval som ho len do manifestu helm chart som uz neaktualizoval.
PDB specifikuje pocet replik aplikacie, min alebo max, ktore musia bezat v Kubernetes clustry ak pride nejaka voluntary disruption request. (voluntary disruption are requests that are made to our kubernetes api that affects pod distribution like node draining, pod eviction etc.)
DPB je sluzba ktora by mala pomahat pri prevadzke HA aplikacii ak su tam caste voluntary disttripbution. Takze implementovat PDB ma zmysel ak poskytujeme highly available aplikaciu a casto robime zmeny ktore mozu ovplivnit dostupnost podov napriklad node draining, alebo ak chceme mat istotu ze niekto nahodou nevymaze vsetky pody z produ a bude zabespecena urcita uroven dostupnosti sluzby.
Este by stalo asi za zmenienku ze ja mam nastavenu iba jednu repliku a kubectl drain by zlihalo v tomto pripade, ale je to len testovacie prostredie, v prode sa predpoklada ze ked chceme HA ze bude replica bude minimalne 2, takze tam by to malo byt ok.
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
