Rental-Property-REST-API


## Query Architecture for `v1/properties` endpoint
Swagger / HTTP Request
        ↓
Controller
        ↓
Read query parameters
        ↓
validators package
        ↓
PropertyFilters
        ↓
Service
        ↓
Filter source records
        ↓
Transform
        ↓
Limit
        ↓
Result