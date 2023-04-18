import Cookies from 'js-cookie';
import { SANDBOX } from '../constants';

type DeleteSandboxResp = {
  message: string;
};

export async function deleteSandbox({ id }: { id: string }) {
  const token = Cookies.get('jwt-token');
  const resp = await fetch(`${SANDBOX}/${id}`, {
    method: 'DELETE',
    headers: { Authorization: `Bearer ${token}` },
  });
  const data: DeleteSandboxResp = await resp.json();
  if (!resp.ok) {
    return undefined;
  }
  return data;
}
