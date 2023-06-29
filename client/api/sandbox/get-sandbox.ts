import { z } from 'zod';
import { ZSandbox } from '@/models/sandbox';
import Cookies from 'js-cookie';
import { SANDBOX } from '../constants';

const ZGetAllSandboxesResp = z.object({ data: z.array(ZSandbox), message: z.string() });

export async function getAllSandboxes() {
  const token = Cookies.get('jwt-token');
  const resp = await fetch(SANDBOX, {
    headers: { Authorization: `Bearer ${token}` },
  });
  const data = ZGetAllSandboxesResp.parse(await resp.json());
  if (!resp.ok) {
    return undefined;
  }
  return data.data;
}

const ZGetSandboxResp = z.object({ data: ZSandbox, message: z.string() });

export async function getSandbox(id: string) {
  const token = Cookies.get('jwt-token');
  const resp = await fetch(`${SANDBOX}/${id}`, {
    headers: { Authorization: `Bearer ${token}` },
  });
  const data = ZGetSandboxResp.parse(await resp.json());
  if (!resp.ok) {
    return undefined;
  }
  return data.data;
}
