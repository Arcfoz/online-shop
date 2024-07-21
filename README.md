
# Go-Ecommerce-RestAPI

This project enhances a popular Go REST API tutorial by implementing Domain-Driven Design (DDD). Using the Fiber framework, it adds new features and improves user experience. The application is structured with distinct domain and infrastructure layers, using PostgreSQL for data storage. DDD principles guide the design, with domain objects at the core and business logic encapsulated in aggregates. Postman is used for API testing. This project serves as a practical learning experience in both Golang backend development and DDD implementation.




## Usage

#### To initialize the project, use the following instructions:

```bash
docker-compose up -d
go run .\cmd\api\main.go
```


## API Reference

POSTMAN API : https://elements.getpostman.com/redirect?entityId=28552659-612a3bdc-6cee-471c-8189-20881ac92b7a&entityType=collection

#### Register

```bash
  POST http://localhost:4000/auth/register
```

| Parameter | Type     | 
 :-------- | :------- | 
| `email` | `string` |
| `password` | `string` |

###### response:

```json
  {
    "success": true,
    "message": "register successed",
    "error": ""
  }
```

#### Login

```bash
  POST http://localhost:4000/auth/register
```

| Parameter | Type     | 
 :-------- | :------- | 
| `email` | `string` |
| `password` | `string` |

###### response:

```bash
  {
    "success": true,
    "message": "login successed",
    "payload": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYxZjZkM2FkLTYzNzItNGY5OC1iZjRjLTg4MWFiMTk5NzBiNSIsInJvbGUiOiJ1c2VyIn0.JSvH5uIXvn7pc2EwmuDxmK4WzqjzZw5pg3XXhYBsWTU"
    },
    "error": ""
  }
```

#### Add Product (ROLE : ADMIN)

Change the role in database to admin manualy

```bash
  POST http://localhost:4000/products
```

| Parameter | Type     | 
 :-------- | :------- | 
| `name` | `string` |
| `stock` | `int` |
| `price` | `int` |

| Authorization | Type     | 
 :-------- | :------- | 
| `Bearer Token` | `Token` |

###### response:

```json
  {
    "success": true,
    "message": "create product successed",
    "error": ""
  }
```

#### Find Product SKU


```bash
  GET http://localhost:4000/products/sku/:sku
```
| Path Variables | value     | 
 :-------- | :------- | 
| `sku` | `18beb206-47bb-40d1-95a4-f53dabe6089c` |

###### response:

```json
  {
    "success": true,
    "message": "get product detail successed",
    "payload": {
        "id": 8,
        "sku": "18beb206-47bb-40d1-95a4-f53dabe6089c",
        "name": "Navy T-shirt",
        "stock": 10,
        "price": 100000,
        "created_at": "2024-07-21T15:23:46.676518Z",
        "updated_at": "2024-07-21T15:23:46.676518Z"
    },
    "error": ""
  }
```

#### Pagination Product


```bash
  GET http://localhost:4000/products?cursor=0&size=10
```
| Path Variables | value     | 
 :-------- | :------- | 
| `cursor` | `0` |
| `size` | `10` |

###### response:

```json
  {
    "success": true,
    "message": "get list products successed",
    "payload": [
        {
            "id": 8,
            "sku": "18beb206-47bb-40d1-95a4-f53dabe6089c",
            "name": "Navy T-shirt",
            "stock": 10,
            "price": 100000
        },
    ],
    "query": {
        "cursor": 0,
        "size": 10
    },
    "error": ""
  }
```

#### Checkout Product


```bash
  POST http://localhost:4000/transaction/checkout
```
| Parameter | Type     | 
 :-------- | :------- | 
| `product_sku` | `string` |
| `amount` | `int` |

| Authorization | Type     | 
 :-------- | :------- | 
| `Bearer Token` | `Token` |

###### response:

```json
  {
    "success": true,
    "message": "create transaction successed",
    "error": ""
  }
```

#### History Checkout Product


```bash
  GET http://localhost:4000/transaction/user/histories
```
| Authorization | Type     | 
 :-------- | :------- | 
| `Bearer Token` | `Token` |

###### response:

```json
  {
    "success": true,
    "message": "get transaction history successed",
    "payload": [
        {
            "id": 2,
            "user_public_id": "61f6d3ad-6372-4f98-bf4c-881ab19970b5",
            "product_id": 8,
            "product_price": 100000,
            "amount": 1,
            "subtotal": 100000,
            "platform_fee": 1000,
            "grand_total": 101000,
            "status": "CREATED",
            "created_at": "2024-07-21T15:39:17.737364Z",
            "updated_at": "2024-07-21T15:39:17.737364Z",
            "product": {
                "id": 8,
                "sku": "18beb206-47bb-40d1-95a4-f53dabe6089c",
                "name": "",
                "price": 100000
            }
        }
    ],
    "error": ""
  }
```



## References

Tutorial link: https://youtu.be/YpKq8T7wUjs?si=8xZb58Momj4hqapR

