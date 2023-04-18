import { retrieveUser } from '@/api/user/retrieve-user';
import useSWR from 'swr';

export function useUser() {
  const { data, error, isLoading } = useSWR(`/user`, retrieveUser);

  return {
    user: data,
    isLoading,
    isError: error,
  };
}
