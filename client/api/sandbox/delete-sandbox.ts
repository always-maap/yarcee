import Cookies from 'js-cookie';
import { z } from 'zod';
import { SANDBOX } from '../constants';

const ZDeleteSandboxResp = z.object({ message: z.string() });

export async function deleteSandbox({ id }: { id: string }) {
  const token = Cookies.get('jwt-token');
  const resp = await fetch(`${SANDBOX}/${id}`, {
    method: 'DELETE',
    headers: { Authorization: `Bearer ${token}` },
  });
  const data = ZDeleteSandboxResp.parse(await resp.json());
  if (!resp.ok) {
    return undefined;
  }
  return data;
}
