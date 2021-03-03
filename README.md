# GO RESTful API with GIN + GORM + PostgreSQL + JWT

Contoh master project Restful API menggunakan golang. Ini hanya project belajar, silahkan jika ada yang ingin mengembangkan project ini.

## Instalation

```bash
go mod download
```

## Running

```bash
go run server.go
```

## Testing

Gunakan [insomnia](https://insomnia.rest/), [postman](https://www.postman.com/), [RapidAPI](https://rapidapi.com/products/api-testing/) atau API Tester lainnya untuk melakukan testing pada project ini. Saat testing, wajib memeberikan attribut header berupa Authorization (token hasil generate saat login) dan X-API-KEY (API Key sesuai dengan yang di set pada file .env).

## License
[MIT](https://choosealicense.com/licenses/mit/)