# GeoLocation App

The GeoLocation App is a Go application that allows you to parse CSV files containing geolocation data and provides a RESTful API to retrieve information based on IP addresses.

## Features

- Parse CSV files containing geolocation data and store it in a database.
- Expose a RESTful API to retrieve geolocation information by IP address.
- Supports both SQLite and MySQL databases.

## Getting Started

### Prerequisites

Before running the GeoLocation App, ensure you have the following prerequisites installed:

- Go (at least Go 1.16)
- Docker and Docker Compose (if using MySQL)
- Git (for cloning the repository)

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/zkapezov/geolocation.git
2. Run docker-compose
   ```
   docker-compose --env-file .env.dev up 
3. Copy csv file with geolocation data into `data/` folder
4. Parse geolocation data and persist in database 
   ```
   docker exec -it geolocation-app-1 /app/geolocation-parser /app/data/data_dump.csv

# API Endpoints:
The API will be accessible at http://localhost:8080.

### GET http://localhost:8080/geolocations/127.0.0.1
Retrieve geolocation information based on an IP address.

**Parameters**

|          Name | Required |  Type   | Description                                                        |
| --------------|----------|---------|--------------------------------------------------------------------|
|     `ipadress`| required | string  | The IP address for which you want to retrieve geolocation data.    |


**Response**
```
{
"IPAddress": "127.0.0.1",
"CountryCode": "CY",
"Country": "Equatorial Guinea",
"City": "South Austyn",
"Latitude": -36.74979952,
"Longitude": 145.71556407,
"MysteryValue": 6003526404
}
```


