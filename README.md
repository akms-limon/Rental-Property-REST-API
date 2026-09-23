# Rental Property REST API

### Query Architecture for `GET /v1/properties`

```text
GET /v1/properties?feed=11&min_price=50&limit=10
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

### Query Architecture for `Get /v1/properties/:id`
