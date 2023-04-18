'use client';

import YARCEE from '@/components/yarcee';
import Container from '@/components/ui/container';
import Link from 'next/link';
import { useUser } from '@/hooks/useUser';
import Cookies from 'js-cookie';
import { mutate } from 'swr';

export default function Header() {
  const { user, isLoading } = useUser();

  function onSignOut() {
    Cookies.remove('jwt-token');
    mutate('/user');
  }

  return (
    <header className="border-b border-b-neutral-600 py-4" style={{ height: 'var(--nav-height)' }}>
      <Container>
        <div className="flex justify-between items-center">
          <Link href="/">
            <YARCEE />
          </Link>
          {!isLoading && (
            <ul className="flex gap-6 text-neutral-400">
              {user ? (
                <>
                  <li className="text-xs">
                    <button onClick={onSignOut}>Sign out</button>
                  </li>
                  <li className="text-xs">
                    <Link href="/dashboard" className="btn-primary">
                      Dashboard
                    </Link>
                  </li>
                </>
              ) : (
                <>
                  <li className="text-xs">
                    <Link href="signin">Sign In</Link>
                  </li>
                  <li className="text-xs">
                    <Link href="/signup" className="btn-primary">
                      Try for free
                    </Link>
                  </li>
                </>
              )}
            </ul>
          )}
        </div>
      </Container>
    </header>
  );
}
