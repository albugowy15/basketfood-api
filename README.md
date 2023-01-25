# basketfood-api

A public API build with GORM and Gin Web Framework

## PDM
![BasketFood-2023-01-08_17-46](https://user-images.githubusercontent.com/49820990/214455894-72c8fabc-3e2f-4973-9849-762923c9799f.png)

## How to run this project
### 1. Clone this project
```sh
git clone https://github.com/albugowy15/basketfood-api.git
```
### 2. Use your own database
This API using PostgreSQL database. You can create your own database on your local machine, then save all its credential to .env folder.

Create `.env` file inside root directory, fill this variable with your own database credentials
```.env
DB_HOST=<your database host>
DB_PASSWORD=<your database password>
DB_USER=<your database user>
DB_NAME=<your database name>
DB_PORT=<your database port>
```
### 3. Install all dependencies
Install all dependencies or library used in this project with command
```sh
go get -u -v ./... 
```
### 4. Run the project
Run this project with command
```sh
go run main.go
```
You can see that your project is running at `http://localhost:8080`
