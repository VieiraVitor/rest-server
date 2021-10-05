import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
  http.post('http://localhost:8080/transaction/v1/send');
  sleep(1);
}