# Google App Engine and Cloud SQL/Storage Demonstration

## What this app does
This all will allow students to comment on other people's resume and be able to view them

## How we did this
This app is created using NextJS/React and Golang for the API. The database of choice will be Cloud SQL/MySQL. The bloc storage of choice will be Cloud Storage

## Tips:
- *How do I make data public?*
- You have to add Users-allUser-Reader to the permissions of the document

## Setup:
- Upload and publicize 4 resume's from online using the naming convention `resume_{n}.pdf`
- During the demo, upload 1 more resume and publicize in front of everyone
- Throw connection string of Cloud SQL instance int he env file
- Connect to Cloud SQL from some type of IDE so that they can see the SQL statement for creating the database outside of the terminal
