'use client';

import { useRouter } from 'next/navigation';

import { createSandbox } from '@/api/sandbox/create-sandbox';
import ProjectCard from './project-card';
import { projectName } from './project-name';
import { TEMPLATES } from './constants';

export default function Templates() {
  const { push } = useRouter();

  async function onTemplate(template: string) {
    const name = projectName();
    const code = TEMPLATES.find((t) => t.name === template)?.code || '';
    const resp = await createSandbox({ name, language: template, code: code });
    if (resp) {
      await push(`/edit/${name}`);
    }
  }

  return (
    <ul className="grid grid-cols-[repeat(auto-fill,_minmax(260px,_1fr))] auto-rows-[minmax(154px,_1fr)] gap-4">
      {TEMPLATES.map((template) => (
        <li onClick={() => onTemplate(template.name)} key={template.name}>
          <ProjectCard name={template.name} icon={template.icon} />
        </li>
      ))}
    </ul>
  );
}
