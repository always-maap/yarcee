import { SIGN_UP } from '../constants';

export async function signUp({
  name,
  username,
  password,
}: {
  name: string;
  username: string;
  password: string;
}) {
  const data = { name, username, password };
  const resp = await fetch(SIGN_UP, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  });

  return resp;
}
