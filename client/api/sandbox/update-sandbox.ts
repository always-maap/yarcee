import { ZSandbox } from '@/models/sandbox';
import Cookies from 'js-cookie';
import { z } from 'zod';
import { SANDBOX } from '../constants';

const ZUpdateSandboxesResp = z.object({ data: ZSandbox, message: z.string() });

export async function updateSandbox({
  id,
  name,
  code,
  language,
}: {
  id: string;
  name: string;
  code: string;
  language: string;
}) {
  const token = Cookies.get('jwt-token');
  const body = { name, code, language };
  const resp = await fetch(`${SANDBOX}/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${token}` },
    body: JSON.stringify(body),
  });
  const data = ZUpdateSandboxesResp.parse(await resp.json());
  if (!resp.ok) {
    return undefined;
  }
  return data;
}
