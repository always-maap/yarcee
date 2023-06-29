import { z } from 'zod';
import { SIGN_IN } from '../constants';

const ZSignInResp = z.object({ data: z.string(), message: z.string() });

export async function signIn({ username, password }: { username: string; password: string }) {
  const body = { username, password };
  const resp = await fetch(SIGN_IN, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  });
  const data = ZSignInResp.parse(await resp.json());
  return data;
}
