import Cookies from 'js-cookie';
import { USER } from '../constants';

export async function retrieveUser() {
  const token = Cookies.get('jwt-token');
  const resp = await fetch(USER, {
    headers: { Authorization: `Bearer ${token}` },
  });
  const data = await resp.json();
  return data;
}
