# Helm Chart for Django
This helm chart install django in a kubernetes cluster.

### Install Helm
Read and follow the [Helm installation guide](https://helm.sh/docs/intro/install/).

Note: The charts in this repository require Helm version 3.x or later.

### Add the Django Helm Chart repo
In order to be able to use the charts in this repository, add the name and URL to your Helm client:
```bash
helm repo add django https://gitlab.org/api/v4/projects/1/packages/helm/stable --username xxxxxxxx --password xxxxxxxxxxx
helm repo update
```

### Install charts
```bash
helm install my-release django/django -f values.yaml
```

### Configuration
django-helm includes

* django-server
* django-config
* django-volumes
* django-celery-beat
* django-celery-work
* django-celery-flower
* django-grpc-server
* django-migrate-job
* django-collectstatic-job
* django-log

see [values.yaml](https://github.com/jvdiago/django-helm-template/src/master/values.yaml)

### For production
Integrated alicloud hpa, sls.
