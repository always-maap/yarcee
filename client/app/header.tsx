import Logo from '@/components/logo';
import Container from '@/components/ui/container';
import Link from 'next/link';

export default function Header() {
  return (
    <header>
      <Container>
        <div className="flex justify-between">
          <Link href="/">
            <Logo />
          </Link>
          <ul className="flex gap-6">
            <li>Sign In</li>
            <li>Try for free</li>
          </ul>
        </div>
      </Container>
    </header>
  );
}
