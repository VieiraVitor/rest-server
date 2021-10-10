import http from 'k6/http';
import { sleep } from 'k6';
import { check } from 'k6';
export let options = {
	vus: 10000,
	iterations: 100000,
}

const url = 'http://127.0.0.1:8080/transaction/v1/send';
const headers = {
	headers: { 'Content-Type': 'application/json' }
}
const data = {
	"id": 1,
	"amount": 1000,
	"description": "Novo produto",
	"creditCard": {
		"number": "1234",
		"name": "Vitor",
		"cvv": 123
	}
}


export default function () {
	const res = http.post(url, JSON.stringify(data), { headers });

	check(res, {
		'is status 200': (r) => r.status === 200,
	});

	sleep(1);
}