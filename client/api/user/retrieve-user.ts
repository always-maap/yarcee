import Cookies from 'js-cookie';
import { USER } from '../constants';

type UserResp =
  | {
      id: string;
      name: string;
      username: string;
    }
  | undefined;

export async function retrieveUser() {
  const token = Cookies.get('jwt-token');
  const resp = await fetch(USER, {
    headers: { Authorization: `Bearer ${token}` },
  });
  const data: UserResp = await resp.json();

  if (!data?.id) {
    return undefined;
  }

  return data;
}
