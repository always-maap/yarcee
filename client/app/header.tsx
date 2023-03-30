import Logo from '@/components/logo';
import Container from '@/components/ui/container';
import Link from 'next/link';

export default function Header() {
  return (
    <header className="border-b border-b-neutral-600 py-4" style={{ height: 'var(--nav-height)' }}>
      <Container>
        <div className="flex justify-between items-center">
          <Link href="/">
            <Logo />
          </Link>
          <ul className="flex gap-6 text-neutral-400">
            <li className="text-xs">Sign In</li>
            <li className="text-xs">Try for free</li>
          </ul>
        </div>
      </Container>
    </header>
  );
}
