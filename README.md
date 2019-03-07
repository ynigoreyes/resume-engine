# Google App Engine and Cloud SQL/Storage Demonstration

## What this app does
This all will allow students to comment on and view other people's resumes.

## How we did this
This app is created using NextJS/React for the frontend and Golang for the API. The database of choice will be Google Cloud SQL's MySQL offering. The blob storage provider (for resumes) of choice will be Google Cloud Storage.

## Database Setup
### Part 1 - In Google Cloud Console
1. After creating a Google Cloud project and enabling billing (or using credits), go to the `Cloud SQL` service and click `Create instance`. 
2. Select which SQL engine you would like; `MySQL` or `PostgreSQL`.
3. After selecting a SQL engine, input an `Instance ID` and `Root password`. Document the root password as you will need it to connect in Go. You may leave the `Region` and `Zone` on the default options.
4. After a few minutes or so, the instance should be created. Click on the instance to view its dashboard. 
5. Click on the `DATABASES` tab and create two new databases. Call one `prod` (for the live, production app) and one called `dev` (used during development).
6. Under `Connect to this instance`, copy the `Instance connection name`. It follows the form: project-id:zone:instance-name
### Part 2 - In Go code
1. Create a main.go file like the one in the following Gist: https://gist.github.com/erans/795f7616049f310b2a71a50e5df4529a
2. Install gorm and the cloudsql-proxy modules
```
go get -u github.com/jinzhu/gorm
go get -u github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql
```
Note: If you are using Go modules (i.e. `go mod`) to track your dependencies, set `GO111module=off` variable when installing these two. You may include them in go.mod later by setting `GO111module=on`.

3. Change the `dbConnectionString` variable to include your credentials. For example:
  - The default user name is `root`
  - My password is: `supersecretpassword`
  - My instance connection name is: `cool-project:us-east:cool-project-database`
  - The database that I want to connect to is: `dev`
  - The connection string will look like this: 
```go
var dbConnectionString = "root:supersecretpassword@cloudsql(cool-project:us-east:cool-project-database)/dev?charset=utf8&parseTime=True&loc=UTC"
```
Note: If you selected PostgreSQL as your preferred engine, the first parameter of `gorm.Open` should be set to `"postgres"` not `"mysql"`

4. Run `go run main.go` to ensure the call to `gorm.Open` successfully connects to the database within the Cloud SQL instance and does not return any errors. 

5. Before commiting the code to GitHub, make sure the credentials (username, password, instance connection name) are not public. To do this, you may use a `.env` file along with the [godotenv](https://github.com/joho/godotenv) dependency.
```
TODO: Link code in database.Create as an example
```

Optional: If you prefer not to use the root account to connect to the Cloud SQL instance, go back to the Cloud SQL console and select the `USERS` tab. Click `Create user account` and replace the username and password in the connection string. Make sure that `Allow any host(%)` is checked. Make sure the username and password are seperated by a colon in the connection string (`username:password`). 

## Tips:
- *How do I make data public?*
  - You have to add `Users-allUser-Reader` to the permissions of the document

## Setup:
- Upload and publicize 4 resume's from online using the naming convention `resume_{n}.pdf`
- During the demo, upload 1 more resume and publicize in front of everyone
- Throw connection string of Cloud SQL instance int he env file
- Connect to Cloud SQL from some type of IDE so that they can see the SQL statement for creating the database outside of the terminal
