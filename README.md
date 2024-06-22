# Como Usar?

Defina a variável de ambiente:
```bash
export DIGITALOCEAN_ACCESS_TOKEN=your_access_token_here
```
Execute a aplicação:

````bash
go run main.go
````

Ele deve exibir como saída em formato tabelado as informações solicitadas:
````
+---------------------------------+-----------+-------------+-----------------+---------------------------------------+----------------+
|              NAME               |    ID     | INTERNAL IP |    PUBLIC IP    |                VOLUMES                | COST/MONTH ($) |
+---------------------------------+-----------+-------------+-----------------+---------------------------------------+----------------+
| nome da maquina                 |  00000000 |             | 000.111.222.333 | ID:                                   |         102.86 |
+---------------------------------+-----------+-------------+-----------------+---------------------------------------+----------------+
````


# o que esse Script faz?

- Usa para a autenticação com a API da DigitalOcean um token de acesso.
- Lista todas as droplets da sua conta.
- Para cada droplet, obtém o nome, IPs internos e públicos, volumes anexados e calcula o custo mensal baseado no preço por hora.

> Certifique-se de ter o token de API configurado corretamente e ajuste o código conforme necessário para se adequar às suas necessidades específicas.