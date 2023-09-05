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
./bin/air main.go
```
## Dale una estrella ⭐

![Fiber-crud-api](https://github.com/agustfricke/fiber-crud-api/assets/110266171/759c96a0-1152-4a64-9cb0-20badcda7a1a)
