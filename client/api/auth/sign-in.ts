import { SIGN_IN } from '../constants';

export async function signIn({ username, password }: { username: string; password: string }) {
  const data = { username, password };
  const resp = await fetch(SIGN_IN, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  });

  return resp;
}
