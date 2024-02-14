# simple-web-app-nasa

## Starting App
To start backend:

Download dependencies using
```
go mod download
```
NOTE: may need to do "go mod init" before this command if module is not set up

Start database in docker, must install docker if not already running
Database will be created with needed tables
```
docker-compose up -d
```

From server directory run command:

```
go run ./cmd/api/*.go
```

To start frontend:

From client directory:

Download dependencies using npm install

Start client with 
```
npm run dev
```

## Using App

Welcome page
localhost:3000

Load Data into Database (this will load first 10 nasa data points into database)
localhost:3000/loadNasaData

Look at all NasaData
localhost:3000/nasaData

Look at specific data
localhost:3000/nasaData/some-id