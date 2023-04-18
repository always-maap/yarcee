import { SIGN_IN } from '../constants';

type SignInResp = {
  data: string;
  message: string;
};

export async function signIn({ username, password }: { username: string; password: string }) {
  const body = { username, password };
  const resp = await fetch(SIGN_IN, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  });
  const data: SignInResp = await resp.json();
  return data;
}
