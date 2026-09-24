# Rental Property REST API

A REST API built with Go and Beego for retrieving and filtering rental property data.

## Project overview
The API loads the provided rental property JSON file into memory when the application starts. It provides two read-only endpoints:
###### 1. `GET v1/properties`
###### 2. `GET v1/properties/:id`
And has these features:
- rental property listing
- property filtering with query parameters
- property lookup by ID
- JSON data loading and in-memory data handling
- structured API response mapping
- Swagger API documentation


## Project Structure
```
Rental-Property-REST-API/
├── conf/
│   └── app.conf
├── controllers/
│   └── property.go
├── data/
│   └── rental_properties.json
├── models/
│   ├── filter.go
│   ├── response.go
│   └── source.go
├── routers/
│   ├── commentsRouter.go
│   └── router.go
├── services/
│   └── property.go
├── utils/
│   └── query.go
├── validators/
│   └── property.go
├── swagger/
├── main.go
├── go.mod
├── go.sum
└── README.md
```

## Requirements and Instructions for project set up
Install these:
- Git
- Go 1.22 or higher
- Bee CLI
- CURL

### Install Go

```bash
sudo apt update
sudo apt install golang-go -y
````

### Install cURL

```bash
sudo apt install curl -y
```

### Install Bee CLI

```bash
go install github.com/beego/bee/v2@latest
```
### Add Go's binary directory to the PATH:

```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

### Clone the Project Repository

```bash
git clone https://github.com/akms-limon/Rental-Property-REST-API
```

### Install and synchronize project dependencies:

```bash
cd Rental-Property-REST-API
go mod download
go mod tidy
```

### Configuration

```text
conf/app.conf
```

The property data path is configured as:

```text
property_data_path = data/rental_properties.json
```

### Run the Application

Start the server using Bee:

```bash
bee run
```

The API will run on:

```text
http://localhost:8080
```

## Swagger API Documentation

After starting the application:

```bash
bee run
```

Open Swagger UI in your browser:

```text
http://localhost:8080/swagger/
```
If it is not working then try
```text
http://localhost:8080/swagger/index.html
```

Swagger can be used to view and test the available API endpoints.

## API Endpoints
### 1. `GET v1/properties/:id`
This will return single property based on that id.
Example:

```bash
curl "http://localhost:8080/v1/properties/BC-1000001"
```
### 2. `GET v1/properties`
Returns rental properties with optional filters.

```
   Architecture of GET /v1/properties?query_params
                    │
                    ▼
                 Router
                    │
                    ▼
               Controller
                    │
                    ▼
          Any filter parameter?
             /              \
           NO                YES
           │                  │
           │                  ▼
           │        ┌─────────────────────┐
           │        │ Validator 1         │
           │        │ Parameter validation│
           │        └──────────┬──────────┘
           │                   │
           │            Are parameters
           │               supported?
           │              /          \
           │            NO            YES
           │            │              │
           │            ▼              ▼
           │         400 Error     Validator 2
           │                       Type validation
           │                            │
           │                     Are values valid?
           │                       /           \
           │                     NO             YES
           │                     │               │
           │                     ▼               ▼
           │                  400 Error       Service
           │                                      │
           └──────────────────────────────────────┤
                                                  ▼
                                        Filter source records
                                                  │
                                                  ▼
                                      ┌─────────────────────┐
                                      │ Matching properties?│
                                      └──────────┬──────────┘
                                           /            \
                                         NO              YES
                                         │                │
                                         ▼                ▼
                                  No result found     Transform
                                                          │
                                                          ▼
                                                Is `limit` provided?
                                                   /          \
                                                 NO            YES
                                                 │              │
                                                 │        Apply limit
                                                 │              │
                                                 └──────┬───────┘
                                                        ▼
                                                      Result
                                                        │
                                                        ▼
                                                    Controller
                                                        │
                                                        ▼
                                                      JSON
```

### Query Parameters

The `GET /v1/properties` endpoint supports:

* `min_price`
* `max_price`
* `min_star_rating`
* `min_review_score`
* `min_reviews`
* `published`
* `property_type`
* `feed`
* `min_bedroom`
* `amenities`
* `limit`

### Filter Examples

#### Filter by Feed

```bash
curl "http://localhost:8080/v1/properties?feed=11"
```

#### Filter by Price Range

```bash
curl "http://localhost:8080/v1/properties?min_price=50&max_price=150"
```

#### Filter by Minimum Star Rating

```bash
curl "http://localhost:8080/v1/properties?min_star_rating=4"
```

#### Filter by Minimum Review Score

```bash
curl "http://localhost:8080/v1/properties?min_review_score=7"
```

#### Filter by Minimum Reviews

```bash
curl "http://localhost:8080/v1/properties?min_reviews=20"
```

#### Filter by Published Status

```bash
curl "http://localhost:8080/v1/properties?published=true"
```

#### Filter by Property Type

```bash
curl "http://localhost:8080/v1/properties?property_type=Hotel"
```

#### Filter by Minimum Bedroom

```bash
curl "http://localhost:8080/v1/properties?min_bedroom=2"
```

#### Filter by Amenities

Multiple amenities can be provided as a comma-separated list:

```bash
curl "http://localhost:8080/v1/properties?amenities=Internet,Parking"
```

#### Apply a Result Limit

```bash
curl "http://localhost:8080/v1/properties?limit=10"
```

#### Combine Multiple Filters

```bash
curl "http://localhost:8080/v1/properties?feed=11&published=false"
```

#### Combine Price and Property Type

```bash
curl "http://localhost:8080/v1/properties?min_price=50&max_price=150&property_type=Hotel"
```

#### Combine Filters and Amenities

```bash
curl "http://localhost:8080/v1/properties?feed=11&amenities=Internet,Parking"
```

#### Combine Filters with a Limit

```bash
curl "http://localhost:8080/v1/properties?feed=11&min_price=50&min_star_rating=4&limit=10"
```

## Unit Tests

<!-- Add unit test information here after implementing the tests. -->

Run the service tests:

```bash
go test ./... -v
```

Run static analysis:

```bash
go vet ./...
```
