import YARCEE from '@/components/yarcee';
import Container from '@/components/ui/container';
import Link from 'next/link';

export default function Header() {
  return (
    <header className="border-b border-b-neutral-600 py-4" style={{ height: 'var(--nav-height)' }}>
      <Container>
        <div className="flex justify-between items-center">
          <Link href="/">
            <YARCEE />
          </Link>
          <Link href="signin">Sign out</Link>
        </div>
      </Container>
    </header>
  );
}
