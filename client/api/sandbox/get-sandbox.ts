import Cookies from 'js-cookie';
import { SANDBOX } from '../constants';

type GetSandboxResp = {
  data: {
    code: string;
    createdAt: string;
    id: number;
    language: string;
    name: string;
    updatedAt: string;
    userRefer: number;
  }[];
  message: string;
};

export async function getAllSandboxes() {
  const token = Cookies.get('jwt-token');
  const resp = await fetch(SANDBOX, {
    headers: { Authorization: `Bearer ${token}` },
  });
  const data: GetSandboxResp = await resp.json();
  if (!resp.ok) {
    return undefined;
  }
  return data.data;
}
