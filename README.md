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
