# Architecture Diagram

List of architecture diagrams.

## Sequence Diagram

```mermaid
sequenceDiagram
    client->>+auth: POST /auth (login to get token)
    auth->>-client: return token
    client->>+user: POST /users (create user)
    user->>-client: return success create user
    client->>+user: GET /users (list users)
    user->>-client: return list of users
```
