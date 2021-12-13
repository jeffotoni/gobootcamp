import http from 'k6/http';
import { sleep } from 'k6';

const headers = { 'Content-Type': 'application/json' };

// export let options = {
//     vus: 10,
//     duration: '1m',
// };

export default function() {

    http.get(`${__ENV.DOMAIN}/ping`, { headers: headers });
    sleep(1);
    http.get(`${__ENV.DOMAIN}/api/hulk`, { headers: headers });
    sleep(1);

    // const data = {  };
    // http.put(`${__ENV.DOMAIN}/api/hulk`,JSON.stringify(data), { headers: headers } );
}