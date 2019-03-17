# Google Cloud SQL

#### Set up - Google Cloud
1. After creating a Google Cloud project and enabling billing (or using credits), go to the `Cloud SQL` service and click `Create instance`.
2. Select which SQL engine you would like; `MySQL` or `PostgreSQL`.
3. After selecting a SQL engine, input an `Instance ID` and `Root password`. Document the root password as you will need it to connect in Go. You may leave the `Region` and `Zone` on the default options.
4. After a few minutes or so, the instance should be created. Click on the instance to view its dashboard.
5. Click on the `DATABASES` tab and create two new databases. Call one `prod` (for the live, production app) and one called `dev` (used during development).
6. Under `Connect to this instance`, copy the `Instance connection name`. It follows the form: project-id:zone:instance-name

#### Set up - Go Code
1. Install gorm and the cloudsql-proxy modules
```
go get -u github.com/jinzhu/gorm
go get -u github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql
```
Note: If you are using Go modules (i.e. `go mod`) to track your dependencies, set `GO111module=off` variable when installing these two. You may include them in go.mod later by setting `GO111module=on`.

2. Create a function to create a connection to the database and return the database instance
```go
func Create() *gorm.DB {
	log.Println("Creating new Google Cloud SQL connection...")

	username := os.Getenv("SQL_USERNAME")
	password := os.Getenv("SQL_PASSWORD")
	connectionName := os.Getenv("SQL_CONNECTION_NAME")
	databaseName := os.Getenv("SQL_DATABASE")

	dataSourceName := fmt.Sprintf("%s:%s@cloudsql(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, connectionName, databaseName)

	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	return db
}
```
Note: If you selected PostgreSQL as your preferred engine, the first parameter of `gorm.Open` should be set to `"postgres"` not `"mysql"`
3. Change the `dbConnectionString` variable to include your credentials. For example:
  - The default user name is `root`
  - My password is: `supersecretpassword`
  - My instance connection name is: `cool-project:us-east:cool-project-database`
  - The database that I want to connect to is: `dev`
  - The connection string will look like this:
```go
dbConnectionString := "root:supersecretpassword@cloudsql(cool-project:us-east:cool-project-database)/dev?charset=utf8&parseTime=True&loc=UTC"
```

4. Run the api to ensure the call to `gorm.Open` successfully connects to the database within the Cloud SQL instance and does not return any errors.

5. Before commiting the code to GitHub, make sure the credentials (username, password, instance connection name) are not public. To do this, you may use a `.env` file along with the [godotenv](https://github.com/joho/godotenv) dependency.
Optional: If you prefer not to use the root account to connect to the Cloud SQL instance, go back to the Cloud SQL console and select the `USERS` tab. Click `Create user account` and replace the username and password in the connection string. Make sure that `Allow any host(%)` is checked. Make sure the username and password are seperated by a colon in the connection string (`username:password`).

## Tips:
- *How do I connect to Cloud SQL locally?*
	1. Use [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/quickstart-proxy-test)
	2. Take note of the instance connection name of your target [Cloud SQL Instance](https://console.cloud.google.com/sql/instances) in the `Instance Connection Name` section
	3. Install [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/quickstart-proxy-test) for your machine
	4. Make sure that you do not have any process running on port 3306 by running `lsof -i tcp:3306` or `netstat -ano -p tcp` on windows
	5. Run the `cloud_sql_proxy` executable with some flags
		- `cloud_sql_proxy -instances={instance connection name}=tcp:3306`
	6. Connect using command line or [MySQLWorkbench](https://www.mysql.com/products/workbench/) as it should be able to find the server automatically
	7. You can check out your databases in the `Schemas` section, and run queries in the main text editor. To run them, hit the thunderbolt symbol. Make sure your query is acting on the right database
