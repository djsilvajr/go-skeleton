# 📘 API Response Pattern

## 🎯 Purpose

Define a strict and consistent response format for all API endpoints.

---

## 📦 Standard Response

All responses MUST follow this structure:

```json
{
  "data": {}
}
```

---

## ✅ Success Response

* Always return `data`
* Never return raw objects outside `data`

### Example:

```json
{
  "data": {
    "id": 1,
    "name": "Admin",
    "email": "admin@example.com",
    "role": "admin",
    "createdAt": "2026-04-06T10:26:43.536Z",
    "updatedAt": "2026-04-06T10:26:43.536Z",
    "deletedAt": null
  }
}
```

---

## 📚 List Response

For collections:

```json
{
  "data": [
    {}
  ]
}
```

---

## ❌ Error Response (MANDATORY)

All errors MUST follow:

```json
{
  "error": {
    "code": 422,
    "message": "Validation error",
    "details": {}
  }
}
```

### Rules:

* `code` → HTTP status code
* `message` → short description
* `details` → validation errors or additional info

---

## ⚠️ Validation Error Example

```json
{
  "error": "error message"
}
```

---

## 🗑️ Delete Pattern

* ALWAYS use **Soft Deletes**
* Never permanently delete unless explicitly requested

### Delete Response Example:

```json
{
  "data": {
    "deleted": true
  }
}
```

---

## 🧠 Naming Conventions (IMPORTANT)

* Use **camelCase** for JSON fields
* Never use:

  * PascalCase (`CreatedAt`)
  * snake_case (`created_at`)

---

## 🚫 Forbidden

* Returning raw JSON without `data`
* Mixing success and error structures
* Inconsistent field naming
* Missing error structure

---

## ✅ Goal

Provide a predictable, stable, and frontend-friendly API response pattern.
