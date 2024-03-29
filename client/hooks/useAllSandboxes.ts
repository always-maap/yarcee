import { getAllSandboxes } from '@/api/sandbox/get-sandbox';
import useSWR from 'swr';

export function useAllSandboxes() {
  const { data, error, isLoading } = useSWR(`/all-sandboxes`, getAllSandboxes);

  return {
    sandboxes: data,
    isLoading,
    isError: error,
  };
}
