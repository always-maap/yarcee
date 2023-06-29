import { getAllSandboxes, getSandbox } from '@/api/sandbox/get-sandbox';
import useSWR from 'swr';

export function useSandbox(id: string) {
  const { data, error, isLoading } = useSWR(`/sandbox/${id}`, () => getSandbox(id));

  return {
    sandbox: data,
    isLoading,
    isError: error,
  };
}
