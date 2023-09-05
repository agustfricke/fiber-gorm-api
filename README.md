##### CRUD API CON FIBER 
- Simple CRUD API de Fiber(Go) con Postgres SQL

- Nesesitas tener instalado Go y Docker en tu PC
- Video de como instalar Go en linux -> https://www.youtube.com/watch?v=9LSpTlCZ8xQ&t=52s
- Video instalar Docker en linux -> https://www.youtube.com/watch?v=tb1wAOssVds&t=45s

### Instalar y usar
```bash
git clone https://github.com/agustfricke/fiber-crud-api
cd fiber-crud-api
mv .env.example .env
docker run --name some-postgres -e POSTGRES_USER=agust -e 
POSTGRES_PASSWORD=agust -e POSTGRES_DB=agust -p 5432:5432 -d postgres
go run main.go
```
## Dale una estrella ⭐

![Fiber-rest-api](https://github.com/agustfricke/fiber-crud-api/assets/110266171/10f83ec6-119b-427f-b55c-1d8fcd681483)
