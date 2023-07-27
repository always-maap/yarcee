'use client';

import { deleteSandbox } from '@/api/sandbox/delete-sandbox';
import { useAllSandboxes } from '@/hooks/useAllSandboxes';
import Link from 'next/link';
import { mutate } from 'swr';
import { TEMPLATES } from './constants';
import ProjectCard from './project-card';

export default function AllSandboxes() {
  const { sandboxes, isLoading } = useAllSandboxes();

  if (isLoading || !sandboxes) {
    return <div>loading...</div>;
  }

  async function onDelete(id: string) {
    await deleteSandbox({ id });
    mutate('/all-sandboxes');
  }

  return (
    <ul className="grid grid-cols-[repeat(auto-fill,_minmax(260px,_1fr))] auto-rows-[minmax(154px,_1fr)] gap-4">
      {sandboxes.map((sandbox) => {
        const icon = TEMPLATES.find((t) => t.abbr === sandbox.language)?.icon;

        return (
          <li key={sandbox.id}>
            <ProjectCard
              name={<Link href={`/edit/${sandbox.id}`}>{sandbox.name}</Link>}
              icon={icon}
              actions={<button onClick={() => onDelete(String(sandbox.id))}>delete</button>}
            />
          </li>
        );
      })}
    </ul>
  );
}
