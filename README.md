# learning-docker
Just learning docker

## TODO

- [ ] Make a simple go app for making gophers in a container
  - [x] Simple CRUD
  - [x] Auth
  - [ ] Upload (should persist)
  - [ ] Test
- [x] Make a postgres container
- [ ] ?Make an nginx container?

<br>
---

### To build up (first time)

```bash
docker-compose build
```

### To run

```bash
docker-compose up
```

### To stop

```bash
docker-compose down
```

<br>
---

### To login into adminer

Go to `localhost:2020`

Input into form:

|  Field   |     Value    |
|----------|--------------|
| System   | `PostgreSQL` |
| Server   | `postgres`   |
| Username | `mygo`       |
| Password | `secret`     |
| Database | `crm`        |