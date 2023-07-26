'use client';

import { useRouter } from 'next/navigation';

import { createSandbox } from '@/api/sandbox/create-sandbox';
import ProjectCard from './project-card';
import { generateUniqueProjectName } from './project-name';
import { TEMPLATES } from './constants';

export default function Templates() {
  const { push } = useRouter();

  async function onTemplate(templateAbbr: string) {
    const name = generateUniqueProjectName();
    const code = TEMPLATES.find((t) => t.abbr === templateAbbr)?.code || '';
    const resp = await createSandbox({ name, language: templateAbbr, code: code });
    if (resp) {
      await push(`/edit/${resp.id}`);
    }
  }

  return (
    <ul className="grid grid-cols-[repeat(auto-fill,_minmax(260px,_1fr))] auto-rows-[minmax(154px,_1fr)] gap-4">
      {TEMPLATES.map((template) => (
        <li onClick={() => onTemplate(template.abbr)} key={template.name}>
          <ProjectCard name={template.name} icon={template.icon} />
        </li>
      ))}
    </ul>
  );
}
