import Cookies from 'js-cookie';
import { SANDBOX } from '../constants';

type CreateSandboxResp = {
  code: string;
  createdAt: string;
  id: number;
  language: string;
  name: string;
  updatedAt: string;
  userRefer: number;
};

export async function createSandbox({
  name,
  code,
  language,
}: {
  name: string;
  code: string;
  language: string;
}) {
  const token = Cookies.get('jwt-token');
  const body = { name, code, language };
  const resp = await fetch(SANDBOX, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${token}` },
    body: JSON.stringify(body),
  });
  const data: CreateSandboxResp = await resp.json();
  if (!resp.ok) {
    return undefined;
  }
  return data;
}
