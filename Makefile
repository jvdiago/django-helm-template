package:
	helm package . -d charts; \
    helm repo index charts --url https://gitlab.org/api/v4/projects/1/packages/helm/stable;
