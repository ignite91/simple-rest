import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  vus: 10,
  duration: '10s',
};
export default function () {
  http.get('http://localhost:8080/user');
  sleep(1);
}
