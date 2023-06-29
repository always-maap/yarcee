import { ZUser } from '@/models/user';
import Cookies from 'js-cookie';
import { z } from 'zod';
import { USER } from '../constants';

const ZUserResp = z.union([ZUser, z.undefined()]);

export async function retrieveUser() {
  const token = Cookies.get('jwt-token');
  const resp = await fetch(USER, {
    headers: { Authorization: `Bearer ${token}` },
  });
  const data = ZUserResp.parse(await resp.json());

  if (!data?.id) {
    return undefined;
  }

  return data;
}
