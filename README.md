lib go-chi:
https://github.com/go-chi/chi

Hey benchmark test:
hey -n 10000 -c 1000 -m POST -d '{"id": 1, "amount": 1000, "description": "Novo produto", "creditCard": { "number": "1234", "name": "Vitor", "cvv": 123 }}' http://localhost:8080/transaction/v1/send

hey -n 100 -c 10 -m POST -d '{"id": 1, "amount": 1000, "description": "Novo produto", "creditCard": { "number": "1234", "name": "Vitor", "cvv": 123 }}' http://localhost:50052/transaction/v1/send

teste1:
hey -n 1000 -c 100 -m POST -d '{"id": 1, "amount": 1000, "description": "Novo produto", "creditCard": { "number": "1234", "name": "Vitor", "cvv": 123 }}' http://localhost:50052/transaction/v1/send

teste2:
hey -n 10000 -c 1000 -m POST -d '{"id": 1, "amount": 1000, "description": "Novo produto", "creditCard": { "number": "1234", "name": "Vitor", "cvv": 123 }}' http://localhost:50052/transaction/v1/send

teste3:
hey -n 100000 -c 10000 -m POST -d '{"id": 1, "amount": 1000, "description": "Novo produto", "creditCard": { "number": "1234", "name": "Vitor", "cvv": 123 }}' http://localhost:50052/transaction/v1/send

teste4:
hey -z 10s -c 1000 -m POST -d '{"id": 1, "amount": 1000, "description": "Novo produto", "creditCard": { "number": "1234", "name": "Vitor", "cvv": 123 }}' http://localhost:50052/transaction/v1/send