# gogat-api
Api gateway em go

Com o que a Api Gateway lida?

1. Roteamento de Requisições (Request Routing)

Descrição: A API Gateway deve ser capaz de rotear as requisições para o serviço backend correto com base na URL, método HTTP, cabeçalhos, ou outros parâmetros.
Funcionamento na Prática: Quando uma requisição chega à API Gateway, ela analisa a rota solicitada e, com base em uma tabela de roteamento configurada, encaminha a requisição para o serviço apropriado. Por exemplo, uma requisição GET /users pode ser roteada para o serviço de usuários.

```routes:
  - path: /users
    method: GET
    service: user-service
    backend_url: http://user-service.local/api/users

2. Autenticação e Autorização (Authentication and Authorization)
Descrição: A API Gateway deve validar a identidade do cliente (Autenticação) e verificar se ele tem permissão para acessar o recurso solicitado (Autorização).
Funcionamento na Prática: A autenticação pode ser feita usando tokens JWT, onde a API Gateway valida o token, verifica sua validade e extrai as informações de identidade. A autorização pode ser baseada em papéis (RBAC - Role-Based Access Control) ou escopos definidos no token.

```authentication:
  type: jwt
  public_key: /path/to/public.key
  algorithms: [RS256]

authorization:
  roles:
    admin:
      - /admin/*
    user:
      - /users/*

