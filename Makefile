dev-api:
	gin --appPort 8080 --port 8081 run main.go

dev-client:
	yarn dev

deploy:
	gcloud app deploy

deploy-local:
	dev_appserver.py app.yaml

