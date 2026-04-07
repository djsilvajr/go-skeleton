# 📘 API Documentation System (AI Contract)

## 🎯 Purpose

This system defines how API documentation must be generated and rendered.

The documentation is **mandatory for every route** and must follow a strict and consistent structure.

---

## 📁 Architecture

* The documentation is a separate Laravel project located in `/docs`
* It must NOT contain business logic
* It only consumes the API endpoints
* Access is restricted to **admin users only**

---

## 🧩 Domain Organization

* All routes must be grouped by **domain**
* Each domain must have its own page

### Examples:

* Auth
* Users
* Roles
* Payments

---

## 🎨 UI / UX नियम (MANDATORY)

The UI must use **Semantic UI** and follow this layout:

### Layout Structure

* Left Sidebar → Domain navigation
* Top Bar:

  * API Name
  * Environment (dev / staging / prod)
  * Token input (for authenticated requests)
* Main Content → Route documentation

---

## 🧱 Component System (STRICT)

The UI must be built using reusable components:

* RouteCard
* RequestTable
* ResponseTable
* ErrorTable
* TryItPanel

⚠️ The AI must NEVER generate raw layout from scratch — only fill data.

---

## 📌 Route Documentation Structure (MANDATORY)

Each route MUST follow this exact structure:

---

### 🔹 1. Header

* Route Name
* Description
* HTTP Method
* Endpoint URL

---

### 🔹 2. Request

#### Headers

| Name | Value | Required |
| ---- | ----- | -------- |

#### Query Params

| Name | Type | Required | Description |

#### Body (JSON)

If no body:

```
No body required
```

---

### 🔹 3. Example Request

Must include:

* JSON
* cURL

---

### 🔹 4. Response

#### Success

* Status Code
* JSON response
* Field description table

---

### 🔹 5. Errors (MANDATORY)

| Code | Message | Description |

⚠️ ALL possible errors must be listed

---

### 🔹 6. Try It (Execution)

Each route must include:

* Input fields based on params/body
* "Send Request" button
* Response viewer:

  * Status
  * Body
  * Execution time

---

## ⚙️ AI Rules (CRITICAL)

The AI MUST:

* ALWAYS generate documentation when creating a new route
* FOLLOW this structure strictly
* NEVER skip sections
* NEVER invent fields
* ALWAYS include:

  * Example request
  * Example response
  * Errors

---

## 🔄 Data Contract (IMPORTANT)

The AI must generate documentation as structured JSON.

### Example:

```json
{
  "domain": "Users",
  "name": "Create User",
  "method": "POST",
  "endpoint": "/api/users",
  "description": "Create a new user",
  "request": {
    "headers": [
      {
        "name": "Authorization",
        "value": "Bearer {token}",
        "required": true
      }
    ],
    "query": [],
    "body": {
      "name": "string",
      "email": "string"
    }
  },
  "response": {
    "success": {
      "status": 201,
      "body": {
        "id": "number",
        "name": "string",
        "email": "string"
      }
    }
  },
  "errors": [
    {
      "code": 401,
      "message": "Unauthorized",
      "description": "Invalid or missing token"
    },
    {
      "code": 422,
      "message": "Validation error",
      "description": "Invalid input data"
    }
  ]
}
```

---

## 🚀 Rendering Rule

* The frontend (Laravel + Blade + Semantic UI) is responsible for rendering
* The AI is ONLY responsible for generating structured data

---

## ❌ Forbidden

* Generating inconsistent formats
* Skipping errors
* Mixing layout with data
* Creating undocumented routes

---

## ✅ Goal

Produce a clean, consistent, scalable, and interactive API documentation system similar to professional tools, but fully custom.
