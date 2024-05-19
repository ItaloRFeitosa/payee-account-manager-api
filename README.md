# Payee Account Manager API
REST API to handle payee account management to support payment features and bank account validation
## Sections
1. [How to Run](#how-to-run)
2. [How To Test](#how-to-test)
3. [Use Cases](#use-cases)
4. [Extras](#extras)

## How to Run
### Makefile
### Docker

## How to Test
### Unit Tests
### Integration Tests
### E2E Test

## Use Cases
### Register a new Payee
#### Endpoint
```json
// POST api/v1/payees
// Request Header
// tenant-id: uuid
// Request Body
{
    "name": "Italo Feitosa",
    "cpf_cnpj": "99818083008",
    "email": "italo@feitosa.com",
    "pix_key_type": "CPF",
    "pix_key": "99818083008"
}

// Response 201 Created
{
    "data": {
        "id": "1",
    }
}

```

#### Requirements
* `name` is required, min 2 chars, max 128 chars, min 2 words
* `cpf_cnpj` is required, should follow brazillian CPF and CNPJ validations
* `email` is not required, max length of 140 chars. Regex: `/^[a-z0-9+_.-]+@[a-z0-9.-]+$/`
* When new payee registered, the default status will be **DRAFT**
* `pix_key_type` and `pix_key` are required:

    | Pix Key Type  | Pix Key Pattern |
    |---|---|
    | CPF | `/^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$/` |
    | CNPJ | `/^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$/` |
    | EMAIL | `/^[a-z0-9+_.-]+@[a-z0-9.-]+$/` |
    | TELEFONE | `/^((?:\+?55)?)([1-9][0-9])(9[0-9]{8})$/` |
    | CHAVE_ALEATORIA | `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i`|

### Edit Payee Details
#### Endpoint
```json
// PUT api/v1/payees/:payee_id
// Request Header
// tenant-id: uuid
// Request Body
{
    "name": "Italo Feitosa",
    "cpf_cnpj": "99818083008",
    "email": "italo@feitosa.com",
    "pix_key_type": "CPF",
    "pix_key": "99818083008"
}

// Response 204 No Content
```

#### Requirements
* When Payee status is **DRAFT** same validations in register payee
* When Payee status is **VALID** only `email` can be edited

### List Payees
#### Endpoint
```json
// GET api/v1/payees?page=1&size=1&search=
// Request Header
// tenant-id: uuid

// Response 200 OK
{
    "data": [{
        "id": "1",
        "name": "Italo Feitosa Draft",
        "cpf_cnpj": "99818083008",
        "email": "italo@feitosa.com",
        "pix_key_type": "CPF",
        "pix_key": "99818083008",
        "status": "DRAFT",
        "bank_account": null,
        "created_at": "2024-05-17T20:16:29.666Z",
        "updated_at": "2024-05-17T20:16:29.666Z",
    },{
        "id": "2",
        "name": "Italo Feitosa Valid",
        "cpf_cnpj": "99818083008",
        "email": "italo@feitosa.com",
        "pix_key_type": "CPF",
        "pix_key": "99818083008",
        "status": "VALID",
        "bank_account": {
            "account_type": "CONTA_CORRENTE",
            "account_number": "65465465",
            "account_digit": "5",
            "branch_number": "0001",
            "bank_code": "1",
            "bank_ispb": "54545"
        },
        "created_at": "2024-05-17T20:16:29.666Z",
        "updated_at": "2024-05-17T20:16:29.666Z",
    }],
    "metadata": {
        "total_items": 50,
        "total_pages": 5,
        "page": 1,
        "page_size": 5
    }
}
```
#### Requirements
* Should be paginated
* Searchable by name, cpf_cnpj, branch_number, account_number, status, pix_key_type, pix_key
* Page default size is 10

### Delete Payees
#### Endpoint
```json
// DELETE api/v1/payees
// Request Header
// tenant-id: uuid
// Request Body
{
    "ids": ["1", "2"]
}

// Response 204 No Content
```
#### Requirements
* Should mark payees as deleted (soft delete)


## Extras
### Project Structure
### Swagger
### High Level Architecure
### ER Diagram
### API Conventions
### Decisions
