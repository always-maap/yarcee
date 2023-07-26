import { z } from 'zod';
import { ZSandbox } from '@/models/sandbox';
import Cookies from 'js-cookie';
import { SANDBOX } from '../constants';

const ZExecuteSandboxesResp = z.object({ data: ZSandbox, message: z.string() });

export async function executeSandbox({ id, code }: { id: string; code: string }) {
  const token = Cookies.get('jwt-token');
  const body = { code };
  const resp = await fetch(`${SANDBOX}/${id}/execute`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${token}` },
    body: JSON.stringify(body),
  });
  const data = ZExecuteSandboxesResp.parse(await resp.json());
  if (!resp.ok) {
    return undefined;
  }
  return data;
}
