import Container from '@/components/ui/container';
import Header from './header';

export default function Home() {
  return (
    <>
      <Header />
      <main>
        <div style={{ minHeight: 'calc(100vh - var(--nav-height) * 2 - 90px)' }}>
          <Container>
            <h1 className="capitalize text-7xl md:text-9xl xl:text-[12.5rem] leading-none py-8 tracking-tight">
              <div>code.</div>
              <div>execute.</div>
              <div className="uppercase">yarcee.</div>
            </h1>
          </Container>
        </div>
        <div>Python - C++ - C - Nodejs</div>
      </main>
    </>
  );
}
