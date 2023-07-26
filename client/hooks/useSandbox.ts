import { getAllSandboxes, getSandbox } from '@/api/sandbox/get-sandbox';
import { useState } from 'react';
import useSWR from 'swr';

export function useSandbox(id: string) {
  const [isSandboxPending, setIsSandboxPending] = useState(true);
  const { data, error, isLoading } = useSWR(`/sandbox/${id}`, () => getSandbox(id), {
    refreshInterval: isSandboxPending ? 100 : 0,
    onSuccess: (data) => {
      if (data?.status === 'done') {
        setIsSandboxPending(false);
      }
    },
  });

  return {
    sandbox: data,
    isLoading,
    isError: error,
  };
}
