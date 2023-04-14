'use client';

import Container from '@/components/ui/container';
import Header from './header';
import ProjectCard from './project-card';
import Nodejs from '@/components/nodejs';
import Image from 'next/image';
import { useEffect } from 'react';
import { retrieveUser } from '@/api/user/retrieve-user';

const TEMPLATES = [
  { name: 'Node.js', icon: <Nodejs /> },
  { name: 'Python', icon: <Image width={20} height={20} src="/logo/python.png" alt="python-logo" /> },
];

export default function Dashboard() {
  useEffect(() => {
    (async () => {
      const resp = await retrieveUser();
      console.log(resp);
    })();
  }, []);

  return (
    <>
      <Header />
      <main className="my-4">
        <Container className="flex flex-col gap-12">
          <section>
            <h2>Pick up where you left off</h2>
          </section>
          <section>
            <h2 className="mb-4">Start from a template</h2>
            <ul className="grid grid-cols-[repeat(auto-fill,_minmax(260px,_1fr))] auto-rows-[minmax(154px,_1fr)] gap-4">
              {TEMPLATES.map((template) => (
                <li key={template.name}>
                  <ProjectCard name={template.name} icon={template.icon} />
                </li>
              ))}
            </ul>
          </section>
        </Container>
      </main>
    </>
  );
}
