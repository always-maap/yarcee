'use client';

import YARCEE from '@/components/yarcee';
import Container from '@/components/ui/container';
import Link from 'next/link';
import Cookies from 'js-cookie';
import { mutate } from 'swr';

export default function Header() {
  function onSignOut() {
    Cookies.remove('jwt-token');
    mutate('/user');
  }

  return (
    <header className="border-b border-b-neutral-600 py-4" style={{ height: 'var(--nav-height)' }}>
      <Container>
        <div className="flex justify-between items-center">
          <Link href="/dashboard">
            <YARCEE />
          </Link>
          <button className="text-xs text-neutral-400" onClick={onSignOut}>
            Sign out
          </button>
        </div>
      </Container>
    </header>
  );
}
