
# Maximumsoft Exam

- Solotion 1 : https://github.com/sSlotty/maximumsoft_exam/tree/main/solution1
- Solotion 2 : https://github.com/sSlotty/maximumsoft_exam/tree/main/solution2





## Solotion 1 API Reference

#### Login

```http
  POST /login
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required**. Username  |
| `password` | `string` | **Required**.  Password|

#### Refresh Token

```http
  POST /refreshToken
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `refreshToken`      | `string` | **Required**. Refresh token |

#### Create Employee

```http
  POST /api/v1/employee
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `assess_token`| `Authorization` | **Required**. assess_token |
| `full_name` | `string` | **Required**. Full name of employee |
| `address` | `string` | **Required**. address of employee |
| `email` | `string` | **Required**. email of employee |
| `phone` | `string` | **Required**. phone of employee |
| `salary` | `int` | **Required**. salary of employee |

#### Get employee by id

```http
  GET /api/v1/employee/{$id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `assess_token`| `Authorization` | **Required**. assess_token |
| `id` | `string` | **Required**. ID of employee |

#### Get all employee

```http
  GET /api/v1/employees
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `assess_token`| `Authorization` | **Required**. assess_token |

#### Get all employee

```http
  PUT /api/v1/employee/{$id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `assess_token`| `Authorization` | **Required**. assess_token |
| `full_name` | `string` | **Required**. Full name of employee |
| `address` | `string` | **Required**. address of employee |
| `email` | `string` | **Required**. email of employee |
| `phone` | `string` | **Required**. phone of employee |
| `salary` | `int` | **Required**. salary of employee |

#### DELETE employee by id

```http
  DELETE /api/v1/employee/{$id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `assess_token`| `Authorization` | **Required**. assess_token |
| `id` | `string` | **Required**. ID of employee |

