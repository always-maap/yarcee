import Container from '@/components/ui/container';
import AllSandboxes from './all-sandboxes';
import Header from './header';
import Templates from './templates';

export default function Dashboard() {
  return (
    <>
      <Header />
      <main className="my-4">
        <Container className="flex flex-col gap-12">
          <section>
            <h2 className="mb-4">Pick up where you left off</h2>
            <AllSandboxes />
          </section>
          <section>
            <h2 className="mb-4">Start from a template</h2>
            <Templates />
          </section>
        </Container>
      </main>
    </>
  );
}
