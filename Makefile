dev:
	gin --appPort 8080 --port 8081 run main.go

deploy:
	gcloud app deploy

dev-db-proxy:
	cloud_sql_proxy -instances=resume-engine:us-central1:resume-engine-main=tcp:3306
