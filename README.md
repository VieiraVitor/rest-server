lib go-chi:
https://github.com/go-chi/chi

Hey benchmark test:
hey -n 10000 -c 1000 -m POST -d '{"id": 1, "amount": 1000, "description": "Novo produto", "creditCard": { "number": "1234", "name": "Vitor", "cvv": 123 }}' http://localhost:8080/transaction/v1/send